// Copyright (c) 2017 Author Name
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate stringer -type=Type

package echo

import (
	"fmt"
	"strings"
)

type Type int

const (
	Simple Type = iota
	Reverse
	end
)

var echoers = []Echoer{
	Simple:  &simpleEchoer{},
	Reverse: &reverseEchoer{},
}

func NewEchoer(typ Type) Echoer {
	return echoers[typ]
}

func TypeFrom(typ string) (Type, error) {
	for curr := Simple; curr < end; curr++ {
		if strings.ToLower(typ) == strings.ToLower(curr.String()) {
			return curr, nil
		}
	}
	return end, fmt.Errorf("unrecognized type: %s", typ)
}

type simpleEchoer struct{}

func (e *simpleEchoer) Echo(in string) string {
	return in
}

type reverseEchoer struct{}

func (e *reverseEchoer) Echo(in string) string {
	out := make([]byte, len(in))
	for i := 0; i < len(out); i++ {
		out[i] = in[len(in)-1-i]
	}
	return string(out)
}
