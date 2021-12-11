// Copyright 2021 FerretDB Inc.
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

package internal

import (
	"sort"
	"strings"
)

type TestsConfig struct {
	Pass []string `yaml:"pass"`
}

type CompareResult struct { //nolint:govet // we care about the fields order more than about alignment
	UnexpectedFail map[string]TestResult
	ExpectedPass   []string
	Fail           map[string]TestResult
	Rest           map[string]TestResult
}

func (c *TestsConfig) Compare(results *Results) *CompareResult {
	compareResult := CompareResult{
		UnexpectedFail: make(map[string]TestResult),
		Fail:           make(map[string]TestResult),
		Rest:           make(map[string]TestResult),
	}

	for test, testRes := range results.TestResults {
		var found bool
		for _, pass := range c.Pass {
			if strings.HasPrefix(test, pass) {
				switch testRes.Result {
				case Pass:
					compareResult.ExpectedPass = append(compareResult.ExpectedPass, test)
				case Unknown:
					fallthrough
				case Fail:
					fallthrough
				case Skip:
					compareResult.UnexpectedFail[test] = testRes
				}

				found = true
				break
			}
		}

		if found {
			continue
		}

		if testRes.Result == Fail {
			compareResult.Fail[test] = testRes
			continue
		}

		compareResult.Rest[test] = testRes
	}

	sort.Strings(compareResult.ExpectedPass)

	return &compareResult
}