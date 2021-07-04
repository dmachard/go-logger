package logger

import (
	"bytes"
	"regexp"
	"sync"
	"testing"
)

const Pattern_ts = `\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2}.\d{6} `

func TestLogInfo(t *testing.T) {
	// prepare a buffer instead of stdout and a string to search
	const msg = "hello world"
	var o bytes.Buffer

	// init the logger and redirect output to buffer
	l := New(true)
	l.SetOutput(&o)

	// log info message and check format output with regexp
	l.Info(msg)
	var pattern = `^` + prefixInfo + Pattern_ts + msg + `\n$`
	if regexp.MustCompile(pattern).MatchString(o.String()) != true {
		t.Errorf("expect level info msg: %q", o.String())
	}
}

func TestLogError(t *testing.T) {
	// prepare a buffer instead of stdout and a string to search
	const msg = "hello world"
	var o bytes.Buffer
	var pattern = `^` + prefixError + Pattern_ts + msg + `\n$`

	// init the logger and redirect output to buffer
	l := New(true)
	l.SetOutput(&o)

	// log error and check format output with regexp
	l.Error(msg)
	if regexp.MustCompile(pattern).MatchString(o.String()) != true {
		t.Errorf("expect level error msg: %q", o.String())
	}
}

func TestVerbose(t *testing.T) {
	// prepare a buffer instead of stdout and a string to search
	const msg = "hello world"
	var o bytes.Buffer

	// init the logger with verbose to false and redirect output to buffer
	l := New(false)
	l.SetOutput(&o)

	// log info
	// because verbose mode is disabled, no log should appears
	l.Info(msg)
	if o.String() != "" {
		t.Errorf("expect no log message: %q", o.String())
	}

	// log error
	// verbose mode is disabled, error must always appears
	l.Error(msg)
	if o.String() == "" {
		t.Errorf("expect one message: %q", o.String())
	}
}

func TestLogInfoRace(t *testing.T) {
	var b bytes.Buffer
	var wg sync.WaitGroup

	l := New(true)
	l.SetOutput(&b)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			for z := 0; z < 100; z++ {
				l.Info("%d - %d", j, z)
			}
		}(i)
	}

	wg.Wait()

}

func BenchmarkLogInfo(b *testing.B) {
	var buf bytes.Buffer
	l := New(true)
	l.SetOutput(&buf)
	for i := 0; i < b.N; i++ {
		buf.Reset()
		l.Info("test")
	}
}
