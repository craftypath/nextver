// Copyright The nextver Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package version

import (
	"testing"
)

type testcase = struct {
	current          string
	incrementPattern string
}

func FuzzNext(f *testing.F) {
	testcases := []testcase{
		{
			current:          " ",
			incrementPattern: " ",
		}, {
			current:          "0.2.4",
			incrementPattern: "x.0.0",
		}, {
			current:          "0.0.0",
			incrementPattern: "0.0.x-alpha.x",
		}, {
			current:          "0.0.0-alpha+build",
			incrementPattern: "0.0.x-alpha+build",
		}, {
			current:          "1.5.2",
			incrementPattern: "?.x.0",
		}, {
			current:          "1.0.0",
			incrementPattern: "v1.x.0",
		}, {
			current:          "v1.0.0",
			incrementPattern: "1.x.0",
		},
	}
	for _, tc := range testcases {
		f.Add(tc.current, tc.incrementPattern)
	}
	f.Fuzz(func(t *testing.T, current string, incrementPattern string) {
		nextVersion, err := Next(current, incrementPattern)
		if err != nil {
			return
		}
		_, err = Next(nextVersion, incrementPattern)
		if err != nil {
			t.Errorf("Result of Next(...) with current %v and incrementPattern %v was not a valid SemVer although no error was returned: %v", current, incrementPattern, nextVersion)
		}
	})
}
