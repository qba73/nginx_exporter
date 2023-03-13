package main

import (
	"os"

	nginx "github.com/qba73/nginx_exporter"
)

func main() {
	os.Exit(nginx.Main())
}
