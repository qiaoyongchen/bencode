package token

import (
	"errors"
	"strconv"
)

// Int Int
type Int struct {
	v int
}

// NewInt NewInt
func NewInt(v int) *Int {
	return &Int{v: v}
}

// GetType GetType
func (p *Int) GetType() Type {
	return TypeInt
}

// Encode Encode
func (p *Int) Encode() string {
	return "i" + strconv.Itoa(p.v) + "e"
}

// Decode Decode
func (p *Int) Decode(strpointer *string) error {
	if (*strpointer)[0] != 'i' {
		return errors.New("can't decode to int, begining of the integer '^^^" + (*strpointer) + "'")
	}

	i := 1
	strrst := ""
	for IsDigit((*strpointer)[i : i+1]) {
		strrst += (*strpointer)[i : i+1]
		i = i + 1
	}

	if (*strpointer)[i] != 'e' {
		return errors.New("can't decode to int, ending of the integer '^^^" + (*strpointer) + "'")
	}

	irst, _ := strconv.Atoi(strrst)
	p.v = irst
	*strpointer = (*strpointer)[i+1:]
	return nil
}

// GetValue GetValue
func (p *Int) GetValue() interface{} {
	return p.v
}
