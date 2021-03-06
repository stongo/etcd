// Copyright 2015 CoreOS, Inc.
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

package testutil

import (
	"net/url"
	"runtime"
	"testing"
)

// WARNING: This is a hack.
// Remove this when we are able to block/check the status of the go-routines.
func ForceGosched() {
	// possibility enough to sched up to 10 go routines.
	for i := 0; i < 10000; i++ {
		runtime.Gosched()
	}
}

func MustNewURLs(t *testing.T, urls []string) []url.URL {
	if urls == nil {
		return nil
	}
	var us []url.URL
	for _, url := range urls {
		u := MustNewURL(t, url)
		us = append(us, *u)
	}
	return us
}

func MustNewURL(t *testing.T, s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		t.Fatalf("parse %v error: %v", s, err)
	}
	return u
}
