package token

import (
	"errors"
)

// Dict Dict
type Dict struct {
	m map[Token]Token
}

// NewDict NewDict
func NewDict(kvmap map[Token]Token) *Dict {
	if kvmap == nil {
		kvmap = map[Token]Token{}
	}
	return &Dict{m: kvmap}
}

// GetType GetType
func (p *Dict) GetType() Type {
	return TypeDict
}

// Encode Encode
func (p *Dict) Encode() string {
	str := "d"
	for k, v := range p.m {
		str += k.Encode() + v.Encode()
	}
	return str + "e"
}

// Decode Decode
func (p *Dict) Decode(strpointer *string) error {
	if (*strpointer)[0] != 'd' {
		return errors.New("can't decode to dictionary, beginning of '" + (*strpointer) + "'")
	}
	*strpointer = (*strpointer)[1:]
	for {
		var firstToken, secondToken Token
		var firstError, secondError error
		if (*strpointer)[0] == 'i' {
			firstToken = NewInt(0)
			firstError = firstToken.Decode(strpointer)
		} else if (*strpointer)[0] == 'l' {
			firstToken = NewList(nil)
			firstError = firstToken.Decode(strpointer)
		} else if IsDigit((*strpointer)[0:1]) {
			firstToken = NewString("")
			firstError = firstToken.Decode(strpointer)
		} else if (*strpointer)[0] == 'd' {
			firstToken = NewDict(nil)
			firstError = firstToken.Decode(strpointer)
		} else if (*strpointer)[0] == 'e' {
			*strpointer = (*strpointer)[1:]
			return nil
		}
		if firstError != nil {
			return firstError
		}
		if (*strpointer)[0] == 'i' {
			secondToken = NewInt(0)
			secondError = secondToken.Decode(strpointer)
		} else if (*strpointer)[0] == 'l' {
			secondToken = NewList(nil)
			secondError = secondToken.Decode(strpointer)
		} else if IsDigit((*strpointer)[0:1]) {
			secondToken = NewString("")
			secondError = secondToken.Decode(strpointer)
		} else if (*strpointer)[0] == 'd' {
			secondToken = NewDict(nil)
			secondError = secondToken.Decode(strpointer)
		} else {
			return errors.New("can't decode to dictionary, near to '" + (*strpointer) + "'")
		}

		if secondError != nil {
			return secondError
		}
		p.m[firstToken] = secondToken
	}
}

// GetValue GetValue
func (p *Dict) GetValue() interface{} {
	return p.m
}
