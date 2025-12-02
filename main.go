package debugger

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	Silent          = false         // Prevents any output from this package
	Verbose         = false         // Enables output from the Debug and Debugf functions
	Timestamps      = true          // Prefixes messages with a timestamp. The format can be set by changing the [TimestampFormat] variable in the module importing this
	TimestampFormat = time.DateTime // The string that timestamp prefixes will be formatted against. Does nothing if the Timestamps variable is set to false
)

// Output will print the message with the provided format regardless of he Verbose seting. It will respect the Silent seting.
func Output(format string, args ...any) {
	if Silent {
		return
	}
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	format = AddTimestamp(format)
	fmt.Printf(format, args...)
}

// Debug is an alias for Debugf
func Debug(format string, args ...any) {
	Debugf(format, args...)
}

// If [Verbose] is set to true, it will print the provided message to stdout
func Debugf(format string, args ...any) {
	if Verbose {
		Output(format, args...)
	}
}

// If [Timestamps] is set to true, AddTimestamp will prefix the provided format string with a formatted timestamp, as defined by [TimestampFormat]
func AddTimestamp(s string) string {
	if Timestamps {
		s = time.Now().Format(TimestampFormat) + ": " + s
	}
	return s
}

// DumpErrorStack will unwrap an error until no errors remain in the interface, presenting each layer along the way. If the initial error i non-empty, it will return 1. If the initial error is empty, it will return 0. This can be used for an exit code.
//
// It does not include timestamps in the output because those will not reflect when the error occured, but when the error is being read/printed. It ignores the Verbose setting. It does not ignore the Silent setting
func DumpErrorStack(err error) (i int) {
	var prefix string
	var errorDepth int
	for err != nil {
		i = 1
		if errorDepth == 0 {
			prefix = "Error stack: "
		} else {
			prefix = "           └"
		}
		errorDepth += 1
		Output("%s%v", prefix, err)
		err = errors.Unwrap(err)
	}
	return
}
