# go-zhconv


This repository is a library that supports character conversion in the Go language. It performs mutual conversion between full-width and half-width characters and kana.

## installtion

```sh
go get github.com/suwakei/go-zhconv
```

## Usage
### convert from FullWidth to HalfWidth
```go
package main

import (
    "fmt"
    "github.com/suwakei/go-zhconv"
)

func main() {
    result := zhconv.Z2h("Ｈｅｌｌｏ， ｗｏｒｌｄ！")
    fmt.Println(result) // "Hello, world!"
}
```


### convert from HalfWidth to FullWidth
```go
package main

import (
    "fmt"
    "github.com/suwakei/go-zhconv"
)

func main() {
    result := zhconv.H2z("Hello, world!")
    fmt.Println(result) // "Ｈｅｌｌｏ， ｗｏｒｌｄ！"
}
```