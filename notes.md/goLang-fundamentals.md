# 🧠 Golang Fundamentals Guide for JS Developers

Welcome, Manish! This is your end‑to‑end Go learning guide, tailored for someone coming from **JavaScript/TypeScript + Node/Express**. It’s optimized for **VS Code** with:

- ✅ Clickable **Table of Contents**
- ✅ **Syntax‑highlighted** Go code blocks (` ```go `)
- ✅ Clear explanations + **JS ↔ Go** comparisons
- ✅ Practical examples and tips

> **Tip:** In VS Code, press **Ctrl/⌘ + Shift + V** to open Markdown preview. To auto-open `.md` files in preview, add this to `settings.json`:
>
> ```jsonc
> "workbench.editorAssociations": { "*.md": "vscode.markdown.preview.editor" }
> ```

---

## 📚 Table of Contents

- [1. Introduction](#1-introduction)
- [2. Setup & Environment](#2-setup--environment)
- [3. Language Basics](#3-language-basics)
- [4. Control Structures](#4-control-structures)
- [5. Functions](#5-functions)
- [6. Collections](#6-collections)
- [7. Structs & Interfaces](#7-structs--interfaces)
- [8. Error Handling](#8-error-handling)
- [9. Concurrency](#9-concurrency)
- [10. I/O and Networking](#10-io-and-networking)
- [11. Testing](#11-testing)
- [12. Tooling](#12-tooling)
- [13. Advanced Topics](#13-advanced-topics)
- [14. Best Practices](#14-best-practices)
- [15. Projects & Practice](#15-projects--practice)

---

## 1. Introduction

**Go (Golang)** is a statically typed, compiled language designed at Google for simplicity, speed, and **built‑in concurrency**.

**Why Go?**

- ⚡️ Fast build & runtime
- 🧩 Simple, minimal language surface
- 🧲 Strong standard library (HTTP, JSON, crypto, etc.)
- 🧵 First‑class concurrency (goroutines, channels)
- 🧰 Excellent tooling (format, vet, test, modules)

**JS ↔ Go quick view**

| Concept     | JavaScript/TypeScript           | Go                                               |
| ----------- | ------------------------------- | ------------------------------------------------ |
| Typing      | Dynamic (TS: erased at runtime) | **Static** (checked at compile time)             |
| Errors      | `throw/catch`                   | **`error` values** (return + handle)             |
| Concurrency | Event loop + Promises           | **Goroutines** + **channels** + `context`        |
| Collections | Arrays, Objects, Map            | **slices**, **maps**, **structs**                |
| OOP         | Classes/Prototypes              | **Structs + methods + interfaces** (composition) |

---

## 2. Setup & Environment

### Install Go

- Download from **https://go.dev/dl** and follow the installer.
- Verify:

```bash
go version
```

### Create your first module (project)

```bash
mkdir hello-go && cd hello-go
# Initialize a module (replace path with yours)
go mod init github.com/you/hello-go

cat > main.go <<'EOF'
package main
import "fmt"
func main() { fmt.Println("Hello, Go!") }
EOF

# Run the module
go run .
```

### Useful environment

```bash
go env        # inspect go env
```

### VS Code essentials

- Install **Go** extension (by Go Team at Google)
- Formatting & imports:

```jsonc
// settings.json
{
  "[go]": { "editor.defaultFormatter": "golang.go" },
  "go.formatTool": "goimports",
  "editor.formatOnSave": true,
  "editor.codeActionsOnSave": { "source.organizeImports": true }
}
```

---

## 3. Language Basics

### Packages, `main`, imports

```go
package main // executables must be package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Now:", time.Now())
}
```

### Variables, constants, zero values

```go
var a int        // 0
var s string     // ""
var ok bool      // false

x := 42          // short declare (func scope)
const Pi = 3.14159
```

### Types & conversions

```go
var n int = 10
var f float64 = float64(n) // explicit cast

var r rune = 'न'     // Unicode code point (int32)
var b byte = 'A'     // alias for uint8
```

### Strings & runes

```go
s := "नमस्ते"
for i, ch := range s { // ch is rune
    fmt.Printf("%d: %c\n", i, ch)
}
```

### Pointers (mutate / avoid copies)

```go
func inc(n *int) { *n++ }

func main(){
    x := 1
    inc(&x)
}
```

### Type alias vs defined type

```go
type CustomerID = string    // alias (same type)
type UserID string          // new distinct type
```

### `iota` (enum-like)

```go
type Role int
const (
    Admin Role = iota
    Editor
    Viewer
)
```

---

## 4. Control Structures

### If / For / Switch

```go
if v := compute(); v > 10 { fmt.Println(v) }

for i := 0; i < 3; i++ { fmt.Println(i) }     // classic
for condition() { /* while */ }
for { break }                                 // infinite

switch day := weekday(); day {
case "Mon", "Tue": fmt.Println("Weekday")
case "Sat", "Sun": fmt.Println("Weekend")
default: fmt.Println("Unknown")
}
```

### `range` loops

```go
nums := []int{10,20,30}
for i, v := range nums { fmt.Println(i, v) }
```

> **Note:** Modern Go makes **per‑iteration copies** of `i`/`v` in `range`, preventing the classic closure capture bug when launching goroutines inside loops.

### Defer, Panic, Recover (use sparingly)

```go
func safe() {
    defer func(){
        if r := recover(); r != nil { fmt.Println("recovered:", r) }
    }()
    mayPanic()
}
```

---

## 5. Functions

### Multiple returns & named results

```go
func splitHostPort(addr string) (host string, port int, err error) {
    // ...parse...
    return "localhost", 8080, nil // plain return uses named vars
}
```

### Variadic & higher‑order functions

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums { total += n }
    return total
}

func mapInts(xs []int, f func(int) int) []int {
    out := make([]int, 0, len(xs))
    for _, x := range xs { out = append(out, f(x)) }
    return out
}
```

### Closures (with loop)

```go
func main(){
    var fns []func()
    for i := 0; i < 3; i++ {
        v := i // capture a copy per iteration
        fns = append(fns, func(){ fmt.Println(v) })
    }
    for _, fn := range fns { fn() } // 0 1 2
}
```

---

## 6. Collections

### Arrays vs Slices

- **Array**: fixed length `[3]int{1,2,3}` (value type)
- **Slice**: dynamic view `[]int{1,2,3}` (most common)

```go
s := []int{1,2,3}
s = append(s, 4)            // may reallocate

buf := make([]byte, 2, 4)   // len=2 cap=4
b := make([]int, len(s))
copy(b, s)
```

**Three-index slicing** (limit capacity of subslice)

```go
x := []int{1,2,3,4,5}
sub := x[1:3:3] // len=2 cap=2
sub = append(sub, 9) // forces new array; x unaffected
```

### Maps

```go
m := map[string]int{"a":1}
m["b"] = 2
v, ok := m["c"] // ok=false if missing
delete(m, "a")
```

> Iteration order is **not** guaranteed.

### Sets with maps

```go
set := map[string]struct{}{}
set["go"] = struct{}{}
_, present := set["go"]
```

---

## 7. Structs & Interfaces

### Structs & methods

```go
type User struct {
    ID   int
    Name string
}

func (u *User) Rename(to string) { u.Name = to } // pointer receiver mutates
func (u User) Greeting() string { return "Hi, " + u.Name }
```

### Embedding (composition)

```go
type Logger struct{ *log.Logger }

type Service struct {
    Logger // methods promoted: s.Printf(...)
}
```

### Interfaces (implicit)

```go
type Reader interface { Read(p []byte) (int, error) }

type File struct{}
func (File) Read(p []byte) (int, error) { return 0, io.EOF }
// File now implements Reader without explicit declaration
```

### Type assertions & switches

```go
var any interface{} = 42
n, ok := any.(int)

switch v := any.(type) {
case int: fmt.Println("int", v)
case string: fmt.Println("string", v)
default: fmt.Println("unknown")
}
```

---

## 8. Error Handling

- Errors are **values** (`error` interface)
- Prefer `fmt.Errorf("...: %w", err)` to wrap
- Use `errors.Is` / `errors.As` for checks
- Use `panic` only for programmer bugs; recover at boundaries

```go
var ErrNotFound = errors.New("not found")

type User struct{ ID int }

func find(id int) (User, error) {
    // ...
    return User{}, ErrNotFound
}

u, err := find(7)
if err != nil {
    if errors.Is(err, ErrNotFound) { /* 404 */ }
}
```

```go
func do() error {
    if err := step1(); err != nil { return fmt.Errorf("step1: %w", err) }
    if err := step2(); err != nil { return fmt.Errorf("step2: %w", err) }
    return nil
}
```

---

## 9. Concurrency

### Goroutines

```go
func worker(id int){ fmt.Println("worker", id) }

func main(){
    for i := 0; i < 3; i++ { go worker(i) }
    time.Sleep(50 * time.Millisecond) // demo only
}
```

### Channels

```go
ch := make(chan int)

go func(){
    defer close(ch)
    for i := 0; i < 5; i++ { ch <- i }
}()

for v := range ch { fmt.Println(v) }
```

### `select` & timeouts

```go
select {
case v := <-ch:
    fmt.Println(v)
case <-time.After(500 * time.Millisecond):
    fmt.Println("timeout")
}
```

### `sync` & `context`

```go
var wg sync.WaitGroup
wg.Add(1)
go func(){ defer wg.Done(); /* work */ }()
wg.Wait()
```

```go
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
resp, err := http.DefaultClient.Do(req)
```

### Worker pool (bounded parallelism)

```go
jobs := make(chan int)
results := make(chan int)
var wg sync.WaitGroup

for w := 0; w < 4; w++ {
    wg.Add(1)
    go func(){
        defer wg.Done()
        for j := range jobs { results <- j*j }
    }()
}

go func(){ for i:=0;i<10;i++{ jobs <- i }; close(jobs) }()

go func(){ wg.Wait(); close(results) }()

for r := range results { fmt.Println(r) }
```

---

## 10. I/O and Networking

### Files

```go
b, err := os.ReadFile("in.txt")
if err != nil { log.Fatal(err) }
if err := os.WriteFile("out.txt", b, 0o644); err != nil { log.Fatal(err) }
```

### Streaming with `bufio`

```go
f, _ := os.Open("large.txt")
defer f.Close()
sc := bufio.NewScanner(f)
for sc.Scan() { fmt.Println(sc.Text()) }
```

### JSON

```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

u := User{ID:1, Name:"Manish"}
js, _ := json.Marshal(u)
var u2 User
_ = json.Unmarshal(js, &u2)
```

### HTTP server & client

```go
mux := http.NewServeMux()
mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"msg":"hello"})
})

log.Fatal(http.ListenAndServe(":8080", mux))
```

```go
resp, err := http.Get("http://localhost:8080/hello")
if err != nil { log.Fatal(err) }
defer resp.Body.Close()
var m map[string]string
json.NewDecoder(resp.Body).Decode(&m)
```

---

## 11. Testing

### Table‑driven tests

```go
// file: mathx/mathx_test.go
package mathx
import "testing"

func TestAdd(t *testing.T){
    tests := []struct{ a,b, want int }{
        {1,2,3}, {2,2,4}, {0,0,0},
    }
    for _, tt := range tests {
        if got := Add(tt.a, tt.b); got != tt.want {
            t.Fatalf("Add(%d,%d)=%d want %d", tt.a, tt.b, got, tt.want)
        }
    }
}
```

### Benchmarks

```go
func BenchmarkAdd(b *testing.B){
    for i := 0; i < b.N; i++ { _ = Add(1,2) }
}
```

### Run

```bash
go test ./...
go test -bench=. -benchmem ./...
```

---

## 12. Tooling

- **Format**: `go fmt ./...` (or `gofmt -s -w .`)
- **Vet**: `go vet ./...`
- **Build/Run**: `go build`, `go run .`
- **Doc**: `go doc`, `go list -m all`
- **Modules**:

```bash
go get github.com/google/uuid@latest
go mod tidy
```

- **Coverage**:

```bash
go test -cover -coverprofile=cover.out ./...
go tool cover -html=cover.out
```

- **Generate**: add `//go:generate` comments and run `go generate ./...`

**VS Code settings (Go)**

```jsonc
{
  "[go]": { "editor.defaultFormatter": "golang.go" },
  "go.formatTool": "goimports",
  "editor.formatOnSave": true,
  "editor.codeActionsOnSave": { "source.organizeImports": true },
  "editor.semanticHighlighting.enabled": true
}
```

---

## 13. Advanced Topics

### Generics (type parameters & constraints)

```go
// Map applies f to each element of in
func Map[T any, U any](in []T, f func(T) U) []U {
    out := make([]U, 0, len(in))
    for _, v := range in { out = append(out, f(v)) }
    return out
}

// Comparable set
type Set[T comparable] map[T]struct{}
```

### Type constraints example

```go
type Number interface { ~int | ~int64 | ~float64 }

func Sum[T Number](xs ...T) T {
    var s T
    for _, v := range xs { s += v }
    return s
}
```

### Reflection (use sparingly)

```go
v := reflect.ValueOf(anyVal)
if v.Kind() == reflect.Struct { /* ... */ }
```

### Profiling

```go
import _ "net/http/pprof"
// then visit /debug/pprof when server is running
```

---

## 14. Best Practices

- Keep the **happy path left‑aligned**; handle errors early
- **Accept interfaces, return structs** (small interfaces)
- **Context** as the first param in long‑running ops: `func (s *Svc) Do(ctx context.Context, ...)`
- Prefer **composition over inheritance** (embedding)
- Avoid global state; inject dependencies via interfaces
- Use pointer receivers when methods mutate or type is large
- Do **not** rely on map iteration order
- Pre‑size slices/maps when possible for performance
- Prefer `strings.Builder` / `bytes.Buffer` for string concat in loops
- Add comments for exported symbols (good for `go doc`)

---

## 15. Projects & Practice

Try these to solidify concepts:

1. **CLI Todo**: parse flags, read/write JSON file, list/add/remove tasks
2. **REST API**: `net/http` server with `/users` (GET/POST), JSON, context timeouts
3. **Worker Pool**: fetch URLs concurrently with rate limit and timeout
4. **File Processor**: stream a big file, count lines, output stats
5. **Caching Layer**: map + mutex with TTL, benchmark with `testing` pkg

**Next steps**

- Add table‑driven tests + benchmarks to each project
- Use `errgroup` (from `golang.org/x/sync/errgroup`) for goroutine lifecycles
- Add graceful shutdown with `http.Server.Shutdown`

---

### Appendix: Common One‑Liners

```go
// String ops
strings.Contains("hello", "he")
strings.Split("a,b,c", ",")
strings.Join([]string{"a","b"}, ",")

// Convert
n, _ := strconv.Atoi("42")
s := strconv.Itoa(42)

// Sort
sort.Ints([]int{3,1,2})

// Timers
<-time.After(500 * time.Millisecond)

// Env
home := os.Getenv("HOME")

// Quick HTTP GET
resp, _ := http.Get("https://example.com")
```

Happy Go hacking! 🏃‍♂️💨
