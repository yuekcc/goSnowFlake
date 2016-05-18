package goSnowFlake

import (
	"fmt"
	"testing"
)

func TestSnowFlake(t *testing.T) {
	fmt.Println("start generate")
	iw, _ := NewIdWorker(2)
	for i := 0; i < 100; i++ {
		id, err := iw.NextId()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(id)
		}
	}
	fmt.Println("end generate")
}

func TestSnowFlakeUni(t *testing.T) {
	result := make(map[int64]bool)
	max := 100

	iw, _ := NewIdWorker(3)
	for i := 0; i < max; i++ {
		id, _ := iw.NextId()
		result[id] = true
	}

	if len(result) != max {
		t.Fatalf("map result size should be %v", max)
	}
}
