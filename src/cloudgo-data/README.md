# 用xorm构建golang数据服务

---

> 详细文档：http://blog.csdn.net/pmlpml/article/details/78602290

## **一、概述**
> - 本部分的目标是使用 golang database/sql 写出易于阅读、扩展和可维护的数据库服务。重点是掌握经典的 “entity - dao - service” 层次结构编程模型
> - [orm][1]是个复杂的话题，其反射技术、牺牲性能获得易用性。如果你做面向程序猿的系统应用而不是面向客户的应用，database/sql 是你的第一选择；相反，**orm 可以让你获得开发效率，orm 使得你不需要编写 dao 服务！**

## **二、任务内容**
> 1. 使用 xorm 或 gorm 实现本文的程序，从编程效率、程序结构、服务性能等角度对比 database/sql 与 orm 实现的异同！ 
> - orm 是否就是实现了 dao 的自动化？
> - 使用 ab 测试性能
> 2. 参考 Java JdbcTemplate 的设计思想，设计 GoSqlTemplate 的原型, 使得 sql 操作对于爱写 sql 的程序猿操作数据库更容易。 
> - 轻量级别的扩展，程序员的最爱
> - 程序猿不怕写 sql ，怕的是线程安全处理和错误处理
> - sql 的 CRUD 操作 database/sql 具有强烈的模板特征，适当的回调可以让程序员自己编写 sql 语句和处理 RowMapping
> - 建立在本文 SQLExecer 接口之上做包装，直观上是有利的选择
> - 暂时不用考虑占位符等数据库移植问题，方便使用 mysql 或 sqlite3 就可以
> - 参考资源：github.com/jmoiron/sqlx


## **三、 相关说明**

 - 本程序暂时仅实现（采用[xorm][2]）了第一部分的内容
 - 使用以下命令获取源码，
  `git clone https://github.com/my937889621/Go-Project.git`  
  或者直接下载zip文件并解压放到`$GOPATH/github.com/my937889621/`下，
  而后进入项目文件夹（cloudgo-data），执行`go run main.go`即可
 - 请务必使用[上文][3]的数据库设置。
 - 需将将`userinfo`表单的`created`字段更名为`createat`
## **三. 结果**
### **1. ab压力测试**
 - 往数据库里填入了33条数据，用FindAll指令进行测试。

```
lkf@ubuntu:~$ ab -n 10000 -c 1000 http://localhost:8080/service/userinfo?userid=
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


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /service/userinfo?userid=
Document Length:        4078 bytes

Concurrency Level:      1000
Time taken for tests:   16.838 seconds
Complete requests:      10000
Failed requests:        4188
   (Connect: 0, Receive: 0, Length: 4188, Exceptions: 0)
Non-2xx responses:      4188
Total transferred:      42922520 bytes
HTML transferred:       41838076 bytes
Requests per second:    593.91 [#/sec] (mean)
Time per request:       1683.755 [ms] (mean)
Time per request:       1.684 [ms] (mean, across all concurrent requests)
Transfer rate:          2489.47 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0  272 746.8      0    7019
Processing:    36 1299 694.9   1176    5598
Waiting:        1  724 622.9    564    3324
Total:         36 1571 1076.0   1435    9213

Percentage of the requests served within a certain time (ms)
  50%   1435
  66%   1856
  75%   2057
  80%   2177
  90%   2702
  95%   3505
  98%   4974
  99%   5179
 100%   9213 (longest request)
```
### **2. 总结**

 1. 从编程效率的角度来看，orm比database/sql要好，其提供了齐全的api，几乎可以实现数据库的全部操作。
 2. 从程序结构的角度来看，orm更符合“entity - dao - service” 层次结构编程模型， service层几乎就只是调用dao层的实现就行了，而dao层的实现也仅包含少量的逻辑判断，其余操作均由接口来完成。
 3. 从程序结构的角度来看，orm就不如database/sql了，因为其可以说是database/sql的一层封装，使用起来是很便利，但多一层封装，性能自然也就下降了。从上面的压力测试结果也可以看得出来：发送了10000条请求，其中有4188条失败了，近乎一半，这个数据还是有点大的。


  [1]: http://blog.csdn.net/zhanghongjie0302/article/details/47344417
  [2]: http://xorm.io/
  [3]: http://blog.csdn.net/pmlpml/article/details/78602290
