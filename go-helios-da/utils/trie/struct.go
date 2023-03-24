package trie

type TrieTreeUtil struct {
}

type TrieTree struct {
	TrieMap map[rune]*TrieTree
	IsEnd   bool
	Data    []*interface{}
}

type IndexConf struct {
	IndexKey    string `toml:"index_key"`
	IndexType   string `toml:"index_type"`
	IndexFormat string `toml:"index_format"`
	URL         string `toml:"url"`
	LRUSize     int    `toml:"index_lru_size"`
	LRUTime     int    `toml:"index_lru_time"`
	Mini        [][]string
}

type IndexNeedInfo struct {
	IndexConf
	DataMap  []map[string]interface{}
	DataList []string
}

var TrieRootMap *map[string]*TrieTree = &map[string]*TrieTree{}

var SUG_MAX_NUM = 10
