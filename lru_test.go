package liblru

import (
	"math"
	"reflect"
	"strconv"
	"fmt"
	"testing"
)
const (
	OperatorGet = "get"
	OperatorSet = "set"
)
type testCase struct{
	operator string
	data    []string
	respect []string
}
func getTestCases()[]*testCase{
	return []*testCase{
		{
			operator:OperatorSet,
			data: []string{"key01", "1"},
			respect: []string{"1"},
		},
		{
			operator:OperatorSet,
			data: []string{"key02", "2"},
			respect: []string{"2","1"},
		},
		{
			operator:OperatorSet,
			data: []string{"key03", "3"},
			respect: []string{"3","2","1"},
		},
		{
			operator:OperatorGet,
			data: []string{"key02","2",},
			respect: []string{"2","3","1"},
		},
		{
			operator:OperatorSet,
			data: []string{"key04", "4"},
			respect: []string{"4","2","3","1"},
		},
		{
			operator:OperatorSet,
			data: []string{"key05", "5"},
			respect: []string{"5","4","2","3"},
		},
		{
			operator:OperatorGet,
			data: []string{"key02", "2"},
			respect: []string{"2","5","4","3"},
		},
	}
}

var (
	benchFunc = []struct{
		builder func(max int)(LRU, error)
		name    string
	}{
		{
			builder: NewSliceLRU,
			name :   "slicelru",
		},
		{
			builder: NewLinklistLRU,
			name :   "linklistlru",
		},
		{
			builder: NewConcurrentSafeLRU,
			name :   "concurrentlru",
		},
	}
)

const MinPower = 4
const MaxPower = 16

func TestLinkListLRU(t *testing.T) {
	lru,_ := NewLinklistLRU(4)
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
		result := lru.(*LinkListLRU).list.toList()
		if len(result) != len(test.respect){
			t.Errorf("result length is not equal respect length")
		}
		items := make([]string,0)
		for _, item := range result{
			items = append(items, item.(string))
		}
		if !reflect.DeepEqual(items, test.respect){
			t.Errorf("%v != %v", items, test.respect)
		}
	}
}


//BenchmarkLinkListLRU_Set-6   	 5410965	       220 ns/op
//BenchmarkLinkListLRU_Get-6   	53544406	        21.7 ns/op
func BenchmarkLRU_Set(b *testing.B) {
	for _, lruBuilder := range benchFunc{
		for k := MinPower; k <= MaxPower; k++ {
			n := int(math.Pow(2, float64(k)))
			lru, _ := lruBuilder.builder(n)
			b.Run(fmt.Sprintf("%s-%d", lruBuilder.name, n), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					lru.Set(strconv.Itoa(i),i)
				}
			})
		}
	}

}
func BenchmarkLRU_GetInLRU(b *testing.B) {
	for _, lruBuilder := range benchFunc {
		for k := MinPower; k <= MaxPower; k++ {
			n := int(math.Pow(2, float64(k)))
			lru, _ := lruBuilder.builder(n)
			for i := 0; i < n; i++ {
				lru.Set(strconv.Itoa(i), i)
			}
			b.Run(fmt.Sprintf("%s-%d", lruBuilder.name, n), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					lru.Get(strconv.Itoa(i % n))
				}
			})
		}
	}
}

func BenchmarkLRU_GetNotInLRU(b *testing.B) {
	for _, lruBuilder := range benchFunc {
		for k := MinPower; k <= MaxPower; k++ {
			n := int(math.Pow(2, float64(k)))
			lru, _ := lruBuilder.builder(n)
			for i := 0; i < n; i++ {
				lru.Set(strconv.Itoa(i), i)
			}
			b.Run(fmt.Sprintf("%s-%d", lruBuilder.name, n), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					lru.Get(strconv.Itoa(i%n + n))
				}
			})
		}
	}
}
