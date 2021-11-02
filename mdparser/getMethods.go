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
		wmd := &wotoMarkDown{
			_value: str,
		}

		return wmd
	}

	return nil
}

func (m *wotoMarkDown) AppendThis(v WMarkDown) WMarkDown {
	m.setValue(m.getValue() + v.getValue())

	return m
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

func (m *wotoMarkDown) AppendNormalThis(v string) WMarkDown {
	if strongStringGo.IsEmpty(&v) {
		return m
	}

	m.setValue(m.getValue() + toNormal(v))

	return m
}

func (m *wotoMarkDown) AppendBold(v string) WMarkDown {
	if strongStringGo.IsEmpty(&v) {
		return m
	}

	return toWotoMD(m.getValue() + toBold(v))
}

func (m *wotoMarkDown) AppendBoldThis(v string) WMarkDown {
	if strongStringGo.IsEmpty(&v) {
		return m
	}

	m.setValue(m.getValue() + toBold(v))

	return m
}

func (m *wotoMarkDown) AppendItalic(v string) WMarkDown {
	if strongStringGo.IsEmpty(&v) {
		return m
	}

	return toWotoMD(m.getValue() + toItalic(v))
}

func (m *wotoMarkDown) AppendItalicThis(v string) WMarkDown {
	if strongStringGo.IsEmpty(&v) {
		return m
	}

	m.setValue(m.getValue() + toItalic(v))

	return m
}

func (m *wotoMarkDown) AppendMono(v string) WMarkDown {
	if strongStringGo.IsEmpty(&v) {
		return m
	}

	return toWotoMD(m.getValue() + toMono(v))
}

func (m *wotoMarkDown) AppendMonoThis(v string) WMarkDown {
	if strongStringGo.IsEmpty(&v) {
		return m
	}

	m.setValue(m.getValue() + toMono(v))

	return m
}

func (m *wotoMarkDown) AppendHyperLink(text, url string) WMarkDown {
	if strongStringGo.IsEmpty(&text) || strongStringGo.IsEmpty(&url) {
		return m
	}

	return toWotoMD(m.getValue() + toHyperLink(text, url))
}

func (m *wotoMarkDown) AppendHyperLinkThis(text, url string) WMarkDown {
	if strongStringGo.IsEmpty(&text) || strongStringGo.IsEmpty(&url) {
		return m
	}

	m.setValue(m.getValue() + toHyperLink(text, url))

	return m
}

func (m *wotoMarkDown) AppendMention(text string, id int64) WMarkDown {
	if strongStringGo.IsEmpty(&text) || id == strongStringGo.BaseIndex {
		return m
	}

	return toWotoMD(m.getValue() + toUserMention(text, id))
}

func (m *wotoMarkDown) AppendMentionThis(text string, id int64) WMarkDown {
	if strongStringGo.IsEmpty(&text) || id == strongStringGo.BaseIndex {
		return m
	}

	m.setValue(m.getValue() + toUserMention(text, id))

	return m
}

// El method appends a new line (Endline) to the markdown value and returns
// a new instance of WMarkDown.
func (m *wotoMarkDown) El() WMarkDown {
	return m.AppendNormal("\n")
}

// ElThis method appends a new line (Endline) to the current markdown value and
// returns the current instance of WMarkDown.
func (m *wotoMarkDown) ElThis() WMarkDown {
	return m.AppendNormalThis("\n")
}

// Space method appends a space to the markdown value and returns
// a new instance of WMarkDown.
func (m *wotoMarkDown) Space() WMarkDown {
	return m.AppendNormal(" ")
}

// SpaceThis method appends a new line to the current markdown value and
// returns the current instance of WMarkDown.
func (m *wotoMarkDown) SpaceThis() WMarkDown {
	return m.AppendNormalThis(" ")
}

// Tab method appends a tab ("\t") to the markdown value and returns
// a new instance of WMarkDown.
func (m *wotoMarkDown) Tab() WMarkDown {
	return m.AppendNormal("\t")
}

// TabThis method appends a tab ("\t") to the current markdown value and
// returns the current instance of WMarkDown.
func (m *wotoMarkDown) TabThis() WMarkDown {
	return m.AppendNormalThis("\t")
}

func (m *wotoMarkDown) getValue() string {
	return m._value
}
