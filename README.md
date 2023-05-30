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
