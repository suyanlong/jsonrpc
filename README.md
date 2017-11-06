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
    ```

##### 其他

* 项目结构
    
    ----
        |- httpserver

        |- rpc

        |- mqserver

        |- snappy

        |- libproto

        |- example


##### Benchmark
* 发送获取高度:cita_blockNumber
* jsonrpc Go版本的监听地址:http://127.0.0.1:1337
* jsonrpc Rust版本的监听地址:http://127.0.0.1:1338
* 以下测试数据:

    ```
    wrk -c 300 -d 30 -t 3 -s post.lua http://127.0.0.1:1337
    Running 30s test @ http://127.0.0.1:1337
      3 threads and 300 connections
      Thread Stats   Avg      Stdev     Max   +/- Stdev
        Latency    63.82ms   22.43ms 361.83ms   97.60%
        Req/Sec     1.61k   280.06     2.02k    78.53%
      144420 requests in 30.04s, 24.24MB read
    Requests/sec:   4807.17
    Transfer/sec:    826.23KB

    wrk -c 300 -d 30 -t 3 -s post.lua http://127.0.0.1:1338
    Running 30s test @ http://127.0.0.1:1338
      3 threads and 300 connections
      Thread Stats   Avg      Stdev     Max   +/- Stdev
        Latency    59.35ms   11.87ms 144.15ms   68.41%
        Req/Sec     1.13k   211.36     1.72k    65.67%
      101083 requests in 30.07s, 11.38MB read
    Requests/sec:   3361.99
    Transfer/sec:    387.42KB

    -----------------------------------------------------------------------

    wrk -c 30 -d 30 -t 3 -s post.lua http://127.0.0.1:1337
    Running 30s test @ http://127.0.0.1:1337
      3 threads and 30 connections
      Thread Stats   Avg      Stdev     Max   +/- Stdev
        Latency     6.47ms    3.85ms 110.29ms   95.45%
        Req/Sec     1.63k   247.69     1.84k    94.78%
      145950 requests in 30.01s, 24.50MB read
    Requests/sec:   4863.81
    Transfer/sec:    835.97KB


    wrk -c 30 -d 30 -t 3 -s post.lua http://127.0.0.1:1338
    Running 30s test @ http://127.0.0.1:1338
      3 threads and 30 connections
      Thread Stats   Avg      Stdev     Max   +/- Stdev
        Latency    31.96ms    6.45ms  78.38ms   78.28%
        Req/Sec   314.25     36.15   400.00     73.67%
      28181 requests in 30.03s, 3.17MB read
    Requests/sec:    938.40
    Transfer/sec:    108.14KB

    ```