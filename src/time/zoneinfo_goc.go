// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build js,wasm,goc

package time

import (
	"runtime"
)

var zoneSources = []string{
	"/usr/share/zoneinfo/",
	"/usr/share/lib/zoneinfo/",
	"/usr/lib/locale/TZ/",
	runtime.GOROOT() + "/lib/time/zoneinfo.zip",
}

func getTimezoneOffset() int

func initLocal() {
	localLoc.name = "Local"

	z := zone{}
	//TODO: Use and implement getTimezoneOffset. /aj
	offset := 0 //getTimezoneOffset()
	z.offset = offset * 60
	z.name = "UTC"
	if offset < 0 {
		z.name += "-"
		offset *= -1
	} else {
		z.name += "+"
	}
	z.name += itoa(offset / 60)
	min := offset % 60
	if min != 0 {
		z.name += ":" + itoa(min)
	}
	localLoc.zone = []zone{z}
}

// itoa is like strconv.Itoa but only works for values of i in range [0,99].
// It panics if i is out of range.
func itoa(i int) string {
	if i < 10 {
		return digits[i : i+1]
	}
	return smallsString[i*2 : i*2+2]
}

const smallsString = "00010203040506070809" +
	"10111213141516171819" +
	"20212223242526272829" +
	"30313233343536373839" +
	"40414243444546474849" +
	"50515253545556575859" +
	"60616263646566676869" +
	"70717273747576777879" +
	"80818283848586878889" +
	"90919293949596979899"
const digits = "0123456789"
