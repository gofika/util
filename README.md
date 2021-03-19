[![codecov](https://codecov.io/gh/leaker/util/branch/main/graph/badge.svg)](https://codecov.io/gh/leaker/util)
[![Build Status](https://github.com/leaker/util/workflows/build/badge.svg)](https://github.com/leaker/util)
[![go.dev](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/leaker/util)
[![Go Report Card](https://goreportcard.com/badge/github.com/leaker/util)](https://goreportcard.com/report/github.com/leaker/util)
[![Licenses](https://img.shields.io/badge/license-bsd-orange.svg)](https://opensource.org/licenses/BSD-2-Clause)
[![donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](about::blank)

# util
golang utils for common use

## Basic Usage

### Installation

To get the package, execute:

```bash
go get github.com/leaker/util
```

### FileUtil
```go
package main

import (
	"fmt"
	"github.com/leaker/util/fileutil"
)

func main() {
	name := "foo/bar.txt"
	data := []byte("Hello")
	// write data to file. will create dir: foo
	err := fileutil.WriteFile(name, data)
	if err != nil {
		fmt.Printf("Write file failed. err: %s\n", err.Error())
		return
    }
    // check file exist
	if !fileutil.IsExist(name) {
		fmt.Printf("file %s not exist.\n", name)
		return
    }
    // clear temp file
    fileutil.DeleteAll("foo")
}
```

### GobUtil
```go
package main

import (
	"fmt"
	"github.com/leaker/util/gobutil"
)

func main() {
	type Foo struct {
		Name  string
		Value int
	}
	type Bar struct {
		Name  string
		Value int
	}
	foo := &Foo{"Jason", 100}
	// deep copy for different struct
	var bar Bar
	err := gobutil.DeepCopy(&bar, foo)
	if err != nil {
		fmt.Printf("DeepCopy failed. err: %s\n", err.Error())
		return
    }
	fmt.Printf("bar.Name: %s\n", bar.Name)
	fmt.Printf("bar.Value: %d\n", bar.Value)
}
```

### JsonUtil
```go
package main

import (
	"fmt"
	"github.com/leaker/util/jsonutil"
)

type Foo struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func main() {
	foo := &Foo{
		Name:  "Jason",
		Value: 100,
	}
	name := "foo.json"
	// write struct to file
	err := jsonutil.WriteFile(name, foo)
	if err != nil {
		fmt.Printf("WriteFile failed. err: %s\n", err.Error())
		return
	}
	// read struct from file
	var bar Foo
	err = jsonutil.ReadFile(name, &bar)
	if err != nil {
		fmt.Printf("ReadFile failed. err: %s\n", err.Error())
		return
    }
	fmt.Printf("bar.Name: %s\n", bar.Name)
	fmt.Printf("bar.Value: %d\n", bar.Value)
}
```