package common

import (
	"errors"
	"sync/atomic"
)

var (
	ErrAlreadySet   = errors.New("flag already set")
	ErrAlreadyUnset = errors.New("flag already unset")
)

func NewFlag(val bool) Flag {
	if val {
		return Flag(100)
	}
	return Flag(0)
}

type Flag uint64

func (f *Flag) SetFlag() error {
	if atomic.CompareAndSwapUint64((*uint64)(f), 0, 100) {
		return nil
	}
	return ErrAlreadySet
}

func (f *Flag) UnsetFlag() error {
	if atomic.CompareAndSwapUint64((*uint64)(f), 100, 0) {
		return nil
	}
	return ErrAlreadyUnset
}

func (f *Flag) GetFlag() bool {
	if atomic.CompareAndSwapUint64((*uint64)(f), 100, 100) {
		return true
	}
	return false
}
