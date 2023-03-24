package lru

import (
	"context"
	"fmt"
	"time"
)

func (r *LRUUtil) LRUInit(ctx context.Context, index string) (err error) {
	s := LRUNode{}
	e := LRUNode{}
	s.next = &e
	e.pre = &s
	root := LRURoot{
		head: &s,
		tail: &e,
		size: 5,
		t:    5,
		kMap: map[string]*LRUNode{},
	}
	// 将该root注册到LRU表
	(*LruM)[index] = &root
	return nil
}

func (r *LRUUtil) LRUUtilHasIndex(ctx context.Context, index string) (b bool) {
	if _, ok := (*LruM)[index]; ok {
		return true
	}
	return false
}

func (r *LRUUtil) GetAllByIndex(ctx context.Context, index string) (data interface{}, err error) {
	if _, ok := (*LruM)[index]; ok {
		return (*LruM)[index], nil
	}
	return nil, nil
}

func (r *LRUUtil) GetLRUByKeyAndIndex(ctx context.Context, index, key string) (data interface{}, err error) {
	if _, ok := (*LruM)[index]; ok {
		return getLRUNodeData(ctx, key, (*LruM)[index])
	}
	err = fmt.Errorf("root is nil")
	return nil, err
}

func (r *LRUUtil) PutLRUByKeyAndIndex(ctx context.Context, index, key string, data interface{}) (err error) {
	if _, ok := (*LruM)[index]; ok {
		return putLRUNode(ctx, key, data, (*LruM)[index])
	}
	err = fmt.Errorf("root is nil")
	return err
}

// 基础方法 查询LRU
func getLRUNodeData(ctx context.Context, key string, root *LRURoot) (data interface{}, err error) {

	if root == nil {
		err = fmt.Errorf("root is nil")
		return nil, err
	}

	if _, ok := root.kMap[key]; !ok {
		// 不存在
		return nil, nil
	}

	node := root.kMap[key]

	// 检查时间
	if node.overTime < time.Now().Unix() {
		// 删除当前节点
		removeNode(node)
		// 删除当前map_k
		delete(root.kMap, node.key)
		return nil, nil
	}

	// 如果存在，需要将他重新放到头节点
	data = node.data
	moveToHead(node, root)

	return data, nil
}

// 基础方法 增加节点LRU
func putLRUNode(ctx context.Context, key string, data interface{}, root *LRURoot) (err error) {
	if root == nil {
		err = fmt.Errorf("root is nil")
		return err
	}
	if node, ok := root.kMap[key]; ok {
		// 如果存在，需要将他重新放到头节点
		moveToHead(node, root)
		return nil
	} else {
		// 不存在需要添加到头部并且
		node := LRUNode{
			key:      key,
			data:     data,
			overTime: time.Now().Add(time.Minute * time.Duration(root.t)).Unix(),
		}
		addToHead(&node, root)
		root.kMap[key] = &node
		if len(root.kMap) > root.size {
			// 长度超长的时候，需要删除最后一个
			removeTail(root)
		}
	}
	return
}

func addToHead(node *LRUNode, root *LRURoot) {
	// 列表头节点
	sN := root.head.next
	// 交换列表头节点和当前数据节点
	sN.pre = node
	node.next = sN
	// 交换头节点位置
	node.pre = root.head
	root.head.next = node
}

func removeNode(node *LRUNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func moveToHead(node *LRUNode, root *LRURoot) {
	removeNode(node)
	addToHead(node, root)
}

func removeTail(root *LRURoot) {
	// 尾节点
	tailNode := root.tail.pre
	removeNode(tailNode)
	delete(root.kMap, tailNode.key)
}
