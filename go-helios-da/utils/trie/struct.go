package trie

type TrieTreeDao struct {
}

type TrieTree struct {
	TrieMap map[rune]*TrieTree
	IsEnd   bool
	Data    []*interface{}
}

type QueryTrieTree struct {
	TrieMap map[rune]*TrieTree
	IsEnd   bool
	/*Data    []*interface{}*/
}

var TrieRootMap *map[string]*TrieTree = &map[string]*TrieTree{}

var SUG_MAX_NUM = 10
