package lru

import "context"

type LRUUtilInterface interface {
	LRUInit(ctx context.Context, index string, size, lruTime int) (err error)
	LRUUtilHasIndex(ctx context.Context, index string) (b bool)
	GetAllByIndex(ctx context.Context, index string) (data interface{}, err error)
	GetLRUByKeyAndIndex(ctx context.Context, index, key string) (data interface{}, err error)
	PutLRUByKeyAndIndex(ctx context.Context, index, key string, data interface{}) (err error)
}

var r LRUUtilInterface

func init() {
	Register(&LRUUtil{})
}

func Register(m *LRUUtil) {
	r = m
}

func LRUInit(ctx context.Context, index string, size, lruTime int) (err error) {
	return r.LRUInit(ctx, index, size, lruTime)
}

func LRUUtilHasIndex(ctx context.Context, index string) (b bool) {
	return r.LRUUtilHasIndex(ctx, index)
}

func GetAllByIndex(ctx context.Context, index string) (data interface{}, err error) {
	return r.GetAllByIndex(ctx, index)
}

func GetLRUByKeyAndIndex(ctx context.Context, key, index string) (data interface{}, err error) {
	return r.GetLRUByKeyAndIndex(ctx, key, index)
}

func PutLRUByKeyAndIndex(ctx context.Context, index, key string, data interface{}) (err error) {
	return r.PutLRUByKeyAndIndex(ctx, index, key, data)
}
