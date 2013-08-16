package echo

// #cgo CFLAGS: -DSALT_OPTION
// #include "echo.h"
import "C"
import (
    "hash"
    "errors"
)

const SALT_SIZE = 16

type echo struct {
    hashSize int
    hashState C.hashState
}

func New(hashSize int) hash.Hash {
    e := new(echo)
    e.hashSize = hashSize
    e.Reset()
    return e
}

func New128() hash.Hash { return New(128) }
func New160() hash.Hash { return New(160) }
func New192() hash.Hash { return New(192) }
func New224() hash.Hash { return New(224) }
func New256() hash.Hash { return New(256) }
func New384() hash.Hash { return New(384) }
func New512() hash.Hash { return New(512) }

func (e *echo) Reset() {
    if C.Init(&e.hashState, C.int(e.hashSize)) == C.BAD_HASHBITLEN {
        panic("Bad hashSize, must be 128 to 512 bits, inclusive")
    }
}

func (e echo) BlockSize() int {
    return 1
}


func (e echo) Size() int {
    return e.hashSize / 8
}

func (e *echo) Write(data []byte) (nn int, err error) {
    nn = len(data)
    if nn == 0 {
        return
    }
    C.Update(&e.hashState, (*C.BitSequence)(&data[0]), C.DataLength(nn*8))
    return
}


func (e echo) Sum(in []byte) []byte {
    result := make([]byte, e.hashSize/8)
    e0 := e
    C.Final(&e0.hashState, (*C.BitSequence)(&result[0]))
    return append(in, result...)
}

// Optional: May be called after Reset or NewXXX *but before either Write or Sum*
func (e *echo) SetSalt(salt []byte) (err error) {
    if SALT_SIZE != len(salt) {
        err = errors.New("Invalid salt")
        return
    }
    C.SetSalt(&e.hashState, (*C.BitSequence)(&salt[0]))
    return
}
