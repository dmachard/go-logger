# go-logger

Basic logger library with the following features: 
- timestamp
- level support (INFO, ERROR and FATAL)
- verbose mode

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

logger := logger.New(true)
logger.Info("just a basic message!")
```

Output example:

```
INFO: 2021/07/04 14:18:22.270971 just a basic message
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