package trie

import "context"

type TrieTreeDaoInterface interface {
	PopTrieRoot(ctx context.Context) *map[string]*TrieTree
	BuildTrieTreeBySet(ctx context.Context, index string, dataMap map[string][]interface{})
	BuildTrieTree(ctx context.Context, index string, word string, data []interface{})

	// 查询某个索引是否存在
	KeyIsExistByIndex(ctx context.Context, index string, key string) (b bool, err error)
	GetDataByKey(ctx context.Context, index string, key string) (data []interface{}, err error)
	SugQueryBySubWord(ctx context.Context, index, subQuery string, maxNum int) (list []string, err error)
	SugDataListBySubWord(ctx context.Context, index, subQuery string, maxNum int) (dataList []interface{}, err error)
}

var r TrieTreeDaoInterface

func init() {
	Register(&TrieTreeDao{})
	//r = &Search{}
}

func Register(m *TrieTreeDao) {
	r = m
}

func PopTrieRoot(ctx context.Context) *map[string]*TrieTree {
	return r.PopTrieRoot(ctx)
}

func BuildTrieTreeBySet(ctx context.Context, index string, dataMap map[string][]interface{}) {
	r.BuildTrieTreeBySet(ctx, index, dataMap)
}

func BuildTrieTree(ctx context.Context, index string, word string, data []interface{}) {
	r.BuildTrieTree(ctx, index, word, data)
}

func KeyIsExistByIndex(ctx context.Context, index string, word string) (b bool, err error) {
	return r.KeyIsExistByIndex(ctx, index, word)
}

func GetDataByKey(ctx context.Context, index string, key string) (data []interface{}, err error) {
	return r.GetDataByKey(ctx, index, key)
}

func SugQueryBySubWord(ctx context.Context, index, subQuery string, maxNum int) (list []string, err error) {
	return r.SugQueryBySubWord(ctx, index, subQuery, maxNum)
}

func SugDataListBySubWord(ctx context.Context, index, subQuery string, maxNum int) (dataList []interface{}, err error) {
	return r.SugDataListBySubWord(ctx, index, subQuery, maxNum)
}
