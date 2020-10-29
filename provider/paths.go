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

// Package provider includes helpers for implementing provider logic.
package provider

import (
	"fmt"
	"path/filepath"
)

// JoinPaths is a safer way to append a filename to a base path. This should be
// used when constructing subpaths of targetPath for writing files to avoid path
// traversal conditions.
func JoinPaths(base, filename string) (string, error) {
	part := filepath.Base(filename)
	if part != filename {
		// filename includes path separator
		return "", fmt.Errorf("invalid path: %q", filename)
	}
	return filepath.Join(base, filename), nil
}
