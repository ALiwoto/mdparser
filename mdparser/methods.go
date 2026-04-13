// mdparser library Project
// Copyright (C) 2021-2026 ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package mdparser

import (
	"fmt"
	"reflect"
	"strings"
)

func (m *wotoMarkDown) Append(v WMarkDown) WMarkDown {
	if isNilValue(v) {
		return m
	}

	if md, ok := v.(*wotoMarkDown); ok {
		return m.appendSegments(md._segments...)
	}

	return m.appendSegment(newRawSegment(v.ToString()))
}

func (m *wotoMarkDown) ReplaceMd(md1, md2 WMarkDown) WMarkDown {
	return m.replaceRendered(md1.ToString(), md2.ToString(), -1)
}

func (m *wotoMarkDown) ReplaceMdN(md1, md2 WMarkDown, n int) WMarkDown {
	return m.replaceRendered(md1.ToString(), md2.ToString(), n)
}

func (m *wotoMarkDown) ReplaceToNew(text1, text2 string) WMarkDown {
	return m.replaceRendered(toNormal(text1), toNormal(text2), -1)
}

func (m *wotoMarkDown) ReplaceToNewN(text1, text2 string, n int) WMarkDown {
	return m.replaceRendered(toNormal(text1), toNormal(text2), n)
}

func (m *wotoMarkDown) Clone() WMarkDown {
	return &wotoMarkDown{_segments: cloneSegments(m._segments)}
}

func (m *wotoMarkDown) ToString() string {
	return renderSegments(m._segments)
}

func (m *wotoMarkDown) ToUnformattedString() string {
	return renderUnformattedSegments(m._segments)
}

func (m *wotoMarkDown) Normal(values ...any) WMarkDown {
	return m.appendFormattedText(segmentNormal, values...)
}

func (m *wotoMarkDown) Bold(values ...any) WMarkDown {
	return m.appendFormattedText(segmentBold, values...)
}

func (m *wotoMarkDown) Italic(values ...any) WMarkDown {
	return m.appendFormattedText(segmentItalic, values...)
}

func (m *wotoMarkDown) Mono(values ...any) WMarkDown {
	return m.appendFormattedText(segmentMono, values...)
}

func (m *wotoMarkDown) Styled(text string, styles ...TextStyle) WMarkDown {
	if text == "" {
		return m
	}

	return m.appendSegment(newStyledSegment(text, styles...))
}

func (m *wotoMarkDown) CodeBlock(values ...any) WMarkDown {
	return m.appendFormattedText(segmentCodeBlock, values...)
}

func (m *wotoMarkDown) CodeBlockLang(lang, text string) WMarkDown {
	if text == "" {
		return m
	}

	return m.appendSegment(newCodeBlockLangSegment(lang, text))
}

func (m *wotoMarkDown) Strike(values ...any) WMarkDown {
	return m.appendFormattedText(segmentStrike, values...)
}

func (m *wotoMarkDown) Underline(values ...any) WMarkDown {
	return m.appendFormattedText(segmentUnderline, values...)
}

func (m *wotoMarkDown) HyperLink(text, url string) WMarkDown {
	return m.appendPairSegment(text, url, newHyperLinkSegment)
}

func (m *wotoMarkDown) Link(text, url string) WMarkDown {
	return m.appendPairSegment(text, url, newHyperLinkSegment)
}

// Mention, mentions a user.
func (m *wotoMarkDown) Mention(text string, id int64) WMarkDown {
	return m.appendMentionSegment(text, id)
}

// UserMention, mentions a user.
func (m *wotoMarkDown) UserMention(text string, id int64) WMarkDown {
	return m.appendMentionSegment(text, id)
}

func (m *wotoMarkDown) Spoiler(values ...any) WMarkDown {
	return m.appendFormattedText(segmentSpoiler, values...)
}

// El method appends a new line (End-line) to the markdown value.
func (m *wotoMarkDown) El() WMarkDown {
	return m.Normal("\n")
}

// Space method appends a space to the markdown value.
func (m *wotoMarkDown) Space() WMarkDown {
	return m.Normal(" ")
}

// Tab method appends a tab ("\t") to the markdown value.
func (m *wotoMarkDown) Tab() WMarkDown {
	return m.Normal("\t")
}

func (m *wotoMarkDown) Replace(text1, text2 string) WMarkDown {
	return m.replaceRendered(toNormal(text1), toNormal(text2), -1)
}

func (m *wotoMarkDown) ReplaceN(text1, text2 string, n int) WMarkDown {
	return m.replaceRendered(toNormal(text1), toNormal(text2), n)
}

func (m *wotoMarkDown) appendSegment(segment markdownSegment) WMarkDown {
	m._segments = append(m._segments, segment)
	return m
}

func (m *wotoMarkDown) appendSegments(segments ...markdownSegment) WMarkDown {
	if len(segments) == 0 {
		return m
	}

	m._segments = append(m._segments, cloneSegments(segments)...)
	return m
}

func (m *wotoMarkDown) appendTextSegment(kind markdownSegmentKind, text string) WMarkDown {
	if text == "" {
		return m
	}

	return m.appendSegment(newTextSegment(kind, text))
}

func (m *wotoMarkDown) appendFormattedText(kind markdownSegmentKind, values ...any) WMarkDown {
	if len(values) == 0 {
		return m
	}

	flush := func(chunk []any) WMarkDown {
		if len(chunk) == 0 {
			return m
		}

		return m.appendTextSegment(kind, formatMarkdownValues(chunk))
	}

	var chunk []any
	for _, current := range values {
		if md, ok := asMarkdownValue(current); ok {
			flush(chunk)
			chunk = chunk[:0]

			if md == nil {
				continue
			}

			m.Append(md)
			continue
		}

		chunk = append(chunk, current)
	}

	return flush(chunk)
}

func (m *wotoMarkDown) appendPairSegment(text, extra string, factory func(string, string) markdownSegment) WMarkDown {
	if text == "" || extra == "" {
		return m
	}

	return m.appendSegment(factory(text, extra))
}

func (m *wotoMarkDown) appendMentionSegment(text string, id int64) WMarkDown {
	if text == "" || id == baseIndex {
		return m
	}

	return m.appendSegment(newMentionSegment(text, id))
}

func (m *wotoMarkDown) replaceRendered(oldValue, newValue string, n int) WMarkDown {
	rendered := strings.Replace(m.ToString(), oldValue, newValue, n)
	if rendered == "" {
		m._segments = nil
		return m
	}

	m._segments = []markdownSegment{newRawSegment(rendered)}
	return m
}

func formatMarkdownValues(values []any) string {
	if len(values) == 0 {
		return ""
	}

	format, ok := values[0].(string)
	if ok && strings.Contains(format, "%") {
		return fmt.Sprintf(format, values[1:]...)
	}

	return fmt.Sprint(values...)
}

func asMarkdownValue(value any) (WMarkDown, bool) {
	md, ok := value.(WMarkDown)
	if !ok {
		return nil, false
	}

	if isNilValue(md) {
		return nil, true
	}

	return md, true
}

func isNilValue(value any) bool {
	if value == nil {
		return true
	}

	current := reflect.ValueOf(value)
	switch current.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return current.IsNil()
	default:
		return false
	}
}
