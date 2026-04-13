// mdparser library Project
// Copyright (C) 2021-2026 ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package mdparser

// WMarkDown is the mutable markdown builder exposed by this package.
// Every instance method mutates the current value unless you explicitly
// call Clone first.
type WMarkDown interface {
	// Append appends another markdown value to the current markdown.
	Append(md WMarkDown) WMarkDown
	// Clone returns a copy of the current markdown value.
	Clone() WMarkDown
	// ReplaceMd replaces all matching markdown fragments with another markdown value.
	ReplaceMd(md1, md2 WMarkDown) WMarkDown
	// ReplaceMdN replaces up to n matching markdown fragments with another markdown value.
	ReplaceMdN(md1, md2 WMarkDown, n int) WMarkDown
	// ToString returns the final markdown string.
	ToString() string

	// ReplaceToNew replaces all matching plain-text fragments with escaped plain-text replacements.
	ReplaceToNew(text1, text2 string) WMarkDown
	// ReplaceToNewN replaces up to n matching plain-text fragments with escaped plain-text replacements.
	ReplaceToNewN(text1, text2 string, n int) WMarkDown

	// Normal appends plain escaped text built from mixed values.
	Normal(values ...any) WMarkDown
	// Bold appends bold markdown text built from mixed values.
	Bold(values ...any) WMarkDown
	// Italic appends italic markdown text built from mixed values.
	Italic(values ...any) WMarkDown
	// Mono appends inline code wrapped in single backticks built from mixed values.
	Mono(values ...any) WMarkDown
	// Styled appends text wrapped in one or more nestable text styles.
	Styled(text string, styles ...TextStyle) WMarkDown
	// CodeBlock appends a fenced code block built from mixed values.
	CodeBlock(values ...any) WMarkDown
	// CodeBlockLang appends a fenced code block with a language hint.
	CodeBlockLang(lang, text string) WMarkDown
	// HyperLink appends a markdown hyperlink.
	HyperLink(text, url string) WMarkDown
	// Link appends a markdown hyperlink.
	Link(text, url string) WMarkDown
	// Mention appends a Telegram user mention.
	Mention(text string, id int64) WMarkDown
	// UserMention appends a Telegram user mention.
	UserMention(text string, id int64) WMarkDown
	// Spoiler appends spoiler markdown built from mixed values.
	Spoiler(values ...any) WMarkDown
	// Strike appends strike-through markdown built from mixed values.
	Strike(values ...any) WMarkDown
	// Underline appends underline markdown built from mixed values.
	Underline(values ...any) WMarkDown

	// El appends a newline.
	El() WMarkDown
	// Space appends a space.
	Space() WMarkDown
	// Tab appends a tab.
	Tab() WMarkDown
	// Replace replaces all matching escaped plain-text fragments in the current value.
	Replace(text1, text2 string) WMarkDown
	// ReplaceN replaces up to n matching escaped plain-text fragments in the current value.
	ReplaceN(text1, text2 string, n int) WMarkDown
}

// TextStyle identifies a nestable Telegram text style.
type TextStyle uint8

type markdownSegmentKind uint8

type markdownSegment struct {
	kind   markdownSegmentKind
	text   string
	meta   string
	styles TextStyle
}

type wotoMarkDown struct {
	_segments []markdownSegment
}

// secretContainer contains information about a secret value which needs
// to be replaced by its name when it's being appended to a markdown value.
type secretContainer struct {
	name  string
	value string
}
