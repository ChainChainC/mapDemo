# mapDemo

### source
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

### style
The naming follow the style in https://google.github.io/styleguide/go/best-practices

### 用户信息
不保存用户信息

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
