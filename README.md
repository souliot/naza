<p align="center">
<br>
Go语言基础库
<br><br>
<a title="TravisCI" target="_blank" href="https://www.travis-ci.org/q191201771/naza"><img src="https://www.travis-ci.org/q191201771/naza.svg?branch=master"></a>
<a title="codecov" target="_blank" href="https://codecov.io/gh/q191201771/naza"><img src="https://codecov.io/gh/q191201771/naza/branch/master/graph/badge.svg?style=flat-square"></a>
<a title="goreportcard" target="_blank" href="https://goreportcard.com/report/github.com/souliot/naza"><img src="https://goreportcard.com/badge/github.com/souliot/naza?style=flat-square"></a>
<a title="sourcegraph" target="_blank" href="https://sourcegraph.com/github.com/souliot/naza"><img src="https://sourcegraph.com/github.com/souliot/naza/-/badge.svg"></a>
<br>
<a title="codeline" target="_blank" href="https://github.com/souliot/naza"><img src="https://sloc.xyz/github/q191201771/naza/?category=code"></a>
<a title="license" target="_blank" href="https://github.com/souliot/naza/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square"></a>
<a title="lastcommit" target="_blank" href="https://github.com/souliot/naza/commits/master"><img src="https://img.shields.io/github/commit-activity/m/q191201771/naza.svg?style=flat-square"></a>
<a title="commitactivity" target="_blank" href="https://github.com/souliot/naza/graphs/commit-activity"><img src="https://img.shields.io/github/last-commit/q191201771/naza.svg?style=flat-square"></a>
<br>
<a title="pr" target="_blank" href="https://github.com/souliot/naza/pulls"><img src="https://img.shields.io/github/issues-pr-closed/q191201771/naza.svg?style=flat-square&color=FF9966"></a>
<a title="hits" target="_blank" href="https://github.com/souliot/naza"><img src="https://hits.b3log.org/q191201771/naza.svg?style=flat-square"></a>
<a title="language" target="_blank" href="https://github.com/souliot/naza"><img src="https://img.shields.io/github/languages/count/q191201771/naza.svg?style=flat-square"></a>
<a title="toplanguage" target="_blank" href="https://github.com/souliot/naza"><img src="https://img.shields.io/github/languages/top/q191201771/naza.svg?style=flat-square"></a>
<a title="godoc" target="_blank" href="https://godoc.org/github.com/souliot/naza"><img src="http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square"></a>
<br><br>
<a title="watcher" target="_blank" href="https://github.com/souliot/naza/watchers"><img src="https://img.shields.io/github/watchers/q191201771/naza.svg?label=Watchers&style=social"></a>&nbsp;&nbsp;
<a title="star" target="_blank" href="https://github.com/souliot/naza/stargazers"><img src="https://img.shields.io/github/stars/q191201771/naza.svg?label=Stars&style=social"></a>&nbsp;&nbsp;
<a title="fork" target="_blank" href="https://github.com/souliot/naza/network/members"><img src="https://img.shields.io/github/forks/q191201771/naza.svg?label=Forks&style=social"></a>&nbsp;&nbsp;
</p>

---

#### 工程目录说明

```
pkg/                    ...... 源码包
    |-- bininfo/        ...... 将编译时源码的git版本信息（当前commit log的sha值和commit message），编译时间，Go版本，平台打入程序中
    |-- log/        ...... 日志库
    |-- assert/         ...... 提供了单元测试时的断言功能，减少一些模板代码
    |-- nazaerrors/     ...... error相关
    |-- fake/           ...... 实现一些常用的接口，hook一些不方便测试的代码
    |-- taskpool/       ...... 非阻塞协程池，协程数量可动态增长，可配置最大协程并发数量，可手动释放空闲的协程
    |-- defertaskthread ...... 执行延时任务
    |-- bele/           ...... 大小端转换操作
    |-- nazabits/       ...... 位操作
    |-- bitrate/        ...... 计算带宽
    |-- ratelimit/      ...... 限流器，令牌桶，漏桶
    |-- connection/     ...... 对net.Conn接口的二次封装
    |-- nazanet/        ...... socket操作相关
    |-- nazahttp/       ...... http操作
    |-- circularqueue   ...... 底层基于切片实现的固定容量大小的FIFO的环形队列
    |-- lru/            ...... LRU缓存
    |-- consistenthash/ ...... 一致性哈希
    |-- nazajson/       ...... json操作
    |-- nazamd5/        ...... md5操作
    |-- nazaatomic/     ...... 原子操作
    |-- slicebytepool/  ...... []byte内存池
    |-- nazastring/     ...... string和[]byte相关的操作
    |-- unique/         ...... 对象唯一ID
    |-- snowflake/      ...... 分布式唯一性ID生成器
    |-- nazareflect/    ...... 利用反射做的一些操作
    |-- filebatch/      ...... 文件批处理操作
    |-- ic/             ...... 将整型切片压缩成二进制字节切片
playground/             ...... Go实验代码片段
demo/                   ...... 示例相关的代码
bin/                    ...... 可执行文件编译输出目录
```

#### 依赖

无任何第三方依赖

#### 相关文档

- pkg/snowflake [分布式 ID 生成算法 snowflake 介绍及 Go 语言实现](https://pengrl.com/p/20041/)
- pkg/bininfo [给 Go 程序加入编译版本时间等信息](https://pengrl.com/p/37397/)
- pkg/nazastring [Go 语言中[]byte 和 string 类型相互转换时的性能分析和优化](https://www.pengrl.com/p/31544/)

- [Go 创建对象时，如何优雅的传递初始化参数](https://pengrl.com/p/60015/)
- playground/p3 [利用 CPU cache 特性优化 Go 程序](https://pengrl.com/p/9125/)
- playground/p4 [老弟在吗，我的 Go 标准库中的二分查找有 bug！](https://pengrl.com/p/20011/)
- [写的一些其他 Go 相关的文章](https://pengrl.com/categories/Go/)

#### 联系我

欢迎扫码加我微信，进行技术交流或扯淡。

<img src="https://pengrl.com/images/yoko_vx.jpeg" width="180" height="180" />

#### 项目名 naza 由来

本仓库主要用于存放我自己写的一些 Go 基础库代码。目前主要服务于我的另一个项目： [lal](https:////github.com/q191201771/lal)

naza 即哪吒（正确拼音为 nezha，我女儿发音读作 naza，少一个字母，挺好~），希望本仓库以后能像三头六臂，有多种武器的哪吒一样，为我提供一个趁手的工具箱。
