package resource

import (
	"go-helios-da/config"
	"go-helios-da/utils/lru"
	"go-helios-da/utils/trie"
)

var RESOURCE_TRIEROOT *trie.TrieTreeUtil

var RESOURCE_LRUROOT *lru.LRUUtil

var RESOURCE_CONF *config.TomlConfig
