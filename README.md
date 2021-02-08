# Human Readable UUID

A "human readable UUID" is (something I just made up, and) simply the first half of a full UUID + two random words at the end.
This makes the IDs easier to read, track and communicate between humans while preserving enough uniqueness.

Examples of output:
- 7e1e2527-62d6-divine-grass
- e6633073-dff0-broken-moon
- 2a761d30-a2c1-weathered-field

The random words are generated by [goombaio/namegenerator](https://github.com/goombaio/namegenerator).

## Usage

```
go get github.com/eaardal/human-uuid
```

```go
// Create a new ID
id := humanuuid.New() // id="e6633073-dff0-broken-moon"

// Validate a string
if humanuuid.IsValid(id) {
    // ...
}
```