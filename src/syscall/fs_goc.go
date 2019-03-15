// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build js,wasm,goc

package syscall

import (
	"sync"
	_ "syscall/js" // Need init to call setEventHandler
	"unsafe"
)

var (
	openFilesMut  sync.Mutex
	handleCounter uint32 = 3
	openFiles            = map[int]int64{}
)

func openFile(path unsafe.Pointer, openmode, perm uint32) int64

func Open(path string, openmode int, perm uint32) (int, error) {
	openFilesMut.Lock()
	defer openFilesMut.Unlock()

	if handleCounter == 0 {
		panic("handle wraparound")
	}

	fp := openFile(unsafe.Pointer(&path), uint32(openmode), perm)
	if fp == 0 {
		//TODO: Should return a real error. /aj
		return 0, EIO
	}

	fd := int(handleCounter)
	openFiles[fd] = fp
	handleCounter++
	return int(fd), nil
}

func closeFile(fd uintptr) uint32

func Close(fd int) error {
	openFilesMut.Lock()
	defer openFilesMut.Unlock()

	if handle, ok := openFiles[fd]; ok {
		delete(openFiles, fd)
		closeFile(uintptr(handle))
	}
	return EIO
}

func CloseOnExec(fd int) {
	// nothing to do - no exec
}

func Mkdir(path string, perm uint32) error {
	return ENOSYS
}

func ReadDirent(fd int, buf []byte) (int, error) {
	return 0, ENOSYS
}

func Stat(path string, st *Stat_t) error {
	return ENOSYS
}

func Lstat(path string, st *Stat_t) error {
	return ENOSYS
}

func Fstat(fd int, st *Stat_t) error {
	return ENOSYS
}

func Unlink(path string) error {
	return ENOSYS
}

func Rmdir(path string) error {
	return ENOSYS
}

func Chmod(path string, mode uint32) error {
	return ENOSYS
}

func Fchmod(fd int, mode uint32) error {
	return ENOSYS
}

func Chown(path string, uid, gid int) error {
	return ENOSYS
}

func Fchown(fd int, uid, gid int) error {
	return ENOSYS
}

func Lchown(path string, uid, gid int) error {
	return ENOSYS
}

func UtimesNano(path string, ts []Timespec) error {
	return ENOSYS
}

func Rename(from, to string) error {
	return ENOSYS
}

func Truncate(path string, length int64) error {
	return ENOSYS
}

func Ftruncate(fd int, length int64) error {
	return ENOSYS
}

func Getcwd(buf []byte) (int, error) {
	return 0, ENOSYS
}

func Chdir(path string) error {
	return ENOSYS
}

func Fchdir(fd int) error {
	return ENOSYS
}

func Readlink(path string, buf []byte) (n int, err error) {
	return 0, ENOSYS
}

func Link(path, link string) error {
	return ENOSYS
}

func Symlink(path, link string) error {
	return ENOSYS
}

func Fsync(fd int) error {
	return ENOSYS
}

func readFile(fd uintptr, p unsafe.Pointer, n int32) int32

func Read(fd int, b []byte) (int, error) {
	ln := int32(len(b))
	if ln == 0 {
		return 0, nil
	}

	if n := readFile(uintptr(fd), unsafe.Pointer(&b[0]), ln); n != ln {
		return int(n), EIO
	}
	return int(ln), nil
}

func writeFile(fd uintptr, p unsafe.Pointer, n int32) int32

func Write(fd int, b []byte) (int, error) {
	ln := int32(len(b))
	if ln == 0 {
		return 0, nil
	}

	if n := writeFile(uintptr(fd), unsafe.Pointer(&b[0]), ln); n != ln {
		return int(n), EIO
	}
	return int(ln), nil
}

func Pread(fd int, b []byte, offset int64) (int, error) {
	return 0, ENOSYS
}

func Pwrite(fd int, b []byte, offset int64) (int, error) {
	return 0, ENOSYS
}

func Seek(fd int, offset int64, whence int) (int64, error) {
	return 0, ENOSYS
}

func Dup(fd int) (int, error) {
	return 0, ENOSYS
}

func Dup2(fd, newfd int) error {
	return ENOSYS
}

func Pipe(fd []int) error {
	return ENOSYS
}
