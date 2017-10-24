package gskiplist

import (
	"math/rand"
	//"strconv"
	//log "github.com/sirupsen/logrus"
	"fmt"
	"time"
)

const maxNumberOfLevels = 32
const p = 0.25

type SkipListLevel struct {
	forward *SkipListNode
	span    int
}

type SkipListNode struct {
	Backward *SkipListNode
	Score    float64
	levels   []SkipListLevel
	Obj      string
}

type SkipList struct {
	head, tail *SkipListNode
	level      int
	length     int
}

type RangeSpec struct {
	min, max     float64
	minex, maxex int32 /* are min or max exclusive? */
}

func CreateSkipList() *SkipList {
	// Create a dummy node first
	header := SkipListNode{nil, -1, make([]SkipListLevel, maxNumberOfLevels), ""}
	for i := 0; i < maxNumberOfLevels; i++ {
		header.levels[i].forward = nil
		header.levels[i].span = 0
	}
	header.Backward = nil
	sl := SkipList{&header, nil, 1, 0}
	return &sl
}

func FreeSkipList(sl *SkipList) {
	sl.head = nil
	sl.tail = nil
	sl.level = 0
	sl.length = 0
	sl = nil
}

func Insert(sl *SkipList, score float64, ele string) *SkipListNode {
	// Cache the node which need to be updated
	var update [maxNumberOfLevels]*SkipListNode
	var rank [maxNumberOfLevels]int

	p := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		// Store rank that is crossed to reach the insert position
		if i == (sl.level - 1) {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}

		for p.levels[i].forward != nil && (p.levels[i].forward.Score < score ||
			(p.levels[i].forward.Score == score && p.levels[i].forward.Obj < ele)) {
			// How long is the span from the first node to the last in the same layer i
			rank[i] += p.levels[i].span
			p = p.levels[i].forward
		}
		// Record the node which should connect its layer i to the new node
		update[i] = p
	}

	// Insert node now
	levelForNewNode := getRandomLevel()
	if levelForNewNode > sl.level {
		for i := sl.level; i < levelForNewNode; i++ {
			rank[i] = 0
			update[i] = sl.head
			update[i].levels[i].span = sl.length
		}
		sl.level = levelForNewNode
	}

	newNode := SkipListNode{nil, score, make([]SkipListLevel, maxNumberOfLevels), ele}
	for i := 0; i < levelForNewNode; i++ {
		newNode.levels[i].forward = update[i].levels[i].forward
		update[i].levels[i].forward = &newNode

		newNode.levels[i].span = update[i].levels[i].span - (rank[0] - rank[i])
		update[i].levels[i].span = rank[0] - rank[i] + 1
	}

	// Increment span for untouched levels
	for i := levelForNewNode; i < sl.level; i++ {
		update[i].levels[i].span += 1
	}

	if update[0] == sl.head {
		newNode.Backward = nil
	} else {
		newNode.Backward = update[0]
	}
	if newNode.levels[0].forward != nil {
		newNode.levels[0].forward.Backward = &newNode
	} else {
		sl.tail = &newNode
	}
	sl.length += 1

	return &newNode
}

func Delete(sl *SkipList, score float64, ele string, node *SkipListNode) bool {
	return true
}

func GetRank(sl *SkipList, score float64, ele string) float64 {
	return 0
}

func FirstInRange(sl *SkipList, rangeSpec *RangeSpec) *SkipListNode {
	return nil
}

func LastInRange(sl *SkipList, rangeSpec *RangeSpec) *SkipListNode {
	return nil
}

// Returns a random level for the new skiplist node we are going to create.
// The return value of this function is between 1 and maxNumberOfLevels
// (both inclusive), with a powerlaw-alike distribution where higher
// levels are less likely to be returned.
// NOTE: the algorithm is from redis
func getRandomLevel() int {
	level := 1
	rand.Seed(time.Now().UnixNano())

	var factor = p * 0xFFFF
	for rand.Int31n(2147483647)&0XFFFF < int32(factor) {
		fmt.Printf("Random counted, set level to: %d\n", level)
		level++
	}

	if level < maxNumberOfLevels {
		fmt.Printf("Get final level for new node: %d\n", level)
		return level
	}
	return maxNumberOfLevels
}

func searchInsertPos(sl *SkipList, score float64, ele string) *SkipListNode {
	return nil
}

// TODO: the last node cannot be printed correctly
func PrintSkipList(sl *SkipList) {
	p := sl.head
	isDummyNode := true
	for p != nil {
		if isDummyNode {
			p = p.levels[0].forward
			isDummyNode = false
			continue
		}

		levels := p.levels
		isLastNode := levels[0].forward == nil
		fmt.Printf("Node { ")
		for j := 0; j < maxNumberOfLevels; j++ {
			level := levels[j]
			if (isLastNode && p.Obj != "") || level.forward != nil {
				fmt.Printf("%g %s, ", p.Score, p.Obj)
			} else {
				fmt.Printf("nil, ")
			}
		}
		fmt.Printf("}\n")

		p = p.levels[0].forward
	}
}
