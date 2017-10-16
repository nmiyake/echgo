package echo

func NewEchoer() Echoer {
	return &simpleEchoer{}
}

type simpleEchoer struct{}

func (e *simpleEchoer) Echo(in string) string {
	return in
}
