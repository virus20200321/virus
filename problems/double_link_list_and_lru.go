package problems

import (
	log "github.com/sirupsen/logrus"

	"fmt"
)

type LruCache struct {
	Cap  int64
	List DoubleList
	Dict map[string]*DoubleListNode
}

func NewLruCache(cap int64) *LruCache {
	lru := &LruCache{
		Cap:  cap,
		List: DoubleList{},
		Dict: map[string]*DoubleListNode{},
	}
	return lru
}

func (lru *LruCache) Print() {
	fmt.Println("list")
	lru.List.Print()
	fmt.Println("cap", lru.Cap)
	fmt.Println("keys", lru.Dict)
	fmt.Println("length", len(lru.Dict), lru.List.Len())

}

func (lru *LruCache) Get(key string) (string, error) {
	if node, ok := lru.Dict[key]; !ok {
		return "", fmt.Errorf("not found")
	} else {
		lru.List.Delete(key)
		lru.List.PushFront(key, node.Value)
		return node.Value, nil
	}

}

func (lru *LruCache) Push(key, value string) {
	if n, ok := lru.Dict[key]; ok {
		lru.List.DeleteNode(n)
		lru.List.PushFront(key, value)
		lru.Dict[key] = lru.List.Head
	} else {
		if lru.Cap == int64(len(lru.Dict)) { //full
			//delete last
			pop := lru.List.Pop()
			log.Debugf("pop key %s", pop.Key)
			delete(lru.Dict, pop.Key)
		}

		lru.List.PushFront(key, value)
		lru.Dict[key] = lru.List.Head
	}

}

type DoubleListNode struct {
	Key   string
	Value string
	Prev  *DoubleListNode
	Next  *DoubleListNode
}
type DoubleList struct {
	Head *DoubleListNode
	Tail *DoubleListNode
}

func (list *DoubleList) PushFront(k string, values ...interface{}) {
	log.Debugf("push front %s %v", k, values)
	v := ""
	if len(values) > 0 {
		v = values[0].(string)
	}
	node := &DoubleListNode{
		Key:   k,
		Value: v,
		Prev:  nil,
		Next:  nil,
	}

	if list.Head == nil {
		list.Head = node
		list.Tail = node
		return
	}
	oldHead := list.Head
	list.Head = node
	list.Head.Next = oldHead
	oldHead.Prev = list.Head
}

func (list *DoubleList) Push(k string, values ...interface{}) {
	log.Infof("push %s %v", k, values)
	v := ""
	if len(values) > 0 {
		v = values[0].(string)
	}
	node := &DoubleListNode{
		Key:   k,
		Value: v,
		Prev:  nil,
		Next:  nil,
	}
	if list.Head == nil {
		list.Head = node
	} else {
		list.Tail.Next = node
		node.Prev = list.Tail
	}
	list.Tail = node
}

func (list *DoubleList) Print() {
	ptr := list.Head
	for ptr != nil {
		fmt.Printf("%s,", ptr.Value)
		ptr = ptr.Next
	}
	fmt.Println()
}

func (list *DoubleList) Len() int {
	i := 0
	ptr := list.Head
	for ptr != nil {
		i++
		ptr = ptr.Next
	}
	return i
}

func (list *DoubleList) Pop() *DoubleListNode {
	del := list.Tail
	list.Tail = list.Tail.Prev
	list.Tail.Next = nil
	return del
}

func (list *DoubleList) ReversePrint() {
	ptr := list.Tail
	for ptr != nil {
		fmt.Println(ptr.Key)
		ptr = ptr.Prev
	}
}
func (list *DoubleList) Find(string2 string) *DoubleListNode {
	ptr := list.Head
	for ptr != nil {
		//fmt.Println("listnode:", ptr.Key, ptr)
		if ptr.Key == string2 {
			//fmt.Println("found")
			return ptr
		}
		ptr = ptr.Next
	}
	return nil
}

func (list *DoubleList) Delete(string2 string) {
	cur := list.Find(string2)

	if cur != nil {
		prev := cur.Prev
		next := cur.Next
		prev.Next = cur.Next
		next.Prev = cur.Prev
	}

}

func (list *DoubleList) DeleteNode(node *DoubleListNode) {
	//delete node
	fmt.Println("delete node")
	if node != nil {
		prev := node.Prev
		next := node.Next
		prev.Next = node.Next
		next.Prev = node.Prev
	}

}
