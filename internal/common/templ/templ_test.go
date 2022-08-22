// Copyright 2022 The Hugoreleaser Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package templ

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestSprintt(t *testing.T) {
	c := qt.New(t)

	c.Assert(MustSprintt("{{ . }}", "foo"), qt.Equals, "foo")
	c.Assert(MustSprintt("{{ . | upper }}", "foo"), qt.Equals, "FOO")
	c.Assert(MustSprintt("{{ . | lower }}", "FoO"), qt.Equals, "foo")
	c.Assert(MustSprintt("{{ . | trimPrefix `v` }}", "v3.0.0"), qt.Equals, "3.0.0")
	c.Assert(MustSprintt("{{ . | trimSuffix `-beta` }}", "v3.0.0-beta"), qt.Equals, "v3.0.0")
}
