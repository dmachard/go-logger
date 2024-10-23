<img src="https://img.shields.io/badge/go%20version-min%201.16-green" alt="Go version"/>

# go-logger

Basic logger library with the following features:

- timestamp
- level support (INFO, ERROR, WARNING and FATAL)
- verbose mode
- log to channel

Info log messages are redirected to stdout, stderr otherwise.

## Installation

```go
go get -u github.com/dmachard/go-logger
```

## Usage example

Create logger with verbose mode enabled.

```go
import (
    "github.com/dmachard/go-logger"
)

lg := logger.New(true)
lg.Info("just a basic message!")
```

Output example:

```bash
INFO: 2021/07/04 14:18:22.270971 just a basic message
```

## Usage example with channel

Create logger with channel

```go
import (
    "github.com/dmachard/go-logger"
)

logsChan := make(chan logger.LogEntry, 10)

lg := logger.New(false)
lg.SetOutputChannel((logsChan))

lg.Info("just a basic message!")

entry := <-logsChan
fmt.Println(entry.Level, entry.Message)
```

## Testing

```bash
$ go test -v
=== RUN   TestLogInfo
--- PASS: TestLogInfo (0.00s)
=== RUN   TestLogError
--- PASS: TestLogError (0.00s)
=== RUN   TestVerbose
--- PASS: TestVerbose (0.00s)
PASS
ok      github.com/dmachard/go-logger   0.002s
```

## Benchmark

```bash
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/dmachard/go-logger
cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
BenchmarkLogInfo-4       4094160               278.4 ns/op
PASS
ok      github.com/dmachard/go-logger   1.449s
```
