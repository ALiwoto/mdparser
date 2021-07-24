// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package mdparser

import (
	"reflect"

	"github.com/ALiwoto/StrongStringGo/strongStringGo"
)

func (m *wotoMarkDown) Append(v WMarkDown) WMarkDown {
	if reflect.TypeOf(v) == reflect.TypeOf(m) {
		md := v.(*wotoMarkDown)
		str := m._value + md._value
		wmd := wotoMarkDown{
			_value: str,
		}

		return &wmd
	}

	return nil
}

func (m *wotoMarkDown) AppendThis(v WMarkDown) {
	m.setValue(m.getValue() + v.getValue())
}

func (m *wotoMarkDown) ToString() string {
	return m._value
}

func (m *wotoMarkDown) AppendNormal(v string) WMarkDown {
	if strongStringGo.IsEmpty(&v) {
		return m
	}

	return toWotoMD(m.getValue() + toNormal(v))
}

func (m *wotoMarkDown) AppendNormalThis(v string) {
	if strongStringGo.IsEmpty(&v) {
		return
	}

	m.setValue(m.getValue() + toNormal(v))
}

func (m *wotoMarkDown) AppendBold(v string) WMarkDown {
	if strongStringGo.IsEmpty(&v) {
		return m
	}

	return toWotoMD(m.getValue() + toBold(v))
}

func (m *wotoMarkDown) AppendBoldThis(v string) {
	if strongStringGo.IsEmpty(&v) {
		return
	}

	m.setValue(m.getValue() + toBold(v))
}

func (m *wotoMarkDown) AppendItalic(v string) WMarkDown {
	if strongStringGo.IsEmpty(&v) {
		return m
	}

	return toWotoMD(m.getValue() + toItalic(v))
}

func (m *wotoMarkDown) AppendItalicThis(v string) {
	if strongStringGo.IsEmpty(&v) {
		return
	}

	m.setValue(m.getValue() + toItalic(v))
}

func (m *wotoMarkDown) AppendMono(v string) WMarkDown {
	if strongStringGo.IsEmpty(&v) {
		return m
	}

	return toWotoMD(m.getValue() + toMono(v))
}

func (m *wotoMarkDown) AppendMonoThis(v string) {
	if strongStringGo.IsEmpty(&v) {
		return
	}

	m.setValue(m.getValue() + toMono(v))
}

func (m *wotoMarkDown) AppendHyperLink(text, url string) WMarkDown {
	if strongStringGo.IsEmpty(&text) || strongStringGo.IsEmpty(&url) {
		return m
	}

	return toWotoMD(m.getValue() + toHyperLink(text, url))
}

func (m *wotoMarkDown) AppendHyperLinkThis(text, url string) {
	if strongStringGo.IsEmpty(&text) || strongStringGo.IsEmpty(&url) {
		return
	}

	m.setValue(m.getValue() + toHyperLink(text, url))
}

func (m *wotoMarkDown) AppendMention(text string, id int64) WMarkDown {
	if strongStringGo.IsEmpty(&text) || id == strongStringGo.BaseIndex {
		return m
	}

	return toWotoMD(m.getValue() + toUserMention(text, id))
}

func (m *wotoMarkDown) AppendMentionThis(text string, id int64) {
	if strongStringGo.IsEmpty(&text) || id == strongStringGo.BaseIndex {
		return
	}

	m.setValue(m.getValue() + toUserMention(text, id))
}

func (m *wotoMarkDown) getValue() string {
	return m._value
}
