
# Go Fundamentals – A Practical Cheatsheet for JavaScript Devs

> **How to use:** Save this as `go-fundamentals-cheatsheet-for-js-devs.md` (Markdown). Open in VS Code for syntax-highlighted Go code, clickable TOC, and easy navigation. Use the built‑in Markdown preview (⌘/Ctrl+Shift+V).

---

## Table of Contents
- [Mindset Shift (JS → Go)](#mindset-shift-js--go)
- [Setup & Project Basics](#setup--project-basics)
- [Language Building Blocks](#language-building-blocks)
  - [Packages, `main`, and imports](#packages-main-and-imports)
  - [Variables, constants, zero values](#variables-constants-zero-values)
  - [Types (numbers, strings, runes, bool)](#types-numbers-strings-runes-bool)
  - [Pointers (why they matter)](#pointers-why-they-matter)
  - [Functions & multiple returns](#functions--multiple-returns)
  - [Errors: `error`, wrapping, `panic/recover`, `defer`](#errors-error-wrapping-panicrecover-defer)
  - [Control flow: `if`, `for`, `switch`, `range` (Go 1.22 note)](#control-flow-if-for-switch-range-go-122-note)
- [Collections](#collections)
  - [Arrays vs Slices](#arrays-vs-slices)
  - [Maps](#maps)
  - [Structs & Methods](#structs--methods)
  - [Interfaces & Composition (no classes)](#interfaces--composition-no-classes)
  - [Generics (type parameters)](#generics-type-parameters)
- [Concurrency Essentials](#concurrency-essentials)
  - [Goroutines](#goroutines)
  - [Channels & `select`](#channels--select)
  - [Context for cancellation/timeouts](#context-for-cancellationtimeouts)
  - [Common concurrency patterns](#common-concurrency-patterns)
- [I/O, JSON, and HTTP](#io-json-and-http)
  - [Reading/Writing files](#readingwriting-files)
  - [JSON marshal/unmarshal](#json-marshalunmarshal)
  - [HTTP server/client](#http-serverclient)
- [Testing & Tooling](#testing--tooling)
  - [Table-driven tests](#table-driven-tests)
  - [Benchmarks](#benchmarks)
  - [Format, lint, vet](#format-lint-vet)
- [Modules & Dependency Management](#modules--dependency-management)
- [Idioms, Gotchas, and Best Practices](#idioms-gotchas-and-best-practices)
- [Mini Reference: Common One-Liners](#mini-reference-common-one-liners)
- [Recommended VS Code Setup](#recommended-vs-code-setup)

---

## Mindset Shift (JS → Go)
- **Compiled & typed**: Go is statically typed and compiled to a single binary. No runtime type changes like JS.
- **Errors are values**: No exceptions for flow control; return `error` explicitly.
- **Simplicity over magic**: Fewer language features, strong standard library, readability matters.
- **Concurrency is built‑in**: Goroutines and channels are first‑class.
- **No classes**: Use **structs + methods** and **interfaces** (implicit implementation) instead of classes/inheritance.

```diff
+ JS mindset: dynamic, prototype, promises, exceptions
+ Go mindset: static types, composition, goroutines, errors-as-values
```

---

## Setup & Project Basics
```bash
# Install Go (https://go.dev/dl/) and verify
go version

# Create a module (project)
mkdir hello-go && cd hello-go
go mod init github.com/you/hello-go

# Add a file and run
cat > main.go <<'EOF'
package main
import "fmt"
func main() { fmt.Println("Hello, Go!") }
EOF

go run .
```

**Project layout tips**
- One module per repo. Use packages under `/pkg` or root folders (`/internal`, `/cmd`, `/pkg`).
- `internal/` makes packages importable **only** from inside this module.
- `cmd/appname/main.go` is a common convention for binaries.

---

## Language Building Blocks

### Packages, `main`, and imports
```go
package main // executable packages must be main

import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Println(rand.Intn(10))
}
```
- Import paths are module-aware. Unused imports/vars **fail** to compile.

### Variables, constants, zero values
```go
var a int          // zero value: 0
var s string       // ""
var ok bool        // false

// Short declaration (function scope only)
x := 42

const Pi = 3.14159
```

### Types (numbers, strings, runes, bool)
```go
var n int        // machine word (usually 64‑bit on 64‑bit)
var n8 int8      // -128..127
var u32 uint32
var f64 float64
var b bool
var s string
var r rune       // alias for int32; holds a Unicode code point
```
- Strings are **immutable**. Indexing a string yields a **byte**; convert to `rune` for Unicode.

```go
s := "नमस्ते"
for _, ch := range s { // ch is rune
    fmt.Printf("%c ", ch)
}
```

### Pointers (why they matter)
- A pointer holds the **address** of a value; used to mutate data or avoid copies.
```go
func inc(n *int) { *n++ }

func main() {
    x := 1
    inc(&x)
}
```
- No pointer arithmetic (unlike C). Garbage‑collected.

### Functions & multiple returns
```go
func splitHostPort(addr string) (host string, port int, err error) {
    // named results; a plain `return` returns current values
    return "localhost", 8080, nil
}

h, p, err := splitHostPort("localhost:8080")
if err != nil { /* handle */ }
```

### Errors: `error`, wrapping, `panic/recover`, `defer`
```go
import (
    "errors"
    "fmt"
)

func readFile(name string) ([]byte, error) {
    if name == "" {
        return nil, errors.New("empty name")
    }
    // ...
    return nil, fmt.Errorf("read %s: %w", name, errors.New("disk fail"))
}

// defer runs at function exit (LIFO); great for closing resources
f, err := os.Open(path)
if err != nil { /* handle */ }
defer f.Close()
```
- Use `panic` for truly exceptional programmer errors; recover at boundaries only.

### Control flow: `if`, `for`, `switch`, `range` (Go 1.22 note)
```go
if v := compute(); v > 10 { /* scoped v */ }

for i := 0; i < 3; i++ { /* classic for */ }
for condition() { /* while */ }
for { break } // infinite

switch x := kind(); x {
case 1, 2: /* ... */
case 3:     /* ... */
default:    /* ... */
}

// range over slice/map/string
for i, v := range []int{1,2,3} { _ = i; _ = v }
```
> **Note:** Modern Go makes a fresh copy of the range loop variables each iteration, which avoids common closure bugs when capturing `i`/`v` inside goroutines.

---

## Collections

### Arrays vs Slices
- **Array**: fixed length, value type: `[3]int{1,2,3}`
- **Slice**: dynamic view over an array: `[]int{1,2,3}` (most common)

```go
s := []int{1,2,3}
s = append(s, 4)

// make: len=2, cap=4
buf := make([]byte, 2, 4)

// Copying
b := make([]int, len(s))
copy(b, s)
```

**Slice gotchas**
- Slices share backing arrays. Modifying a subslice may mutate the original until a new allocation happens.
```go
x := []int{1,2,3,4}
y := x[:2]   // shares storage
y[0] = 99    // x[0] also becomes 99
```

### Maps
```go
m := map[string]int{"a": 1}
m["b"] = 2
v, ok := m["c"]    // ok=false if missing
delete(m, "a")

// Initialization with capacity hint
m2 := make(map[string]int, 16)
```
- Map iteration order is **not** guaranteed.

### Structs & Methods
```go
type User struct {
    ID   int
    Name string
}

// Method with pointer receiver (mutates)
func (u *User) Rename(to string) { u.Name = to }

// Value receiver (copy)
func (u User) Greeting() string { return "Hi, " + u.Name }
```

### Interfaces & Composition (no classes)
- Interfaces are **implicitly** satisfied by any type that implements the methods.
```go
type Reader interface { Read(p []byte) (int, error) }

type File struct{}
func (File) Read(p []byte) (int, error) { return 0, io.EOF }

// File now implements Reader without explicit declaration.
```
- Compose behavior via **embedding**:
```go
type Logger struct{ *log.Logger }

type Service struct {
    Logger // embedded; methods are promoted
}
```

### Generics (type parameters)
```go
func Map[T any, U any](in []T, f func(T) U) []U {
    out := make([]U, 0, len(in))
    for _, v := range in { out = append(out, f(v)) }
    return out
}

nums := []int{1,2,3}
squares := Map(nums, func(n int) int { return n*n })
```

---

## Concurrency Essentials

### Goroutines
```go
func worker(id int) { fmt.Println("worker", id) }

func main() {
    for i := 0; i < 3; i++ {
        go worker(i) // lightweight concurrent function
    }
    time.Sleep(100 * time.Millisecond) // demo only; prefer sync primitives
}
```

### Channels & `select`
```go
ch := make(chan int)

go func() {
    defer close(ch)
    for i := 0; i < 5; i++ { ch <- i }
}()

for v := range ch { fmt.Println(v) }
```

```go
// Multiplexing with select
done := make(chan struct{})
out := make(chan string)

go func() {
    for {
        select {
        case out <- "tick":
        case <-done:
            return
        }
    }
}()

select {
case v := <-out:
    fmt.Println(v)
case <-time.After(1 * time.Second):
    fmt.Println("timeout")
}
```

### Context for cancellation/timeouts
```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
res, err := http.DefaultClient.Do(req)
```
- Pass `ctx context.Context` as the **first** parameter.

### Common concurrency patterns
- **Worker pool** (bounded parallelism)
- **Fan‑out/Fan‑in** (distribute + gather)
- **Errgroup** simplifies goroutine lifecycles with cancel on first error

```go
// Worker pool sketch
jobs := make(chan int)
results := make(chan int)
var wg sync.WaitGroup

for w := 0; w < 4; w++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        for j := range jobs { results <- j*j }
    }()
}

go func() { for i:=0;i<10;i++{ jobs <- i }; close(jobs) }()

go func() { wg.Wait(); close(results) }()

for r := range results { fmt.Println(r) }
```

---

## I/O, JSON, and HTTP

### Reading/Writing files
```go
b, err := os.ReadFile("input.txt")
if err != nil { /* handle */ }

err = os.WriteFile("out.txt", b, 0o644)
```

### JSON marshal/unmarshal
```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

u := User{ID: 1, Name: "Manish"}

// Go → JSON
jb, _ := json.Marshal(u)      // {"id":1,"name":"Manish"}

// JSON → Go
var u2 User
_ = json.Unmarshal(jb, &u2)
```

### HTTP server/client
```go
// Server
mux := http.NewServeMux()
mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"msg": "hello"})
})

srv := &http.Server{Addr: ":8080", Handler: mux}
log.Println("listening on :8080")
log.Fatal(srv.ListenAndServe())
```

```go
// Client
resp, err := http.Get("http://localhost:8080/hello")
if err != nil { /* handle */ }
defer resp.Body.Close()
var m map[string]string
json.NewDecoder(resp.Body).Decode(&m)
```

---

## Testing & Tooling

### Table-driven tests
```go
// file: mathx/mathx_test.go
package mathx

import "testing"

func TestAdd(t *testing.T) {
    tests := []struct{
        a, b, want int
    }{{1,2,3}, {2,2,4}}

    for _, tt := range tests {
        got := Add(tt.a, tt.b)
        if got != tt.want {
            t.Fatalf("Add(%d,%d)=%d want %d", tt.a, tt.b, got, tt.want)
        }
    }
}
```
```bash
# Run tests
go test ./...
```

### Benchmarks
```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ { _ = Add(1,2) }
}
```
```bash
go test -bench=. -benchmem ./...
```

### Format, lint, vet
```bash
go fmt ./...        # or `gofmt -s -w .`
go vet ./...
```

---

## Modules & Dependency Management
```bash
# Add a dependency (import it in code first)
go get github.com/google/uuid@latest

go mod tidy     # prune unused + add needed

go list -m all  # view dependency graph
```
- Pin versions with `@v1.2.3`. Use semantic import versioning (v2+ has `/v2` in import path).
- `replace` in `go.mod` is useful for local development.

---

## Idioms, Gotchas, and Best Practices
- Prefer small, focused interfaces ("accept interfaces, return structs").
- Handle errors immediately; keep the happy path left-aligned.
- Avoid global state. Use dependency injection via interfaces.
- Use pointer receivers when the method mutates or the type is large; value receivers for small immutable-like types.
- `time.Duration` is in **nanoseconds**; express with helpers: `500 * time.Millisecond`.
- Log with `log` or `slog`. Use `context` for request-scoped values and cancellation.
- Map lookups: check `ok` when presence matters.
- Do not rely on map iteration order.
- Prefer `bytes.Buffer` or `strings.Builder` for string concatenation in loops.
- For performance, pre-size slices/maps when possible.
- In `for range`, loop variables are per-iteration (modern Go), but still copy values (be mindful with large structs).

---

## Mini Reference: Common One-Liners
```go
// Convert string↔int
n, _ := strconv.Atoi("42")
s := strconv.Itoa(42)

// String contains, split, join
strings.Contains("hello", "he")
strings.Split("a,b,c", ",")
strings.Join([]string{"a","b"}, ",")

// Sort
sort.Ints([]int{3,1,2})

// Filesystem walk
filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error { return nil })

// Timers
<-time.After(500 * time.Millisecond)

t := time.NewTicker(time.Second)
defer t.Stop()
for range t.C { break }

// Random
rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10)

// Env
os.Getenv("HOME")

// HTTP quick GET
http.Get("https://example.com")
```

---

## Recommended VS Code Setup
- Install **Go** extension (by Go Team at Google). It brings `gopls` (language server), linting, code actions.
- Enable format-on-save and organize imports:
```jsonc
// .vscode/settings.json
{
  "go.formatTool": "gofmt",
  "editor.formatOnSave": true,
  "go.lintOnSave": "package",
  "go.vetOnSave": "package",
  "editor.codeActionsOnSave": {
    "source.organizeImports": true
  }
}
```
- Optional: add tasks for `go test`, `go vet`, `go build`.

---

## JS ↔ Go Quick Mapping (mental model)
```diff
 JS                      | Go
-------------------------+------------------------------
 let/const (block scope) | var/const; short := in funcs
 Objects                 | structs (typed fields)
 Classes/Inheritance     | structs + methods + interfaces
 Promises/async/await    | goroutines + channels + context
 Exceptions              | explicit error returns; panic for bugs
 Array                   | slices (resizable) / arrays (fixed)
 Map / Object            | map[key]value (unordered)
 JSON.parse/stringify    | json.Unmarshal/Marshal
 npm/yarn/pnpm           | go modules (go.mod), `go get`, `go run`
 console.log             | log.Println / fmt.Printf / slog
 TypeScript              | Go's built-in static types, generics
```

---

### Next steps to practice (15–30 mins)
1. Build a CLI: parse flags with `flag` package; print colored output.
2. Write a REST server: `/users` GET/POST with `net/http` and JSON.
3. Add table-driven tests and a benchmark.
4. Introduce `context` for request timeouts.

Happy Go hacking! 🏃‍♂️💨
