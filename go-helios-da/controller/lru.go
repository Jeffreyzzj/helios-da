package controller

import (
	"fmt"
	"go-helios-da/resource"
	"net/http"

	"github.com/gin-gonic/gin"
)

// lru方法不应当直接暴露接口，而是为基础工程的相关内容进行提速
// 这里仅仅用来测试相关功能

func LRUGetData(ctx *gin.Context) {
	d := ctx.Query("d")
	data, err := resource.RESOURCE_LRUROOT.GetLRUByKeyAndIndex(ctx, key, d)
	if nil != err {
		fmt.Printf("GetLRUByKeyAndIndex has err %s", err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func LRUPutData(ctx *gin.Context) {
	d := ctx.Query("d")
	ds := req{
		Data: d,
	}

	hasIndex := resource.RESOURCE_LRUROOT.LRUUtilHasIndex(ctx, key)
	if !hasIndex {
		resource.RESOURCE_LRUROOT.LRUInit(ctx, key)
	}

	err := resource.RESOURCE_LRUROOT.PutLRUByKeyAndIndex(ctx, key, d, ds)
	if nil != err {
		ctx.JSON(http.StatusOK, err.Error())
	}
	ctx.JSON(http.StatusOK, "ok")
}

var key = "test"

type req struct {
	Data string
}
