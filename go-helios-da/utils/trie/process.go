package trie

import (
	"context"
	"encoding/json"
	"fmt"
	"go-helios-da/config"
	"go-helios-da/global"
	"go-helios-da/utils/lru"

	"github.com/BurntSushi/toml"
)

func (t *TrieTreeUtil) PopTrieRoot(ctx context.Context) *map[string]*TrieTree {
	return TrieRootMap
}

func (t *TrieTreeUtil) TrieRootInit(ctx context.Context) (err error) {
	// 读取需要处理的conf文件
	var indexConf config.TomlConfig
	filePath := global.DA_CONF_PATH
	if _, err := toml.DecodeFile(filePath, &indexConf); err != nil {
		err = fmt.Errorf("read toml has err %s", err.Error())
		return err
	}

	// 取出IndexConfigs列表，开始准备构建相应的倒排索引
	for _, v := range indexConf.HeliosInitConfig.IndexConfigs {
		err = buildIndexByIndexConf(ctx, v)
		if nil != err {
			fmt.Printf("BuildIndexByIndexConf key[%s] has err[%s] \n", v.Conf, err.Error())
		}
	}
	return nil
}

func (t *TrieTreeUtil) BuildTrieTreeBySet(ctx context.Context, index string, dataMap map[string][]interface{}) {
	root := &TrieTree{
		TrieMap: map[rune]*TrieTree{},
		IsEnd:   false,
		Data:    []*interface{}{},
	}
	for k, v := range dataMap {
		addTrieTreeNodeByNewRoot(root, []rune(k), v)
	}

	if (*TrieRootMap)[index] == nil {
		(*TrieRootMap)[index] = &TrieTree{
			TrieMap: map[rune]*TrieTree{},
		}
	}
	(*TrieRootMap)[index] = root
	//return
}

func (t *TrieTreeUtil) BuildTrieTree(ctx context.Context, index string, word string, data []interface{}) {
	// 获得当前索引的位置
	//addTrieTreeNode(index, []rune(word), data)

	root := &TrieTree{
		TrieMap: map[rune]*TrieTree{},
		IsEnd:   false,
		Data:    []*interface{}{},
	}
	addTrieTreeNodeByNewRoot(root, []rune(word), data)

	if (*TrieRootMap)[index] == nil {
		(*TrieRootMap)[index] = &TrieTree{
			TrieMap: map[rune]*TrieTree{},
		}
	}
	(*TrieRootMap)[index] = root
	//return
}

func (t *TrieTreeUtil) KeyIsExistInIndex(ctx context.Context, index string, key string) (b bool, err error) {
	node, err := getNodeByKey(ctx, index, []rune(key))
	if nil != err {
		err = fmt.Errorf("getNodeByKey key[%s] has err %s", key, err.Error())
		return
	} else if node == nil {
		return false, nil
	}
	return node.IsEnd, nil
}

func (t *TrieTreeUtil) GetDataByKey(ctx context.Context, index string, key string) (data []interface{}, err error) {
	node, err := getNodeByKey(ctx, index, []rune(key))
	if nil != err {
		err = fmt.Errorf("getNodeByKey key[%s] has err %s", key, err.Error())
		return nil, err
	} else if node == nil {
		return nil, nil
	}
	resList := []interface{}{}
	for _, v := range node.Data {
		resList = append(resList, *v)
	}
	return resList, nil
}

func (t *TrieTreeUtil) SugQueryBySubWord(ctx context.Context, index, subQuery string, maxNum int) (list []string, err error) {
	// 使用lru加速
	data, err := lru.GetLRUByKeyAndIndex(ctx, index, subQuery)
	if nil != err {
		fmt.Printf("GetLRUByKeyAndIndex has err %s", err.Error())
	} else if data != nil {
		dBtyeList, err := json.Marshal(data)
		if nil == err {
			dataList := []string{}
			err = json.Unmarshal(dBtyeList, &dataList)
			if nil == err {
				return dataList, nil
			}
		}
		fmt.Printf("GetLRUByKeyAndIndex's data json to byte has err %s, \n", err.Error())
	}

	sugList, err := SugBySubWord(ctx, index, []rune(subQuery), maxNum)
	if nil != err {
		err = fmt.Errorf("SugBySubWord has error %s", err.Error())
		return list, err
	}
	for _, v := range sugList {
		list = append(list, string(v))
	}

	// 保存lru
	lru.PutLRUByKeyAndIndex(ctx, index, subQuery, sugList)

	return
}

func (t *TrieTreeUtil) SugDataListBySubWord(ctx context.Context, index, subQuery string, maxNum int) (dataList []interface{}, err error) {
	datas, err := sugDataListBySubWord(ctx, index, []rune(subQuery), maxNum)
	if nil != err {
		err = fmt.Errorf("sugDataListBySubWord has err %s", err.Error())
		return
	}
	// 处理数据，将数据丢出
	dataList = []interface{}{}
	for i := 0; i < len(datas); i++ {
		dataList = append(dataList, *datas[i])
	}
	return dataList, nil
}

// 建立倒排索引
func buildIndexByIndexConf(ctx context.Context, conf config.IndexConf) (err error) {
	indexInfoParams, err := getIndexInfo(ctx, conf)
	if nil != err {
		err = fmt.Errorf("getIndexInfo has err %s ", err.Error())
		return err
	}

	miniIndexs := formatTrieTree(ctx, indexInfoParams.IndexFormat, indexInfoParams)

	// 将数据加入倒排索引
	BuildTrieTreeBySet(ctx, indexInfoParams.IndexKey, miniIndexs)

	return err
}

// 建立倒排索引前，统一处理数据格式
func formatTrieTree(ctx context.Context, formatType string, indexInfo IndexNeedInfo) (m map[string][]interface{}) {
	miniIndexs := map[string][]interface{}{}
	switch formatType {
	case global.INDEX_FORMAT_JSON:
		for _, v := range indexInfo.DataMap {
			if v == nil {
				break
			}
			for _, miniList := range indexInfo.Mini {
				miniIndex := ""
				for _, m := range miniList {
					if _, ok := v[m]; !ok {
						break
					}
					miniIndex = fmt.Sprintf("%s%s", miniIndex, v[m])
				}
				miniIndexs[miniIndex] = append(miniIndexs[miniIndex], v)
			}
		}
	case global.INDEX_FORMAT_ARRAY:
		for _, v := range indexInfo.DataList {
			miniIndexs[v] = nil
		}
	}
	return miniIndexs
}

func getIndexInfo(ctx context.Context, conf config.IndexConf) (info IndexNeedInfo, err error) {
	//获得构建倒排的字段
	var indexConf IndexConf
	if _, err := toml.DecodeFile(conf.Conf, &indexConf); err != nil {
		err = fmt.Errorf("read toml has err %s", err.Error())
		return IndexNeedInfo{}, err
	}

	resInfos := IndexNeedInfo{
		IndexConf: indexConf,
	}

	// 如果需要使用lru
	if indexConf.LRUSize != 0 {
		lru.LRUInit(ctx, indexConf.IndexKey, indexConf.LRUSize, indexConf.LRUTime)
	}

	// 如果是网络文件，需要将文件下载到本地
	/*
		if indexConf.IndexType == global.INDEX_RESOURCE_TYPE_NET {
			// todo 目前仅支持将数据scp到机器上，后续支持上传或其他的网络方式
		}
	*/

	// 如果是本地json
	/*if indexConf.IndexType == global.INDEX_RESOURCE_TYPE_LOCAL && indexConf.IndexFormat == global.INDEX_FORMAT_JSON {
		//获得构建倒排的数据
		data, err := readFileToJsonMap(conf.DataConf)
		if nil != err {
			err = fmt.Errorf("readFileToJsonMap has err %s", err.Error())
			return IndexNeedInfo{}, err
		}
		resInfos.DataMap = data
	} else if indexConf.IndexType == global.INDEX_RESOURCE_TYPE_LOCAL && indexConf.IndexFormat == global.INDEX_FORMAT_ARRAY {
		//获得构建倒排的数据
		arrList, err := readFileToStringList(ctx, conf.DataConf)
		if nil != err {
			err = fmt.Errorf("readFileToStringList has err %s", err.Error())
			return IndexNeedInfo{}, err
		}
		for _, v := range arrList {
			resInfos.DataList = append(resInfos.DataList, v)
		}
	}*/
	if indexConf.IndexFormat == global.INDEX_FORMAT_JSON {
		//获得构建倒排的数据
		data, err := readFileToJsonMap(conf.DataConf)
		if nil != err {
			err = fmt.Errorf("readFileToJsonMap has err %s", err.Error())
			return IndexNeedInfo{}, err
		}
		resInfos.DataMap = data
	} else if indexConf.IndexFormat == global.INDEX_FORMAT_ARRAY {
		//获得构建倒排的数据
		arrList, err := readFileToStringList(ctx, conf.DataConf)
		if nil != err {
			err = fmt.Errorf("readFileToStringList has err %s", err.Error())
			return IndexNeedInfo{}, err
		}
		/*for _, v := range arrList {
			resInfos.DataList = append(resInfos.DataList, v)
		}*/
		resInfos.DataList = append(resInfos.DataList, arrList...)
	}

	return resInfos, nil
}
