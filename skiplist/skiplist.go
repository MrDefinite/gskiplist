package gskiplist

import (
	"math/rand"
	"strconv"
	log "github.com/sirupsen/logrus"
)

const maxNumberOfLevels = 32
const p = 0.25

type SkipListLevel struct {
	forward *SkipListNode
	span    uint32
}

type SkipListNode struct {
	Backward *SkipListNode
	score    float64
	levels   []SkipListLevel
	obj      string
}

type SkipList struct {
	head, tail *SkipListNode
	level      int
	length     uint32
}

type RangeSpec struct {
	min, max     float64
	minex, maxex int32 /* are min or max exclusive? */
}

func CreateSkipList() *SkipList {
	// Creat a dummy node first
	dNode := SkipListNode{nil, 0, make([]SkipListLevel, maxNumberOfLevels), ""}
	for i := 0; i < maxNumberOfLevels; i++ {
		dNode.levels[i].forward = nil
		dNode.levels[i].span = 0
	}
	sl := SkipList{&dNode, nil, 1, 0}
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
	var rank [maxNumberOfLevels]uint32

	p := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		// Store rank that is crossed to reach the insert position
		if i == (sl.level - 1) {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}

		currentLevel := p.levels[i]
		forwardNode := currentLevel.forward
		for forwardNode != nil && (forwardNode.score < score ||
			(forwardNode.score == score && forwardNode.obj < ele)) {
			p = p.levels[i].forward
		}
		update[i] = p
	}
	// p is the place to insert new node, if it is not a duplicated one
	if p.score == score && p.obj == ele {
		return p
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

	newNode := SkipListNode{nil, score, []SkipListLevel{}, ele}
	for i := 0; i < levelForNewNode; i++ {
		newNode.levels[i].forward = update[i].levels[i].forward
		update[i].levels[i].forward = &newNode

		newNode.levels[i].span = update[i].levels[i].span - (rank[0] - rank[i])
		update[i].levels[i].span = rank[0] - rank[i] + 1
	}

	// Increment span for untouched levels
	for i := levelForNewNode; i < sl.level; i++ {
		update[i].levels[i].span++
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

	var factor float64 = p * 0xFFFF
	for (rand.Int31n(0X7FFFFFFF) & 0XFFFF) < int32(factor) {
		log.Debugln("Random counted, set level to: " + strconv.Itoa(level))
		level++
	}

	if level < maxNumberOfLevels {
		return level
	}
	return maxNumberOfLevels
}

func searchInsertPos(sl *SkipList, score float64, ele string) *SkipListNode {
	return nil
}

func printSkipList(sl *SkipList) {

}
