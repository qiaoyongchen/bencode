# Bencode 介绍

BEncoding是一种编码方式，比如种子文件就是就是采用这种编码方式。

Bencode有4种类型数据:

## 1. String

"12345" => 5:12345

## 2. Int

12345 => i12345e

## 3. List

List<"abced", 12345> => l5:abcdei12345ee

## 4. dictionary

Dictionary<{"abced":"abced"},{"abc":123}> => d5:abced5:abced3:abci23ee

# 使用方式

## 安装

```
github.com/qiaoyongchen/bencode
```

## 编码

### string

```
package main

import (
	"fmt"

	"github.com/qiaoyongchen/bencode"
)

func main() {
	bc := bencode.NewBenCode()
	str, _ := bc.Encode("abcde")
	fmt.Println(str)
}

// 5:abcde
```

### int

```
package main

import (
	"fmt"

	"github.com/qiaoyongchen/bencode"
)

func main() {
	bc := bencode.NewBenCode()
	str, _ := bc.Encode(12345)
	fmt.Println(str)
}

// i12345e
```

### List

```
package main

import (
	"fmt"

	"github.com/qiaoyongchen/bencode"
)

func main() {
	bc := bencode.NewBenCode()
	str, _ := bc.Encode([]string{"123", "abc"})
	fmt.Println(str)
}

// d3:4563:4563:1233:123e
```

### map

```
package main

import (
	"fmt"

	"github.com/qiaoyongchen/bencode"
)

func main() {
	bc := bencode.NewBenCode()
	str, _ := bc.Encode(map[string]string{
		"123": "123",
		"456": "456",
	})
	fmt.Println(str)
}

// d3:4563:4563:1233:123e
```

### struct

```
package main

import (
	"fmt"

	"github.com/qiaoyongchen/bencode"
)

func main() {
	bc := bencode.NewBenCode()
	str, _ := bc.Encode(struct {
		Name string
		Age  int
	}{
		Name: "qiaoyongchen",
		Age:  12,
	})
	fmt.Println(str)
}

// d4:Name12:qiaoyongchen3:Agei12ee
```

### struct with tag

```
package main

import (
	"fmt"

	"github.com/qiaoyongchen/bencode"
)

func main() {
	bc := bencode.NewBenCode("bc")
	str, _ := bc.Encode(struct {
		Name string `bc:"another_name"`
		Age  int    `bc:"another_age"`
	}{
		Name: "qiaoyongchen",
		Age:  12,
	})
	fmt.Println(str)
}

// d12:another_name12:qiaoyongchen11:another_agei12ee
```

## 解码

```
package main

import (
	"fmt"

	"github.com/qiaoyongchen/bencode"
)

func main() {
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
	rst, _ := bc.Encode(r)

	fmt.Println(rst)

	becodestr := "d5:Phone11:159XXXXXXXX6:Emailsld6:Remark4:home7:Address24:qiaoyongchen@hotmail.comed6:Remark4:work7:Address24:qiaoyongchen@hotmail.comee4:Name13:Qiao Yongchene"
	rr := &Result{}
	bc.Decode(becodestr, rr)
	fmt.Println(rr)
}

```