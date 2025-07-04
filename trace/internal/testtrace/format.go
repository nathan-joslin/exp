// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by "gen.bash" from internal/trace; DO NOT EDIT.

//go:build go1.23

package testtrace

import (
	"bytes"
	"fmt"
	"io"

	"github.com/nathan-joslin/exp/trace/internal/raw"
	"github.com/nathan-joslin/exp/trace/internal/version"
	"golang.org/x/tools/txtar"
)

// ParseFile parses a test file generated by the testgen package.
func ParseFile(testPath string) (io.Reader, version.Version, *Expectation, error) {
	ar, err := txtar.ParseFile(testPath)
	if err != nil {
		return nil, 0, nil, fmt.Errorf("failed to read test file for %s: %v", testPath, err)
	}
	if len(ar.Files) != 2 {
		return nil, 0, nil, fmt.Errorf("malformed test %s: wrong number of files", testPath)
	}
	if ar.Files[0].Name != "expect" {
		return nil, 0, nil, fmt.Errorf("malformed test %s: bad filename %s", testPath, ar.Files[0].Name)
	}
	if ar.Files[1].Name != "trace" {
		return nil, 0, nil, fmt.Errorf("malformed test %s: bad filename %s", testPath, ar.Files[1].Name)
	}
	tr, err := raw.NewTextReader(bytes.NewReader(ar.Files[1].Data))
	if err != nil {
		return nil, 0, nil, fmt.Errorf("malformed test %s: bad trace file: %v", testPath, err)
	}
	var buf bytes.Buffer
	tw, err := raw.NewWriter(&buf, tr.Version())
	if err != nil {
		return nil, 0, nil, fmt.Errorf("failed to create trace byte writer: %v", err)
	}
	for {
		ev, err := tr.ReadEvent()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, 0, nil, fmt.Errorf("malformed test %s: bad trace file: %v", testPath, err)
		}
		if err := tw.WriteEvent(ev); err != nil {
			return nil, 0, nil, fmt.Errorf("internal error during %s: failed to write trace bytes: %v", testPath, err)
		}
	}
	exp, err := ParseExpectation(ar.Files[0].Data)
	if err != nil {
		return nil, 0, nil, fmt.Errorf("internal error during %s: failed to parse expectation %q: %v", testPath, string(ar.Files[0].Data), err)
	}
	return &buf, tr.Version(), exp, nil
}
