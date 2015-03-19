# Ambrosio XKCD

XKCD plugin for Ambrosio.

Inspired by [Hubot's XKCD plugin](https://github.com/github/hubot-scripts/blob/master/src/scripts/xkcd.coffee)

### Version
0.5

### Pattern
```
xkcd latest
```
```
xkcd <number> // e.g. xkcd 999
```
### Usage

Teach the behaviours to Ambrosio:

```golang
import (
    "github.com/opedromiranda/ambrosio"
    "github.com/opedromiranda/ambrosio-xkcd"
)

func main() {
	steve := ambrosio.NewAmbrosio("Steve")

    steve.Teach(ambrosio_xkcd.Behaviours)

	steve.Listen(3000)
}

```

### Todo

Mock HTTP requests in unit tests

Create behaviour for "xkcd random"



License
----

MIT
