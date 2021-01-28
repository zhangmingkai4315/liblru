package liblru

import (
	"fmt"
	"testing"
)

func TestNewDoubleLinkList(t *testing.T) {
	dll := NewDoubleLinkList()
	for i:=0; i<100;i++{
		key := fmt.Sprintf("key_%d",i)
		dll.Push(&Node{key:key, value:i})
	}
	if dll.Size() != 100{
		t.Error("double link list size should be 100")
		return
	}
	for i:= 0;i< 100;i++{
		item, err := dll.Pop()
		if err != nil{
			t.Error("item pop should success")
			return
		}
		if item.value != i {
			t.Errorf("pop value is %d != %d", item.value, i)
		}
	}
	v, err := dll.Pop()
	if err == nil {
		t.Errorf("pop from empty list should return error but got no err and got : %v", v)
	}
}


func TestDoubleLinkList_Remove(t *testing.T) {
	dll := NewDoubleLinkList()
	hash := make(map[string]*Node)
	for i:=0; i<100;i++{
		key := fmt.Sprintf("key_%d",i)
		node := &Node{key:key, value:i}
		hash[key] = node
		dll.Push(node)
	}
	if dll.Size() != 100{
		t.Error("double linklist should have 100 items")
	}
	for _, v := range hash{
		dll.Remove(v)
	}
	if dll.Size() != 0{
		t.Error("double linklist should remove all items and return size=0")
	}

}