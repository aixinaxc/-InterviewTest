package algorithm

import (
	"fmt"
)

func main() {
	LRUCatch(3)
	put("1", 1111)
	put("1", 2222)
	put("2", 1111)
	put("3", 3333)
	put("2", 2222)
	put("5", 5555)
	put("5", 6666)
	fmt.Println("nodes:", nodes)
	fmt.Println("first:", first)
	fmt.Println("last:", last)
	f(nodes)
}

type LRUCatchNode struct {
	Prev  *LRUCatchNode
	Next  *LRUCatchNode
	Key   string
	Value interface{}
}

var catchSize int
var currentSize int
var nodes *LRUCatchNode
var first string
var last string

//初始化lru大小
func LRUCatch(i int) {
	currentSize = 0
	catchSize = i
}

//像lru中添加数据
func put(key string, value interface{}) {
	if first == "" {
		first = key
	}
	last = key

	if h(key, nodes) == "" {
		if currentSize >= catchSize {
			nodes = Delete(nodes)
		} else {
			currentSize++
		}
	} else {
		nodes = DeleteSameKey(key, nodes)
	}
	add(key, value)
}
//检查此key是否在lru链表中
func h(key string, node *LRUCatchNode) string {
	if node == nil {
		return ""
	}
	if node.Key == key {
		return key
	}
	return h(key, node.Prev)
}
//像链表中添加数据
func add(key string, value interface{}) {
	nodes = w(nodes, nil, key, value)
}
//像链表中添加数据
func w(node *LRUCatchNode, n *LRUCatchNode, key string, value interface{}) *LRUCatchNode {
	if node == nil {
		nN := new(LRUCatchNode)
		nN.Key = key
		nN.Value = value
		nN.Prev = n
		if n != nil {
			n.Next = nN
		}
		node = nN
		return node
	}else {
		if node.Key == key {
			node.Value = value
			return node
		}
	}
	return w(node.Next, node, key, value)
}
//删除相同key的链表数据
func DeleteSameKey(key string, node *LRUCatchNode) *LRUCatchNode {
	if node.Key == key {
		if currentSize == 1 {
			return nil
		}
		if last == key {
			return node
		}
		if node.Prev != nil {
			node.Prev.Next = node.Next
		}
		if node.Next != nil {
			node.Next.Prev = node.Prev
			if key == first {
				first = node.Next.Key
			}
		}
		return node
	}
	return DeleteSameKey(key, node.Prev)
}
//当新key不存在于链表时，删除第一条数据
func Delete(node *LRUCatchNode) *LRUCatchNode {
	if node.Key == first {
		node.Next.Prev = nil
		first = node.Next.Key
		return node
	}
	return Delete(node.Prev)
}
//打印所以数据
func f(node *LRUCatchNode) *LRUCatchNode {
	fmt.Println(node)
	if node == nil {
		return nil
	}
	return f(node.Prev)
}
