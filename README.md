
# go-binary-extract

  Extract a value from a json blob without parsing the whole thing.

  Read the [docs](http://godoc.org/github.com/juliangruber/go-binary-extract).

## Installation

```bash
$ go get github.com/juliangruber/go-binary-extract
```

## Example

```go
import "github.com/juliangruber/go-binary-extract"

value, err := Extract([]byte("{\"foo\":\"bar\"}"), "foo")
if err != nil {
  panic(err)
}
fmt.Println(value) // "bar"
```

## Perf

  With the object from `extract_test.go`, `Extract()` is 20x faster than
  `json.Unmarshal()`.

## License

  MIT

