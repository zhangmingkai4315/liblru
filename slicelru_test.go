package liblru

import (
	"reflect"
	"testing"
)

func TestSliceLRU(t *testing.T) {
	lruInterface,_ := NewSliceLRU(4)
	lru := lruInterface.(*SliceLRU)
	testCase := getTestCases()
	for _, test := range testCase{
		if test.operator == OperatorSet{
			lru.Set(test.data[0], test.data[1])
		}else {
			v, err := lru.Get(test.data[0])
			if err != nil {
				t.Errorf("should not return error")
				return
			}
			if v.(string) != test.data[1] {
				t.Errorf("should return %s but got %s", test.data[1], v.(string))
			}
		}
		result := lru.data
		if len(result) != len(test.respect){
			t.Errorf("result length is not equal respect length")
		}
		items := make([]string,0)
		for i:=len(result)-1; i>=0;i--{
			items = append(items, result[i].value.(string))
		}
		if !reflect.DeepEqual(items, test.respect){
			t.Errorf("%v != %v", items, test.respect)
		}
	}
}

