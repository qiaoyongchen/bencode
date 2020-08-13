package token

import (
	"errors"
	"strconv"
)

// String 对应 bencode 中的string
type String struct {
	v string
}

// NewString 创建
func NewString(v string) *String {
	return &String{v: v}
}

// GetType 类型
func (p *String) GetType() Type {
	return TypeString
}

// Encode 编码 token -> string
func (p *String) Encode() string {
	return strconv.Itoa(len(p.v)) + ":" + p.v
}

// Decode 解码 string -> token
func (p *String) Decode(strpointer *string) error {
	i := 0
	lenstr := ""
	for IsDigit((*strpointer)[i : i+1]) {
		lenstr += (*strpointer)[i : i+1]
		i = i + 1
	}

	if (*strpointer)[i:i+1] != ":" {
		return errors.New("can't decode to string, near to '^^^" + (*strpointer)[i:] + "'")
	}

	len, _ := strconv.Atoi(lenstr)
	p.v = (*strpointer)[i+1 : i+1+len]

	*strpointer = (*strpointer)[i+1+len:]
	return nil
}

// GetValue GetValue
func (p *String) GetValue() interface{} {
	return p.v
}
