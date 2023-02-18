b64id

Package `b64id` Generates url-safe Base64 IDs like YouTube IDs.

# Getting started

## Installation

```shell
go get -u github.com/sdqri/b64id
```

## Example

```go
package main

import (
	"fmt"

	"github.com/sdqri/b64id"
)

func main() {
	id := b64id.GenerateB64ID(b64id.IDL11) //IDL11 for eleven character length IDSs
	fmt.Println(id) // // Will print id like ~"49HOWolBqrE"
}
```

## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](issues).

## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

This is distributed under the MIT License, see LICENSE for more information.
