package gtoolkit

import (
	"errors"
	"strings"
	"sync"
)

// ErrBuf accumulates and stores the errors in an internal buffer
type ErrBuf struct {
        builder *strings.Builder
	lock sync.RWMutex
}

// NewErrBuf returns an initialized ErrBuf instance
func NewErrBuf() *ErrBuf {
        return &ErrBuf{
                builder: new(strings.Builder),
        }
}

func (b *ErrBuf) Append(e error) {
        if e != nil {
		b.lock.Lock()
		defer b.lock.Unlock()
                b.builder.WriteString(e.Error() + "\n")
        }
}

func (b *ErrBuf) AppendString(s string) {
	if len(s) > 0 {
		b.lock.Lock()
		defer b.lock.Unlock()
		b.builder.WriteString(s + "\n")
	}
}

func (b *ErrBuf) ToError() error {
        if b.builder.Len() == 0 {
                return nil
        }

        return errors.New(b.builder.String())
}

func (b *ErrBuf) ToString() string {
        return b.builder.String()
}
