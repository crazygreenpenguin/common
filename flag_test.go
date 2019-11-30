package common

import "testing"

func TestNewFlag(t *testing.T) {
	f := NewFlag(true)
	if f.GetFlag() != true {
		t.Fail()
	}
	f = NewFlag(false)
	if f.GetFlag() != false {
		t.Fail()
	}
}

func BenchmarkNewFlag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewFlag(true)
	}
}

func TestFlag_SetFlag(t *testing.T) {
	f := NewFlag(false)
	if err := f.SetFlag(); err != nil {
		t.Log(err)
		t.Fail()
	}
	if err := f.SetFlag(); err != ErrAlreadySet {
		t.Log(err)
		t.Fail()
	}
}

func BenchmarkFlag_SetFlag(b *testing.B) {
	f := NewFlag(false)
	if err := f.SetFlag(); err != nil {
		b.Log(err)
		b.Fail()
	}
	for i := 0; i < b.N; i++ {
		if err := f.SetFlag(); err != ErrAlreadySet {
			b.Log(err)
			b.Fail()
		}
	}
}
