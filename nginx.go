package nginx

import (
	"io"
	"os"
)

func runCLI(wr, ew io.Writer) int {
	return 0
}

func Main() int {
	return runCLI(os.Stdout, os.Stderr)
}
