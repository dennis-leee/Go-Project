# 开发 web 服务程序

---

## **一、概述**

> 设计一个 web 小应用，展示静态文件服务、js 请求支持、模板输出、表单处理、Filter 中间件设计等方面的能力。（不需要数据库支持）

## **二、任务要求**
> 编程 web 应用程序 cloudgo-io。 请在项目 README.MD 给出完成任务的证据！
### **基本要求**
>  1. 支持静态文件服务
>  2. 支持简单 js 访问
>  3. 提交表单，并输出一个表格
>  4. 对 `/unknown` 给出开发中的提示，返回码 `5xx`

## **三、 相关说明**
- 本程序实现了一个简单的注册&显示用户功能
- 仅实现了非法格式判断，无数据库
- 运行时，需将项目文件夹（ex-cloudgo-inout）放到$GOPATH/下，而后用命令行进入该文件夹，执行`go run main.go`即可
- `assets`文件夹存放了`css`, `js`, `image`文件
- `serivice` 存放服务器代码
- `templates`存放模板文件
- `outImages`存放程序预览图
- 建议用`Chrome`浏览器进行访问
## **四、 测试结果**
### **1. 静态文件访问**
目录结构如下：  
![静态文件目录结构][1]  
访问结果：  
![访问结果][2]  
可以正常访问。
### **2. 简单js访问**
本程序用js实现了非法格式的检测，当存在非法格式时，无法提交，且会给出相应的错误提示。  
![支持js][3]  
从上图中可以看到，客户端成功获取到了js文件，并正常执行。
### **3. 提交表单，输出表格**
输入正确的格式并提交：  
![注册][4]  
结果：  
![结果][5]
### **4. 对`/unknown`的处理**
![unknown][6]


  [1]: https://github.com/my937889621/Go-Project/blob/master/src/ex-cloudgo-inout/outImages/%E9%9D%99%E6%80%81%E6%96%87%E4%BB%B6%E7%9B%AE%E5%BD%95%E7%BB%93%E6%9E%84.PNG
  [2]: https://github.com/my937889621/Go-Project/blob/master/src/ex-cloudgo-inout/outImages/%E8%AE%BF%E9%97%AE%E9%9D%99%E6%80%81%E6%96%87%E4%BB%B6.PNG
  [3]: https://github.com/my937889621/Go-Project/blob/master/src/ex-cloudgo-inout/outImages/js%E8%AE%BF%E9%97%AE.PNG
  [4]: https://github.com/my937889621/Go-Project/blob/master/src/ex-cloudgo-inout/outImages/%E6%8F%90%E4%BA%A4-1.PNG
  [5]: https://github.com/my937889621/Go-Project/blob/master/src/ex-cloudgo-inout/outImages/%E6%8F%90%E4%BA%A4-2.PNG
  [6]: https://github.com/my937889621/Go-Project/blob/master/src/ex-cloudgo-inout/outImages/unknown.PNG
