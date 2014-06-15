# cakeday

[![Build Status](https://drone.io/github.com/darthlukan/cakeday/status.png)](https://drone.io/github.com/darthlukan/cakeday/latest)

> Author: Brian Tomlinson <brian.tomlinson@linux.com>

## Description

> Cakeday is a small library that provides a [Reddit](www.reddit.com) user's "Cake Day" given a Reddit username
> as input.  It returns this date in the format "d M".

> *Example:* 6 Jun

## Installation

```bash
    $ go get github.com/darthlukan/cakeday
```

## Usage

```go
    package main

    import (
        "fmt"
        "github.com/darthlukan/cakeday"
    )

    func main() {
        username := "darthlukan"
        response, err := cakeday.Get(username)
        if err != nil {
            panic(err)
        }

        fmt.Printf("%v\n", response)
    }
```

> This outputs:

```
    Reddit Cake Day for darthlukan is: 6 Jun
```

## License

> GPLv3, see LICENSE file.
