package keycache

/*
	keycache keeps providers in memory and check sing before reading
*/

import (
	"bytes"
	"errors"
	"sync"
	"sync/atomic"
)

type IItem interface {
	ID() int
	Sign() []byte

	Decode([]byte) error
	Encode() ([]byte, error)
}

type IKeyCache interface {
	Count() int
	Get(item IItem) error
	Set(item IItem) error
	Del(item IItem)
	Exists(item IItem) bool
}

type KeyCache struct {
	sync.RWMutex
	data    map[int][]byte
	maxSize int
	counter *int32
	checker []int
}

func New(maxSize ...int) IKeyCache {
	out := &KeyCache{
		data: map[int][]byte{},
	}

	if len(maxSize) > 0 && maxSize[0] > 1 {
		out.checker = make([]int, maxSize[0], maxSize[0])
		out.counter = new(int32)
		out.maxSize = maxSize[0]
	}

	return out
}

func CheckSign(b, sign []byte) bool {
	signLength := int(b[0])
	if len(sign) != signLength {
		return false
	}

	for i := 0; i < signLength; i++ {
		if sign[i] != b[i+1] {
			return false
		}
	}

	return true
}

func Encode(item IItem) ([]byte, error) {
	sign := item.Sign()
	if len(sign) == 0 {
		return nil, errors.New("empty result of Sign()")
	}

	l := len(sign)

	// 255 is max uint8
	if l > 255 {
		return nil, errors.New("sign should be shorter than 256")
	}

	data, err := item.Encode()
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(append([]byte{uint8(l)}, sign...))
	buf.Write(data)

	return buf.Bytes(), nil
}

func (cache *KeyCache) Count() int {
	cache.RLock()
	defer cache.RUnlock()

	return len(cache.data)
}

func (cache *KeyCache) replace(id int) {
	next := int(atomic.AddInt32(cache.counter, 1)) % cache.maxSize

	cache.Lock()
	delId := cache.checker[next]
	cache.checker[next] = id
	if _, find := cache.data[delId]; find {
		delete(cache.data, delId)
	}
	cache.Unlock()
}

func (cache *KeyCache) Set(item IItem) error {
	if cache.Exists(item) {
		return nil
	}

	b, err := Encode(item)
	if err == nil {
		cache.Lock()
		cache.data[item.ID()] = b
		cache.Unlock()
		if cache.maxSize > 2 {
			go cache.replace(item.ID())
		}
	}

	return err
}

func (cache *KeyCache) Del(item IItem) {
	cache.del(item.ID())
}

func (cache *KeyCache) del(id int) {
	cache.Lock()
	defer cache.Unlock()

	if _, find := cache.data[id]; find {
		delete(cache.data, id)
	}
}

func (cache *KeyCache) Get(item IItem) error {
	cache.RLock()
	data, find := cache.data[item.ID()]
	cache.RUnlock()

	if !find {
		return nil
	}

	if !CheckSign(data, item.Sign()) {
		return nil
	}

	if len(data) <= int(data[0])+1 {
		return errors.New("data is empty")
	}

	return item.Decode(data[int(data[0])+1:])
}

func (cache *KeyCache) Exists(item IItem) bool {
	cache.RLock()
	data, find := cache.data[item.ID()]
	cache.RUnlock()

	return find && CheckSign(data, item.Sign())
}
