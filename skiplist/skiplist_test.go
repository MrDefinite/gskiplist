package gskiplist

import (
	"testing"
	"fmt"
)


func TestCreateSkipList(t *testing.T) {
	slp := CreateSkipList()
	if slp.length != 0 {
		fmt.Errorf("length init failed")
	}
	if slp.level != 0 {
		fmt.Errorf("level init failed")
	}
	if slp.tail != nil {
		fmt.Errorf("tail init failed")
	}
	if slp.head != nil {
		fmt.Errorf("head init failed")
	}
}
