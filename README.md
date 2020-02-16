# Mitra Reptir for Go

## Installation

1. Get the library.
```bash
go get -u github.com/rohmanhm/go-mitra-reptir
```
2. Use it in your application.
```go
package main

import (
	"fmt"
	"log"

	"github.com/rohmanhm/go-mitra-reptir/mitrareptir"
)

func main() {
	client = mitrareptir.New("YOUR_REPTIR_AUTH_KEY")
	account, err := client.GetAccount()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account)
}
```


### TODO

- [x] Account
- [x] Product List
- [x] Single List
- [x] Buy Product
- [ ] PPOB List
- [ ] PPOB Inquiry
- [ ] PPOB Payment

### How to run example

1. Go to `example` directory.
```bash
cd example
```
2. Replace file `.env.example` to `.env` and fill `REPTIR_AUTH_KEY` with your auth key.
3. Run.
```bash
go run [example-you-need-to-run]/*.go
```

## License

[Apache-2.0](google.com) - © 2020 [Rohman Habib M](mailto:mhrohman@live.com)
