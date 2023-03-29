# helios-da

## 介绍
简介：简单、轻量、安全、可自定义的搜索引擎
描述：
1. 小型搜索引擎，目前支持基础的搜索引擎功能
2. 支持自定义配置LRU，加速相关查询
3. 与ES等大型开源搜索引擎相比，helios-da没有最低的机器配置要求
4. 无插件依赖、让搜索引擎简单快速

## 安装
+ 源代码方式
   1. 配置要求：
    > + golang：12+（建议使用18+，以支持helios-da后续的范型迭代）
    > + 建议Linux机器环境
   2. 克隆此仓库：
    > + git clone git@github.com:Jeffreyzzj/helios-da.git
   3. 进入项目目录（以下简称根目录）: 
    > + cd ./go-helios-da
   4. 执行以下命令，进行检查和更新适合自己机器的相关配置。
    > + go mod tidy
+ 可执行文件
   1. linux:
   2. mac:
   3. windows:

## 使用
1. 配置helios-da，配置文件在根目录：
  > ./da_conf/helios_da_conf.toml
2. 添加自己服务的相关内容，建议放在：./da_conf/search/
3. 配置文件包括两部分
   + 索引配置文件 
   ```toml
        # 倒排索引index主键
        index_key = "music"
        # 倒排索引文件来源类型：目前包含net,local
        index_type = "local"
        # 倒排索引数据的格式目前包含array,json
        index_format = "json"
        # 倒排索引对于LRU的需求,当前索引是否使用lru,如果size为0，表示不使用
        index_lru_size = 5
        # 倒排索引对于LRU的需求,当前索引是否使用lru
        index_lru_time = 5
        # 数据中需要构建倒排索引的字段
        mini = [["title","author"], ["author", "title"]]
   ```
   + 索引数据文件：配置文件index_format = "json"举例
    ```toml
        [
            {
                "loc": "abcdefg1",
                "title": "缘分一道桥",
                "author": "王力宏",
                "fileName": "缘分一道桥-王力宏-ab3ceel.mp3"
            },
            {
                "loc": "abcdef2g",
                "title": "大城小爱",
                "author": "王力宏",
                "fileName": "大城小爱-王力宏-abce2el.mp3"
            },
            {
                "loc": "abcde3fc",
                "title": "大城小爱",
                "author": "王力宏",
                "fileName": "大城小爱-王力宏-现场版-abce2el.mp3"
            },
            {
                "loc": "abcde4fg",
                "title": "浮夸",
                "author": "陈奕迅",
                "fileName": "浮夸-陈奕迅-abcee1l.mp3"
            }
        ]
   ```
   + 索引数据文件：配置文件index_format = "array"举例
   ```toml
        藏红花
        曾舜晞
        曾子杰
        王者荣耀
        闻王昌龄左迁龙标遥有此寄
        枣糕图片
        中国人事考试网
        王楚然
   ```
   
4. 运行
+ 源代码
  > - 方式一：在根目录下运行: go run main.go
  > - 方式二：修改 start.sh 进行启动

+ linux
  > - 方式一：执行 ./go-helios-da
  > - 方式二：修改 start.sh 进行启动

5. 压测
+ mac版本压测 
   ```
    ab -n 16000 -c 200 -s 60 "http://127.0.0.1:9609/helios/sugQ?query=%E6%97%A9%E5%AE%89&index=test"

    Server Software:
    Server Hostname:        127.0.0.1
    Server Port:            9609

    Document Path:          /helios/sugQ?query=%E6%97%A9%E5%AE%89&index=test
    Document Length:        61 bytes

    Concurrency Level:      200
    Time taken for tests:   2.100 seconds
    Complete requests:      16000
    Failed requests:        0
    Total transferred:      8160000 bytes
    HTML transferred:       976000 bytes
    Requests per second:    7618.37 [#/sec] (mean)
    Time per request:       26.252 [ms] (mean)
    Time per request:       0.131 [ms] (mean, across all concurrent requests)
    Transfer rate:          3794.30 [Kbytes/sec] received

    Connection Times (ms)
                  min  mean[+/-sd] median   max
    Connect:        0    0   1.0      0      18
    Processing:     0   26  12.8     25     117
    Waiting:        0   25  12.7     25     117
    Total:          0   26  12.7     25     117

    Percentage of the requests served within a certain time (ms)
      50%     25
      66%     27
      75%     29
      80%     31
      90%     40
      95%     51
      98%     58
      99%     65
      100%    117 (longest request)
   ```
+ linux版本压测(red hot: 4he)
   ```
    ab -n 16000 -c 200 -s 60 http://127.0.0.1:9609/helios/sugQ?query=%E6%9B%BE&index=test

    Server Software:        
    Server Hostname:        127.0.0.1
    Server Port:            9609

    Document Path:          /helios/sugQ?query=%E6%9B%BE
    Document Length:        31 bytes

    Concurrency Level:      200
    Time taken for tests:   1.756 seconds
    Complete requests:      16000
    Failed requests:        0
    Write errors:           0
    Total transferred:      7680000 bytes
    HTML transferred:       496000 bytes
    Requests per second:    9109.94 [#/sec] (mean)
    Time per request:       21.954 [ms] (mean)
    Time per request:       0.110 [ms] (mean, across all concurrent requests)
    Transfer rate:          4270.28 [Kbytes/sec] received

    Connection Times (ms)
                  min  mean[+/-sd] median   max
    Connect:        0    4   2.5      4      15
    Processing:     0   18  38.2      6     217
    Waiting:        0   16  38.2      4     215
    Total:          0   22  38.2     10     222

    Percentage of the requests served within a certain time (ms)
      50%     10
      66%     12
      75%     13
      80%     14
      90%     81
      95%     120
      98%     170
      99%     213
      100%    222 (longest request)
   ```


## 功能

列出该项目的功能列表。

## 应用
+ 在[索需搜索<suoxu>](http://suoxu.top)中已使用
