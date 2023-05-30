# ctp-go
上期技术期货交易协议CTP的Go版本，使用方法和C++基本一致。

具体接口文档请参考上期CTP官方文档：http://www.sfit.com.cn/5_2_DocumentDown_2.htm （非工作时间可能无法访问）

支持Windows(x86/x64)、Linux(x64)，虽然上期官方已经更新了MacOS版本，但是由于我没有Mac，所以暂时不支持。
ctp-go基于Go 1.18，更早版本请自行修改mod文件测试看看是否可用。

## 安装说明
### 第一步：安装GCC
由于ctp-go使用了CGO，所以需要GCC或类GCC编译器，例如LLVM/Clang。
#### Linux下直接安装GCC或者LLVM/Clang (二选一即可)：
```
# Debian系
sudo apt install gcc

# RedHat系
sudo yum install gcc
```
**建议直接使用Clang，编译速度更快**
```
# Debian系
sudo apt install clang 

# RedHat系
sudo yum install clang

# 通过修改环境变量CC指定clang作为编译器
# export命令只会临时修改当前会话的环境变量，长久修改请自行修改系统或用户的profile文件。
export CC=clang
```
#### Windows下安装Mingw w64或者使用LLVM (二选一即可)

Mingw w64：https://github.com/niXman/mingw-builds-binaries/releases/

**建议使用llvm-mingw，编译速度更快！**

llvm-mingw：https://github.com/mstorsjo/llvm-mingw/releases

**记得将解压后的bin目录添加到系统环境变量。**

### 第二步：下载对应的版本
由于CTP接口区分不同版本，所以本包采用多个分支对应多个版本.

* 6.3.15
* 6.6.7

安装的时候需要确定服务器使用的版本。

例如CTP柜台要求是6.3.15版本，那么就安装6.3.15分支:
```
go get github.com/kkqy/ctp-go@6.3.15
```
如果柜台要求是6.6.7版本，那么就安装6.6.7：
```
go get github.com/kkqy/ctp-go@6.6.7
```

一般小版本是能兼容的。

## 使用说明
这里以交易接口为例，使用API查询合约的价格最小变动(PriceTick)。行情接口用法类似。
### 第一步：导入ctp-go包
```
import (
    "fmt"
    "time"

    ctp "github.com/kkqy/ctp-go"
)
```
### 第二步：创建交易SPI，用于响应回调。
```
type MyTraderSpi struct {
    ctp.TraderSpi //TraderSpi里定义了交易Spi的所有回调函数，这些回调函数不做任何操作，仅仅用作默认“兜底”。
}

func (p *MyTraderSpi) OnFrontConnected() {
    fmt.Println("交易前置已连接！")
}

func (p *MyTraderSpi) OnRspAuthenticate(pRspAuthenticateField ctp.CThostFtdcRspAuthenticateField, pRspInfo ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
    errID := pRspInfo.GetErrorID()
    if errID == 0 {
        fmt.Println("客户端认证成功！")
    } else {
        fmt.Println("认证失败！错误代码：", errID)
    }
}
func (p *MyTraderSpi) OnRspUserLogin(pRspUserLogin ctp.CThostFtdcRspUserLoginField, pRspInfo ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
    errID := pRspInfo.GetErrorID()
    if errID == 0 {
        fmt.Println("用户登陆")
    } else {
        fmt.Println("登陆失败！错误代码：", errID)
    }
}

func (p *MyTraderSpi) OnRspQryInstrument(pInstrument ctp.CThostFtdcInstrumentField, pRspInfo ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
    fmt.Println("接收到合约：", pInstrument.GetInstrumentID(), "价格最小变动：", pInstrument.GetPriceTick())
    if bIsLast {
        fmt.Println("合约信息接收完成")
    }
}
```
**此处应特别注意：按照CTP接口的设计，SPI是通过纯虚函数实现的回调，而由于Go没有C++中纯虚函数的概念，编译时也无法检查这些回调函数定义是否存在，当回调被调用时，如果Go里面没有定义相应的回调函数，则会产生空指针/引用的panic。**

举个例子：如果服务器断开，则OnFrontDisconnected会被回调，而此时若没有定义OnFrontDisconnected这个回调函数，就会造成空指针/引用的panic。

但是，大多数时候我们其实只关心一小部分接口的回调，如果为了去实现全部的回调，写一大堆空的回调方法，代码就显得冗长了，所以ctp-go提供了“空的”**TraderSpi**和**MarketDataSpi**，里面包含了SPI所有的响应函数实现，只不过它们都不做任何事情。而利用Go语言结构体的“组合”特性，我们能让这个“空的”TraderSpi作为“兜底”，我们只需要“覆盖”我们关心的回调方法即可。

### 第三步：创建交易API，并且查询合约信息。
```
func main() {
    flowPath := "./"                                 // 流文件存储文件夹，CTP不会自动创建文件夹，所以必须提前创建，否则会报错
    traderFrontAddr := "tcp://180.168.146.187:10201" // 交易前置服务器地址
    brokerID := "9999"                               // BrokerID
    appID := "simnow_client_test"                    // APPID
    authCode := "0000000000000000"                   // 授权码
    userProductionInfo := ""                         // 客户端信息
    userID := "888888"                               // 用户ID
    password := "simnowpassword"                     // 用户密码

    traderApi := ctp.CThostFtdcTraderApiCreateFtdcTraderApi(flowPath)         //创建API
    traderApi.RegisterSpi(ctp.NewDirectorCThostFtdcTraderSpi(&MyTraderSpi{})) // 注册响应Spi
    traderApi.RegisterFront(traderFrontAddr)                                  // 注册交易前置
    traderApi.SubscribePrivateTopic(ctp.THOST_TERT_RESTART)                   // 订阅私有流
    traderApi.SubscribePublicTopic(ctp.THOST_TERT_RESTART)                    // 订阅公共流
    traderApi.Init()                                                          // 初始化API
    time.Sleep(time.Second * 5)                                               // 等待连接完成

    reqAuthenticateField := ctp.NewCThostFtdcReqAuthenticateField() // 创建ReqAuthenticateField
    reqAuthenticateField.SetBrokerID(brokerID)
    reqAuthenticateField.SetUserProductInfo(userProductionInfo)
    reqAuthenticateField.SetAuthCode(authCode)
    reqAuthenticateField.SetAppID(appID)
    traderApi.ReqAuthenticate(reqAuthenticateField, 1)             // 请求认证
    ctp.DeleteCThostFtdcReqAuthenticateField(reqAuthenticateField) // 释放reqAuthenticateField
    time.Sleep(time.Second * 5)                                    // 等待授权认证完成

    userLoginField := ctp.NewCThostFtdcReqUserLoginField() // 创建UserLoginField
    userLoginField.SetUserID(userID)
    userLoginField.SetPassword(password)
    userLoginField.SetBrokerID(brokerID)
    traderApi.ReqUserLogin(userLoginField, 2)             // 请求登陆
    ctp.DeleteCThostFtdcReqUserLoginField(userLoginField) // 释放userLoginField
    time.Sleep(time.Second * 5)                           // 等待用户登陆完成

    qryInstrumentField := ctp.NewCThostFtdcQryInstrumentField()
    traderApi.ReqQryInstrument(qryInstrumentField, 3)          // 请求查询合约
    ctp.DeleteCThostFtdcQryInstrumentField(qryInstrumentField) // 释放qryInstrumentField

    traderApi.Join()
}
```
### 最终代码
```
package main

import (
    "fmt"
    "time"

    ctp "github.com/kkqy/ctp-go"
)


type MyTraderSpi struct {
    ctp.TraderSpi //TraderSpi里定义了交易Spi的所有回调函数，这些回调函数不做任何操作，仅仅用作默认“兜底”。
}

func (p *MyTraderSpi) OnFrontConnected() {
    fmt.Println("交易前置已连接！")
}

func (p *MyTraderSpi) OnRspAuthenticate(pRspAuthenticateField ctp.CThostFtdcRspAuthenticateField, pRspInfo ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
    errID := pRspInfo.GetErrorID()
    if errID == 0 {
        fmt.Println("客户端认证成功！")
    } else {
        fmt.Println("认证失败！错误代码：", errID)
    }
}
func (p *MyTraderSpi) OnRspUserLogin(pRspUserLogin ctp.CThostFtdcRspUserLoginField, pRspInfo ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
    errID := pRspInfo.GetErrorID()
    if errID == 0 {
        fmt.Println("用户登陆")
    } else {
        fmt.Println("登陆失败！错误代码：", errID)
    }
}

func (p *MyTraderSpi) OnRspQryInstrument(pInstrument ctp.CThostFtdcInstrumentField, pRspInfo ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
    fmt.Println("接收到合约：", pInstrument.GetInstrumentID(), "价格最小变动：", pInstrument.GetPriceTick())
    if bIsLast {
        fmt.Println("合约信息接收完成")
    }
}

func main() {
    flowPath := "./"                                 // 流文件存储文件夹，CTP不会自动创建文件夹，所以必须提前创建，否则会报错
    traderFrontAddr := "tcp://180.168.146.187:10201" // 交易前置服务器地址
    brokerID := "9999"                               // BrokerID
    appID := "simnow_client_test"                    // APPID
    authCode := "0000000000000000"                   // 授权码
    userProductionInfo := ""                         // 客户端信息
    userID := "888888"                               // 用户ID
    password := "simnowpassword"                     // 用户密码

    traderApi := ctp.CThostFtdcTraderApiCreateFtdcTraderApi(flowPath)     //创建API
    traderApi.RegisterSpi(ctp.NewDirectorCThostFtdcTraderSpi(&MyTraderSpi{})) // 注册响应Spi
    traderApi.RegisterFront(traderFrontAddr)                              // 注册交易前置
    traderApi.SubscribePrivateTopic(ctp.THOST_TERT_RESTART)               // 订阅私有流
    traderApi.SubscribePublicTopic(ctp.THOST_TERT_RESTART)                // 订阅公共流
    traderApi.Init()                                                      // 初始化API
    time.Sleep(time.Second * 5)                                           // 等待连接完成

    reqAuthenticateField := ctp.NewCThostFtdcReqAuthenticateField() // 创建ReqAuthenticateField
    reqAuthenticateField.SetBrokerID(brokerID)
    reqAuthenticateField.SetUserProductInfo(userProductionInfo)
    reqAuthenticateField.SetAuthCode(authCode)
    reqAuthenticateField.SetAppID(appID)
    traderApi.ReqAuthenticate(reqAuthenticateField, 1)             // 请求认证
    ctp.DeleteCThostFtdcReqAuthenticateField(reqAuthenticateField) // 释放reqAuthenticateField
    time.Sleep(time.Second * 5)                                    // 等待授权认证完成

    userLoginField := ctp.NewCThostFtdcReqUserLoginField() // 创建UserLoginField
    userLoginField.SetUserID(userID)
    userLoginField.SetPassword(password)
    userLoginField.SetBrokerID(brokerID)
    traderApi.ReqUserLogin(userLoginField, 2)             // 请求登陆
    ctp.DeleteCThostFtdcReqUserLoginField(userLoginField) // 释放userLoginField
    time.Sleep(time.Second * 5)                           // 等待用户登陆完成

    qryInstrumentField := ctp.NewCThostFtdcQryInstrumentField()
    traderApi.ReqQryInstrument(qryInstrumentField, 3)          // 请求查询合约
    ctp.DeleteCThostFtdcQryInstrumentField(qryInstrumentField) // 释放qryInstrumentField

    traderApi.Join()
}
```
## 将交易API/SPI、行情API/SPI封装到一起的示例代码
在CTP接口中，API负责请求，SPI负责响应，而且同时又区分交易接口和行情接口，有一些零散。
借助Golang的“组合”特性，我们可以巧妙地将CTP接口的API和SPI全部组合到一起。
在这个示例程序里，交易接口和行情接口的API和SPI都被一起封装到一个名叫Gateway的结构里。
最终实现通过交易接口查询合约代码，然后再通过行情接口订阅行情，并且输出最新的价格。
此示例只是一个抛砖引玉的作用，具体项目具体考虑。

```
package main

import (
    "fmt"
    "sync"
    "sync/atomic"

    ctp "github.com/kkqy/ctp-go"
)
type Gateway struct {
    wg *sync.WaitGroup // 使用WaitGroup来同步连接、认证、登陆的顺序

    ctp.CThostFtdcTraderApi // 交易API
    ctp.CThostFtdcMdApi     // 行情API
    ctp.TraderSpi           // TraderSpi里定义了交易Spi的所有回调函数，这些回调函数不做任何操作，仅仅用作默认“兜底”。
    ctp.MarketDataSpi       // MarketDataSpi里定义了行情Spi的所有回调函数，功能同上。
    requestID               atomic.Uint64
    instrumentIDs           []string // 用来存储合约列表
}

func (p *Gateway) OnFrontConnected() {
    p.wg.Done()
}

func (p *Gateway) OnRspAuthenticate(pRspAuthenticateField ctp.CThostFtdcRspAuthenticateField, pRspInfo ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
    errID := pRspInfo.GetErrorID()
    if errID == 0 {
        p.wg.Done()
    } else {
        fmt.Println("认证失败！错误代码：", errID)
    }
}
func (p *Gateway) OnRspUserLogin(pRspUserLogin ctp.CThostFtdcRspUserLoginField, pRspInfo ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
    errID := pRspInfo.GetErrorID()
    if errID == 0 {
        p.wg.Done()
    } else {
        fmt.Println("登陆失败！错误代码：", errID)
    }
}

func (p *Gateway) OnRspQryInstrument(pInstrument ctp.CThostFtdcInstrumentField, pRspInfo ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
    p.instrumentIDs = append(p.instrumentIDs, pInstrument.GetInstrumentID())
    if bIsLast {
        p.wg.Done()
    }
}

func (p *Gateway) OnRtnDepthMarketData(pDepthMarketData ctp.CThostFtdcDepthMarketDataField) {
    fmt.Println("合约代码：", pDepthMarketData.GetInstrumentID(), "最新价格：", pDepthMarketData.GetLastPrice())
}

func (p *Gateway) Join() {
    p.CThostFtdcTraderApi.Join()
    p.CThostFtdcMdApi.Join()
}
func (p *Gateway) Release() {
    p.CThostFtdcTraderApi.Release()
    p.CThostFtdcMdApi.Release()
}

func NewGateway(flowPath, traderFrontAddr, marketDataFrontAddr, brokerID, userID, password, appID, authCode, userProductionInfo string) *Gateway {

    p := &Gateway{
        wg: &sync.WaitGroup{},

        CThostFtdcTraderApi: ctp.CThostFtdcTraderApiCreateFtdcTraderApi(flowPath),
        CThostFtdcMdApi:     ctp.CThostFtdcMdApiCreateFtdcMdApi(flowPath),
        instrumentIDs:       make([]string, 0),
    }

    p.CThostFtdcTraderApi.RegisterSpi(ctp.NewDirectorCThostFtdcTraderSpi(p))
    p.CThostFtdcMdApi.RegisterSpi(ctp.NewDirectorCThostFtdcMdSpi(p))

    p.CThostFtdcTraderApi.RegisterFront(traderFrontAddr) //注创建册交易前置
    p.CThostFtdcMdApi.RegisterFront(marketDataFrontAddr) //注册行情前置

    p.SubscribePrivateTopic(ctp.THOST_TERT_RESTART) //订阅私有流
    p.SubscribePublicTopic(ctp.THOST_TERT_RESTART)  //订阅公共流
    fmt.Println("网关订阅流完成")

    p.wg.Add(2) //由于交易和行情接口都会调用OnFrontConnected，所以会Done两次
    p.CThostFtdcTraderApi.Init()
    p.CThostFtdcMdApi.Init()
    p.wg.Wait()
    fmt.Println("网关初始化完成")

    p.wg.Add(1)
    reqAuthenticateField := ctp.NewCThostFtdcReqAuthenticateField() //创建认证字段
    reqAuthenticateField.SetBrokerID(brokerID)
    reqAuthenticateField.SetUserProductInfo(userProductionInfo)
    reqAuthenticateField.SetAuthCode(authCode)
    reqAuthenticateField.SetAppID(appID)
    p.ReqAuthenticate(reqAuthenticateField, int(p.requestID.Add(1))) // 交易API发送认证请求
    ctp.DeleteCThostFtdcReqAuthenticateField(reqAuthenticateField)   // 记得释放
    p.wg.Wait() // 等待认证完成
    fmt.Println("网关已认证")

    p.wg.Add(2) //由于交易和行情接口都会调用OnRspUserLogin，所以会Done两次
    userLoginField := ctp.NewCThostFtdcReqUserLoginField()                       //创建登陆字段
    userLoginField.SetUserID(userID)
    userLoginField.SetPassword(password)
    userLoginField.SetBrokerID(brokerID)
    p.CThostFtdcTraderApi.ReqUserLogin(userLoginField, int(p.requestID.Add(1))) // 交易API发送登陆请求
    p.CThostFtdcMdApi.ReqUserLogin(userLoginField, 0)                           // 行情API发送登陆请求
    ctp.DeleteCThostFtdcReqUserLoginField(userLoginField)                       // 记得释放
    p.wg.Wait() // 等待登陆完成
    fmt.Println("网关已登录")

    p.wg.Add(1)
    qryInstrumentField := ctp.NewCThostFtdcQryInstrumentField()
    p.ReqQryInstrument(qryInstrumentField, int(p.requestID.Add(1)))
    ctp.DeleteCThostFtdcQryInstrumentField(qryInstrumentField)
    p.wg.Wait()
    fmt.Println("网关已获取合约列表")

    p.SubscribeMarketData(p.instrumentIDs)
    fmt.Println("网关已订阅所有合约")
    return p
}

func main() {
    // 替换你自己的登陆信息
    gateway := NewGateway("./", "tcp://180.168.146.187:10201", "tcp://180.168.146.187:10211", "9999", "888888", "simnowpassword", "simnow_client_test", "0000000000000000", "")
    gateway.Join()
}
```

## 注意事项（避坑指南）
* 使用ctp.NewXXX 创建的CTP的任何结构，使用完以后一定要用对应的释放函数ctp.DeleteXXX 进行释放，不然可能会造成内存泄漏。例如:
```
userLoginField := ctp.NewCThostFtdcReqUserLoginField() // 创建UserLoginField
userLoginField.SetUserID(userID)
userLoginField.SetPassword(password)
userLoginField.SetBrokerID(brokerID)
traderApi.ReqUserLogin(userLoginField, 2)             // 请求登陆
ctp.DeleteCThostFtdcReqUserLoginField(userLoginField) // 需要手动释放userLoginField，不然会造成内存泄漏！

// 如果担心忘记释放，可以借助Go的defer。
userLoginField := ctp.NewCThostFtdcReqUserLoginField() // 创建UserLoginField
defer ctp.DeleteCThostFtdcReqUserLoginField(userLoginField) // New完马上通过defer来Delete就不容易忘记。如果就这样了还是记不住，那泄露就泄露吧，如果内存占用没炸，那就“无所谓”，如果内存占用炸了，那肯定会去找原因然后长记性对吧？
```
* 回调函数的非GO原生类型的参数不要传递到回调函数外面去，因为这些参数内部的指针指向的内存在回调结束后会被释放掉，那时指向的数据是不确定的。例如：
```
func (p *Gateway) OnRspUserLogin(pRspUserLogin ctp.CThostFtdcRspUserLoginField, pRspInfo ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
    OuterVar1 = pRspUserLogin // 不可以！pRspUserLogin内部指针指向的内存是CTP接口在C/C++运行时的堆里分配的，不受Go管理，所以回调函数结束以后这些参数指向的内存就会被释放。
    go func(){
        _ = pRspUserLogin.GetTradingDay() // 不可以！因为Go并不能保证goroutine执行顺序，也许在调用GetTradingDay之前，OnRspUserLogin就已经执行完，pRspUserLogin内部指针指向的内存已经被释放了。
    }()
    OuterVar2 = pRspUserLogin.GetTradingDay() // 可以，因为GetTradingDay()返回的是Go的string类型数据，由Go负责内存管理和回收。
    OuterVar3 = nRequestID // 可以但没必要，因为int是Go的原生类型，由Go负责内存管理和回收。
    OuterVar4 = bIsLast // 可以但没必要，因为bool是Go的原生类型，由Go负责内存管理和回收。
}
```
* 使用的包的版本必须和动态链接库严格一致，例如：使用6.6.x的包进行编译就必须用6.6.x的动态链接库，否则虽然程序也许运行得起来，但是会出现稀奇古怪的、不可预料的、捉摸不透的错误。
* 由于CTP的字符串使用的是GB2312/GBK编码，而Go的字符串使用的是UTF8编码，所以CTP中出现的中文字符串在Go中会乱码。CTP接口众多，考虑到运行效率问题，最佳实践应该是按需转码，所以ctp-go并没有直接集成编码转换，需要自行处理。

### Linux 平台
* 编译后的可执行文件依赖libthostmduserapi_se.so、libthosttraderapi_se.so这两个动态链接库，这两个文件在包目录下Linux文件夹中可以找到，官方下载的API里这两个动态链接库文件名缺少“lib”前缀，需要手动修改一下文件名。由于编译的时候指定了库的搜索路径，所以在本机上运行的时候能直接运行。但是如果打包发布，则需要附带这两个动态链接库文件一起发布。

### Windows 平台
* 编译后的可执行文件依赖thostmduserapi_se.dll、thosttraderapi_se.dll、wrap.dll这三个动态链接库，这三个文件在包目录下Windows文件夹中可以找到（需要区分x86和x64版本），需要将这三个文件与可执行文件放在同一目录，才能正常运行，否则会出现“由于找不到 xxx.dll，无法继续执行代码。重新安装程序可能会解决此问题。”的错误。

## 常见问题
### 为什么使用ctp-go以后我的程序编译那么慢？
首次构建的时候，ctp-go需要使用CGO进行编译，此时需要耗费较长时间，但仅仅只是首次比较慢，当再次构建的时候，Go的编译器就仅仅只会编译修改过的模块，而不需要再次编译ctp-go。个人实测，如果使用gcc编译器，首次构建最长可能花费10分钟左右，而如果替换成了LLVM/Clang，首次编译时间直接缩短为20秒不到，通过任务管理器可以对比，LLVM/Clang的CPU和内存利用率比gcc高出许多，所以，如果实在在意首次编译时间，不论是Windows还是Linux，都建议直接使用LLVM/Clang替代GCC作为编译器前端。
LLVM/Clang编译速度比GCC快，且有更友好的错误提示。关于LLVM/Clang与GCC的说明、性能对比以及安装方法可参考网上其它教程，

### 修改GOARCH为386以后，出现undefined: CThostFtdcRspUserLoginField等错误
由于本包依赖CGO，所以需要将CGO_ENABLE设置为1。默认GOARCH下，CGO_ENABLE默认也是打开状态，但是在切换GOARCH为386，也就是编译32位的时候，CGO_ENABLE默认为关闭，需要手动设置为1。
