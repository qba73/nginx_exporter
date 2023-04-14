package nginx

import (
	"fmt"
	"io"
	"os"
)

const parseFormat = `Active connections: %d
server accepts handled requests
 %d %d %d
Reading: %d Writing: %d Waiting: %d`

// Connections represents NGINX OSS connection statistics.
type Connections struct {
	// The current number of active client connections including Waiting connections.
	Active int
	// The total number of accepted client connections.
	Accepted int
	// The total number of handled connections.
	Handled int
	// The total number of client requests.
	Requests int
	// The current number of connections where nginx is reading the request header.
	Reading int
	// The current number of connections where nginx is writing the response back to the client.
	Writing int
	// The current number of idle client connections waiting for a request.
	Waiting int
}

// ParseStats takes a reader and returns server statistics (connections).
func ParseStats(r io.Reader) (Connections, error) {
	var active, accepted, handled, requests, reading, writing, waiting int
	_, err := fmt.Fscanf(r, parseFormat, &active, &accepted, &handled, &requests, &reading, &writing, &waiting)
	if err != nil {
		return Connections{}, err
	}
	return Connections{
		Active:   active,
		Accepted: accepted,
		Handled:  handled,
		Requests: requests,
		Reading:  reading,
		Writing:  writing,
		Waiting:  waiting,
	}, nil
}

func runCLI(wr, ew io.Writer) int {
	return 0
}

func Main() int {
	return runCLI(os.Stdout, os.Stderr)
}
