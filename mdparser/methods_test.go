package mdparser

import "testing"

func TestAppendCombinesMarkdownValues(t *testing.T) {
	base := GetNormal("hello ").(*wotoMarkDown)
	addition := GetBold("world")

	got := base.Append(addition)
	if got == nil {
		t.Fatal("Append returned nil")
	}

	if got == base {
		t.Fatal("Append returned the original receiver")
	}

	if got.ToString() != "hello *world*" {
		t.Fatalf("Append() = %q, want %q", got.ToString(), "hello *world*")
	}

	if base.ToString() != "hello " {
		t.Fatalf("base changed to %q, want %q", base.ToString(), "hello ")
	}
}

func TestAppendRejectsNilMarkdown(t *testing.T) {
	base := GetNormal("hello").(*wotoMarkDown)

	if got := base.Append(nil); got != nil {
		t.Fatalf("Append(nil) = %#v, want nil", got)
	}
}

func TestAppendThisMutatesReceiver(t *testing.T) {
	base := GetNormal("hello ").(*wotoMarkDown)
	got := base.AppendThis(GetItalic("world"))

	if got != base {
		t.Fatal("AppendThis should return the receiver")
	}

	if base.ToString() != "hello _world_" {
		t.Fatalf("AppendThis() = %q, want %q", base.ToString(), "hello _world_")
	}
}

func TestCloneReturnsIndependentCopy(t *testing.T) {
	base := GetNormal("base").(*wotoMarkDown)
	clone := base.Clone()

	if clone == nil {
		t.Fatal("Clone returned nil")
	}

	if clone == base {
		t.Fatal("Clone returned the original receiver")
	}

	if clone.ToString() != "base" {
		t.Fatalf("Clone() = %q, want %q", clone.ToString(), "base")
	}

	clone.Bold("!")

	if clone.ToString() != "base*\\!*" {
		t.Fatalf("mutated clone = %q, want %q", clone.ToString(), "base*\\!*")
	}

	if base.ToString() != "base" {
		t.Fatalf("base changed to %q, want %q", base.ToString(), "base")
	}
}

func TestCloneProvidesLegacyCopyStyleMigrationPath(t *testing.T) {
	cases := []struct {
		name string
		call func(WMarkDown) WMarkDown
		want string
	}{
		{name: "normal", call: func(m WMarkDown) WMarkDown { return m.Clone().Normal("text") }, want: "basetext"},
		{name: "bold", call: func(m WMarkDown) WMarkDown { return m.Clone().Bold("text") }, want: "base*text*"},
		{name: "italic", call: func(m WMarkDown) WMarkDown { return m.Clone().Italic("text") }, want: "base_text_"},
		{name: "mono", call: func(m WMarkDown) WMarkDown { return m.Clone().Mono("text") }, want: "base`text`"},
		{name: "underline", call: func(m WMarkDown) WMarkDown { return m.Clone().Underline("text") }, want: "base__text__"},
		{name: "strike", call: func(m WMarkDown) WMarkDown { return m.Clone().Strike("text") }, want: "base~text~"},
		{name: "hyperlink", call: func(m WMarkDown) WMarkDown { return m.Clone().HyperLink("text", "https://example.com") }, want: "base[text](https://example\\.com)"},
		{name: "mention", call: func(m WMarkDown) WMarkDown { return m.Clone().Mention("user", 42) }, want: "base[user](tg://user?id=42)"},
		{name: "spoiler", call: func(m WMarkDown) WMarkDown { return m.Clone().Spoiler("text") }, want: "base||text||"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			base := GetNormal("base").(*wotoMarkDown)
			got := tc.call(base)

			if got == nil {
				t.Fatal("expected a markdown value")
			}

			if got == base {
				t.Fatal("clone-based call returned the receiver")
			}

			if got.ToString() != tc.want {
				t.Fatalf("%s = %q, want %q", tc.name, got.ToString(), tc.want)
			}

			if base.ToString() != "base" {
				t.Fatalf("base changed to %q, want %q", base.ToString(), "base")
			}
		})
	}
}

func TestFormattingMethodsMutateReceiver(t *testing.T) {
	cases := []struct {
		name string
		call func(*wotoMarkDown) WMarkDown
		want string
	}{
		{name: "normal", call: func(m *wotoMarkDown) WMarkDown { return m.Normal("text") }, want: "basetext"},
		{name: "bold", call: func(m *wotoMarkDown) WMarkDown { return m.Bold("text") }, want: "base*text*"},
		{name: "italic", call: func(m *wotoMarkDown) WMarkDown { return m.Italic("text") }, want: "base_text_"},
		{name: "mono", call: func(m *wotoMarkDown) WMarkDown { return m.Mono("text") }, want: "base`text`"},
		{name: "strike", call: func(m *wotoMarkDown) WMarkDown { return m.Strike("text") }, want: "base~text~"},
		{name: "underline", call: func(m *wotoMarkDown) WMarkDown { return m.Underline("text") }, want: "base__text__"},
		{name: "hyperlink", call: func(m *wotoMarkDown) WMarkDown { return m.HyperLink("text", "https://example.com") }, want: "base[text](https://example\\.com)"},
		{name: "link", call: func(m *wotoMarkDown) WMarkDown { return m.Link("text", "https://example.com") }, want: "base[text](https://example\\.com)"},
		{name: "mention", call: func(m *wotoMarkDown) WMarkDown { return m.Mention("user", 42) }, want: "base[user](tg://user?id=42)"},
		{name: "user-mention", call: func(m *wotoMarkDown) WMarkDown { return m.UserMention("user", 42) }, want: "base[user](tg://user?id=42)"},
		{name: "spoiler", call: func(m *wotoMarkDown) WMarkDown { return m.Spoiler("text") }, want: "base||text||"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			base := GetNormal("base").(*wotoMarkDown)
			got := tc.call(base)

			if got != base {
				t.Fatal("expected formatting method to return the receiver")
			}

			if base.ToString() != tc.want {
				t.Fatalf("%s = %q, want %q", tc.name, base.ToString(), tc.want)
			}
		})
	}
}

func TestFormattingMethodsKeepReceiverForEmptyOrInvalidInput(t *testing.T) {
	cases := []struct {
		name string
		call func(*wotoMarkDown) WMarkDown
	}{
		{name: "normal-empty", call: func(m *wotoMarkDown) WMarkDown { return m.Normal("") }},
		{name: "bold-empty", call: func(m *wotoMarkDown) WMarkDown { return m.Bold("") }},
		{name: "italic-empty", call: func(m *wotoMarkDown) WMarkDown { return m.Italic("") }},
		{name: "mono-empty", call: func(m *wotoMarkDown) WMarkDown { return m.Mono("") }},
		{name: "underline-empty", call: func(m *wotoMarkDown) WMarkDown { return m.Underline("") }},
		{name: "strike-empty", call: func(m *wotoMarkDown) WMarkDown { return m.Strike("") }},
		{name: "spoiler-empty", call: func(m *wotoMarkDown) WMarkDown { return m.Spoiler("") }},
		{name: "hyperlink-empty-text", call: func(m *wotoMarkDown) WMarkDown { return m.HyperLink("", "https://example.com") }},
		{name: "hyperlink-empty-url", call: func(m *wotoMarkDown) WMarkDown { return m.HyperLink("text", "") }},
		{name: "mention-empty-text", call: func(m *wotoMarkDown) WMarkDown { return m.Mention("", 42) }},
		{name: "mention-base-index", call: func(m *wotoMarkDown) WMarkDown { return m.Mention("user", baseIndex) }},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			base := GetNormal("base").(*wotoMarkDown)
			got := tc.call(base)

			if got != base {
				t.Fatal("expected the original receiver to be returned")
			}

			if base.ToString() != "base" {
				t.Fatalf("base changed to %q, want %q", base.ToString(), "base")
			}
		})
	}
}

func TestWhitespaceHelpers(t *testing.T) {
	base := GetNormal("base").(*wotoMarkDown)

	if got := base.El(); got.ToString() != "base\n" || got == base {
		t.Fatalf("El() = %q, want new markdown with newline", got.ToString())
	}

	if base.ToString() != "base" {
		t.Fatalf("base changed after El to %q", base.ToString())
	}

	if got := base.Space(); got.ToString() != "base " || got == base {
		t.Fatalf("Space() = %q, want new markdown with space", got.ToString())
	}

	if got := base.Tab(); got.ToString() != "base\t" || got == base {
		t.Fatalf("Tab() = %q, want new markdown with tab", got.ToString())
	}

	if got := base.ElThis(); got != base || base.ToString() != "base\n" {
		t.Fatalf("ElThis() receiver = %q, want %q", base.ToString(), "base\n")
	}

	if got := base.SpaceThis(); got != base || base.ToString() != "base\n " {
		t.Fatalf("SpaceThis() receiver = %q, want %q", base.ToString(), "base\n ")
	}

	if got := base.TabThis(); got != base || base.ToString() != "base\n \t" {
		t.Fatalf("TabThis() receiver = %q, want %q", base.ToString(), "base\n \t")
	}
}

func TestReplaceMarkdownVariants(t *testing.T) {
	base := GetNormal("alpha beta alpha").(*wotoMarkDown)
	md1 := GetNormal("alpha")
	md2 := GetBold("omega")

	if got := base.ReplaceMd(md1, md2); got.ToString() != "*omega* beta *omega*" || got == base {
		t.Fatalf("ReplaceMd() = %q, want new markdown %q", got.ToString(), "*omega* beta *omega*")
	}

	if base.ToString() != "alpha beta alpha" {
		t.Fatalf("base changed after ReplaceMd to %q", base.ToString())
	}

	if got := base.ReplaceMdN(md1, md2, 1); got.ToString() != "*omega* beta alpha" || got == base {
		t.Fatalf("ReplaceMdN() = %q, want new markdown %q", got.ToString(), "*omega* beta alpha")
	}

	if got := base.ReplaceMdThis(md1, md2); got != base || base.ToString() != "*omega* beta *omega*" {
		t.Fatalf("ReplaceMdThis() receiver = %q, want %q", base.ToString(), "*omega* beta *omega*")
	}

	base = GetNormal("alpha beta alpha").(*wotoMarkDown)
	if got := base.ReplaceMdThisN(md1, md2, 1); got != base || base.ToString() != "*omega* beta alpha" {
		t.Fatalf("ReplaceMdThisN() receiver = %q, want %q", base.ToString(), "*omega* beta alpha")
	}
}

func TestReplaceStringVariants(t *testing.T) {
	base := GetNormal("a*b a*b").(*wotoMarkDown)

	if got := base.ReplaceToNew("a*b", "x+y"); got.ToString() != "x\\+y x\\+y" || got == base {
		t.Fatalf("ReplaceToNew() = %q, want new markdown %q", got.ToString(), "x\\+y x\\+y")
	}

	if base.ToString() != "a\\*b a\\*b" {
		t.Fatalf("base changed after ReplaceToNew to %q", base.ToString())
	}

	if got := base.ReplaceToNewN("a*b", "x+y", 1); got.ToString() != "x\\+y a\\*b" || got == base {
		t.Fatalf("ReplaceToNewN() = %q, want new markdown %q", got.ToString(), "x\\+y a\\*b")
	}

	if got := base.Replace("a*b", "x+y"); got != base || base.ToString() != "x\\+y x\\+y" {
		t.Fatalf("Replace() receiver = %q, want %q", base.ToString(), "x\\+y x\\+y")
	}

	base = GetNormal("a*b a*b").(*wotoMarkDown)
	if got := base.ReplaceN("a*b", "x+y", 1); got != base || base.ToString() != "x\\+y a\\*b" {
		t.Fatalf("ReplaceN() receiver = %q, want %q", base.ToString(), "x\\+y a\\*b")
	}
}
