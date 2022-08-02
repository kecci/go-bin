# go-bin

sdk library for bin

| source | website |
| ------ | ------- |
| binlist | https://binlist.net/ |
| freebinchecker | https://www.freebinchecker.com/ |
| bincheck | https://bincheck.io/ |


## usage binlist

```go
package main

import "github.com/kecci/go-bin/binlist"

func main() {
    res, err := binlist.BinLookup("548988")
    if err != nil {
        return
    }
    fmt.Println(*res)
}
```