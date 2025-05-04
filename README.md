# go-zhconv


This repository is a library that supports character conversion in the Go language. It performs mutual conversion between full-width and half-width characters and kana.

## installtion

```sh
go get -u github.com/suwakei/go-zhconv/zhconv
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
    result := zhconv.H2z("Hello, world!")
    fmt.Println(result) // "Ｈｅｌｌｏ， ｗｏｒｌｄ！"
}
```