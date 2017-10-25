package gskiplist

import (
	"testing"
	"strconv"
)


func TestCreateSkipList(t *testing.T) {
	sl := CreateSkipList()
	if sl.length != 0 {
		t.Errorf("length init failed")
	}
	if sl.level != 1 {
		t.Errorf("level init failed")
	}
	if sl.tail != nil {
		t.Errorf("tail init failed")
	}
	if sl.head == nil {
		t.Errorf("head init failed")
	}

	head := sl.head
	if head.Backward != nil {
		t.Errorf("Header node should not contain backward pointer")
	}
	if head.Score != -1 {
		t.Errorf("Header node should not set score")
	}
	if head.Obj != "" {
		t.Errorf("Header node should not set obj")
	}
	if len(head.levels) != maxNumberOfLevels {
		t.Errorf("Header node should have initialized level to: %d", maxNumberOfLevels)
	}
}

func TestInsert(t *testing.T) {
	sl := CreateSkipList()
	for i := 0; i < 10; i++ {
		Insert(sl, float64(i + 1), "test" + strconv.Itoa(i + 1))
	}

	if sl.length != 10 {
		t.Errorf("Failed to insert 10 nodes to skip list")
	}

	head := sl.head
	if head.Backward != nil {
		t.Errorf("Header node should not contain backward pointer")
	}
	tail := sl.tail
	if tail == nil {
		t.Errorf("Skip list should contain tail node")
	}
}



