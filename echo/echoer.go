package echo

type Echoer interface {
	Echo(in string) string
}
