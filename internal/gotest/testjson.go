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

package gotest

import "time"

type Action string

const (
	// ActionRun means the test has started running.
	ActionRun Action = "run"
	// ActionPause means the test has been paused.
	ActionPause Action = "pause"
	// ActionCont means the test has continued running.
	ActionCont Action = "cont"
	// ActionPass means the test passed.
	ActionPass Action = "pass"
	// ActionBench means the benchmark printed log output but did not fail.
	ActionBench Action = "bench"
	// ActionFail means the test or benchmark failed.
	ActionFail Action = "fail"
	// ActionOutput means the test printed output.
	ActionOutput Action = "output"
	// ActionSkip means the test was skipped or the package contained no tests.
	ActionSkip Action = "skip"
)

type TestEvent struct {
	Time           time.Time `json:"Time"`
	Action         Action    `json:"Action"`
	Package        string    `json:"Package"`
	Test           string    `json:"Test"`
	Output         string    `json:"Output"`
	ElapsedSeconds float64   `json:"Elapsed"`
}

func (te TestEvent) Elapsed() time.Duration {
	return time.Duration(te.ElapsedSeconds * float64(time.Second))
}