package trie

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// 基础方法 索引index添加数据data
func addTrieTreeNode(index string, word []rune, data []interface{}) {
	if (*TrieRootMap)[index] == nil {
		(*TrieRootMap)[index] = &TrieTree{
			TrieMap: map[rune]*TrieTree{},
		}
	}
	root := (*TrieRootMap)[index]

	for _, v := range word {
		if _, ok := root.TrieMap[v]; !ok {
			root.TrieMap[v] = &TrieTree{
				TrieMap: map[rune]*TrieTree{},
				IsEnd:   false,
				Data:    nil,
			}
		}
		root = root.TrieMap[v]
	}
	root.IsEnd = true
	if root.Data == nil {
		root.Data = []*interface{}{}
	}
	trieTreeData := []*interface{}{}
	for i := 0; i < len(data); i++ {
		trieTreeData = append(trieTreeData, &data[i])
	}
	root.Data = append(root.Data, trieTreeData...)
}

// 基础方法 在新指针中添加数据data
func addTrieTreeNodeByNewRoot(root *TrieTree, word []rune, data []interface{}) {

	for _, v := range word {
		if _, ok := root.TrieMap[v]; !ok {
			root.TrieMap[v] = &TrieTree{
				TrieMap: map[rune]*TrieTree{},
				IsEnd:   false,
				Data:    nil,
			}
		}
		root = root.TrieMap[v]
	}
	root.IsEnd = true
	if root.Data == nil {
		root.Data = []*interface{}{}
	}
	trieTreeData := []*interface{}{}
	for i := 0; i < len(data); i++ {
		trieTreeData = append(trieTreeData, &data[i])
	}
	root.Data = append(root.Data, trieTreeData...)
}

// 基础方法 根据key获得node
func getNodeByKey(ctx context.Context, index string, words []rune) (node *TrieTree, err error) {
	if (*TrieRootMap)[index] == nil {
		err = fmt.Errorf("helios hasn't index[%s]", index)
		return nil, err
	}
	root := (*TrieRootMap)[index]

	for _, v := range words {
		if _, ok := root.TrieMap[v]; !ok {
			return nil, nil
		}
		root = root.TrieMap[v]
	}

	return root, nil
}

// 基础方法 在一个索引index中查找以subWord为开头的倒排索引
func SugBySubWord(ctx context.Context, index string, words []rune, maxNum int) (list []string, err error) {
	if (*TrieRootMap)[index] == nil {
		err = fmt.Errorf("helios hasn't index[%s]", index)
		return list, err
	}
	root := (*TrieRootMap)[index]
	// 找到当前的subWord节点
	for _, v := range words {
		if _, ok := root.TrieMap[v]; !ok {
			return list, nil
		}
		root = root.TrieMap[v]
	}

	// 当前root为最后一个节点，开始查询当前节点下的完整倒排索引
	dfsList := &[]string{}
	TrieTreeDFS(ctx, root, words, dfsList, maxNum)

	return *dfsList, nil
}

// 基础方法 在一个索引index中查找以subWord为开头的数据集
func sugDataListBySubWord(ctx context.Context, index string, words []rune, maxNum int) (list []*interface{}, err error) {
	if (*TrieRootMap)[index] == nil {
		err = fmt.Errorf("helios hasn't index[%s]", index)
		return list, err
	}
	root := (*TrieRootMap)[index]
	// 找到当前的subWord节点
	for _, v := range words {
		if _, ok := root.TrieMap[v]; !ok {
			return list, nil
		}
		root = root.TrieMap[v]
	}

	// 当前root为最后一个节点，开始查询当前节点下的完整倒排索引
	dfsList := &[]*interface{}{}
	getDataTrieTreeDFS(ctx, root, dfsList, maxNum)

	return *dfsList, nil
}

func TrieTreeDFS(ctx context.Context, root *TrieTree, words []rune, dfsList *[]string, maxNum int) {

	if root == nil {
		return
	} else if len(*dfsList) >= maxNum {
		return
	}

	if root.IsEnd {
		*dfsList = append(*dfsList, string(words))
	}

	for k, v := range root.TrieMap {
		root = v
		words = append(words, k)
		TrieTreeDFS(ctx, v, words, dfsList, maxNum)
		words = words[0 : len(words)-1]
		if len(*dfsList) >= maxNum {
			return
		}
	}
}

func getDataTrieTreeDFS(ctx context.Context, root *TrieTree, dfsList *[]*interface{}, maxNum int) {

	if root == nil {
		return
	} else if len(*dfsList) >= maxNum {
		return
	}

	if root.IsEnd {
		*dfsList = append(*dfsList, root.Data...)
	}

	for _, v := range root.TrieMap {
		root = v
		getDataTrieTreeDFS(ctx, v, dfsList, maxNum)
		if len(*dfsList) >= maxNum {
			return
		}
	}
}

// 以下为文件处理，后续考虑换到别的文件 start

// 按行读取文件, json
func readFileToStringList(ctx context.Context, path string) (list []string, err error) {
	//获得构建倒排的数据
	fileHanle, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return list, err
	}

	defer fileHanle.Close()

	reader := bufio.NewReader(fileHanle)

	var results []string
	// 按行处理txt
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		results = append(results, string(line))
	}
	return results, nil
}

// 直接将数据全部读出来，数组array
func readFileToJsonMap(path string) (dataMap []map[string]interface{}, err error) {
	//获得构建倒排的数据
	content, err := os.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("read dataConf has err %s", err.Error())
		return []map[string]interface{}{}, err
	}
	var data []map[string]interface{}
	err = json.Unmarshal(content, &data)
	return data, err
}

/*func getFileFromNetByUrlGet(ctx context.Context, url string) {
	imgPath := "C:\\Users\\Asche\\go\\src\\GoSpiderTest\\"
	imgUrl := "http://hbimg.b0.upaiyun.com/32f065b3afb3fb36b75a5cbc90051b1050e1e6b6e199-Ml6q9F_fw320"

	fileName := path.Base(imgUrl)


	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32 * 1024)


	file, err := os.Create(imgPath + fileName)
	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d", written)
}*/

// 以下为文件处理，后续考虑换到别的文件 end
