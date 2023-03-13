package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-helios-da/resource"
	"net/http"
	"strconv"
)

func HeliosStart(ctx *gin.Context) {
	//helios.InitHeliosIndexList(ctx)
	ctx.String(http.StatusOK, "ok")
}

// 判断一个key是否在index中存在
func HeliosHasKey(ctx *gin.Context) {
	index := ctx.DefaultQuery("index", "music")
	key := ctx.DefaultQuery("query", "")
	if key == "" {
		err := fmt.Errorf("query is null")
		info := BuildErrResponse(ctx, nil, err)
		ctx.JSON(http.StatusOK, info)
		return
	}
	isExist, err := resource.RESOURCE_TRIEROOT.KeyIsExistInIndex(ctx, index, key)
	if nil != err {
		err = fmt.Errorf("KeyIsExistByIndex has err : %s", err.Error())
		info := BuildErrResponse(ctx, nil, err)
		ctx.JSON(http.StatusOK, info)
		return
	}
	res := BuildOkResponse(ctx, isExist)
	ctx.JSON(http.StatusOK, res)
}

// 根据一个key获得相应的数据
func HeliosGetDataByKey(ctx *gin.Context) {
	index := ctx.DefaultQuery("index", "music")
	key := ctx.DefaultQuery("query", "")
	if key == "" {
		err := fmt.Errorf("query is null")
		info := BuildErrResponse(ctx, nil, err)
		ctx.JSON(http.StatusOK, info)
		return
	}
	data, err := resource.RESOURCE_TRIEROOT.GetDataByKey(ctx, index, key)
	if nil != err {
		err = fmt.Errorf("SearchOneKeyIsExist has err : %s", err.Error())
		info := BuildErrResponse(ctx, nil, err)
		ctx.JSON(http.StatusOK, info)
		return
	}
	res := BuildOkResponse(ctx, data)
	ctx.JSON(http.StatusOK, res)
}

func HeliosSugQueryByIndexAndWord(ctx *gin.Context) {
	key := ctx.DefaultQuery("query", "")
	if key == "" {
		err := fmt.Errorf("query is null")
		info := BuildErrResponse(ctx, nil, err)
		ctx.JSON(http.StatusOK, info)
		return
	}
	index := ctx.DefaultQuery("index", "music")
	maxNumStr := ctx.DefaultQuery("maxNum", "10")
	maxNum, err := strconv.Atoi(maxNumStr)
	if nil != err {
		maxNum = 10
	}
	list, err := resource.RESOURCE_TRIEROOT.SugQueryBySubWord(ctx, index, key, maxNum)

	if nil != err {
		fmt.Println("SugQueryBySubWord has err : %s", err.Error())
		info := BuildErrResponse(ctx, nil, err)
		ctx.JSON(http.StatusOK, info)
		return
	}
	res := BuildOkResponse(ctx, list)
	ctx.JSON(http.StatusOK, res)
}

// 模糊查询，获得相关的数据集
func HeliosSugDataByIndexAndWord(ctx *gin.Context) {
	key := ctx.DefaultQuery("index", "music")
	word := ctx.DefaultQuery("query", "")
	if word == "" {
		err := fmt.Errorf("query is null")
		info := BuildErrResponse(ctx, nil, err)
		ctx.JSON(http.StatusOK, info)
		return
	}

	maxNumStr := ctx.DefaultQuery("maxNum", "10")
	maxNum, err := strconv.Atoi(maxNumStr)
	if nil != err {
		maxNum = 10
	}

	sugData, err := resource.RESOURCE_TRIEROOT.SugDataListBySubWord(ctx, key, word, maxNum)
	if nil != err {
		fmt.Println("SugDataListBySubWord has err : %s", err.Error())
		info := BuildErrResponse(ctx, nil, err)
		ctx.JSON(http.StatusOK, info)
		return
	}
	res := BuildOkResponse(ctx, sugData)
	ctx.JSON(http.StatusOK, res)
}
