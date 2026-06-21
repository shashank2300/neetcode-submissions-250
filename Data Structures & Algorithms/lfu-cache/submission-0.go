type ListNode struct {
    key  int
    val  int
    freq int
    prev *ListNode
    next *ListNode
}

type LinkedList struct {
    left  *ListNode
    right *ListNode
    size  int
}

func NewLinkedList() *LinkedList {
    left := &ListNode{}
    right := &ListNode{}
    left.next = right
    right.prev = left
    return &LinkedList{
        left:  left,
        right: right,
    }
}

func (ll *LinkedList) length() int {
    return ll.size
}

func (ll *LinkedList) pushRight(node *ListNode) {
    prev := ll.right.prev
    prev.next = node
    node.prev = prev
    node.next = ll.right
    ll.right.prev = node
    ll.size++
}

func (ll *LinkedList) pop(node *ListNode) {
    prev, next := node.prev, node.next
    prev.next = next
    next.prev = prev
    node.prev = nil
    node.next = nil
    ll.size--
}

func (ll *LinkedList) popLeft() *ListNode {
    node := ll.left.next
    ll.pop(node)
    return node
}

type LFUCache struct {
    capacity int
    lfuCount int
    nodeMap  map[int]*ListNode
    listMap  map[int]*LinkedList
}

func Constructor(capacity int) LFUCache {
    return LFUCache{
        capacity: capacity,
        lfuCount: 0,
        nodeMap:  make(map[int]*ListNode),
        listMap:  make(map[int]*LinkedList),
    }
}

func (this *LFUCache) counter(node *ListNode) {
    count := node.freq
    this.listMap[count].pop(node)

    if count == this.lfuCount && this.listMap[count].length() == 0 {
        this.lfuCount++
    }

    node.freq++
    if _, exists := this.listMap[node.freq]; !exists {
        this.listMap[node.freq] = NewLinkedList()
    }
    this.listMap[node.freq].pushRight(node)
}

func (this *LFUCache) Get(key int) int {
    node, exists := this.nodeMap[key]
    if !exists {
        return -1
    }
    this.counter(node)
    return node.val
}

func (this *LFUCache) Put(key int, value int) {
    if this.capacity == 0 {
        return
    }

    if node, exists := this.nodeMap[key]; exists {
        node.val = value
        this.counter(node)
        return
    }

    if len(this.nodeMap) == this.capacity {
        toRemove := this.listMap[this.lfuCount].popLeft()
        delete(this.nodeMap, toRemove.key)
    }

    node := &ListNode{key: key, val: value, freq: 1}
    this.nodeMap[key] = node
    if _, exists := this.listMap[1]; !exists {
        this.listMap[1] = NewLinkedList()
    }
    this.listMap[1].pushRight(node)
    this.lfuCount = 1
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param1 := obj.Get(key);
 * obj.Put(key,value);
 */
