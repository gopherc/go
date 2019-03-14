// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build js,wasm,goc

package rand

func init() {
	Reader = &reader{}
}

type reader struct{}

func getRandomValues(b []byte)

func (r *reader) Read(b []byte) (int, error) {
	getRandomValues(b)
	return len(b), nil
}
