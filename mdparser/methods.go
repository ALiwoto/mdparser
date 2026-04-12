// mdparser library Project
// Copyright (C) 2021-2022 ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package mdparser

import (
	"strings"
)

func (m *wotoMarkDown) Append(v WMarkDown) WMarkDown {
	if v == nil {
		return nil
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

func (m *wotoMarkDown) Normal(text string) WMarkDown {
	return m.appendTextSegment(segmentNormal, text)
}

func (m *wotoMarkDown) Bold(text string) WMarkDown {
	return m.appendTextSegment(segmentBold, text)
}

func (m *wotoMarkDown) Italic(text string) WMarkDown {
	return m.appendTextSegment(segmentItalic, text)
}

func (m *wotoMarkDown) Mono(text string) WMarkDown {
	return m.appendTextSegment(segmentMono, text)
}

func (m *wotoMarkDown) CodeBlock(text string) WMarkDown {
	return m.appendTextSegment(segmentCodeBlock, text)
}

func (m *wotoMarkDown) CodeBlockLang(lang, text string) WMarkDown {
	if text == "" {
		return m
	}

	return m.appendSegment(newCodeBlockLangSegment(lang, text))
}

func (m *wotoMarkDown) Strike(text string) WMarkDown {
	return m.appendTextSegment(segmentStrike, text)
}

func (m *wotoMarkDown) Underline(text string) WMarkDown {
	return m.appendTextSegment(segmentUnderline, text)
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

func (m *wotoMarkDown) Spoiler(text string) WMarkDown {
	return m.appendTextSegment(segmentSpoiler, text)
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
