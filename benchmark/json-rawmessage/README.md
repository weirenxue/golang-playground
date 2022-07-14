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
BenchmarkUnmarshal01
BenchmarkUnmarshal01-16            14325             84804 ns/op           27232 B/op        701 allocs/op
BenchmarkUnmarshal02
BenchmarkUnmarshal02-16            16741             72970 ns/op           24072 B/op        602 allocs/op
BenchmarkUnmarshal03
BenchmarkUnmarshal03-16            10000            122479 ns/op           45688 B/op       1003 allocs/op
BenchmarkUnmarshal04
BenchmarkUnmarshal04-16              704           1640174 ns/op           47392 B/op       9801 allocs/op
BenchmarkUnmarshal05
BenchmarkUnmarshal05-16             1552            720029 ns/op           25840 B/op        602 allocs/op
BenchmarkUnmarshal06
BenchmarkUnmarshal06-16              516           2313543 ns/op           67632 B/op      10103 allocs/op
```

`Unmarshal01` unmarshall json once and gets the full struct, but not `Unmarshal02` and `Unmarshal03`, which use `json.RawMessage` to provide flexibility to process json message twice. The difference between `Unmarshal02` and `Unmarshal03` is that the former is only unmarshal once and does not process `json.RawMessage`, but the latter unmarshal twice (once to `json.RawMessage` and once to processing `json.RawMessage`). `Unmarshal01` is less efficient than `Unmarshal02` but much more efficient than `Unmarshal03`.

The relationship between `Unmarshal04`, `Unmarshal05` and `Unmarshal06` is just like `Unmarshal01`, `Unmarshal02` and `Unmarshal03`, but with unmarshal for longer json and struct.

Unmarshal without `json.RawMessage` can get the full value of the struct and is more efficient than using `json.RawMessage`, so we should not use `json.RawMessage` if it is not necessary.

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
