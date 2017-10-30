
wrk -c 100 -d 10 -t 4 -s post.lua http://127.0.0.1:1337
Running 10s test @ http://127.0.0.1:1337
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    28.92ms   14.37ms 226.35ms   93.45%
    Req/Sec     0.91k   183.10     1.17k    85.35%
  36056 requests in 10.01s, 6.05MB read
Requests/sec:   3600.67
Transfer/sec:    618.86KB


wrk -c 100 -d 10 -t 4 -s post.lua http://127.0.0.1:1338
Running 10s test @ http://127.0.0.1:1338
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    37.09ms   10.63ms 143.45ms   85.20%
    Req/Sec   681.27    117.73   840.00     82.75%
  27154 requests in 10.02s, 3.06MB read
Requests/sec:   2710.40
Transfer/sec:    312.33KB

wrk -c 100 -d 10 -t 4 -s post.lua http://127.0.0.1:1339
Running 10s test @ http://127.0.0.1:1339
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    34.94ms    8.06ms 128.98ms   82.11%
    Req/Sec   721.23     93.49     0.89k    80.25%
  28738 requests in 10.01s, 3.23MB read
Requests/sec:   2871.42
Transfer/sec:    330.89KB


wrk -c 100 -d 10 -t 4 -s post.lua http://127.0.0.1:1337
Running 10s test @ http://127.0.0.1:1337
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    26.35ms    5.39ms  74.06ms   79.00%
    Req/Sec     0.95k   103.24     1.19k    75.50%
  37947 requests in 10.01s, 6.37MB read
Requests/sec:   3790.59
Transfer/sec:    651.51KB

wrk -c 100 -d 20 -t 4 -s post.lua http://127.0.0.1:1337
Running 20s test @ http://127.0.0.1:1337
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    26.29ms    5.28ms  62.66ms   79.02%
    Req/Sec     0.96k   104.14     1.22k    73.62%
  76081 requests in 20.02s, 12.77MB read
Requests/sec:   3800.95
Transfer/sec:    653.29KB

wrk -c 100 -d 10 -t 8 -s post.lua http://127.0.0.1:1337
Running 10s test @ http://127.0.0.1:1337
  8 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    25.09ms    6.03ms  65.86ms   84.41%
    Req/Sec   480.31     99.50     2.67k    90.76%
  38336 requests in 10.10s, 6.43MB read
Requests/sec:   3796.39
Transfer/sec:    652.50KB

wrk -c 100 -d 10 -t 8 -s post.lua http://127.0.0.1:1338
Running 10s test @ http://127.0.0.1:1338
  8 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    34.50ms    6.68ms  79.81ms   74.27%
    Req/Sec   348.53     43.82   460.00     75.50%
  28067 requests in 10.09s, 3.16MB read
Requests/sec:   2781.03
Transfer/sec:    320.47KB

wrk -c 100 -d 10 -t 8 -s post.lua http://127.0.0.1:1339
Running 10s test @ http://127.0.0.1:1339
  8 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    33.99ms    6.88ms  90.08ms   78.68%
    Req/Sec   354.27     38.00   450.00     73.12%
  28256 requests in 10.02s, 3.18MB read
Requests/sec:   2819.86
Transfer/sec:    324.94KB

wrk -c 100 -d 10 -t 8 -s post.lua http://127.0.0.1:1340
Running 10s test @ http://127.0.0.1:1340
  8 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    35.49ms    7.84ms  84.41ms   76.33%
    Req/Sec   339.01     45.71   470.00     75.00%
  27020 requests in 10.02s, 3.04MB read
Requests/sec:   2696.37
Transfer/sec:    310.71KB

wrk -c 1000 -d 10 -t 8 -s post.lua http://127.0.0.1:1340
Running 10s test @ http://127.0.0.1:1340
  8 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    69.40ms   12.40ms 126.68ms   73.34%
    Req/Sec   411.49    268.33     1.17k    68.43%
  28712 requests in 10.08s, 3.23MB read
Requests/sec:   2848.57
Transfer/sec:    328.25KB

wrk -c 1000 -d 10 -t 16 -s post.lua http://127.0.0.1:1340
Running 10s test @ http://127.0.0.1:1340
  16 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    66.68ms   12.06ms 120.30ms   69.48%
    Req/Sec   500.98    349.28     1.24k    55.54%
  30231 requests in 10.09s, 3.40MB read
Requests/sec:   2996.10
Transfer/sec:    345.25KB

wrk -c 1000 -d 10 -t 16 -s post.lua http://127.0.0.1:1337
Running 10s test @ http://127.0.0.1:1337
  16 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   234.18ms   56.81ms 631.97ms   84.22%
    Req/Sec   270.36    194.78     1.10k    57.90%
  38105 requests in 10.08s, 6.40MB read
Requests/sec:   3778.64
Transfer/sec:    649.45KB

wrk -c 1000 -d 10 -t 16 -s post.lua http://127.0.0.1:1337
Running 10s test @ http://127.0.0.1:1337
  16 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   257.34ms   28.37ms 390.61ms   87.28%
    Req/Sec   243.81    181.21   747.00     57.03%
  37313 requests in 10.08s, 6.26MB read
Requests/sec:   3702.78
Transfer/sec:    636.42KB

wrk -c 1000 -d 10 -t 16 -s post.lua http://127.0.0.1:1337
Running 10s test @ http://127.0.0.1:1337
  16 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   253.83ms   21.14ms 306.08ms   83.34%
    Req/Sec   248.29    175.14   620.00     60.13%
  38897 requests in 10.10s, 6.53MB read
Requests/sec:   3852.81
Transfer/sec:    662.20KB

wrk -c 100 -d 10 -t 16 -s post.lua http://127.0.0.1:1337
Running 10s test @ http://127.0.0.1:1337
  16 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    24.92ms    6.04ms  77.43ms   83.41%
    Req/Sec   241.79     36.43   590.00     76.28%
  38694 requests in 10.10s, 6.49MB read
Requests/sec:   3831.33
Transfer/sec:    658.51KB

wrk -c 100 -d 10 -t 32 -s post.lua http://127.0.0.1:1337
Running 10s test @ http://127.0.0.1:1337
  32 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    24.55ms    5.85ms  73.63ms   83.89%
    Req/Sec   122.78     32.92     1.31k    98.81%
  39192 requests in 10.10s, 6.58MB read
Requests/sec:   3880.75
Transfer/sec:    667.00KB

wrk -c 100 -d 10 -t 32 -s post.lua http://127.0.0.1:1337
Running 10s test @ http://127.0.0.1:1337
  32 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    24.74ms    6.61ms  76.48ms   84.08%
    Req/Sec   122.07     18.67   212.00     68.81%
  39136 requests in 10.10s, 6.57MB read
Requests/sec:   3874.61
Transfer/sec:    665.95KB

wrk -c 100 -d 10 -t 32 -s post.lua http://127.0.0.1:1338
Running 10s test @ http://127.0.0.1:1338
  32 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    34.70ms    7.55ms  84.79ms   78.80%
    Req/Sec    86.51     13.24   121.00     63.51%
  27724 requests in 10.10s, 3.12MB read
Requests/sec:   2745.04
Transfer/sec:    316.32KB

wrk -c 300 -d 10 -t 32 -s post.lua http://127.0.0.1:1338
Running 10s test @ http://127.0.0.1:1338
  32 threads and 300 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    65.77ms   12.62ms 141.24ms   72.16%
    Req/Sec   127.03     47.82   282.00     69.92%
  30371 requests in 10.10s, 3.42MB read
Requests/sec:   3007.08
Transfer/sec:    346.52KB

wrk -c 32 -d 10 -t 32 -s post.lua http://127.0.0.1:1338
Running 10s test @ http://127.0.0.1:1338
  32 threads and 32 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    31.99ms    6.54ms  76.96ms   79.11%
    Req/Sec    31.26      4.96    40.00     74.73%
  10023 requests in 10.10s, 1.13MB read
Requests/sec:    992.50
Transfer/sec:    114.37KB

wrk -c 32 -d 10 -t 32 -s post.lua http://127.0.0.1:1337
Running 10s test @ http://127.0.0.1:1337
  32 threads and 32 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     8.70ms    4.16ms  56.31ms   88.05%
    Req/Sec   119.46     17.86   171.00     76.65%
  38291 requests in 10.10s, 6.43MB read
Requests/sec:   3791.11
Transfer/sec:    651.60KB

wrk -c 32 -d 10 -t 32 -s post.lua http://127.0.0.1:1338
Running 10s test @ http://127.0.0.1:1338
  32 threads and 32 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    32.29ms    7.71ms  93.10ms   84.13%
    Req/Sec    31.06      5.46    40.00     71.06%
  9944 requests in 10.02s, 1.12MB read
Requests/sec:    992.77
Transfer/sec:    114.40KB

wrk -c 3200 -d 10 -t 32 -s post.lua http://127.0.0.1:1338
Running 10s test @ http://127.0.0.1:1338
  32 threads and 3200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    67.30ms   14.13ms 137.91ms   74.87%
    Req/Sec   596.61    339.50     1.63k    60.75%
  30134 requests in 10.09s, 3.39MB read
Requests/sec:   2987.86
Transfer/sec:    344.30KB

wrk -c 3200 -d 10 -t 32 -s post.lua http://127.0.0.1:1339
Running 10s test @ http://127.0.0.1:1339
  32 threads and 3200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    69.44ms   12.30ms 151.74ms   71.99%
    Req/Sec   578.08    286.07     1.32k    63.04%
  29197 requests in 10.10s, 3.29MB read
Requests/sec:   2891.15
Transfer/sec:    333.16KB

wrk -c 32 -d 10 -t 32 -s post.lua http://127.0.0.1:1337
Running 10s test @ http://127.0.0.1:1337
  32 threads and 32 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     8.63ms    4.30ms  59.85ms   88.77%
    Req/Sec   120.77     33.10     0.93k    93.45%
  38572 requests in 10.10s, 6.47MB read
Requests/sec:   3819.17
Transfer/sec:    656.42KB


wrk -c 3200 -d 10 -t 32 -s post.lua http://127.0.0.1:1337
Running 10s test @ http://127.0.0.1:1337
  32 threads and 3200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   768.30ms  189.01ms   1.67s    81.48%
    Req/Sec   146.10    168.14     0.99k    85.51%
  37916 requests in 10.09s, 6.36MB read
  Socket errors: connect 0, read 1, write 0, timeout 0
Requests/sec:   3756.43
Transfer/sec:    645.64KB

wrk -c 320 -d 10 -t 4 -s post.lua http://127.0.0.1:1337
Running 10s test @ http://127.0.0.1:1337
  4 threads and 320 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    82.21ms   14.97ms 172.71ms   90.52%
    Req/Sec     0.98k   226.70     1.55k    76.13%
  38797 requests in 10.04s, 6.51MB read
Requests/sec:   3865.16
Transfer/sec:    664.32KB


wrk -c 30 -d 10 -t 4 -s post.lua http://127.0.0.1:1337
Running 10s test @ http://127.0.0.1:1337
  4 threads and 30 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     7.67ms    4.09ms  57.93ms   90.72%
    Req/Sec     0.96k   151.78     1.17k    87.75%
  38158 requests in 10.01s, 6.40MB read
Requests/sec:   3812.49
Transfer/sec:    655.27KB


wrk -c 30 -d 10 -t 4 -s post.lua http://127.0.0.1:1338
Running 10s test @ http://127.0.0.1:1338
  4 threads and 30 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    32.37ms    7.53ms  89.69ms   82.36%
    Req/Sec   217.29     33.67   282.00     74.50%
  8663 requests in 10.02s, 0.97MB read
Requests/sec:    864.35
Transfer/sec:     99.60KB


wrk -c 30 -d 10 -t 4 -s post.lua http://127.0.0.1:1337
Running 10s test @ http://127.0.0.1:1337
  4 threads and 30 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     7.00ms    3.40ms  46.83ms   85.68%
    Req/Sec     0.93k   155.03     1.40k    76.25%
  36920 requests in 10.01s, 6.20MB read
Requests/sec:   3689.64
Transfer/sec:    634.16KB


wrk -c 300 -d 10 -t 12 -s post.lua http://127.0.0.1:1338
Running 10s test @ http://127.0.0.1:1338
  12 threads and 300 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   113.17ms   16.78ms 212.29ms   72.34%
    Req/Sec   220.39     41.40   464.00     77.44%
  26380 requests in 10.10s, 2.97MB read
Requests/sec:   2612.30
Transfer/sec:    301.03KB

wrk -c 300 -d 10 -t 12 -s post.lua http://127.0.0.1:1337
Running 10s test @ http://127.0.0.1:1337
  12 threads and 300 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    81.13ms   16.98ms 241.43ms   88.75%
    Req/Sec   311.97     74.42   505.00     77.12%
  37289 requests in 10.07s, 6.26MB read
Requests/sec:   3703.30
Transfer/sec:    636.50KB


wrk -c 300 -d 10 -t 12 -s post.lua http://127.0.0.1:1337
Running 10s test @ http://127.0.0.1:1337
  12 threads and 300 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    78.91ms    9.26ms 142.70ms   75.80%
    Req/Sec   317.91     82.14   505.00     78.04%
  37884 requests in 10.04s, 6.36MB read
Requests/sec:   3774.14
Transfer/sec:    648.68KB


