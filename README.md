# osu!go
An unofficial wrapper library for the osu! web API made for the Go language.

[![docs](https://godoc.org/github.com/renellc/osugo?status.svg)](https://godoc.org/github.com/renellc/osugo)
[![Go Report Card](https://goreportcard.com/badge/github.com/renellc/osugo)](https://goreportcard.com/report/github.com/renellc/osugo)

## Introduction
osu!go is a wrapper library that allows you to interact and make calls to the osu! web API with
ease within Go. My purpose for making this library was mainly to practice writing Go code but
seeing as how other wrapper libraries for the osu! web API haven't been updated in quite a
while, I thought I would make it public for everyone to use.
  
## Installation
This package was written in Go version 1.12, so having Go 1.12 is recommended. I haven't tested
this package on previous versions of Go, but it should work with any release version of Go.

## Usage
The following shows the general usage of the library. You can import this package by adding
 `import github.com/renellc/osugo` at the top of your file.

```
package main

import "github.com/renellc/osugo"

func main() {
    c := osugo.InitClient("your-osu-api-key-goes-here")
    user, err := c.GetUser(UserBestQuery{
        User: "nathan on osu" // user IDs are valid as well
    })

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    user.Print()
}
```

## Resources
- [osu! API Wiki Page](https://github.com/ppy/osu-api/wiki)

## Support
Having problems or are there bugs? Feel free to open an issue stating what your problem is and
 how you might go about recreating your problem, I'm happy to help.
 
 ## Contributing
 I'm open to anyone forking and creating pull requests if it means improving the library. If the
  pull request is in reference to an already created issue, please reference that in the request.
