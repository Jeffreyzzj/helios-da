package trie

import (
	"context"
	"fmt"
)

// 基础方法 索引index添加数据data
func addTrieTreeNode(index string, word []rune, data []interface{}) {
	if (*TrieRootMap)[index] == nil {
		(*TrieRootMap)[index] = &TrieTree{
			TrieMap: map[rune]*TrieTree{},
		}
	}
	root := (*TrieRootMap)[index]

	for _, v := range word {
		if _, ok := root.TrieMap[v]; !ok {
			root.TrieMap[v] = &TrieTree{
				TrieMap: map[rune]*TrieTree{},
				IsEnd:   false,
				Data:    nil,
			}
		}
		root = root.TrieMap[v]
	}
	root.IsEnd = true
	if root.Data == nil {
		root.Data = []*interface{}{}
	}
	trieTreeData := []*interface{}{}
	for i := 0; i < len(data); i++ {
		trieTreeData = append(trieTreeData, &data[i])
	}
	root.Data = append(root.Data, trieTreeData...)
}

// 基础方法 在新指针中添加数据data
func addTrieTreeNodeByNewRoot(root *TrieTree, word []rune, data []interface{}) {

	for _, v := range word {
		if _, ok := root.TrieMap[v]; !ok {
			root.TrieMap[v] = &TrieTree{
				TrieMap: map[rune]*TrieTree{},
				IsEnd:   false,
				Data:    nil,
			}
		}
		root = root.TrieMap[v]
	}
	root.IsEnd = true
	if root.Data == nil {
		root.Data = []*interface{}{}
	}
	trieTreeData := []*interface{}{}
	for i := 0; i < len(data); i++ {
		trieTreeData = append(trieTreeData, &data[i])
	}
	root.Data = append(root.Data, trieTreeData...)
}

// 基础方法 根据key获得node
func getNodeByKey(ctx context.Context, index string, words []rune) (node *TrieTree, err error) {
	if (*TrieRootMap)[index] == nil {
		err = fmt.Errorf("helios hasn't index[%s]", index)
		return nil, err
	}
	root := (*TrieRootMap)[index]

	for _, v := range words {
		if _, ok := root.TrieMap[v]; !ok {
			return nil, nil
		}
		root = root.TrieMap[v]
	}

	return root, nil
}

// 基础方法 在一个索引index中查找以subWord为开头的倒排索引
func SugBySubWord(ctx context.Context, index string, words []rune, maxNum int) (list [][]rune, err error) {
	if (*TrieRootMap)[index] == nil {
		err = fmt.Errorf("helios hasn't index[%s]", index)
		return list, err
	}
	root := (*TrieRootMap)[index]
	// 找到当前的subWord节点
	for _, v := range words {
		if _, ok := root.TrieMap[v]; !ok {
			return list, nil
		}
		root = root.TrieMap[v]
	}

	// 当前root为最后一个节点，开始查询当前节点下的完整倒排索引
	dfsList := &[][]rune{}
	TrieTreeDFS(ctx, root, words, dfsList, maxNum)

	return *dfsList, nil
}

// 基础方法 在一个索引index中查找以subWord为开头的数据集
func sugDataListBySubWord(ctx context.Context, index string, words []rune, maxNum int) (list []*interface{}, err error) {
	if (*TrieRootMap)[index] == nil {
		err = fmt.Errorf("helios hasn't index[%s]", index)
		return list, err
	}
	root := (*TrieRootMap)[index]
	// 找到当前的subWord节点
	for _, v := range words {
		if _, ok := root.TrieMap[v]; !ok {
			return list, nil
		}
		root = root.TrieMap[v]
	}

	// 当前root为最后一个节点，开始查询当前节点下的完整倒排索引
	dfsList := &[]*interface{}{}
	getDataTrieTreeDFS(ctx, root, dfsList, maxNum)

	return *dfsList, nil
}

func TrieTreeDFS(ctx context.Context, root *TrieTree, words []rune, dfsList *[][]rune, maxNum int) {

	if root == nil {
		return
	} else if len(*dfsList) >= maxNum {
		return
	}

	if root.IsEnd {
		*dfsList = append(*dfsList, words)
	}

	for k, v := range root.TrieMap {
		root = v
		words = append(words, k)
		TrieTreeDFS(ctx, v, words, dfsList, maxNum)
		words = words[0 : len(words)-1]
		if len(*dfsList) >= maxNum {
			return
		}
	}
	return
}

func getDataTrieTreeDFS(ctx context.Context, root *TrieTree, dfsList *[]*interface{}, maxNum int) {

	if root == nil {
		return
	} else if len(*dfsList) >= maxNum {
		return
	}

	if root.IsEnd {
		*dfsList = append(*dfsList, root.Data...)
	}

	for _, v := range root.TrieMap {
		root = v
		getDataTrieTreeDFS(ctx, v, dfsList, maxNum)
		if len(*dfsList) >= maxNum {
			return
		}
	}
	return
}
