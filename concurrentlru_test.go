package liblru

import (
	"sync"
	"testing"
)

func TestConcurrentSafeLRU(t *testing.T) {
	lru, _ := NewConcurrentSafeLRU(100)
	lru.Set("test", 1)
	var wg sync.WaitGroup
	for i:=0;i<1000;i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			v, err := lru.Get("test")
			if err != nil{
				t.Errorf("should return err = nil")
				return
			}
			if v != 1 {
				t.Errorf("should return 1 but got %v", v)
			}

		}()
	}
	wg.Wait()
}

