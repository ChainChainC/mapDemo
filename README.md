# mapDemo

### source
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

### style
The naming follow the style in https://google.github.io/styleguide/go/best-practices

### 用户信息 (目前以可运行为目的，可选的用户资料后续补充)
一、第一版本目前逻辑
1、玩家和后端交互基本信息：
--Uuid用来识别玩家唯一性；
--Pos代表玩家位置，用于更新前后端位置数据；
--Name为玩家nickName；
--jwt信息用来判断玩家是否在线、离线、合法性等（TODO）；
2、玩家需要实现的基本后端逻辑：
--NewPlayer：玩家通过前端登录时，同时在后端登录，玩家加入PlayerIdMap表中（model库中定义）；
--PlayerUpdatePos：玩家连入服务器，前端每隔一段时间请求该接口，后端进行判断，同时更新Pos，如果在房间内，返回其它玩家Pos；
--PlayerJoinRoom：玩家申请加入房间，需要携带房间ID（同为Uuid）
--PlayerQuitRoom：玩家退出房间
上述：玩家登录，加入房间、退出房间逻辑基本完成

# -----------项目结构书写指导-------------
--model中定义基本对象的struct，以及一些全局变量
--每一个请求需要对应一个controller的Req struct，用于接收请求的body数据（统一流程）
------后续结构更改
1、将controller和业务逻辑分离开，controllerReqHandler处理请求进入，并进行一些合法性验证
--业务逻辑迁移到单独的controller/player.go文件下，主要进行业务逻辑实现；

# -----------------------------------------------
jwt签发token, token中包含表名，用户id，地理位置

数据库保存用户id和 昵称 

同房间内的用户数据存在一张表中，每隔一段时间读取组内其他人的状态，

表中可能需要的字段：id 昵称 地理位置 状态（在线/离线）

还是需要用户一张表，不然游戏有可能被入侵

用户一张表：id，昵称，手机号（不做第三方验证，仅验证11位长度），密码

游戏房间一张表：房间id，房间状态（是否使用中）

对应游戏房间数量的N张表：玩家id，玩家状态，玩家地理位置 等。。。

//先不管下面需求
可能需要增加的数据 玩家的历史地理位置 玩家断线重联状态字段

### TODO
数据库设计，目前的是cv来的

Logger规范，增加功能
