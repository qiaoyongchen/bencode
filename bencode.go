package bencode

import (
	"errors"
	"reflect"
	"strings"

	"github.com/qiaoyongchen/bencode/token"
)

var tag = "bencode"

// BenCode 可以自定义需要解析的tag
type BenCode struct {
	tag string
}

// NewBenCode 创建时可以指定tag
func NewBenCode(tags ...string) *BenCode {
	t := ""
	if len(tags) == 0 || tags[0] == "" {
		t = tag
	} else {
		t = tags[0]
	}
	return &BenCode{tag: t}
}

// Encode Encode
func (p *BenCode) Encode(i interface{}) (string, error) {
	t, e := p.struct2token(reflect.ValueOf(i))
	if e != nil {
		return "", e
	}
	return t.Encode(), nil
}

// Decode Decode
func (p *BenCode) Decode(str string, i interface{}) error {
	t, e := p.parseToken(str)
	if e != nil {
		return e
	}

	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Ptr {
		return errors.New("must pass ptr as second parameter")
	}

	e = p.token2struct(t, v.Elem())
	if e != nil {
		return e
	}
	return nil
}

func (p *BenCode) struct2token(v reflect.Value) (token.Token, error) {
	switch v.Kind() {
	case reflect.String:
		return token.NewString(v.String()), nil
	case reflect.Int:
		return token.NewInt(int(v.Int())), nil
	case reflect.Int64:
		return token.NewInt(int(v.Int())), nil
	case reflect.Array:

		// 遍历元素分别解析
		tkns := []token.Token{}
		for i := 0; i < v.Len(); i++ {
			vv := v.Index(i)
			tkn, err := p.struct2token(vv)
			if err != nil {
				return nil, err
			}
			tkns = append(tkns, tkn)
		}
		return token.NewList(tkns...), nil

	case reflect.Slice:

		// 遍历元素分别解析
		tkns := []token.Token{}
		for i := 0; i < v.Len(); i++ {
			vv := v.Index(i)
			tkn, err := p.struct2token(vv)
			if err != nil {
				return nil, err
			}
			tkns = append(tkns, tkn)
		}
		return token.NewList(tkns...), nil

	case reflect.Struct:

		// 遍历每个 filed
		tknsmap := map[token.Token]token.Token{}
		for i := 0; i < v.NumField(); i++ {
			vv := v.Field(i)
			tkn, err := p.struct2token(vv)
			if err != nil {
				return nil, err
			}

			// 获取到tag按tag值存key, 否则按fieldName值存key
			tag := v.Type().Field(i).Tag.Get(p.tag)
			if tag == "" {
				tknsmap[token.NewString(v.Type().Field(i).Name)] = tkn
			} else {
				tknsmap[token.NewString(tag)] = tkn
			}
		}
		return token.NewDict(tknsmap), nil

	case reflect.Map:

		// 遍历每个 key-value
		tknsmap := map[token.Token]token.Token{}
		mr := v.MapRange()
		for mr.Next() {
			keytoken, keyerror := p.struct2token(mr.Key())
			if keyerror != nil {
				return nil, keyerror
			}
			valtoken, valerror := p.struct2token(mr.Value())
			if valerror != nil {
				return nil, valerror
			}
			tknsmap[keytoken] = valtoken
		}
		return token.NewDict(tknsmap), nil

	case reflect.Ptr:

		// 如果是指针,找到指针的实际值再编码
		vv := v.Elem()
		tkn, err := p.struct2token(vv)
		if err != nil {
			return nil, err
		}
		return tkn, nil

	default:
		return nil, errors.New("can't parse struct to token")
	}
}

func (p *BenCode) token2struct(tkn token.Token, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Int:

		// int
		if tkn.GetType() != token.TypeInt {
			return errors.New("kind not match: int")
		}
		v.SetInt(int64(tkn.GetValue().(int)))

	case reflect.String:

		// string
		if tkn.GetType() != token.TypeString {
			return errors.New("kind not match: string")
		}
		v.SetString(tkn.GetValue().(string))

	case reflect.Ptr:

		// pointer
		// 如果是指针,获取指针实际指向的值再解码
		v = v.Elem()
		return p.token2struct(tkn, v)

	case reflect.Array:

		// array
		// array是定长,直接遍历每个条目并解析
		// 超出array长度的剩余token, 直接舍弃
		if tkn.GetType() != token.TypeList {
			return errors.New("kind not match: list(array)")
		}

		tkns := tkn.GetValue().([]token.Token)
		i := 0
		for _, ttkn := range tkns {
			if i >= v.Len() {
				break
			}

			err := p.token2struct(ttkn, v.Index(i))
			if err != nil {
				return err
			}
			i++
		}

	case reflect.Slice:

		// slice
		// slice 是变长, 用反射根据元素的类型新建空slice, 并在遍历tokens的时候解析出来追加进slice
		if tkn.GetType() != token.TypeList {
			return errors.New("kind not match: list(array)")
		}

		tkns := tkn.GetValue().([]token.Token)
		itemvs := v

		for _, ttkn := range tkns {
			itemv := reflect.New(v.Type().Elem())
			err := p.token2struct(ttkn, itemv)
			if err != nil {
				return err
			}
			itemvs = reflect.Append(itemvs, itemv.Elem())
		}
		v.Set(itemvs)

	case reflect.Map:

		// map
		// 同slice
		if tkn.GetType() != token.TypeDict {
			return errors.New("kind not match: dict(map)")
		}

		tknsmap := tkn.GetValue().(map[token.Token]token.Token)
		itemsmap := reflect.MakeMap(reflect.MapOf(v.Type().Key(), v.Type().Elem()))
		for ktkn, vtkn := range tknsmap {
			kitem := reflect.New(v.Type().Key())
			kerror := p.token2struct(ktkn, kitem)
			if kerror != nil {
				return kerror
			}

			vitem := reflect.New(v.Type().Elem())
			verror := p.token2struct(vtkn, vitem)
			if verror != nil {
				return verror
			}

			itemsmap.SetMapIndex(kitem.Elem(), vitem.Elem())
		}
		v.Set(itemsmap)

	case reflect.Struct:

		// struct
		// 根据 token map 找到对应 field, 如果中途出错直接跳过, 继续执行
		if tkn.GetType() != token.TypeDict {
			return errors.New("kind not match: dict(struct)")
		}
		tknsmap := tkn.GetValue().(map[token.Token]token.Token)

		for ktkn, vtkn := range tknsmap {
			// 对应field name 的token必须是string类型
			if ktkn.GetType() != token.TypeString {
				continue
			}

			fieldname := ktkn.GetValue().(string)
			for i := 0; i < v.NumField(); i++ {
				// 根据 tag 查找filed
				tmpfieldtag := v.Type().Field(i).Tag.Get(p.tag)
				if tmpfieldtag == fieldname {
					tmpfield := v.Field(i)
					p.token2struct(vtkn, tmpfield)
					break
				}

				// 根据field name查找 field
				tmpfieldname := v.Type().Field(i).Name
				if tmpfieldname == fieldname || strings.ToLower(tmpfieldname) == strings.ToLower(fieldname) {
					tmpfield := v.Field(i)
					p.token2struct(vtkn, tmpfield)
					break
				}

			}
		}

	default:
		return errors.New("other kind")
	}
	return nil
}

func (p *BenCode) parseToken(str string) (token.Token, error) {
	var tkn token.Token
	var err error
	if str[0] == 'l' {
		tkn = token.NewList(nil)
		err = tkn.Decode(&str)
	} else if str[0] == 'i' {
		tkn = token.NewInt(0)
		err = tkn.Decode(&str)
	} else if str[0] == 'd' {
		tkn = token.NewDict(nil)
		err = tkn.Decode(&str)
	} else if token.IsDigit(str[0:1]) {
		tkn = token.NewString("")
		err = tkn.Decode(&str)
	} else {
		return nil, errors.New("can't parse to token")
	}
	if err != nil {
		return nil, err
	}
	return tkn, err
}
