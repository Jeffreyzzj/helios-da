package lru

type LRUUtil struct {
}

type LRUNode struct {
	pre      *LRUNode    // pre
	next     *LRUNode    // next
	key      string      // key
	data     interface{} // data
	overTime int64       // time
}

type LRURoot struct {
	size int                 // 允许长度
	t    int                 // 缓存时间
	head *LRUNode            // start
	tail *LRUNode            // end
	kMap map[string]*LRUNode // 当前LRU包含那些数据
}

// LRU不需要初始化对象，需要使用时传递值并将结果保存
var LruM *map[string]*LRURoot = &map[string]*LRURoot{}
