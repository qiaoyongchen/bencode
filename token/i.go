package token

// Type TokenType
type Type = uint8

const (
	// TypeString TypeString
	TypeString Type = 1

	// TypeInt TypeInt
	TypeInt Type = 2

	// TypeList TypeList
	TypeList Type = 3

	// TypeDict TypeDict
	TypeDict Type = 4
)

// Token Token
type Token interface {
	GetType() Type
	Encode() string
	Decode(*string) error
	GetValue() interface{}
}
