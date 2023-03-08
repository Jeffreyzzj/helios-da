package trie

import (
	"context"
	"fmt"
)

func (t *TrieTreeDao) PopTrieRoot(ctx context.Context) *map[string]*TrieTree {
	return TrieRootMap
}

func (t *TrieTreeDao) BuildTrieTreeBySet(ctx context.Context, index string, dataMap map[string][]interface{}) {
	root := &TrieTree{
		TrieMap: map[rune]*TrieTree{},
		IsEnd:   false,
		Data:    []*interface{}{},
	}
	for k, v := range dataMap {
		//addTrieTreeNode(index, []rune(k), v)
		addTrieTreeNodeByNewRoot(root, []rune(k), v)
	}

	if (*TrieRootMap)[index] == nil {
		(*TrieRootMap)[index] = &TrieTree{
			TrieMap: map[rune]*TrieTree{},
		}
	}
	(*TrieRootMap)[index] = root
	return
}

func (t *TrieTreeDao) BuildTrieTree(ctx context.Context, index string, word string, data []interface{}) {
	// 获得当前索引的位置
	//addTrieTreeNode(index, []rune(word), data)

	root := &TrieTree{
		TrieMap: map[rune]*TrieTree{},
		IsEnd:   false,
		Data:    []*interface{}{},
	}
	addTrieTreeNodeByNewRoot(root, []rune(word), data)

	if (*TrieRootMap)[index] == nil {
		(*TrieRootMap)[index] = &TrieTree{
			TrieMap: map[rune]*TrieTree{},
		}
	}
	(*TrieRootMap)[index] = root
	return
}

func (t *TrieTreeDao) KeyIsExistByIndex(ctx context.Context, index string, key string) (b bool, err error) {
	node, err := getNodeByKey(ctx, index, []rune(key))
	if nil != err {
		err = fmt.Errorf("getNodeByKey key[%s] has err %s", key, err.Error())
		return
	} else if node == nil {
		return false, nil
	}
	return node.IsEnd, nil
}

func (t *TrieTreeDao) GetDataByKey(ctx context.Context, index string, key string) (data []interface{}, err error) {
	node, err := getNodeByKey(ctx, index, []rune(key))
	if nil != err {
		err = fmt.Errorf("getNodeByKey key[%s] has err %s", key, err.Error())
		return nil, err
	} else if node == nil {
		return nil, nil
	}
	resList := []interface{}{}
	for _, v := range node.Data {
		resList = append(resList, *v)
	}
	return resList, nil
}

func (t *TrieTreeDao) SugQueryBySubWord(ctx context.Context, index, subQuery string, maxNum int) (list []string, err error) {
	sugList, err := SugBySubWord(ctx, index, []rune(subQuery), maxNum)
	if nil != err {
		err = fmt.Errorf("SugBySubWord has error %s", err.Error())
		return list, err
	}
	for _, v := range sugList {
		list = append(list, string(v))
	}

	return
}

func (t *TrieTreeDao) SugDataListBySubWord(ctx context.Context, index, subQuery string, maxNum int) (dataList []interface{}, err error) {
	datas, err := sugDataListBySubWord(ctx, index, []rune(subQuery), maxNum)
	if nil != err {
		err = fmt.Errorf("sugDataListBySubWord has err %s", err.Error())
		return
	}
	// 处理数据，将数据丢出
	dataList = []interface{}{}
	for i := 0; i < len(datas); i++ {
		dataList = append(dataList, *datas[i])
	}
	return dataList, nil
}
