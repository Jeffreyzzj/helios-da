package resource

import (
	"go-helios-da/config"
	"go-helios-da/utils/lru"
	"go-helios-da/utils/trie"

	"go.uber.org/zap"
)

var LOGGER *zap.SugaredLogger

var LOGGER_USER *zap.SugaredLogger

var RESOURCE_TRIEROOT *trie.TrieTreeUtil

var RESOURCE_LRUROOT *lru.LRUUtil

var RESOURCE_CONF *config.TomlConfig
