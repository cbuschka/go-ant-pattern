# go-ant-pattern - An ant pattern parser.

## Usage
```
package main

import (
    "fmt"
	"github.com/cbuschka/go-ant-pattern"
)

func main() {

    path := "/"
	antPattern := ant_pattern.MustCompile("/**")

    if antPattern.Matches(path) {
        fmt.Printf("%s matches %s", path, antPattern.String());
    }  
}
```

## License
Copyright (c) 2021 by [Cornelius Buschka](https://github.com/cbuschka).

[MIT](./license.txt)
