# This powershell script goes through all Go files in the target path and makes sure
# the content inside of them adheres to the following rules:
#  1. *types*.go files should ONLY contain type definitions (structs, interfaces, type aliases, etc.)
#  2. *helpers*.go files should ONLY contain helper functions (functions that are not methods and do not belong to a type)
#  3. *methods*.go files should ONLY contain methods (functions that belong to a type)
#  4. *constants*.go files should ONLY contain constant definitions (const blocks or individual const declarations)
#  5. *vars*.go files should ONLY contain variable definitions (var blocks or individual var declarations)
#  6. handlers.go or handlers_<category>.go files should ONLY contain functions/methods
#     whose name contains "handle"

param (
    [string]$TargetPath = "mdparser"
)

$testsPattern = "*_test.go"

$typesFiles = Get-ChildItem -Path $TargetPath -Recurse -Include "*types*.go" -Exclude $testsPattern
$helpersFiles = Get-ChildItem -Path $TargetPath -Recurse -Include "*helpers*.go" -Exclude $testsPattern
$methodsFiles = Get-ChildItem -Path $TargetPath -Recurse -Include "*methods*.go" -Exclude $testsPattern
$constantsFiles = Get-ChildItem -Path $TargetPath -Recurse -Include "*constants*.go" -Exclude $testsPattern
$varsFiles = Get-ChildItem -Path $TargetPath -Recurse -Include "*vars*.go" -Exclude $testsPattern
$handlersFiles = Get-ChildItem -Path $TargetPath -Recurse -File -Exclude $testsPattern | Where-Object {
    $_.Name -match '^handlers(?:_[A-Za-z0-9]+)?\.go$'
}


# Note: These assume standard formatting (gofmt) where definitions start at the beginning of the line
$pType = '^\s*type\s+\w+'   # Matches "type Name"
$pMethod = '^\s*func\s+\('  # Matches "func (receiver)"
$pHelper = '^\s*func\s+\w+' # Matches "func Name" (no receiver)
$pConst = '^\s*const\s+'    # Matches "const"
$pVar = '^\s*var\s+'        # Matches "var"
$pNonHandleFunc = '^\s*func\s+(\([^)]*\)\s*)?(?!\w*handle\w*)\w+'


function Get-StrippedGoLine {
    param (
        [string]$Line,
        [ref]$State
    )

    $out = New-Object System.Text.StringBuilder
    $i = 0

    while ($i -lt $Line.Length) {
        # If we're inside a block comment, consume until it ends.
        if ($State.Value.InBlockComment) {
            $end = $Line.IndexOf("*/", $i)
            if ($end -lt 0) {
                return $out.ToString()
            }
            $State.Value.InBlockComment = $false
            $i = $end + 2
            continue
        }

        # If we're inside a raw string, consume until the next backtick.
        if ($State.Value.InRawString) {
            $end = $Line.IndexOf('`', $i)
            if ($end -lt 0) {
                return $out.ToString()
            }
            $State.Value.InRawString = $false
            $i = $end + 1
            continue
        }

        $ch = $Line[$i]
        $next = if ($i + 1 -lt $Line.Length) { $Line[$i + 1] } else { [char]0 }

        # Line comment
        if ($ch -eq '/' -and $next -eq '/') {
            break
        }

        # Block comment start
        if ($ch -eq '/' -and $next -eq '*') {
            $State.Value.InBlockComment = $true
            $i += 2
            continue
        }

        # Raw string start
        if ($ch -eq '`') {
            $State.Value.InRawString = $true
            $i++
            continue
        }

        # Interpreted string or rune literal (single line only)
        if ($ch -eq '"' -or $ch -eq "'") {
            $delim = $ch
            $i++
            while ($i -lt $Line.Length) {
                $c = $Line[$i]
                if ($c -eq '\') {
                    $i += 2
                    continue
                }
                if ($c -eq $delim) {
                    $i++
                    break
                }
                $i++
            }
            continue
        }

        $null = $out.Append($ch)
        $i++
    }

    return $out.ToString()
}

# 4. Define the validation logic
function Test-GoRules {
    param (
        [System.IO.FileInfo[]]$Files,
        [string[]]$ForbiddenPatterns,
        [string]$RuleName
    )

    foreach ($file in $Files) {
        $lines = Get-Content -Path $file.FullName
        $lineNumber = 0
        $braceDepth = 0
        $parenDepth = 0
        $state = @{
            InBlockComment = $false
            InRawString    = $false
        }

        foreach ($line in $lines) {
            $lineNumber++
            $stripped = Get-StrippedGoLine -Line $line -State ([ref]$state)

            # Skip empty/whitespace lines after stripping comments/strings
            if ($stripped -match '^\s*$') {
                continue
            }

            # Only check at the top level (not inside any () or {} blocks)
            if ($braceDepth -eq 0 -and $parenDepth -eq 0) {
                foreach ($pattern in $ForbiddenPatterns) {
                    if ($stripped -match $pattern) {
                        $correctPath = $file.FullName.Replace("$PWD", "").Replace("\", "/").TrimStart("/")
                        Write-Host "[$RuleName Violation] $($correctPath):$lineNumber" -ForegroundColor Red
                        Write-Host "  Found forbidden content: $line" -ForegroundColor DarkGray
                    }
                }
            }

            # Update depth counters based on stripped content
            $braceDepth += ([regex]::Matches($stripped, '\{').Count - [regex]::Matches($stripped, '\}').Count)
            $parenDepth += ([regex]::Matches($stripped, '\(').Count - [regex]::Matches($stripped, '\)').Count)

            if ($braceDepth -lt 0) { $braceDepth = 0 }
            if ($parenDepth -lt 0) { $parenDepth = 0 }
        }
    }
}

Write-Host "Starting Go Content Enforcement..." -ForegroundColor Cyan

# ---------------------------------------------------------
# Rule 1: Types files should NOT contain methods, helpers, vars, or consts
# ---------------------------------------------------------
Test-GoRules -Files $typesFiles `
    -ForbiddenPatterns @($pMethod, $pHelper, $pVar, $pConst) `
    -RuleName "*types*.go"

# ---------------------------------------------------------
# Rule 2: Helpers files should NOT contain types, methods, vars, or consts
# ---------------------------------------------------------
Test-GoRules -Files $helpersFiles `
    -ForbiddenPatterns @($pType, $pMethod, $pVar, $pConst) `
    -RuleName "*helpers*.go"

# ---------------------------------------------------------
# Rule 3: Methods files should NOT contain types, helpers, vars, or consts
# ---------------------------------------------------------
Test-GoRules -Files $methodsFiles `
    -ForbiddenPatterns @($pType, $pHelper, $pVar, $pConst) `
    -RuleName "*methods*.go"

# ---------------------------------------------------------
# Rule 4: Constants files should NOT contain types, methods, helpers, or vars
# ---------------------------------------------------------
Test-GoRules -Files $constantsFiles `
    -ForbiddenPatterns @($pType, $pMethod, $pHelper, $pVar) `
    -RuleName "*constants*.go"

# ---------------------------------------------------------
# Rule 5: Vars files should NOT contain types, methods, helpers, or consts
# ---------------------------------------------------------
Test-GoRules -Files $varsFiles `
    -ForbiddenPatterns @($pType, $pMethod, $pHelper, $pConst) `
    -RuleName "*vars*.go"

# ---------------------------------------------------------
# Rule 6: Handlers files should ONLY contain handle* functions/methods
# ---------------------------------------------------------
Test-GoRules -Files $handlersFiles `
    -ForbiddenPatterns @($pType, $pVar, $pConst, $pNonHandleFunc) `
    -RuleName "handlers*.go"

Write-Host "Check complete." -ForegroundColor Green

