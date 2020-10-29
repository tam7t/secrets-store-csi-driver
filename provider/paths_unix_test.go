/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import "testing"

func TestJoinPaths(t *testing.T) {
	base := "/foo/bar"
	cases := []struct {
		in   string
		want string
	}{
		{
			in:   "baz",
			want: "/foo/bar/baz",
		},
		{
			in:   "baz.txt",
			want: "/foo/bar/baz.txt",
		},
		{
			in:   "a:b",
			want: "/foo/bar/a:b",
		},
	}

	for _, tc := range cases {
		got, err := JoinPaths(base, tc.in)
		if err != nil {
			t.Errorf("JoinPaths(%q) failed: %s", tc.in, err)
		}
		if got != tc.want {
			t.Errorf("JoinPaths(%v) = %v, want %v", tc.in, got, tc.want)
		}
	}
}

func TestJoinPathsError(t *testing.T) {
	base := "/foo/bar"
	cases := []string{
		"baz/baz",
		"baz/baz.txt",
		"../foo/bar",
		"../../foo",
		"/foo",
		"./foo",
		"/./foo",
	}

	for _, tc := range cases {
		_, err := JoinPaths(base, tc)
		if err == nil {
			t.Errorf("JoinPaths(%q) succeeded for malformed input, want error", tc)
		}
	}
}
