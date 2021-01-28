package liblru

import (
	"errors"
)

var (
	ErrKeyNotExist = errors.New("query key not exist")
	ErrorValidationForLRUSize = errors.New("size must great than zero")
)

type Node struct {
	key string
	value interface{}
	pre  *Node
	next *Node
}

type LinkListLRU struct{
	hash map[string]*Node
	list *DoubleLinkList
	max  int
}

func NewLinklistLRU(max int) (LRU,error){
	if max <=0 {
		return nil, ErrorValidationForLRUSize
	}
	return &LinkListLRU{
		hash:    make(map[string]*Node),
		list:    NewDoubleLinkList(),
		max:     max,
	}, nil
}

func (lru *LinkListLRU) Get(key string)(interface{}, error){
	node ,exist := lru.hash[key]
	if !exist {
		return nil,  ErrKeyNotExist
	}
	lru.list.Remove(node)
	lru.list.Push(node)
	return node.value, nil
}

func (lru *LinkListLRU) Set(key string, value interface{}){
	node, exist := lru.hash[key]
	if exist{
		lru.list.Remove(node)
	}else{
		if lru.list.Size() == lru.max{
			node, _ := lru.list.Pop()
			delete(lru.hash, node.key)
		}
		node = &Node{
			key:   key,
			value: value,
		}
		lru.hash[key] = node
	}
	node.value = value
	lru.list.Push(node)
}

type LRU interface {
	Set(key string, value interface{})
    Get(key string)(interface{}, error)
}

