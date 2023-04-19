package rotator

import (
	"errors"
	"fmt"
	"sync"
)

type Rotator[T any, C comparable] struct {
	sync.Mutex
	current int
	keyFunc func(T) C
	values  []T
}

func New[T any, C comparable](keyFunc func(T) C) *Rotator[T, C] {
	return &Rotator[T, C]{
		current: 0,
		values:  make([]T, 0),
		keyFunc: keyFunc,
	}
}

func (k *Rotator[T, C]) Add(v ...T) error {
	k.Lock()
	defer k.Unlock()

	k.values = append(k.values, v...)
	return nil
}

func (k *Rotator[T, C]) Rotate() (T, error) {
	var v T

	k.Lock()
	defer k.Unlock()

	if len(k.values) == 0 {
		return v, errors.New("no values to rotate")
	}

	v = k.values[k.current]
	k.current = (k.current + 1) % len(k.values)

	return v, nil
}

func (k *Rotator[T, C]) Get(key C) (T, bool) {
	var v T

	k.Lock()
	defer k.Unlock()

	for _, v := range k.values {
		if k.keyFunc(v) == key {
			return v, true
		}
	}
	return v, false
}

func (k *Rotator[T, C]) Delete(key C) error {

	for i, v := range k.values {
		if k.keyFunc(v) == key {

			k.Lock()
			k.values = append(k.values[:i], k.values[i+1:]...)
			k.Unlock()

			return nil
		}
	}

	return fmt.Errorf("key not found")
}

func (k *Rotator[T, C]) Len() int {
	return len(k.values)
}
