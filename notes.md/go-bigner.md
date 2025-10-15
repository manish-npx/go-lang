# 🧭 Complete Golang Guide for Beginners to Advanced

---

## 🟢 Basic Entry

### 👶 Common Mistakes Newbies Make

- Misunderstanding zero values
- Confusing nil slices vs empty slices
- Improper use of goroutines and channels
- Forgetting to initialize maps
- Misusing `defer` and `panic`
- Not understanding value vs reference types

---

## 🔤 Data Types

- **Nil vs Empty Slice**: Nil slice has no memory allocated; empty slice (`[]T{}`) does.
- **String to Byte Array**: Memory copy occurs.
- **Flip String with Mixed Characters**: Requires rune-level reversal for proper Unicode handling.
- **Large vs Small Slice Copy**: Larger slices consume more memory and time.
- **Uninitialized Map**: Causes runtime panic on assignment.
- **Map Length**: `len(map)` returns the number of key-value pairs.
- **Map Size**: Internally managed, not directly accessible.
- **Map Safety**: Not thread-safe; use `sync.Map` for concurrency.
- **Immutable Strings**: Convert to `[]rune` or `[]byte` to modify.
- **Array Sorted Check**: Iterate and compare adjacent elements.
- **Coroutine Safety Without Locking**: Use channels or atomic operations.
- **Array vs Slice**: Arrays are fixed-size; slices are dynamic.
- **JSON Tag Missing**: Field won’t be marshaled.
- **Reflect & JSON Tags**: Use `reflect.Type.Field(i).Tag.Get("json")`; private fields are not exported.
- **Zero/Empty/Nil Slicing**: Different memory and behavior implications.
- **Slice Copy**: `copy()` does shallow copy; deep copy requires manual handling.
- **Map Expansion**: Triggered by load factor and internal heuristics.
- **Map Scaling Strategy**: Rehashing and bucket expansion.
- **Custom Type ↔ Byte Slice**: Requires manual conversion or encoding packages.
- **make vs new**: `make` for slices, maps, channels; `new` returns pointer.
- **Slice/Map/Channel Creation Parameters**: Capacity, length, etc.
- **Slice len/cap/share/expand**: `len` is current size, `cap` is total capacity.
- **Thread-safe Map**: Use `sync.Map` or mutex.
- **Struct Comparison**: Possible if all fields are comparable.
- **Map Read Order**: Not guaranteed.
- **Set in Go**: Use `map[T]struct{}`.
- **Map Expansion Mechanism**: Internal rehashing.
- **Map with nil Value**: Safe but may cause logic issues.
- **This Pointer**: Implicit receiver in methods.
- **Default Values**: Zero values for types.
- **Reference Types**: Slices, maps, channels, pointers, interfaces.
- **Range Map Order**: Not ordered.
- **Slice Expansion**: Doubling strategy.
- **Pointer Operations**: Dereferencing, address-of.
- **Type Value Change**: Depends on mutability.
- **JSON Parsing Defaults**: Zero values if field missing.
- **Array Parameter Passing**: Passed by value.

---

## 🔁 Process Control

- **For Loop Append**: May cause unexpected behavior due to shared backing array.
- **Closed Channel**: Sending panics; receiving returns zero value.
- **Defer**: Executes after function returns.
- **Select**: Multiplex channel operations.
- **Context Package**: Cancellation, timeout, deadlines.
- **Select Use Cases**: Channel coordination, timeouts.
- **Defer in Goroutine**: Executes when goroutine exits.
- **Switch Case Enforcement**: Use `fallthrough`.
- **Recover from Panic**: Use `recover()` inside deferred function.

---

## 🔷 Advanced

### 📦 Package Management

- **Go Modules (`go mod`)**: Dependency management, versioning, reproducibility.

### ⚙️ Optimizing

- **Memory Escape**: Analyze with `go build -gcflags="-m"`.
- **Memory Leak**: Use `pprof`, trace.
- **Fragmentation**: Caused by uneven allocation/deallocation.
- **Goroutine Leaks**: Often due to blocked channels.
- **sync.Pool**: Object reuse, reduces GC pressure.
- **Go1.13 vs Go1.12**: Improved allocation and GC behavior.

### 🔄 Concurrent Programming

- **Closed Channel Read/Write**: Panic on write, zero value on read.
- **Uninitialized Channel**: Nil channel blocks forever.
- **sync.Map**: Thread-safe, optimized for read-heavy workloads.
- **Wait for Goroutines**: Use `sync.WaitGroup`.
- **Buffered vs Unbuffered Channel**: Buffered allows async sends.
- **Communication Modes**: Channels, shared memory.
- **Channel Internals**: Ring buffer, mutexes.
- **RWLock Implementation**: Reader/writer counters.
- **CSP Model**: Communicating Sequential Processes.
- **Channel Safety**: Built-in synchronization.

### 🧠 Advanced Features

- **uintptr vs unsafe.Pointer**: `uintptr` for arithmetic, `unsafe.Pointer` for conversions.
- **Reflect Tags**: Only exported fields accessible.
- **Coroutine vs Thread**: Goroutines are lightweight, managed by Go runtime.
- **GC Process**: Mark-and-sweep, concurrent GC.
- **Write Barrier**: Ensures GC consistency.
- **Interface Assertion Copy**: Depends on type.
- **Interface Implementation**: Method set matching.
- **GMP Model**: Goroutine, Machine, Processor.
- **GC Pressure**: Many small objects increase GC frequency.
- **Closure Implementation**: Captures outer scope variables.
- **GC Cycles**: Triggered by allocation thresholds.
- **Goroutine Suspension**: Scheduler decision.
- **Data Race Detection**: `go run -race`.
- **Anomaly Scenarios**: Deadlocks, leaks, panics.
- **net/http Long Connection**: Keep-alive, connection pooling.
- **Advanced Interview Questions**: Deep dive into internals.

---

## 🔍 Question Investigation

- **Trace**: Execution timeline.
- **pprof**: CPU, memory profiling.
- **Goroutine Leak**: Unfinished goroutines.
- **Memory Leak in Production**: Use profiling tools, analyze heap.

---

## 📖 Source Code Reading

- `sync.Map`, `net/http`, `mutex`, `channel`, `context`, `select`, `main`, `GC`, `timer`

---

## 🧵 Compendium

- Compilation Introduction: Go compiler phases, SSA form.

---

## 📚 Recommended Books

- _The Go Programming Language_ by Alan Donovan & Brian Kernighan
- _Go in Action_ by William Kennedy
- _Concurrency in Go_ by Katherine Cox-Buday
- _Go Systems Programming_ by Mihalis Tsoukalos

---

## 🎥 Video Tutorials

- [Golang Family](https://www.youtube.com/@golangfamily)
- [JustForFunc](https://www.youtube.com/c/JustForFunc)
- [GoBridge](https://www.youtube.com/c/GoBridge)

---

## 🛠️ Practically Used Tools

---

## 🗄️ Database Connectivity in Go

- **Popular Packages**:

  - `database/sql`: Standard library for SQL databases
  - Drivers:
    - MySQL: `github.com/go-sql-driver/mysql`
    - PostgreSQL: `github.com/lib/pq`
    - SQLite: `github.com/mattn/go-sqlite3`

- **Basic Example** (PostgreSQL):

```go
import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

func main() {
    connStr := "user=username dbname=mydb password=mypassword sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    rows, err := db.Query("SELECT id, name FROM users")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        var id int
        var name string
        err := rows.Scan(&id, &name)
        if err != nil {
            log.Fatal(err)
        }
        log.Println(id, name)
    }
}
```

## Folder Structure

```
/project-root
│
├── cmd/                # Entry points (main apps)
│   └── app/
│       └── main.go
│
├── pkg/                # Library code used by applications
│
├── internal/           # Private application code
│   ├── config/         # Config parsing
│   ├── handlers/       # HTTP handlers/controllers
│   ├── models/         # Data models, structs
│   ├── repository/     # DB access layer
│   ├── service/        # Business logic layer
│   └── utils/          # Utility packages
│
├── api/                # API definitions (OpenAPI/Swagger)
├── migrations/         # DB migration files
├── scripts/            # Build or deployment scripts
├── web/                # Frontend React app (if mono repo)
└── go.mod


```
