package log

import (
	"fmt"
	"os"
)

// StdErr outputs a message & the string form of a data structure (`v`) to the stderr stream
// 	this is helpful, as Concourse consumes the `stdout` stream for metadata
func StdErr(msg string, v interface{}) {
	fmt.Fprintf(os.Stderr, "%s: %s", msg, v)
}
