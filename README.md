# GoGetXueQiu

A gogetxueqiu package for getting stocks and portfolios info from xueqiu.com, in Golang (Go).

本项目是用 Golang 实现的一个针对雪球的爬虫系统，模拟普通用户登录，并从雪球的 json API 获取信息。登录完成后，各接口函数支持 go 并发调用。

### 依赖第三方包

"github.com/bitly/go-simplejson"

### 使用方法

go get "github.com/ynsfsmj/gogetxueqiu"

API 用法请参考 examples/ 文件：

### 开发完成情况

实现了模拟用户登录（解决 CSRF 保护等问题）

初步实现了请求 API 的 Demo

接下来：

1. 实现已知 API 的适配，解决 API 数据获取问题

2. 利用 Go 的优势，实现并发抓取数据，提高效率，但尽量避免高频访问而被封禁

3. 开发实现数据本地缓存，减少抓取数据频率，并提供二次 API，供大家使用

4. 增加数据分析和展示功能

### 项目介绍

本项目是一个针对雪球的爬虫系统，模拟普通用户登录，并从雪球的 json API 获取信息。

在模拟用户登录后，可以访问的信息如下所示：（包括但不限）

* 雪球组合

雪球组合净值变化（按天）
https://xueqiu.com/cubes/nav_daily/all.json?cube_symbol=ZH024581&since=1453555757000&until=1461331757000

组合调仓记录
https://xueqiu.com/cubes/rebalancing/history.json?cube_symbol=ZH024581&count=20&page=1

雪球当日热门组合列表
https://xueqiu.com/cubes/discover/rank/cube/list.json?category=14

雪球组合站内排名
https://xueqiu.com/cubes/data/rank_percent.json?cube_id=24482&market=cn&dimension=annual

雪球组合评分
https://xueqiu.com/cubes/rank/summary.json?symbol=ZH024581&ua=web

雪球组合管理者时间线
https://xueqiu.com/cube/timeline?symbol=ZH024581&page=1&count=20&comment=0&uid=9188557237

* 雪球股票

雪球股票当前状态和价格
https://xueqiu.com/v4/stock/quote.json?code=SZ000625

雪球股票盘口交易数据
https://xueqiu.com/stock/pankou.json?symbol=SZ000625

雪球股票当日分时图
https://xueqiu.com/stock/forchart/stocklist.json?symbol=SZ000625&period=1d&one_min=1

雪球股票后复权日线图
https://xueqiu.com/stock/forchartk/stocklist.json?symbol=SZ000625&period=1day&type=before&begin=1429798115327&end=1461334115327

