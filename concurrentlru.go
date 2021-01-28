package liblru

import "sync"

type ConcurrentSafeLRU struct {
	sync.Mutex
	lru LRU
}

func NewConcurrentSafeLRU(max int)(LRU, error){
	lru, err := NewLinklistLRU(max)
	if err != nil{
		return nil, err
	}
	return &ConcurrentSafeLRU{
		lru:  lru,
	}, nil
}

func (lru *ConcurrentSafeLRU) Get(key string)(interface{}, error){
	lru.Lock()
	defer lru.Unlock()
	return lru.lru.Get(key)
}

func (lru *ConcurrentSafeLRU) Set(key string, value interface{}){
	lru.Lock()
	defer lru.Unlock()
	lru.lru.Set(key, value)
}