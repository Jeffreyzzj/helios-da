package trie

import "context"

type TrieTreeDaoInterface interface {
	TrieRootInit(ctx context.Context) (err error)
	PopTrieRoot(ctx context.Context) *map[string]*TrieTree
	BuildTrieTreeBySet(ctx context.Context, index string, dataMap map[string][]interface{})
	BuildTrieTree(ctx context.Context, index string, word string, data []interface{})

	// 查询某个索引是否存在
	KeyIsExistInIndex(ctx context.Context, index string, key string) (b bool, err error)
	GetDataByKey(ctx context.Context, index string, key string) (data []interface{}, err error)
	SugQueryBySubWord(ctx context.Context, index, subQuery string, maxNum int) (list []string, err error)
	SugDataListBySubWord(ctx context.Context, index, subQuery string, maxNum int) (dataList []interface{}, err error)
}

var r TrieTreeDaoInterface

func init() {
	Register(&TrieTreeUtil{})
	//r = &Search{}
}

func Register(m *TrieTreeUtil) {
	r = m
}

func TrieRootInit(ctx context.Context) (err error) {
	return r.TrieRootInit(ctx)
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

func KeyIsExistInIndex(ctx context.Context, index string, word string) (b bool, err error) {
	return r.KeyIsExistInIndex(ctx, index, word)
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
