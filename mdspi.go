package ctp

type MarketDataSpi struct {
}

// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
func (p *MarketDataSpi) OnFrontConnected() {}

// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
// @nReason param 错误原因
//
//	0x1001 网络读失败
//	0x1002 网络写失败
//	0x2001 接收心跳超时
//	0x2002 发送心跳失败
//	0x2003 收到错误报文
func (p *MarketDataSpi) OnFrontDisconnected(nReason int) {}

// 心跳超时警告。当长时间未收到报文时，该方法被调用。
// @nTimeLapse param 距离上次接收报文的时间
func (p *MarketDataSpi) OnHeartBeatWarning(nTimeLapse int) {}

// 登录请求响应
func (p *MarketDataSpi) OnRspUserLogin(pRspUserLogin CThostFtdcRspUserLoginField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 登出请求响应
func (p *MarketDataSpi) OnRspUserLogout(pUserLogout CThostFtdcUserLogoutField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 错误应答
func (p *MarketDataSpi) OnRspError(pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {}

// 订阅行情应答
func (p *MarketDataSpi) OnRspSubMarketData(pSpecificInstrument CThostFtdcSpecificInstrumentField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 取消订阅行情应答
func (p *MarketDataSpi) OnRspUnSubMarketData(pSpecificInstrument CThostFtdcSpecificInstrumentField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 订阅询价应答
func (p *MarketDataSpi) OnRspSubForQuoteRsp(pSpecificInstrument CThostFtdcSpecificInstrumentField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 取消订阅询价应答
func (p *MarketDataSpi) OnRspUnSubForQuoteRsp(pSpecificInstrument CThostFtdcSpecificInstrumentField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 深度行情通知
func (p *MarketDataSpi) OnRtnDepthMarketData(pDepthMarketData CThostFtdcDepthMarketDataField) {}

// 询价通知
func (p *MarketDataSpi) OnRtnForQuoteRsp(pForQuoteRsp CThostFtdcForQuoteRspField) {}
