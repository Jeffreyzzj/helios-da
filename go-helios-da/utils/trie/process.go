package trie

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"go-helios-da/config"
	"go-helios-da/global"
	"os"
)

func (t *TrieTreeUtil) PopTrieRoot(ctx context.Context) *map[string]*TrieTree {
	return TrieRootMap
}

func (t *TrieTreeUtil) TrieRootInit(ctx context.Context) (err error) {
	// 读取需要处理的conf文件
	var indexConf config.TomlConfig
	filePath := "./da_conf/helios_da_conf.toml"
	if _, err := toml.DecodeFile(filePath, &indexConf); err != nil {
		err = fmt.Errorf("read toml has err %s", err.Error())
		return err
	}

	// 取出IndexConfigs列表，开始准备构建相应的倒排索引
	for _, v := range indexConf.HeliosInitConfig.IndexConfigs {
		err = buildIndexByIndexConf(ctx, v)
		if nil != err {
			fmt.Println("BuildIndexByIndexConf key[%s] has err[%s]", v.Conf, err.Error())
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
	return
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
	return
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
	sugList, err := SugBySubWord(ctx, index, []rune(subQuery), maxNum)
	if nil != err {
		err = fmt.Errorf("SugBySubWord has error %s", err.Error())
		return list, err
	}
	for _, v := range sugList {
		list = append(list, string(v))
	}

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

func buildIndexByIndexConf(ctx context.Context, conf config.IndexConf) (err error) {
	indexInfoParams, err := getIndexInfo(ctx, conf)
	if nil != err {
		err = fmt.Errorf("getIndexInfo has err %s ", err.Error())
		return err
	}

	miniIndexs := map[string][]interface{}{}
	for _, miniList := range indexInfoParams.Mini {
		// 本地文件
		if indexInfoParams.IndexType == global.INDEX_TYPE_LOCAL {
			for _, v := range indexInfoParams.Data {
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
	}

	// 将数据加入倒排索引
	BuildTrieTreeBySet(ctx, indexInfoParams.IndexKey, miniIndexs)

	return err
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

	if indexConf.IndexType == global.INDEX_TYPE_LOCAL {
		//获得构建倒排的数据
		content, err := os.ReadFile(conf.DataConf)
		if err != nil {
			err = fmt.Errorf("read dataConf has err %s", err.Error())
			return IndexNeedInfo{}, err
		}
		var dataList []map[string]interface{}
		err = json.Unmarshal(content, &dataList)
		resInfos.Data = dataList
	}

	return resInfos, nil
}
