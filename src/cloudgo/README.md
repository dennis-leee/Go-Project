# 开发 web 服务程序

---

## **一. 概述**

> 开发简单 web 服务程序 cloudgo，了解 web 服务器工作原理。

### **任务目标**

>  1. 熟悉 go 服务器工作原理
>  2. 基于现有 web 库，编写一个简单 web 应用类似 [cloudgo][1]。
>  3. 使用 curl 工具访问 web 程序
>  4. 对 web 执行压力测试

## **二. 相关说明**

 - 本项目是一个简单的[beego][2]应用，需安装该框架才能正常运行，此外推荐配合[bee工具][3]使用
 - 安装好bee工具后，在命令行模式下进入项目根目录，输入`bee run`即可运行。
 - 默认采用8080端口，通过`localhost:8080`进行访问
 - 主要有两个页面:
 一个beego主页（`localhost:8080/`）
 一个app页（`localhost:8080/puzzle`）
 前者是用beego框架默认生成的，后者是本人之前web课程的一个作业。

## **三. curl测试**

### **1. 简介**

> curl 是 linux 系统自带的一个命令行工具，是 web 开发者最常用的利器。它可以精确控制 HTTP 请求的每一个细节。实战中，配合
> shell 程序，我们可以简单，重复给服务器发送不同的请求序列，调试程序或分析输出。

### **2. 结果**

 - 服务端
```
lkf@ubuntu:~/桌面/go/myProject/src/cloudgo$ bee run
______
| ___ \
| |_/ /  ___   ___
| ___ \ / _ \ / _ \
| |_/ /|  __/|  __/
\____/  \___| \___| v1.9.1
2017/11/14 23:55:21 INFO     ▶ 0001 Using 'cloudgo' as 'appname'
2017/11/14 23:55:21 INFO     ▶ 0002 Initializing watcher...
cloudgo/controllers
cloudgo/routers
cloudgo
2017/11/14 23:55:24 SUCCESS  ▶ 0003 Built Successfully!
2017/11/14 23:55:24 INFO     ▶ 0004 Restarting 'cloudgo'...
2017/11/14 23:55:24 SUCCESS  ▶ 0005 './cloudgo' is running...
2017/11/14 23:55:24 [I] [asm_amd64.s:2337] http server Running on http://:8080
2017/11/14 23:55:37 [D] [server.go:2619] |      127.0.0.1| 200 |   2.816321ms|   match| GET      /     r:/
2017/11/14 23:57:26 [D] [server.go:2619] |      127.0.0.1| 200 |    3.00066ms|   match| GET      /     r:/
2017/11/14 23:58:44 [D] [server.go:2619] |      127.0.0.1| 200 |   8.812095ms|   match| GET      /     r:/
2017/11/15 00:05:02 [D] [server.go:2619] |      127.0.0.1| 200 |   2.666575ms|   match| GET      /puzzle   r:/puzzle
2017/11/15 00:06:00 [D] [server.go:2619] |      127.0.0.1| 200 |   4.124325ms|   match| GET      /puzzle   r:/puzzle
```

 - 客户端：

    localhost:8080：
```
lkf@ubuntu:~/桌面$ curl localhost:8080 > out.txt
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 70163  100 70163    0     0  4446k      0 --:--:-- --:--:-- --:--:-- 4567k
```
    out.txt：
```
<!DOCTYPE html>
<html>
...  //内容有点多，且因采用的是模板，基本全是16进制内容，就不贴上来了
</html>
```
    localhost:8080/puzzle：
```
lkf@ubuntu:~/桌面$ curl localhost:8080/puzzle
<!DOCTYPE HTML>
<html>
    <head>
        <title>Fifteen Puzzle</title>
        <meta http-equiv="content-type" content="text/html; charset=utf-8" />
        <link href="./../static/img/favicon.png" sizes="196x196" rel="icon" />
        <link href="./../static/css/puzzle.css" type="text/css" rel="stylesheet" />
        <script src="./../static/js//puzzle.js" type="text/javascript"></script>
    </head>
    
    <body>
        <div class="myGame">
            
            <div class="header">
                <h1>Fifteen Puzzle</h1>
            </div>
            
            <div class="state">
                <p>Time: <input type="text" readOnly="readonly"  value="0s" id="time" /></p>
                <p>State: <input type="text" readOnly="readonly" value="Welcome" id="state" /></p>
                <p>Steps: <input type="text" readOnly="readonly" value="0" id="steps" /></p>
            </div>
            
            <div class="operating_area">
                <div id="game"></div>
                
                <div class="menu">
                    <button id="start">Start Game</button>
                    <button id="resetting">Resetting</button>
                </div>
            </div>
            
            <div class="show_area">
                <h2>original picture</h2>
                <img src="./../static/img/miku.png" alt="miku" />
            </div>
            
            <div class="tips">
                <h2>tips:</h2>
                <ul>
                    <li>The "Time" for the time you have been using, the "Steps" show the number of steps you walk now.</li>
                    <li>Click "Resetting" to disrupt the puzzle until one you like, and then click "Start Game" to start the game, or click the "Start Game" to play directly.</li>
                    <li>After the game starts, the "Start game" button will become "Stop Game", and if you click it, the game will stop.</li>
                    <li>After the game starts, the "Resetting" button is banned.</li>
                </ul>
            </div>
        </div>
    </body>
</html>
```
###  **3. 解析**
从上面可以看出curl其实扮演的就是我们日常生活中使用的浏览器的角色，只不过浏览器将取得的文档经过各自的渲染后再呈现在与用户眼前，而curl则是直接将原文档呈现出来。

## **四. 压力测试**
### **1.简介**

> ab是Apache超文本传输协议(HTTP)的性能测试工具。其设计意图是描绘当前所安装的Apache的执行性能，主要是显示你安装的Apache每秒可以处理多少个请求。

### **2.结果**
```
lkf@ubuntu:~/桌面$ ab -n 10000 -c 1000 localhost:8080/puzzle
This is ApacheBench, Version 2.3 <$Revision: 1528965 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        beegoServer:1.9.0  //服务器软件
Server Hostname:        localhost  //服务器主机名
Server Port:            8080  //服务器端口号

Document Path:          /puzzle    //请求的资源
Document Length:        2142 bytes  //文档返回的长度，不包括相应头

Concurrency Level:      1000  //并发用户数
Time taken for tests:   19.338 seconds  //总请求时间
Complete requests:      10000  //总请求数
Failed requests:        0  //失败的请求数
Total transferred:      22870000 bytes
HTML transferred:       21420000 bytes
Requests per second:    517.11 [#/sec] (mean)  //吞吐率，即平均每秒的请求数
Time per request:       1933.826 [ms] (mean)  //用户平均请求等待时间，即平均每个请求消耗的时间
Time per request:       1.934 [ms] (mean, across all concurrent requests)  //服务器平均请求等待时间，即上面的请求除以并发数
Transfer rate:          1154.91 [Kbytes/sec] received  //传输速率

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   98 469.9      0    3005
Processing:    15 1749 1699.7   1221    4635
Waiting:       15 1748 1699.6   1221    4635
Total:         15 1847 1801.2   1225    6944

Percentage of the requests served within a certain time (ms)
  50%   1225  //50%的请求都在1225ms内完成
  66%   2993
  75%   3692
  80%   3804
  90%   4241
  95%   4469
  98%   4573
  99%   6415
 100%   6944 (longest request)
```

### **3.解析**
这里我一共发送了10000个数据请求（参数n），并发量（参数c）为1000，从结果中可以看出全部请求均已成功响应，50%的请求都在1225ms内完成。

  [1]: http://blog.csdn.net/pmlpml/article/details/7840483
  [2]: https://beego.me/docs/intro/
  [3]: https://beego.me/docs/install/bee.md