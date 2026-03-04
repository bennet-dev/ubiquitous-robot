# Go Quick Reference

## Basic Types

| Type | Example | Notes |
|------|---------|-------|
| `string` | `"hello"` | immutable |
| `int` | `42` | platform-dependent size |
| `int64` | `42` | explicit 64-bit |
| `float64` | `3.14` | default float |
| `bool` | `true` | true/false |
| `[]byte` | `[]byte("hi")` | byte slice |

## Composite Types

```go
// Slice (dynamic array)
items := []string{"a", "b", "c"}

// Map
m := map[string]int{"key": 1}

// Struct
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

## Print Functions (fmt package)

| Function | Output | Newline |
|----------|--------|---------|
| `Print` | stdout | no |
| `Println` | stdout | yes |
| `Printf` | stdout | no |
| `Fprint` | writer | no |
| `Fprintln` | writer | yes |
| `Fprintf` | writer | no |
| `Sprint` | string | no |
| `Sprintf` | string | no |

```go
fmt.Println("hello")                    // stdout + newline
fmt.Printf("val: %s\n", s)              // stdout formatted
fmt.Fprintf(os.Stderr, "err: %s", err)  // stderr
s := fmt.Sprintf("n: %d", n)            // to string
```

## Format Verbs

| Verb | Use |
|------|-----|
| `%s` | string |
| `%d` | integer |
| `%f` | float |
| `%v` | any (default) |
| `%+v` | struct w/ fields |
| `%T` | type |

## JSON (encoding/json)

```go
// Parse JSON into struct
var data MyStruct
err := json.Unmarshal([]byte(jsonStr), &data)

// Struct to JSON
bytes, err := json.Marshal(data)
```

## File I/O (os package)

```go
content, err := os.ReadFile("path")     // read entire file
err := os.WriteFile("path", data, 0644) // write file
_, err := os.Stat("path")               // check exists
```

## Slices (slices package)

```go
idx := slices.IndexFunc(items, func(x T) bool { return x.ID == 1 })
found := slices.ContainsFunc(items, func(x T) bool { return x.ID == 1 })
```

## Error Handling

```go
if err != nil {
    return err  // or handle it
}
```

## Loops

```go
for i, item := range items {
    // i = index, item = value
}

for key, val := range myMap {
    // iterate map
}
```
