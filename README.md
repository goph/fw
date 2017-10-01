# Application framework

[![Build Status](https://img.shields.io/travis/goph/fw.svg?style=flat-square)](https://travis-ci.org/goph/fw)
[![Go Report Card](https://goreportcard.com/badge/github.com/goph/fw?style=flat-square)](https://goreportcard.com/report/github.com/goph/fw)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/goph/fw)


**A simple application framework with sane defaults.**


## Installation

Since this library uses [Glide](http://glide.sh/) I recommend using it in your
project as well.

```bash
$ glide get github.com/goph/fw
```


## Usage

```go
package main

import "github.com/goph/fw"

func main() {
    app := fw.NewApplication()
    defer app.Close()
    
    // your app logic
}
```

You can also take a look at some [boilerplate](https://github.com/deshboard/boilerplate-service)
code which relies on this framework.


## History

When I first started to work with Go I was amazed by the standard library.
(Almost) Everything I needed was already there. Of course all of the tools
were perfect, so I needed some external libraries as well (logging, error handling, etc),
but the main point is that I didn't need any frameworks to build my applications.

Soon I realized that this "no framework" philosophy requires a lot of copy-pasting.
So I created [boilerplates](https://github.com/deshboard/boilerplate-service) to
make copying easier. But it just didn't feel right either. It became clear
that maintaining 5-6 applications still requires too much time.

So I went back to the table and came up with this library. Although it is a
framework, I tried to build it in a way that supports easy extension. One could
even just copy the whole thing.

Using this library one can avoid copying a lot of (unmodified) code.
Furthermore, unlike the linked boilerplates the components in this library
are fully tested.


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
