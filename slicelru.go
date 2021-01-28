package liblru

type SliceLRU struct{
	hash    map[string]*Node
	data    []*Node
	max     int
}

func NewSliceLRU(max int) (LRU, error){
	if max<=0 {
		return nil, ErrorValidationForLRUSize
	}
	return &SliceLRU{
		hash    : make(map[string]*Node),
		data    : make([]*Node, 0),
		max     : max,
	}, nil
}

func (cache *SliceLRU) Set(key string, value interface{} ){
	if node, exist := cache.hash[key]; exist{
		position := cache.find(node)
		cache.data = append(cache.data[:position], cache.data[position+1:]...)
		delete(cache.hash, node.key)
	}else{
		/// 判断是否超过最大值
		if len(cache.data) == cache.max {
			/// 删除头部的元素 更新计数值
			head := cache.data[0]
			delete(cache.hash, head.key)
			cache.data = cache.data[1:]
		}
	}
	/// 插入新的元素，并更新hash表
	node := &Node{key:key, value:value}
	cache.hash[key] = node
	cache.data = append(cache.data, node)
}

func (cache *SliceLRU) Get(key string) (interface{}, error){
	node, exist := cache.hash[key]
	if !exist{
		return nil, ErrKeyNotExist
	}
	/// 更新位置 追加到slice尾部
	position := cache.find(node)
	cache.data = append(cache.data[:position], cache.data[position+1:]...)
	cache.data = append(cache.data, node)

	return node.value, nil
}

func (cache *SliceLRU) find(target *Node) int{
	for index, node := range cache.data{
		if node.key == target.key{
			return index
		}
	}
	return -1
}
