package mdparser

import "testing"

func resetSecrets(t *testing.T) {
	t.Helper()

	original := secrets
	secrets = nil

	t.Cleanup(func() {
		secrets = original
	})
}

func TestSecretLifecycle(t *testing.T) {
	resetSecrets(t)

	AddSecret("token-1", "$TOKEN")
	AddSecret("token-2", "$PASSWORD")
	AddSecret("token-3", "$KEEP")

	if !SecretValueExists("token-1") {
		t.Fatal("expected token-1 to exist")
	}

	if got := GetSecretIndexByValue("token-1"); got != 0 {
		t.Fatalf("GetSecretIndexByValue(token-1) = %d, want 0", got)
	}

	AddSecret("token-1", "$UPDATED")

	if len(secrets) != 3 {
		t.Fatalf("len(secrets) = %d, want 3", len(secrets))
	}

	if secrets[0].name != "$UPDATED" {
		t.Fatalf("updated secret name = %q, want %q", secrets[0].name, "$UPDATED")
	}

	RemoveSecretByValue("token-2")

	if SecretValueExists("token-2") {
		t.Fatal("expected token-2 to be removed")
	}

	RemoveSecretByName("$UPDATED")

	if len(secrets) != 1 {
		t.Fatalf("len(secrets) = %d, want 1", len(secrets))
	}

	if secrets[0].name != "$KEEP" {
		t.Fatalf("remaining secret name = %q, want %q", secrets[0].name, "$KEEP")
	}

	RemoveSecretByName("$KEEP")

	if len(secrets) != 0 {
		t.Fatalf("len(secrets) = %d, want 0", len(secrets))
	}

	if got := GetSecretIndexByValue("missing"); got != -1 {
		t.Fatalf("GetSecretIndexByValue(missing) = %d, want -1", got)
	}
}

func TestRepairValueCensorsSecretsAndEscapesMarkdown(t *testing.T) {
	resetSecrets(t)
	AddSecret("token-123", "$TOKEN")

	got := repairValue("token-123 *._[]()!~|\\")
	want := "$TOKEN \\*\\.\\_\\[\\]\\(\\)\\!\\~\\|\\\\"

	if got != want {
		t.Fatalf("repairValue() = %q, want %q", got, want)
	}
}

func TestRepairCodeValueCensorsSecretsAndEscapesCodeFenceChars(t *testing.T) {
	resetSecrets(t)
	AddSecret("token-123", "$TOKEN")

	got := repairCodeValue("token-123 `code` \\ path")
	want := "$TOKEN \\`code\\` \\\\ path"

	if got != want {
		t.Fatalf("repairCodeValue() = %q, want %q", got, want)
	}
}

func TestSecretsAreAppliedWhenSegmentsAreCreated(t *testing.T) {
	resetSecrets(t)

	md := GetNormal("token-123")
	AddSecret("token-123", "$TOKEN")

	if got := md.ToString(); got != "token\\-123" {
		t.Fatalf("ToString() after late AddSecret = %q, want %q", got, "token\\-123")
	}
}

func TestNormalizeCodeLanguage(t *testing.T) {
	got := normalizeCodeLanguage("  go\t\r\n")
	want := "go"

	if got != want {
		t.Fatalf("normalizeCodeLanguage() = %q, want %q", got, want)
	}

	got = normalizeCodeLanguage("c`++\\17")
	want = "c\\`++\\\\17"

	if got != want {
		t.Fatalf("normalizeCodeLanguage() escaped = %q, want %q", got, want)
	}
}

func TestFormattingConstructors(t *testing.T) {
	input := "a*b_[c](d)!~|\\"
	repaired := "a\\*b\\_\\[c\\]\\(d\\)\\!\\~\\|\\\\"

	cases := []struct {
		name string
		got  string
		want string
	}{
		{name: "empty", got: GetEmpty().ToString(), want: ""},
		{name: "normal", got: GetNormal(input).ToString(), want: repaired},
		{name: "bold", got: GetBold(input).ToString(), want: "*" + repaired + "*"},
		{name: "italic", got: GetItalic(input).ToString(), want: "_" + repaired + "_"},
		{name: "mono", got: GetMono(input).ToString(), want: "`" + repaired + "`"},
		{name: "styled-bold-italic", got: GetStyled(input, StyleBold, StyleItalic).ToString(), want: "*_" + repaired + "_*"},
		{name: "styled-italic-underline", got: GetStyled(input, StyleItalic, StyleUnderline).ToString(), want: "___" + repaired + "_\r__"},
		{name: "styled-all", got: GetStyled(input, StyleBold, StyleItalic, StyleUnderline, StyleStrike, StyleSpoiler).ToString(), want: "||*~___" + repaired + "_\r__~*||"},
		{name: "styled-empty-styles", got: GetStyled(input).ToString(), want: repaired},
		{name: "code-block", got: GetCodeBlock("fmt.Println(`x`) \\\nnext").ToString(), want: "```\nfmt.Println(\\`x\\`) \\\\\nnext\n```"},
		{name: "code-block-lang", got: GetCodeBlockLang("go", "fmt.Println(`x`)").ToString(), want: "```go\nfmt.Println(\\`x\\`)\n```"},
		{name: "code-block-lang-empty-lang", got: GetCodeBlockLang("", "fmt.Println(`x`)").ToString(), want: "```\nfmt.Println(\\`x\\`)\n```"},
		{name: "code-block-lang-normalized-lang", got: GetCodeBlockLang(" c`++\\17 \n", "code").ToString(), want: "```c\\`++\\\\17\ncode\n```"},
		{name: "spoiler", got: GetSpoiler(input).ToString(), want: "||" + repaired + "||"},
		{name: "underline", got: GetUnderline(input).ToString(), want: "__" + repaired + "__"},
		{name: "strike", got: GetStrike(input).ToString(), want: "~" + repaired + "~"},
		{name: "hyperlink", got: GetHyperLink("a*b", "https://example.com/x_y").ToString(), want: "[a\\*b](https://example\\.com/x\\_y)"},
		{name: "mention", got: GetUserMention("a*b", 42).ToString(), want: "[a\\*b](tg://user?id=42)"},
		{name: "mention-base-index", got: GetUserMention("a*b", baseIndex).ToString(), want: "`a\\*b`"},
		{name: "empty-normal", got: GetNormal("").ToString(), want: ""},
		{name: "empty-bold", got: GetBold("").ToString(), want: ""},
		{name: "empty-italic", got: GetItalic("").ToString(), want: ""},
		{name: "empty-mono", got: GetMono("").ToString(), want: ""},
		{name: "empty-code-block", got: GetCodeBlock("").ToString(), want: ""},
		{name: "empty-code-block-lang", got: GetCodeBlockLang("go", "").ToString(), want: ""},
		{name: "empty-spoiler", got: GetSpoiler("").ToString(), want: ""},
		{name: "empty-underline", got: GetUnderline("").ToString(), want: ""},
		{name: "empty-strike", got: GetStrike("").ToString(), want: ""},
		{name: "empty-hyperlink", got: GetHyperLink("", "https://example.com").ToString(), want: ""},
		{name: "empty-mention", got: GetUserMention("", 42).ToString(), want: ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got != tc.want {
				t.Fatalf("%s = %q, want %q", tc.name, tc.got, tc.want)
			}
		})
	}
}

func TestFormattingConstructorsAcceptMixedValues(t *testing.T) {
	cases := []struct {
		name string
		got  string
		want string
	}{
		{name: "normal-sprintf", got: GetNormal("value=%d/%s", 7, "x").ToString(), want: "value\\=7/x"},
		{name: "bold-mixed-values", got: GetBold("hello", " ", 42).ToString(), want: "*hello 42*"},
		{name: "italic-appends-markdown-between-text-chunks", got: GetItalic("hello ", GetBold("world"), "!", 7).ToString(), want: "_hello _*world*_\\!7_"},
		{name: "mono-skips-typed-nil-markdown", got: func() string {
			var typedNil *wotoMarkDown
			return GetMono("x", typedNil, "y").ToString()
		}(), want: "`x``y`"},
		{name: "code-block-sprintf", got: GetCodeBlock("fmt.Println(%d)", 42).ToString(), want: "```\nfmt.Println(42)\n```"},
		{name: "spoiler-appends-markdown", got: GetSpoiler("pre", GetUnderline("mid"), "post").ToString(), want: "||pre||__mid__||post||"},
		{name: "underline-stringifies-non-string-values", got: GetUnderline(true, 9).ToString(), want: "__true 9__"},
		{name: "strike-sprintf", got: GetStrike("sum=%d", 12).ToString(), want: "~sum\\=12~"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got != tc.want {
				t.Fatalf("%s = %q, want %q", tc.name, tc.got, tc.want)
			}
		})
	}
}

func TestFormattingConstructorsKeepEmptyResultForEmptyOrInvalidInput(t *testing.T) {
	cases := []struct {
		name string
		got  string
	}{
		{name: "normal-no-args", got: GetNormal().ToString()},
		{name: "normal-empty", got: GetNormal("").ToString()},
		{name: "bold-no-args", got: GetBold().ToString()},
		{name: "bold-empty", got: GetBold("").ToString()},
		{name: "italic-no-args", got: GetItalic().ToString()},
		{name: "italic-empty", got: GetItalic("").ToString()},
		{name: "mono-no-args", got: GetMono().ToString()},
		{name: "mono-empty", got: GetMono("").ToString()},
		{name: "code-block-no-args", got: GetCodeBlock().ToString()},
		{name: "code-block-empty", got: GetCodeBlock("").ToString()},
		{name: "spoiler-no-args", got: GetSpoiler().ToString()},
		{name: "spoiler-empty", got: GetSpoiler("").ToString()},
		{name: "underline-no-args", got: GetUnderline().ToString()},
		{name: "underline-empty", got: GetUnderline("").ToString()},
		{name: "strike-no-args", got: GetStrike().ToString()},
		{name: "strike-empty", got: GetStrike("").ToString()},
		{name: "typed-nil-markdown", got: func() string {
			var typedNil *wotoMarkDown
			return GetBold(typedNil).ToString()
		}()},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got != "" {
				t.Fatalf("%s = %q, want empty string", tc.name, tc.got)
			}
		})
	}
}

func TestToUnformattedString(t *testing.T) {
	cases := []struct {
		name string
		md   WMarkDown
		want string
	}{
		{name: "empty", md: GetEmpty(), want: ""},
		{name: "normal", md: GetNormal("a*b"), want: "a*b"},
		{name: "bold", md: GetBold("a*b"), want: "a*b"},
		{name: "italic", md: GetItalic("a*b"), want: "a*b"},
		{name: "mono", md: GetMono("a*b"), want: "a*b"},
		{name: "styled", md: GetStyled("a*b", StyleBold, StyleItalic), want: "a*b"},
		{name: "code-block", md: GetCodeBlock("fmt.Println(`x`)"), want: "fmt.Println(`x`)"},
		{name: "code-block-lang", md: GetCodeBlockLang("go", "fmt.Println(`x`)"), want: "fmt.Println(`x`) (go)"},
		{name: "spoiler", md: GetSpoiler("a*b"), want: "a*b"},
		{name: "underline", md: GetUnderline("a*b"), want: "a*b"},
		{name: "strike", md: GetStrike("a*b"), want: "a*b"},
		{name: "hyperlink", md: GetHyperLink("click", "https://example.com"), want: "click (https://example.com)"},
		{name: "mention", md: GetUserMention("user", 42), want: "user (42)"},
		{name: "mixed-builder", md: GetNormal("hello ").Append(GetBold("world")).Spoiler("!"), want: "hello world!"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.md.ToUnformattedString(); got != tc.want {
				t.Fatalf("%s = %q, want %q", tc.name, got, tc.want)
			}
		})
	}
}

func TestToUnformattedStringReturnsStoredRawSegmentsAsIs(t *testing.T) {
	md := newWotoMD(newRawSegment("*raw*"))

	if got := md.ToUnformattedString(); got != "*raw*" {
		t.Fatalf("ToUnformattedString() = %q, want %q", got, "*raw*")
	}
}

func TestIsSpecial(t *testing.T) {
	if !IsSpecial('*') {
		t.Fatal("expected '*' to be special")
	}

	if !IsSpecial('\\') {
		t.Fatal("expected '\\' to be special")
	}

	if IsSpecial('a') {
		t.Fatal("expected 'a' to be non-special")
	}
}

func TestNormalizeTextStyles(t *testing.T) {
	got := normalizeTextStyles(StyleBold, StyleItalic, StyleBold, TextStyle(0xff))
	want := StyleBold | StyleItalic | StyleUnderline | StyleStrike | StyleSpoiler

	if got != want {
		t.Fatalf("normalizeTextStyles() = %v, want %v", got, want)
	}
}

func TestRenderStyledValue(t *testing.T) {
	got := renderStyledValue("a*b", StyleBold|StyleItalic)
	want := "*_a\\*b_*"

	if got != want {
		t.Fatalf("renderStyledValue() = %q, want %q", got, want)
	}

	got = renderStyledValue("a*b", StyleItalic|StyleUnderline)
	want = "___a\\*b_\r__"

	if got != want {
		t.Fatalf("renderStyledValue() italic+underline = %q, want %q", got, want)
	}
}
