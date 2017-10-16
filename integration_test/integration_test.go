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

package integration_test

import (
	"os/exec"
	"testing"

	"github.com/palantir/godel/pkg/products"
)

func TestInvalidType(t *testing.T) {
	echgoPath, err := products.Bin("echgo")
	if err != nil {
		panic(err)
	}
	cmd := exec.Command(echgoPath, "-type", "invalid", "foo")
	output, err := cmd.CombinedOutput()
	gotOutput := string(output)
	if err == nil {
		t.Errorf("expected command %v to fail. Output: %s", cmd.Args, gotOutput)
	}
	wantOutput := "invalid echo type: invalid\n"
	if wantOutput != gotOutput {
		t.Errorf("invalid output: want %q, got %q", wantOutput, gotOutput)
	}
	wantErr := "exit status 1"
	gotErr := err.Error()
	if wantErr != gotErr {
		t.Errorf("invalid error output: want %q, got %q", wantErr, gotErr)
	}
}
