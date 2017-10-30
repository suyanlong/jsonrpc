# go jsonrpc for CITA


#### 安装依赖

* go 版本
    ```
    go version go1.9 linux/amd64
    ```

* 官方包依赖管理工具dep:
    ```
    go get -u github.com/golang/dep/cmd/dep
    ```

   模仿rust的包管理器实现的，目前未正式发布，预计go1.20版本发布。初始化项目命令：

    ```
    dep ensure
    ```

##### JSONRPC

* 模拟调试数据：

    ``` json
    curl -X POST --data '{"jsonrpc":"2.0","method":"cita_blockNumber","params":[],"id":74}' 127.0.0.1:1337 | jq

    curl -X POST --data '{"jsonrpc":"2.0","method":"peerCount","params":[],"id":74}' 127.0.0.1:1337 | jq

    curl -X POST --data '{"jsonrpc":"2.0","method":"cita_sendTransaction","params":["..."],"id":1}' 127.0.0.1:1337 | jq

    curl -X POST --data '{"jsonrpc":"2.0","method":"cita_getBlockByNumber","params":["0xF9", true],"id":1}' 127.0.0.1:1337 | jq

    ```

##### 其他

* 项目结构
    
    ----
        |- httpserver

        |- proto

        |- rpc

        |- mqserver

        |- example



