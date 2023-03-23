package lru

type LRUUtil struct {
}

type LRUNode struct {
	pre  *LRUNode    // pre
	next *LRUNode    // next
	key  string      // key
	data interface{} // data
	time int64       // time
}

type LRURoot struct {
	size int                 // 长度
	head LRUNode             // start
	tail LRUNode             // end
	kMap map[string]*LRUNode // 当前LRU包含那些数据
}

// LRU不需要初始化对象，需要使用时传递值并将结果保存
var LruM *map[string]*LRURoot = &map[string]*LRURoot{}
