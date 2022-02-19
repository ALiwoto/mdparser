// mdparser library Project
// Copyright (C) 2021-2022 ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package mdparser

const (
	_TG_USER_ID = `tg://user?id=`
)

const (
	_CHAR_S1  = '\\'
	_CHAR_S2  = '\''
	_CHAR_S3  = '*'
	_CHAR_S4  = '_'
	_CHAR_S5  = '{'
	_CHAR_S6  = '}'
	_CHAR_S7  = '['
	_CHAR_S8  = ']'
	_CHAR_S9  = '('
	_CHAR_S10 = ')'
	_CHAR_S11 = '#'
	_CHAR_S12 = '+'
	_CHAR_S13 = '-'
	_CHAR_S14 = '.'
	_CHAR_S15 = '!'
	_CHAR_S16 = '`'
	_CHAR_S17 = '='
	_CHAR_S18 = '>'
	_CHAR_S19 = '<'
	_CHAR_S20 = '~'
	_CHAR_S21 = '|'
)

var _sChars = map[rune]bool{
	_CHAR_S1:  true,
	_CHAR_S2:  true,
	_CHAR_S3:  true,
	_CHAR_S4:  true,
	_CHAR_S5:  true,
	_CHAR_S6:  true,
	_CHAR_S7:  true,
	_CHAR_S8:  true,
	_CHAR_S9:  true,
	_CHAR_S10: true,
	_CHAR_S11: true,
	_CHAR_S12: true,
	_CHAR_S13: true,
	_CHAR_S14: true,
	_CHAR_S15: true,
	_CHAR_S16: true,
	_CHAR_S17: true,
	_CHAR_S18: true,
	_CHAR_S19: true,
	_CHAR_S20: true,
	_CHAR_S21: true,
}
