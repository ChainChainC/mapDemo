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

--jwt信息用来判断玩家是否在线、离线、合法性等（TODO），并且可以存储信息；

2、玩家需要实现的基本后端逻辑：
--NewPlayer：玩家通过前端登录时，同时在后端登录，玩家加入PlayerIdMap表中（model库中定义）；

--PlayerUpdatePos：玩家连入服务器，前端每隔一段时间请求该接口，后端进行判断，同时更新Pos，如果在房间内，返回其它玩家Pos；

--PlayerJoinRoom：玩家申请加入房间，需要携带房间ID（同为Uuid）

--PlayerQuitRoom：玩家退出房间

上述：玩家登录，加入房间、退出房间逻辑基本完成

3、游戏开始
--PlayerUpdatePosInSight：玩家更新位置，同时获取视野内玩家坐标

--PlayerTypeChange：玩家身份变化（本地需要有玩家当前type信息）

--Map需要定时更新：

### -------请求处理统一流程--------
--玩家请求（目前所有请求都是玩家发来的）
1、req进来，进行对应req字段获取（）
2、尝试获取玩家对象
3、验证玩家jwt合法性



### -----------项目结构书写指导-------------


### -----------------------------------------------
jwt签发token, token中包含用户id

数据库保存用户id和 昵称 

同房间内的用户数据存在一张表中，每隔一段时间读取组内其他人的状态，

表中可能需要的字段：id 昵称 地理位置 状态（在线/离线）

还是需要用户一张表，不然游戏有可能被入侵

用户一张表：id，昵称，手机号（不做第三方验证，仅验证11位长度），密码

游戏房间一张表：房间id，房间状态（是否使用中）

对应游戏房间数量的N张表：玩家id，玩家状态，玩家地理位置 等。。。

//先不管下面需求
可能需要增加的数据 玩家的历史地理位置 玩家断线重联状态字段
### DONE 6.12
0、请求携带参数初步确定：

1、玩家加入房间：Jwt，RoomId，Pos

2、玩家退出房间：Jwt，RoomId

3、玩家更新位置：jwt，pos，Type && RoomId

4、新玩家登录：code，jwt

5、新建房间：jwt，Pos

### TODO 6.12--
1、验证redis缓存版本的基础逻辑是否有问题

2、完善获取视野内玩家坐标逻辑函数

3、前端需要按需要缓存信息：jwt，type（玩家标识uint8），roomId（房间号）暂定

4、给server加缓存（减少redis读写压力）

5、增加redis分布式锁（虽然似乎不太需要）

持续：优化代码性能和美观

### Redis目前存储的数据
------玩家信息------
hash表：玩家信息也用hash表存储
{
    "PlayerType": 0,
	"RoomId":     "",
}

-----玩家坐标-------
string：玩家坐标直接全局存储，面临频繁的读写
Pos -> str -> 存入

-----房间信息-------
set：房间存储了当前在房间中的所有玩家
：玩家uuid直接存入set中，代表房间内玩家