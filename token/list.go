package token

import (
	"errors"
)

// List 对应bencode中list类型
type List struct {
	ts []Token
}

// NewList 创建如果tokens不为空取tokens[0]
func NewList(tokens ...Token) *List {
	if len(tokens) == 0 || tokens[0] == nil {
		tokens = nil
	}
	return &List{ts: tokens}
}

// GetType 获得类型
func (p *List) GetType() Type {
	return TypeList
}

// Encode 编码 token -> string
// 对列表中每个元素分别编码,以“l” + SumSubItem + 'e'返回结果
func (p *List) Encode() string {
	str := "l"
	for _, t := range p.ts {
		str += t.Encode()
	}
	return str + "e"
}

// Decode 解码 string -> token
func (p *List) Decode(strpointer *string) error {
	if (*strpointer)[0] != 'l' {
		return errors.New("can't decode to list, beginning of '" + (*strpointer) + "'")
	}

	*strpointer = (*strpointer)[1:]

	var tkn Token
	var err error
	for {
		if (*strpointer)[0] == 'i' {
			tkn = NewInt(0)
			err = tkn.Decode(strpointer)
		} else if (*strpointer)[0] == 'l' {
			tkn = NewList(nil)
			err = tkn.Decode(strpointer)
		} else if IsDigit((*strpointer)[0:1]) {
			tkn = NewString("")
			err = tkn.Decode(strpointer)
		} else if (*strpointer)[0] == 'd' {
			tkn = NewDict(nil)
			err = tkn.Decode(strpointer)
		} else if (*strpointer)[0] == 'e' {
			*strpointer = (*strpointer)[1:]
			return nil
		} else {
			return errors.New("can't decode to dictionary, near to '" + (*strpointer) + "'")
		}

		p.ts = append(p.ts, tkn)

		if err != nil {
			return err
		}
	}
}

// GetValue GetValue
func (p *List) GetValue() interface{} {
	return p.ts
}
