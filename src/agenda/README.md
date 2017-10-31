# CLI 命令行实用程序开发实战 - Agenda
---
### **小组成员**
#### 梁耀飞  15331185
#### 李开锋  15331155
#### 李近朱  15331153
---

### 需求描述

- 业务需求：见后面需求

- 功能需求： 设计一组命令完成 agenda 的管理，例如：
>agenda help ：列出命令说明
>agenda register -uUserName --password pass -email=a@xxx.com ：注册用户
>agenda help register ：列出 register 命令的描述
>agenda cm ... : 创建一个会议
>
>原则上一个命令对应一个业务功能

- 持久化要求：
>使用 json 存储 User 和 Meeting 实体
>当前用户信息存储在 curUser.txt 中

- 开发需求:
>团队：2-3人，一人作为 master 创建程序框架，其他人 fork 该项目，所有人同时开发。团队 不能少于 2 人
时间：两周完成

- 项目目录
>cmd ：存放命令实现代码
>entity ：存放 User 和 Meeting 对象读写与处理逻辑

- 其他目录 ： 自由添加
- 日志服务
>使用 log 包记录命令执行情况

---

### Agenda 业务需求

- 用户注册
注册新用户时，用户需设置一个唯一的用户名和一个密码。另外，还需登记邮箱及电话信息。
如果注册时提供的用户名已由其他用户使用，应反馈一个适当的出错信息；成功注册后，亦应反馈一个成功注册的信息。

- 用户登录
用户使用用户名和密码登录 Agenda 系统。
用户名和密码同时正确则登录成功并反馈一个成功登录的信息。否则，登录失败并反馈一个失败登录的信息。

- 用户登出
已登录的用户登出系统后，只能使用用户注册和用户登录功能。

- 用户查询
已登录的用户可以查看已注册的所有用户的用户名、邮箱及电话信息。

- 用户删除

1.已登录的用户可以删除本用户账户（即销号）。

2.操作成功，需反馈一个成功注销的信息；否则，反馈一个失败注销的信息。

3.删除成功则退出系统登录状态。删除后，该用户账户不再存在。

4.用户账户删除以后：

>以该用户为发起者的会议将被删除
>  
>以该用户为参与者的会议将从参与者列表中移除该用户。若因此造成会议参与者人数为0，则会议也将被删除。

- 创建会议

1.已登录的用户可以添加一个新会议到其议程安排中。会议可以在多个已注册
用户间举行，不允许包含未注册用户。添加会议时提供的信息应包括：

>会议主题(title)（在会议列表中具有唯一性）
>
> 会议参与者(participator)
>
> 会议起始时间(start time)
>
> 会议结束时间(end time)
  
2.注意，任何用户都无法分身参加多个会议。如果用户已有的会议安排（作为发起者或参与者）与将要创建的会议在时间上重叠 （允许仅有端点重叠的情况），则无法创建该会议。

3.用户应获得适当的反馈信息，以便得知是成功地创建了新会议，还是在创建过程中出现了某些错误。

- 增删会议参与者。

1.已登录的用户可以向 自己发起的某一会议增加/删除 参与者 。

2.增加参与者时需要做 时间重叠 判断（允许仅有端点重叠的情况）。

3.删除会议参与者后，若因此造成会议 参与者 人数为0，则会议也将被删除。

- 查询会议

1.已登录的用户可以查询自己的议程在某一时间段(time interval)内的所有会议安排。

2.用户给出所关注时间段的起始时间和终止时间，返回该用户议程中在指定时间范围内找到的所有会议安排的列表。

3.在列表中给出每一会议的起始时间、终止时间、主题、以及发起者和参与者。

4.注意，查询会议的结果应包括用户作为 发起者或参与者 的会议。

- 取消会议

1.已登录的用户可以取消 自己发起 的某一会议安排。

2.取消会议时，需提供唯一标识：会议主题（title）。

- 退出会议

1.已登录的用户可以退出 自己参与 的某一会议安排。

2.退出会议时，需提供一个唯一标识：会议主题（title）。若因此造成会议 参与者人数为0，则会议也将被删除。

- 清空会议

已登录的用户可以清空 自己发起 的所有会议安排。

---

## 测试

> $ go get -u github.com/my937889621/Go-Project/tree/master/src/agenda

> $GOPATH/bin/agenda -h
>
>Agenda is a meeting manager based on CLI using cobra library.
>It supports different operation on meetings including register, create meeting, query and so on.
>It's a cooperation homework assignment for service computing.
>
>Usage:
>  Agenda [command]
>
>
```
Available Commands:
  cancel              Cancel your own meeting by specifying title name.
  changeParticipators Change your own meetings' participators.
  clear               Clear all meetings you attended or created.
  createMeetings      Create meetings.
  delete              A brief description of your command
  help                Help about any command
  list
  listMeetingsCmd     List all of your own meetings during a time interval.
  login               Login
  logout              Logout
  quit                Quit meetings.
  register            Register user.
```

>Flags:
  -d, --debug   display log message
  -h, --help    help for Agenda
>
>Use "Agenda [command] --help" for more information about a command.


## 测试用例

#### 1.注册

>./agenda register -u dennis -p 123456 -m 123@qq.com -t 147258369


#### 2.再次注册已存在账户

>there's another user with username Tomy


#### 3.登录

>./agenda login -u dennis -p 123456


#### 4.密码错误

>Authentication Fail


#### 5.未退出再次登录

>Action login requires an logout state


#### 6.创建会议

>./agenda createMeetings -t Halloween -p dennis -s 2017-10-31 -e 2017-10-31

创建后的几种情况输出：
>meeting hosted

>there's another meeting with title: ABC_Meeting

>there are time conflict of some participants

>meeting should end later than start

>invalid time format: 2018-12-1


#### 7.更改参与者

>./agenda changeParticipators -t Halloween -p AAA

更改后的几种情况输出：

>user 'AAA' is already a participant of meeting 'Halloween'

>meeting doesn't exist: PPP


#### 8.取消会议

>./agenda cancel -t Halloween

会议不存在：
>meeting doesn't exist: ZZZ


#### 9.退出会议

>./agenda quit -t Halloween

不是参与者无法退出会议：

>user 'Li' is not a participant of meeting 'Halloween'


#### 10.列出数据

列出所有注册用户：
>./agenda list

```
Username Email Phone
'dennis' '123@qq.com' '147258369'
'jack' '456@qq.com' '159357852'
'simon' '789@qq.com' '456852369'
```

根据时间列出会议

>./agenda listMeetings -s 2017-10-20 -e 2017-12-01

```
title: Halloween
  host: dennis
  time: 2017-10-31 to 2017-12-31
  participants: jack
```

列出会议时情况输出

>meeting should end later than start

>invalid time format: 2017-10-1


#### 11.清除会议

>./agenda clear


#### 12.删除会议

>./agenda delete
