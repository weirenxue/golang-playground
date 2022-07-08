# Benchmark json.RawMessage

This project attempts to understand the impact of json.RawMessage on unmarshall json performance.

## Result

```sh
make
go test -v -bench=. -run=none . -benchmem
goos: linux
goarch: amd64
pkg: benchmark
cpu: 11th Gen Intel(R) Core(TM) i9-11900 @ 2.50GHz
BenchmarkUnmarshallWithoutRawMessage
BenchmarkUnmarshallWithoutRawMessage-16              145           8390009 ns/op         3040070 B/op      80001 allocs/op
BenchmarkUnmarshallWithRawMessage01
BenchmarkUnmarshallWithRawMessage01-16               160           7675215 ns/op         3120071 B/op      80001 allocs/op
BenchmarkUnmarshallWithRawMessage02
BenchmarkUnmarshallWithRawMessage02-16               100          12540608 ns/op         5440077 B/op     130001 allocs/op
```

The difference between `UnmarshallWithRawMessage01` and `UnmarshallWithRawMessage02` is that `01` is only unmarshal once and does not process `json.RawMessage`, but `02` unmarshal twice (once to `json.RawMessage` and once to processing `json.RawMessage`). `UnmarshallWithoutRawMessage` is less efficient than `UnmarshallWithRawMessage01` but much more efficient than `UnmarshallWithRawMessage02`.

Unmarshal without `json.RawMessage` can get the full value of the struct, so we should not use `json.RawMessage` if it is not necessary.

### When do we need json.RawMessage?

A common way to handle messages from the client in WebSocket is to use the `type` field to route the request.

For example, the message might look like this

```json
{
    "type": "get_user_data",
    "data": {
        "user_id": "user_1"
    }
}
```

or this

```json
{
    "type": "get_product_data",
    "data": {
        "product_id": "product_1"
    }
}
```

We can't just create a struct to unmarshal these two message. So, how do we declare the structs?

```go
type MessageIn struct {
    Type string          `json:"type"`
    Data json.RawMessage `json:"data"`
}

type  GetUserData {
    ID string `json:"user_id"`
}

type  GetProductData {
    ID string `json:"product_id"`
}
```

And use them like this

```go
msg := json.RawMessage(`{
    "type": "get_user_data",
    "data": {
        "user_id": "1"
    }
}`)

var msgIn MessageIn
json.Unmarshal(msg, &msgIn)

switch msgIn.Type {
case "get_user_data":
    var data GetUserData
    json.Unmarshal(msgIn.Data, &data)
    fmt.Println(data.ID)
case "get_product_data":
    var data GetProductData
    json.Unmarshal(msgIn.Data, &data)
    fmt.Println(data.ID)
}
```
