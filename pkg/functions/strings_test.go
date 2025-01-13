package functions_test

import (
	"testing"
	"os"
	"github.com/onsi/ginkgo/v2/reporters"
	"github.com/onsi/gomega/gexec"

	fn "github.com/kloudlite/api/pkg/functions"
)

var junitReporter *reporters.JUnitReporter
var session *gexec.Session

func TestReverse(t *testing.T) {
	type args struct {
		x string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "sample",
			args: args{x: "sample"},
			want: "elpmas",
		},
		{
			name: "sample2133",
			args: args{x: "sample2133"},
			want: "3312elpmas",
		},
		{
			name: "sample.2133435",
			args: args{x: "sample.2133435"},
			want: "5343312.elpmas",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.StringReverse(tt.args.x); got != tt.want {
				t.Errorf("StringReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
