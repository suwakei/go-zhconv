# go-zhconv

[![Build Status](https://github.com/suwakei/go-zhconv/actions/workflows/build.yml/badge.svg)](https://github.com/suwakei/go-zhconv/actions/workflows/build.yml)
[![Test Status](https://github.com/suwakei/go-zhconv/actions/workflows/test.yml/badge.svg)](https://github.com/suwakei/go-zhconv/actions/workflows/test.yml)
[![Lint Status](https://github.com/suwakei/go-zhconv/actions/workflows/lint.yml/badge.svg)](https://github.com/suwakei/go-zhconv/actions/workflows/lint.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/suwakei/go-zhconv)](https://goreportcard.com/report/github.com/suwakei/go-zhconv)
[![codecov](https://codecov.io/gh/suwakei/go-zhconv/graph/badge.svg?token=3XKGD5O102)](https://codecov.io/gh/suwakei/go-zhconv)

## Overview
This repository is a library that supports character conversion in the Go language. It performs mutual conversion between full-width and half-width characters and kana.

### zhconv written in other langage
- TypeScript: https://github.com/suwakei/deno-zhconv


## Installation

```sh
go get -u github.com/suwakei/go-zhconv/zhconv
```

## Features
```go
// H2z converts half-width characters (hankaku) in a string to full-width characters (zenkaku).
// It handles ASCII, Katakana, digits, and Katakana with dakuten/handakuten.
func H2z(string) string


// H2zAt returns string that converted from half-width to full-width.
// Conversion string can be selected with the second argument.
func H2zAt(string, ...int) string


// Z2h converts full-width characters (zenkaku) in a string to half-width characters (hankaku).
// It handles ASCII, Katakana, digits, and Katakana with dakuten/handakuten.
func Z2h(string) string


// Z2hAt returns string that converted from full-width to half-width.
// Conversion string can be selected with the second argument.
func Z2hAt(string, ...int) string
```
## Usage
### convert from FullWidth to HalfWidth
```go
package main

import (
    "fmt"
    "github.com/suwakei/go-zhconv/zhconv"
)

func main() {
    result := zhconv.Z2h("Ｈｅｌｌｏ， ｗｏｒｌｄ！")
    fmt.Println(result) // "Hello, world!"

    result = zhconv.Z2h("ＡＢＣｄｅｆ　ＸＹＺ！＃＄％＆’（）＊＋，－．／：；＜＝＞？＠［￥］＾＿‘｛｜｝～")
    fmt.Println(result) // ABCdef XYZ!#$%&'()*+,-./:;<=>?@[¥]^_`{|}~

    result = zhconv.Z2h("０１２３４５６７８９")
    fmt.Println(result) // 0123456789

    result = zhconv.Z2h("アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲン")
    fmt.Println(result) // ｱｲｳｴｵｶｷｸｹｺｻｼｽｾｿﾀﾁﾂﾃﾄﾅﾆﾇﾈﾉﾊﾋﾌﾍﾎﾏﾐﾑﾒﾓﾔﾕﾖﾗﾘﾙﾚﾛﾜｦﾝ

	result = zhconv.Z2h("ァィゥェォッャュョ")
    fmt.Println(result) // ｧｨｩｪｫｯｬｭｮ ヮ is not converted because there is no corresponding character for half-width.

    result = zhconv.Z2h("。、・ー「」")
    fmt.Println(result) // ｡､･ｰ｢｣

    result = zhconv.Z2h("ガギグゲゴザジズゼゾダヂヅデドバビブベボヴ")
    fmt.Println(result) // ｶﾞｷﾞｸﾞｹﾞｺﾞｻﾞｼﾞｽﾞｾﾞｿﾞﾀﾞﾁﾞﾂﾞﾃﾞﾄﾞﾊﾞﾋﾞﾌﾞﾍﾞﾎﾞｳﾞ

    result = zhconv.Z2h("パピプペポ")
    fmt.Println(result) // ﾊﾟﾋﾟﾌﾟﾍﾟﾎﾟ

    result = zhconv.Z2h("　スペース　") // convert from Full width space to half width space
    fmt.Println(result) //  ｽﾍﾟｰｽ 

    result = zhconv.Z2h("①②③㈱㈲") // It is assumed that environment dependent characters will not be converted.
    fmt.Println(result) // ①②③㈱㈲
}
```


### convert from HalfWidth to FullWidth
```go
package main

import (
    "fmt"
    "github.com/suwakei/go-zhconv/zhconv"
)

func main() {
	result = zhconv.H2z("Hello, world!")
	fmt.Println(result) // Ｈｅｌｌｏ， ｗｏｒｌｄ！.

	result = zhconv.H2z("") // Empty string.
	fmt.Println(result) // "".

	result = zhconv.H2z("ＡＢＣ１２３アイウガパ") // No conversion needed (Zenkaku).
	fmt.Println(result) // ＡＢＣ１２３アイウガパ.

	result = zhconv.H2z("ABCdef XYZ!#$%&'()*+,-./:;<=>?@[¥]^_`{|}~ \\")
	fmt.Println(result) // ＡＢＣｄｅｆ　ＸＹＺ！＃＄％＆’（）＊＋，－．／：；＜＝＞？＠［￥］＾＿‘｛｜｝～　＼.

	result = zhconv.H2z("0123456789") // Hankaku Digits to Zenkaku.
	fmt.Println(result) // ０１２３４５６７８９.

	result = zhconv.H2z("ｱｲｳｴｵｶｷｸｹｺｻｼｽｾｿﾀﾁﾂﾃﾄﾅﾆﾇﾈﾉﾊﾋﾌﾍﾎﾏﾐﾑﾒﾓﾔﾕﾖﾗﾘﾙﾚﾛﾜｦﾝ") // Hankaku Katakana to Zenkaku.
	fmt.Println(result) // アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲン.

	result = zhconv.H2z("ｧｨｩｪｫｯｬｭｮ") // Hankaku Katakana (Small) to Zenkaku.
	fmt.Println(result) // ァィゥェォッャュョ.

	result = zhconv.H2z("｡､･ｰ｢｣")// Hankaku Katakana (Symbols) to Zenkaku.
	fmt.Println(result) // 。、・ー「」.

	result = zhconv.H2z("ｶﾞｷﾞｸﾞｹﾞｺﾞｻﾞｼﾞｽﾞｾﾞｿﾞﾀﾞﾁﾞﾂﾞﾃﾞﾄﾞﾊﾞﾋﾞﾌﾞﾍﾞﾎﾞｳﾞ") // Hankaku Katakana with Dakuten to Zenkaku".
	fmt.Println(result) // ガギグゲゴザジズゼゾダヂヅデドバビブベボヴ.

	result = zhconv.H2z( "ﾊﾟﾋﾟﾌﾟﾍﾟﾎﾟ") // Hankaku Katakana with Handakuten to Zenkaku.
	fmt.Println(result) // パピプペポ.

	result = zhconv.H2z("これはﾃｽﾄです｡123 ABC ｱｲｳ ｶﾞｷﾞｸﾞ ﾊﾟﾋﾟﾌﾟ!") // Mixed Hankaku/Zenkaku/Other.
	fmt.Println(result) // これはテストです。１２３　ＡＢＣ　アイウ　ガギグ　パピプ！.

	result = zhconv.H2z(" ｽﾍﾟｰｽ ") // convert from Half width space to Full width space.
	fmt.Println(result) //  　スペース　.

	result = zhconv.H2z("ｱﾞｲﾟﾝﾞ") // Dakuten/Handakuten cannot be applied.
	fmt.Println(result) // ア゛イ゜ン゛ Converted to full-width characters as separated( (ｱ->ア, ﾞ->ﾞ).

	result = zhconv.H2z("①②③㈱㈲") // Not convertible symbols.
	fmt.Println(result) // It is assumed that environment dependent characters will not be converted.

	result = zhconv.H2z("1ﾊﾞｲﾄ文字と2ﾊﾞｲﾄ文字が混在するﾃｷｽﾄ｡ ABC 123 ｶﾞｷﾞｸﾞﾊﾟﾋﾟﾌﾟ!?") // Long string with various conversions.
	fmt.Println(result) // １バイト文字と２バイト文字が混在するテキスト。　ＡＢＣ　１２３　ガギグパピプ！？.
}
```