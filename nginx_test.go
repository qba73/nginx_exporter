package nginx_test

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	nginx "github.com/qba73/nginx_exporter"
)

func TestParseStats_ParsesStatsOnValidInput(t *testing.T) {
	t.Parallel()

	got, err := nginx.ParseStats(strings.NewReader(valid_stats_data))
	if err != nil {
		t.Fatal(err)
	}
	want := nginx.Connections{
		Active:   291,
		Accepted: 16630948,
		Handled:  16630948,
		Requests: 31070465,
		Reading:  6,
		Writing:  179,
		Waiting:  106,
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

var (
	valid_stats_data = `Active connections: 291
server accepts handled requests
 16630948 16630948 31070465
Reading: 6 Writing: 179 Waiting: 106`
)
