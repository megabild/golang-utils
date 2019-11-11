# Utils

Some handy utils for the go programming language.

## Installation

```shell script
$ go get github.com/megabild/golang-utils
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/megabild/golang-utils/mbpath"
	"github.com/megabild/golang-utils/mbstring"
)

func main() {
	// prints "Hello World!"
	fmt.Println(mbstring.Coalesce("  ","Hello World!"))
}
```

## License
[MIT](https://choosealicense.com/licenses/mit/)