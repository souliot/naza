#### v0.17.0

- [feat] package nazaerrors: 增加 Wrap 函数，用于封装 error
- [chore] 最低 Go 依赖版本提高至 1.13

#### v0.16.1

- [feat] package nazastring，增加 SubSliceSafety 函数，安全的获取切片的子切片
- [feat] package nazaerrors，增加 CombineErrors 函数，将多个 error 组合成一个

#### v0.16.0

- [refactor] package nazahttp: 删除函数 ReadHTTPRequest，新增函数 ReadHTTPRequestMessage，函数 ReadHTTPResponseMessage

#### v0.15.5

- [fix] package taskpool: 全局对象解析 task 参数错误
- [feat] package nazahttp: 新增函数 ReadHTTPRequest，读取 HTTP 请求并解析

#### v0.15.4

- [feat] 新增 package defertaskthread，用于执行延时任务

#### v0.15.3

- [feat] func UnmarshalRequestJsonBody 反序列化 http json body
- [feat] func PostJson 序列化 json 数据做 http post

#### v0.15.2

- [fix] package nazanet: 使用 net.UDPConn 方式初始化 UDPConnection 时，可以同时初始化 RAddr 远端地址

#### v0.15.1

- [feat] package connection: 增加 GetStat 函数，用于获取连接读取、发送数据的总字节数

#### v0.15.0

- [feat] package nazabits: 新增函数 ReadBits64
- [feat] package nazanet: UDPConnection 支持 IPv6
- [refactor] package nazanet: UDPConnection 构造函数支持更多配置

#### v0.14.0

- package nazanet:
  - [refactor] 重构 UDPConnection，既可作为服务端连接，又可作为客户端连接
- package nazastring:
  - [feat] 增加 func DumpSliceByte

#### v0.13.4

- package circularqueue:
  - [feat] 新增 package，底层基于切片实现的固定容量大小的 FIFO 的环形队列
- package nazanet:
  - [feat] 这是一个新增的 package，其中 struct AvailUDPConnPool，可以从指定的 UDP 端口范围内，寻找可绑定监听的端口，绑定监听并返回
  - [feat] 新增 struct UDPConnection，对 UDP 连接对象的简易封装
- package bele:
  - [feat] 新增 func BEPutUint16, BEPutUint64

#### v0.13.3

- package nazahttp:
  - [fix] 当 request line 和 status line 存在多个空格时，解析错误
- package bininfo:
  - [feat] 增加 git tag 信息

#### v0.13.2

- package nazareflect:
  - [feat] 新增 package，提供三个函数 IsNil，Equal，EqualInteger

#### v0.13.1

- package connection:
  - [feat] 增加连接关闭标志，使用 channel 发送数据时，如果连接已关闭，可以向调用方返回错误
  - [feat] 增加 Option.WriteChanFullBehavior，使用 channel 发送数据时，如果 channel 满了，可以配置是阻塞还是返回错误
  - [fix] 设置 wChan 大小时，应该使用 WriteChanSize 而不是 WriteBufSize
  - [refactor] 不同错误返回不同的错误值
  - [refactor] 去除一些 debug 日志

#### v0.13.0

- package nazabits:
  - [feat] BitReader 的所有函数增加读取越界检查
  - [feat] 增加 BitReader::ReadGolomb 函数，读取 0 阶指数哥伦布编码
- package nazahttp:
  - [feat] nazahttp: 增加函数 ReadHTTPHeader，ParseHTTPRequestLine，ParseHTTPStatusLine，读取 HTTP 头部信息
  - [refactor] 函数 GetHttpFile，DownloadHTttpFile 重命名为 GetHTTPFile，DownloadHTTPFile

#### v0.12.3

- package lru:
  - [feat] 新增 package，一个基础的 LRU 缓存
- package nazahttp:
  - [feat] 新增函数 GetHttpFile，用于下载 HTTP 文件

#### v0.12.2

- package nazabits:
  - [fix] BitWriter::WriteBit 如果原数据不为零值时，会错误覆盖非写入的位

#### v0.12.1

- package nazabits:
  - [feat] 新增 BitReader::ReadBits16，ReadBits32，ReadBytes 函数

#### v0.12.0

- package nazabits:
  - [refactor] BitReader::ReadBits 重命名为 ReadBits8，BitWriter::WriteBits 重命名为 WriteBits8，WriteBits16
  - [fix] BitWriter::WriteBit 传入的值不为 0 和 1 时，只取最低位

#### v0.11.0

- package nazabits:
  - [feat] 增加 BitReader 用于按位读取，增加 BitWriter 用于按位写入
  - [refactor] GetBit8 等函数修改整型类型

#### v0.10.0

- package log:
  - [feat] 新增 WithPrefix 函数，用于支持设置前缀，并且前缀可叠加，使得可以按 repo ，package，对象等维度添加不同的前缀
  - [feat] 新增 Println 等函数，方便替换标准库日志
  - [perf] 减小锁粒度
  - [test] 测试覆盖率增加至 100%
  - [refactor] 删除 FatalIfErrorNotNil, PanicIfErrorNotNil, Outputf, writeShortFile 四个函数函数
- package nazabits:
  - [feat] 新增 GetBit16，GetBits16 函数
- package fake:
  - [feat] 新增 WithRecover 函数
  - [feat] 新增 time.Now hook 相关的接口
  - [refactor] 重新命名 os.Exit hook 相关的接口

#### v0.9.1

- package nazabits:
  - [feat] 新增 package，提供一些位运算函数

#### v0.9.0

- package taskpool:
  - [feat] 增加 Option.MaxWorkerNum，可配置最大协程并发数量
  - [feat] task 任务函数可携带参数
- package log:
  - [feat] 新增 Assert 函数，并可配置断言失败后的行为
- package bele:
  - [feat] 增加一些 ReadXXX 函数，从 io.Reader 中读取数据
  - [feat] 新增 BEUint64 函数

#### v0.8.0

- package ratelimit:
  - [feat] 新增漏桶 LeakyBucket 和令牌桶 TokenBucket，把以前的 RateLimit 删了
- package log:
  - [feat] 新增 Sync 函数，将日志刷盘
  - [feat] panic 和 fatal 前调用 Sync
  - [fix] 全局的 Panic 和 Panicf 忘记调用标准库中的 panic
  - [fix] 使用 IsRotateDaily 控制日志是否按天翻转，之前没有判断这个标志，导致总是按天翻转
- package nazahttp:
  - [feat] 新增 DownloadHttpFile 函数，下载 http 保持为本地文件

#### v0.7.1

- package snowflake
  - [feat] 新增的包，分布式唯一性 64 位 ID 生成器

#### v0.7.0

- package consistenthash:
  - [feat] Nodes 接口返回 node 对应的 point 范围，供调用方判断 hash 是否均匀
  - [feat] hash 函数可由外部配置
  - [refactor] 增加 struct Option
  - [refactor] 内部 point 类型 int -> uint32
- package bitrate:
  - [feat]可配置 bitrate 返回时的单位
  - [feat] unix 时间戳可选择由外部传入
  - [refactor] struct Bitrate -> interface Bitrate
  - [fix] 遍历切片时删除了元素导致崩溃
- package fake:
  - [feat] 添加 func Exit，它是对 os.Exit 的封装，便于其他代码做单元测试
- package log:
  - [test] 使用 fake.Exit

#### v0.6.0

- 新增 package ratelimit：限速器，令牌桶
- 新增 package bitrate：计算带宽
- 新增 package fake
- 删除 package mockwriter
- 删除 package mockserver
- `demo/add_blog_license`：行尾增加两个空格，便于部分 markdown 解析器解析为为换行

#### v0.5.1

- package ic:
  - 新增的包，将整型切片压缩序列化成二进制字符切片
- package bininfo:
  - 增加注释
- package assert:
  - 增加注释

#### v0.5.0

- package filebatch:
  - 遍历读取文件发生错误时，不退出遍历，而是将错误在回调中返回给上层
- package connection:
  - bugfix，初始化 write chan 相关的信息是通过 write chan 的配置，而不是 write buf 的配置
- package slicebytepool:
  - 新增的包，一个 []byte 内存池
- package nazamd5:
  - 新增的包，md5 操作
- package consistenthash:
  - 新增的包，一致性 hash
- package bufferpool:
  - 删除 bufferpool 包
- demo/myapp:
  - 用于演示 package bininfo 的使用
- `demo/add_blog_license`:
  - 修改 license 内容
- 其他：
  - 统一 error 变量的命名方式及内容格式，涉及到的 package：filebatch, connection, taskpool, log

#### v0.4.3

- package bufferpool
  - 新增的包，bytes.Buffer 池
- package nazaatomic
  - 新增的包，对 sync.atomic 的再封装
- package taskpool
  - 新增的包，协程池
- test.sh
  - 做更多的 go tools 检查
- `demo/add_blog_license`
  - 修改许可证样式
  - 检查许可证是否存在时，只检查声明两个字
- `demo/add_go_license`
  - 用户名和邮箱由命令行参数传入

#### v0.4.2

- package filebatch:
  - 新增的包，用于文件批处理操作
- 新增 demo/add_go_license：给 Go 仓库的所有 go 源码文件添加 MIT 许可证
- 新增 demo/add_blog_license：给我自己博客的所有文章尾部添加声明

#### v0.4.1

- package nazastring:
  - 新增 func SliceByteToStringTmp 和 func StringToSliceByteTmp，用于无拷贝的做 string 和[]byte 的转换

#### v0.4.0

- package log:
  - rename -> package log
  - mkdir 0777 and create file 0666, append if file exist
  - 配置使用 Option
  - 配置默认值修改： 打印至控制台开关默认打开，打印源码文件行号开关默认打开
- package connection:
  - erase func Printf
  - 配置使用 Option
- package nazajson:
  - 新增包，作为系统包 json 的补充
- 其它:
  - repo name nezha -> naza

#### v0.3.0

- package connection:
  - 可配置使用 channel 进行异步发送：Config 中增加 WChanSize。增加 Flush, Done, ModWriteChanSize 三个方法
- package log:
  - 增加 panic 相关的方法
- 其它:
  - test.sh 中添加 gofmt 检查

#### v0.2.0

- package log:
  - 去除了对标准库中 log 的依赖
  - 日志支持按天翻转
  - 增加 ShortFileFlag 可配置是否打印源码文件及行号的信息
  - 添加一个 fatal 日志级别，打印完后 exit 程序
  - 当同时打印至控制台和文件时，打印至文件中的 level 字段也带颜色属性
  - 增加 Out 接口
  - 日志不再支持按固定大小翻转 [不兼容]
  - 日志级别从 0 -> 1 开始 [不兼容]

#### v0.1.0

- 删除 /pkg/errors [不兼容]
- package log:
  - 增加 FatalIfErrorNotNil 接口函数，打印错误并退出程序
  - 日志内容中的级别字段右对齐
  - 日志内容中的源码文件名和行号放在整行日志的末尾
  - 增加一些 benchmark
- package assert: 打印正确的源码文件名和行号信息 [bugfix]
- package bele: 增加一些 benchmark
- package unique:
  - 不同的 key 使用不同的自增计数
  - 增加一些 benchmark
- package mockserver: 模拟一些服务端，用于快速测试其它代码
- package mockwriter: 模拟 Writer 接口，用于快速测试其它代码
- 删除 /demo/connstat
- test.sh 脚本只测试 /pkg 目录下的源码

#### v0.0.1

第一个版本
