package lru_cache

//运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
//获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
//写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
//
//进阶:
//
//你是否可以在 O(1) 时间复杂度内完成这两种操作？
//
//示例:
//
//LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得密钥 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得密钥 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/lru-cache
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//题解：为保证在O(1)时间复杂度内完成Get和Set，需采用双向链表存储常用项，并采用哈希表获取值。

type Node struct {
	Next  *Node
	Pre   *Node
	key   int
	Value int
}

type LRUCache struct {
	m      map[int]*Node
	length int
	cap    int
	Head   *Node
	Tail   *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m:   make(map[int]*Node, capacity),
		cap: capacity,
	}
}

func (l *LRUCache) Get(key int) (val int) {
	n, ok := l.m[key]
	if !ok {
		val = -1
		return
	}
	val = n.Value

	l.Fix(n)
	return
}

func (l *LRUCache) Put(key int, value int) {
	if l == nil {
		return
	}

	n, ok := l.m[key]
	if ok {
		n.Value = value
		l.Fix(n)
		return
	}

	if l.length == l.cap {
		//remove one
		delete(l.m, l.Tail.key)
		l.Tail = l.Tail.Pre
		if l.Tail != nil {
			l.Tail.Next = nil
		}
		l.length--
	}
	n = &Node{
		Next:  l.Head,
		key:   key,
		Value: value,
	}
	if l.Head != nil {
		l.Head.Pre = n
	}
	l.Head = n
	if l.Tail == nil {
		l.Tail = l.Head
	}

	l.m[key] = n
	l.length++
}

func (l *LRUCache) Fix(n *Node) {
	if n == l.Head {
		return
	}

	if n == l.Tail {
		l.Tail = n.Pre
		if l.Tail != nil {
			l.Tail.Next = nil
		}
	} else {
		if n.Next != nil {
			n.Next.Pre = n.Pre
		}
		if n.Pre != nil {
			n.Pre.Next = n.Next
		}
	}

	n.Next = l.Head
	l.Head.Pre = n
	n.Pre = nil
	l.Head = n
}
