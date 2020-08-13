package test

import (
	"bencode"
	"fmt"
	"testing"
)

func Test_encode(t *testing.T) {
	// string test
	/*
		bc := bencode.NewBenCode()
		str, err := bc.Encode("123")
		fmt.Println(str)
		fmt.Println(err)
	*/

	// int test
	/*
		bc := bencode.NewBenCode()
		str, err := bc.Encode(123)
		fmt.Println(str)
		fmt.Println(err)
	*/

	// slice test
	/*
		bc := bencode.NewBenCode()
		str, err := bc.Encode([]int{123, 456, 789})
		fmt.Println(str)
		fmt.Println(err)
	*/

	// map test
	/*
		bc := bencode.NewBenCode()
		str, err := bc.Encode(map[string]int{
			"onetwothree": 123,
			"fourfivesix": 456,
		})
		fmt.Println(str)
		fmt.Println(err)
	*/

	// struct test
	/*
		bc := bencode.NewBenCode()
		str, err := bc.Encode(struct {
			a string `bencode:"aa"`
			b int    `bencode:"bb"`
		}{
			a: "123",
			b: 123,
		})
		fmt.Println(str)
		fmt.Println(err)
	*/

	//pointer test
	/*
		bc := bencode.NewBenCode()
		str, err := bc.Encode(&struct {
			a string `bencode:"aa"`
			b int    `bencode:"bb"`
		}{
			a: "123",
			b: 123,
		})
		fmt.Println(str)
		fmt.Println(err)
	*/

	// change tag
	/*
		bc := bencode.NewBenCode("bc")
		str, err := bc.Encode(&struct {
			a string `bc:"aa"`
			b int    `bc:"bb"`
		}{
			a: "123",
			b: 123,
		})
		fmt.Println(str)
		fmt.Println(err)
	*/
}

func Test_decode(t *testing.T) {
	// int test
	/*
		bc := bencode.NewBenCode()
		var i int
		err := bc.Decode("i123e", &i)
		fmt.Println(i)
		fmt.Println(err)
	*/

	// string test
	/*
		bc := bencode.NewBenCode()
		var str string
		err := bc.Decode("3:abc", &str)
		fmt.Println(str)
		fmt.Println(err)
	*/

	// array test
	/*
		bc := bencode.NewBenCode()
		var arr [3]int
		err := bc.Decode("li123ei123ei123ee", &arr)
		fmt.Println(arr)
		fmt.Println(err)
	*/

	// slice test
	/*
		bc := bencode.NewBenCode()
		var arr []int
		err := bc.Decode("li123ei123ei123ei123ei123ei123ee", &arr)
		fmt.Println(arr)
		fmt.Println(err)
	*/

	// map test
	/*
		bc := bencode.NewBenCode()
		var arr map[int]int
		err := bc.Decode("di123ei124ei125ei126ei127ei128ee", &arr)
		fmt.Println(arr)
		fmt.Println(err)
	*/

	// struct test with tag
	/*
		bc := bencode.NewBenCode()
		s := struct {
			Abc int
			Abd int
		}{}
		err := bc.Decode("d3:Abci124e3:Abdi125ee", &s)
		fmt.Println(s)
		fmt.Println(err)
	*/

	// struct test with tag
	/*
		bc := bencode.NewBenCode("bc")
		s := struct {
			Abc int `bc:"Abd"`
			//Abd int
		}{}
		err := bc.Decode("d3:Abci124e3:Abdi125ee", &s)
		fmt.Println(s)
		fmt.Println(err)
	*/

	type Email struct {
		Remark  string
		Address string
	}
	type Result struct {
		Name   string
		Phone  string
		Emails []Email
	}
	r := Result{
		"Qiao Yongchen",
		"159XXXXXXXX",
		[]Email{
			Email{"home", "qiaoyongchen@hotmail.com"},
			Email{"work", "qiaoyongchen@hotmail.com"},
		},
	}
	bc := bencode.NewBenCode()
	rst, rsterr := bc.Encode(r)

	fmt.Println(rst)
	fmt.Println(rsterr)

	becodestr := "d5:Phone11:159XXXXXXXX6:Emailsld6:Remark4:home7:Address24:qiaoyongchen@hotmail.comed6:Remark4:work7:Address24:qiaoyongchen@hotmail.comee4:Name13:Qiao Yongchene"
	rr := &Result{}
	bc.Decode(becodestr, rr)
	fmt.Println(rr)
}
