package main

import (
	"../skiplist"
	_ "fmt"
	"strconv"
)

func main()  {
	sl := gskiplist.CreateSkipList()

	for i := 0; i < 10; i++ {
		gskiplist.Insert(sl, float64(i + 1), "test" + strconv.Itoa(i))
	}
	gskiplist.PrintSkipList(sl)
}


