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
		{name: "code-block", got: GetCodeBlock("fmt.Println(`x`) \\\nnext").ToString(), want: "```\nfmt.Println(\\`x\\`) \\\\\nnext\n```"},
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

func TestToWotoMD(t *testing.T) {
	if got := toWotoMD(""); got != nil {
		t.Fatalf("toWotoMD(\"\") = %#v, want nil", got)
	}

	if got := toWotoMD("value"); got == nil || got.ToString() != "value" {
		t.Fatalf("toWotoMD(\"value\") = %#v, want markdown with value", got)
	}
}

func TestInternalFormattingHelpersWithEmptyInput(t *testing.T) {
	cases := []struct {
		name string
		got  string
		want string
	}{
		{name: "toNormal", got: toNormal(""), want: ""},
		{name: "toBold", got: toBold(""), want: ""},
		{name: "toItalic", got: toItalic(""), want: "__"},
		{name: "toMono", got: toMono(""), want: ""},
		{name: "toCodeBlock", got: toCodeBlock(""), want: ""},
		{name: "toSpoiler", got: toSpoiler(""), want: ""},
		{name: "toUnderline", got: toUnderline(""), want: ""},
		{name: "toStrike", got: toStrike(""), want: ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got != tc.want {
				t.Fatalf("%s = %q, want %q", tc.name, tc.got, tc.want)
			}
		})
	}
}
