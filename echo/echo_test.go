package echo_test

import (
	"testing"

	"github.com/nmiyake/echgo/echo"
)

func TestEcho(t *testing.T) {
	echoer := echo.NewEchoer()
	for i, tc := range []struct {
		in   string
		want string
	}{
		{"foo", "foo"},
		{"foo bar", "foo bar"},
	} {
		if got := echoer.Echo(tc.in); got != tc.want {
			t.Errorf("case %d failed: want %q, got %q", i, tc.want, got)
		}
	}
}
