# go实现selpg

---
## 简介
selpg 是从文本输入选择页范围的实用程序。该输入可以来自作为最后一个命令行参数指定的文件，在没有给出文件名参数时也可以来自标准输入。

selpg 首先处理所有的命令行参数。在扫描了所有的选项参数（也就是那些以连字符为前缀的参数）后，如果 selpg 发现还有一个参数，则它会接受该参数为输入文件的名称并尝试打开它以进行读取。如果没有其它参数，则 selpg 假定输入来自标准输入。

## 使用说明
该程序有6个参数，其中两个（s,e）为必需参数，其余为可选（dd参数未实现），使用时按照如下格式在命令行中输入即可。

     selpg [-s start page(>= 1)] [-e end page(>= s)] [-l page length] [-f use form-feed (l and f can only choose one)] [-d printer destination] [input file]
        
              -d string
                	The name of the printer to use
              -e int
                	The end page of the extracted page range (default -1)
              -f	Whether used form-feed character to delimit the pages
              -l int
                	The number of rows per page (default 72)
              -s int
                	The start page of the extracted page range (default -1)

##设计说明
该程序主要可以分为3部分

 1. 获取输入
    采用flag包的函数flag.*Var()实现
 2. 解析参数
    主要判断参数格式是否合法
 3. 执行
    根据有无inputFile，从不同的流中读取输入
    而后再根据采用的分页模式（f、l）进行分页输出

##测试
测试文件为 testFile.相关结果可在./bin/out中查看

    testFile:
    123 456
    qwe ert
    zxc cvb
    asd
    dfg
    yui iiop
    asd456
    79234
    ok

 1. `$ selpg -s1 -e1 input_file`
 2. `$ selpg -s1 -e1 < input_file`
 3. `$ selpg -s10 -e20 input_file >output_file`
 4. `$ selpg -s10 -e20 input_file 2>error_file`
 5. `$ selpg -s10 -e20 input_file >output_file 2>error_file`
 6. `$ selpg -s10 -e20 input_file >output_file 2>/dev/null`
 7. `$ selpg -s10 -e20 input_file >/dev/null`
 8. `$ selpg -s10 -e20 input_file | other_command`
 9. `$ selpg -s10 -e20 input_file 2>error_file | other_command`
 10. `$ selpg -s10 -e20 -l66 input_file`
 
**注意：**其中一些命令在Windows下无法使用，需在Linux下运行。






