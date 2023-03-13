package resource

import (
	"go-helios-da/utils/trie"
	"go.uber.org/zap"
)

var LOGGER *zap.SugaredLogger

var LOGGER_USER *zap.SugaredLogger

var RESOURCE_TRIEROOT *trie.TrieTreeUtil
