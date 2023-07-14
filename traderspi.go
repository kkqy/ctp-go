package ctp

type TraderSpi struct {
}

// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
func (p *TraderSpi) OnFrontConnected() {}

// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
// @nReason param 错误原因
//
//	0x1001 网络读失败
//	0x1002 网络写失败
//	0x2001 接收心跳超时
//	0x2002 发送心跳失败
//	0x2003 收到错误报文
func (p *TraderSpi) OnFrontDisconnected(nReason int) {}

// 心跳超时警告。当长时间未收到报文时，该方法被调用。
// @nTimeLapse param 距离上次接收报文的时间
func (p *TraderSpi) OnHeartBeatWarning(nTimeLapse int) {}

// 客户端认证响应
func (p *TraderSpi) OnRspAuthenticate(pRspAuthenticateField CThostFtdcRspAuthenticateField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 登录请求响应
func (p *TraderSpi) OnRspUserLogin(pRspUserLogin CThostFtdcRspUserLoginField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 登出请求响应
func (p *TraderSpi) OnRspUserLogout(pUserLogout CThostFtdcUserLogoutField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 用户口令更新请求响应
func (p *TraderSpi) OnRspUserPasswordUpdate(pUserPasswordUpdate CThostFtdcUserPasswordUpdateField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 资金账户口令更新请求响应
func (p *TraderSpi) OnRspTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate CThostFtdcTradingAccountPasswordUpdateField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 查询用户当前支持的认证模式的回复
func (p *TraderSpi) OnRspUserAuthMethod(pRspUserAuthMethod CThostFtdcRspUserAuthMethodField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 获取图形验证码请求的回复
func (p *TraderSpi) OnRspGenUserCaptcha(pRspGenUserCaptcha CThostFtdcRspGenUserCaptchaField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 获取短信验证码请求的回复
func (p *TraderSpi) OnRspGenUserText(pRspGenUserText CThostFtdcRspGenUserTextField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 报单录入请求响应
func (p *TraderSpi) OnRspOrderInsert(pInputOrder CThostFtdcInputOrderField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 预埋单录入请求响应
func (p *TraderSpi) OnRspParkedOrderInsert(pParkedOrder CThostFtdcParkedOrderField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 预埋撤单录入请求响应
func (p *TraderSpi) OnRspParkedOrderAction(pParkedOrderAction CThostFtdcParkedOrderActionField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 报单操作请求响应
func (p *TraderSpi) OnRspOrderAction(pInputOrderAction CThostFtdcInputOrderActionField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 查询最大报单数量响应
func (p *TraderSpi) OnRspQueryMaxOrderVolume(pQueryMaxOrderVolume CThostFtdcQueryMaxOrderVolumeField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 投资者结算结果确认响应
func (p *TraderSpi) OnRspSettlementInfoConfirm(pSettlementInfoConfirm CThostFtdcSettlementInfoConfirmField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 删除预埋单响应
func (p *TraderSpi) OnRspRemoveParkedOrder(pRemoveParkedOrder CThostFtdcRemoveParkedOrderField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 删除预埋撤单响应
func (p *TraderSpi) OnRspRemoveParkedOrderAction(pRemoveParkedOrderAction CThostFtdcRemoveParkedOrderActionField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 执行宣告录入请求响应
func (p *TraderSpi) OnRspExecOrderInsert(pInputExecOrder CThostFtdcInputExecOrderField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 执行宣告操作请求响应
func (p *TraderSpi) OnRspExecOrderAction(pInputExecOrderAction CThostFtdcInputExecOrderActionField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 询价录入请求响应
func (p *TraderSpi) OnRspForQuoteInsert(pInputForQuote CThostFtdcInputForQuoteField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 报价录入请求响应
func (p *TraderSpi) OnRspQuoteInsert(pInputQuote CThostFtdcInputQuoteField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 报价操作请求响应
func (p *TraderSpi) OnRspQuoteAction(pInputQuoteAction CThostFtdcInputQuoteActionField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 批量报单操作请求响应
func (p *TraderSpi) OnRspBatchOrderAction(pInputBatchOrderAction CThostFtdcInputBatchOrderActionField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 期权自对冲录入请求响应
func (p *TraderSpi) OnRspOptionSelfCloseInsert(pInputOptionSelfClose CThostFtdcInputOptionSelfCloseField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 期权自对冲操作请求响应
func (p *TraderSpi) OnRspOptionSelfCloseAction(pInputOptionSelfCloseAction CThostFtdcInputOptionSelfCloseActionField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 申请组合录入请求响应
func (p *TraderSpi) OnRspCombActionInsert(pInputCombAction CThostFtdcInputCombActionField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询报单响应
func (p *TraderSpi) OnRspQryOrder(pOrder CThostFtdcOrderField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询成交响应
func (p *TraderSpi) OnRspQryTrade(pTrade CThostFtdcTradeField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询投资者持仓响应
func (p *TraderSpi) OnRspQryInvestorPosition(pInvestorPosition CThostFtdcInvestorPositionField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询资金账户响应
func (p *TraderSpi) OnRspQryTradingAccount(pTradingAccount CThostFtdcTradingAccountField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询投资者响应
func (p *TraderSpi) OnRspQryInvestor(pInvestor CThostFtdcInvestorField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询交易编码响应
func (p *TraderSpi) OnRspQryTradingCode(pTradingCode CThostFtdcTradingCodeField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询合约保证金率响应
func (p *TraderSpi) OnRspQryInstrumentMarginRate(pInstrumentMarginRate CThostFtdcInstrumentMarginRateField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询合约手续费率响应
func (p *TraderSpi) OnRspQryInstrumentCommissionRate(pInstrumentCommissionRate CThostFtdcInstrumentCommissionRateField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询交易所响应
func (p *TraderSpi) OnRspQryExchange(pExchange CThostFtdcExchangeField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询产品响应
func (p *TraderSpi) OnRspQryProduct(pProduct CThostFtdcProductField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询合约响应
func (p *TraderSpi) OnRspQryInstrument(pInstrument CThostFtdcInstrumentField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询行情响应
func (p *TraderSpi) OnRspQryDepthMarketData(pDepthMarketData CThostFtdcDepthMarketDataField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询投资者结算结果响应
func (p *TraderSpi) OnRspQrySettlementInfo(pSettlementInfo CThostFtdcSettlementInfoField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询转帐银行响应
func (p *TraderSpi) OnRspQryTransferBank(pTransferBank CThostFtdcTransferBankField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询投资者持仓明细响应
func (p *TraderSpi) OnRspQryInvestorPositionDetail(pInvestorPositionDetail CThostFtdcInvestorPositionDetailField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询客户通知响应
func (p *TraderSpi) OnRspQryNotice(pNotice CThostFtdcNoticeField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询结算信息确认响应
func (p *TraderSpi) OnRspQrySettlementInfoConfirm(pSettlementInfoConfirm CThostFtdcSettlementInfoConfirmField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询投资者持仓明细响应
func (p *TraderSpi) OnRspQryInvestorPositionCombineDetail(pInvestorPositionCombineDetail CThostFtdcInvestorPositionCombineDetailField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 查询保证金监管系统经纪公司资金账户密钥响应
func (p *TraderSpi) OnRspQryCFMMCTradingAccountKey(pCFMMCTradingAccountKey CThostFtdcCFMMCTradingAccountKeyField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询仓单折抵信息响应
func (p *TraderSpi) OnRspQryEWarrantOffset(pEWarrantOffset CThostFtdcEWarrantOffsetField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询投资者品种/跨品种保证金响应
func (p *TraderSpi) OnRspQryInvestorProductGroupMargin(pInvestorProductGroupMargin CThostFtdcInvestorProductGroupMarginField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询交易所保证金率响应
func (p *TraderSpi) OnRspQryExchangeMarginRate(pExchangeMarginRate CThostFtdcExchangeMarginRateField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询交易所调整保证金率响应
func (p *TraderSpi) OnRspQryExchangeMarginRateAdjust(pExchangeMarginRateAdjust CThostFtdcExchangeMarginRateAdjustField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询汇率响应
func (p *TraderSpi) OnRspQryExchangeRate(pExchangeRate CThostFtdcExchangeRateField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询二级代理操作员银期权限响应
func (p *TraderSpi) OnRspQrySecAgentACIDMap(pSecAgentACIDMap CThostFtdcSecAgentACIDMapField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询产品报价汇率
func (p *TraderSpi) OnRspQryProductExchRate(pProductExchRate CThostFtdcProductExchRateField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询产品组
func (p *TraderSpi) OnRspQryProductGroup(pProductGroup CThostFtdcProductGroupField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询做市商合约手续费率响应
func (p *TraderSpi) OnRspQryMMInstrumentCommissionRate(pMMInstrumentCommissionRate CThostFtdcMMInstrumentCommissionRateField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询做市商期权合约手续费响应
func (p *TraderSpi) OnRspQryMMOptionInstrCommRate(pMMOptionInstrCommRate CThostFtdcMMOptionInstrCommRateField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询报单手续费响应
func (p *TraderSpi) OnRspQryInstrumentOrderCommRate(pInstrumentOrderCommRate CThostFtdcInstrumentOrderCommRateField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询资金账户响应
func (p *TraderSpi) OnRspQrySecAgentTradingAccount(pTradingAccount CThostFtdcTradingAccountField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询二级代理商资金校验模式响应
func (p *TraderSpi) OnRspQrySecAgentCheckMode(pSecAgentCheckMode CThostFtdcSecAgentCheckModeField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询二级代理商信息响应
func (p *TraderSpi) OnRspQrySecAgentTradeInfo(pSecAgentTradeInfo CThostFtdcSecAgentTradeInfoField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询期权交易成本响应
func (p *TraderSpi) OnRspQryOptionInstrTradeCost(pOptionInstrTradeCost CThostFtdcOptionInstrTradeCostField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询期权合约手续费响应
func (p *TraderSpi) OnRspQryOptionInstrCommRate(pOptionInstrCommRate CThostFtdcOptionInstrCommRateField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询执行宣告响应
func (p *TraderSpi) OnRspQryExecOrder(pExecOrder CThostFtdcExecOrderField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询询价响应
func (p *TraderSpi) OnRspQryForQuote(pForQuote CThostFtdcForQuoteField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询报价响应
func (p *TraderSpi) OnRspQryQuote(pQuote CThostFtdcQuoteField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询期权自对冲响应
func (p *TraderSpi) OnRspQryOptionSelfClose(pOptionSelfClose CThostFtdcOptionSelfCloseField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询投资单元响应
func (p *TraderSpi) OnRspQryInvestUnit(pInvestUnit CThostFtdcInvestUnitField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询组合合约安全系数响应
func (p *TraderSpi) OnRspQryCombInstrumentGuard(pCombInstrumentGuard CThostFtdcCombInstrumentGuardField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询申请组合响应
func (p *TraderSpi) OnRspQryCombAction(pCombAction CThostFtdcCombActionField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询转帐流水响应
func (p *TraderSpi) OnRspQryTransferSerial(pTransferSerial CThostFtdcTransferSerialField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询银期签约关系响应
func (p *TraderSpi) OnRspQryAccountregister(pAccountregister CThostFtdcAccountregisterField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 错误应答
func (p *TraderSpi) OnRspError(pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {}

// 报单通知
func (p *TraderSpi) OnRtnOrder(pOrder CThostFtdcOrderField) {}

// 成交通知
func (p *TraderSpi) OnRtnTrade(pTrade CThostFtdcTradeField) {}

// 报单录入错误回报
func (p *TraderSpi) OnErrRtnOrderInsert(pInputOrder CThostFtdcInputOrderField, pRspInfo CThostFtdcRspInfoField) {
}

// 报单操作错误回报
func (p *TraderSpi) OnErrRtnOrderAction(pOrderAction CThostFtdcOrderActionField, pRspInfo CThostFtdcRspInfoField) {
}

// 合约交易状态通知
func (p *TraderSpi) OnRtnInstrumentStatus(pInstrumentStatus CThostFtdcInstrumentStatusField) {}

// 交易所公告通知
func (p *TraderSpi) OnRtnBulletin(pBulletin CThostFtdcBulletinField) {}

// 交易通知
func (p *TraderSpi) OnRtnTradingNotice(pTradingNoticeInfo CThostFtdcTradingNoticeInfoField) {}

// 提示条件单校验错误
func (p *TraderSpi) OnRtnErrorConditionalOrder(pErrorConditionalOrder CThostFtdcErrorConditionalOrderField) {
}

// 执行宣告通知
func (p *TraderSpi) OnRtnExecOrder(pExecOrder CThostFtdcExecOrderField) {}

// 执行宣告录入错误回报
func (p *TraderSpi) OnErrRtnExecOrderInsert(pInputExecOrder CThostFtdcInputExecOrderField, pRspInfo CThostFtdcRspInfoField) {
}

// 执行宣告操作错误回报
func (p *TraderSpi) OnErrRtnExecOrderAction(pExecOrderAction CThostFtdcExecOrderActionField, pRspInfo CThostFtdcRspInfoField) {
}

// 询价录入错误回报
func (p *TraderSpi) OnErrRtnForQuoteInsert(pInputForQuote CThostFtdcInputForQuoteField, pRspInfo CThostFtdcRspInfoField) {
}

// 报价通知
func (p *TraderSpi) OnRtnQuote(pQuote CThostFtdcQuoteField) {}

// 报价录入错误回报
func (p *TraderSpi) OnErrRtnQuoteInsert(pInputQuote CThostFtdcInputQuoteField, pRspInfo CThostFtdcRspInfoField) {
}

// 报价操作错误回报
func (p *TraderSpi) OnErrRtnQuoteAction(pQuoteAction CThostFtdcQuoteActionField, pRspInfo CThostFtdcRspInfoField) {
}

// 询价通知
func (p *TraderSpi) OnRtnForQuoteRsp(pForQuoteRsp CThostFtdcForQuoteRspField) {}

// 保证金监控中心用户令牌
func (p *TraderSpi) OnRtnCFMMCTradingAccountToken(pCFMMCTradingAccountToken CThostFtdcCFMMCTradingAccountTokenField) {
}

// 批量报单操作错误回报
func (p *TraderSpi) OnErrRtnBatchOrderAction(pBatchOrderAction CThostFtdcBatchOrderActionField, pRspInfo CThostFtdcRspInfoField) {
}

// 期权自对冲通知
func (p *TraderSpi) OnRtnOptionSelfClose(pOptionSelfClose CThostFtdcOptionSelfCloseField) {}

// 期权自对冲录入错误回报
func (p *TraderSpi) OnErrRtnOptionSelfCloseInsert(pInputOptionSelfClose CThostFtdcInputOptionSelfCloseField, pRspInfo CThostFtdcRspInfoField) {
}

// 期权自对冲操作错误回报
func (p *TraderSpi) OnErrRtnOptionSelfCloseAction(pOptionSelfCloseAction CThostFtdcOptionSelfCloseActionField, pRspInfo CThostFtdcRspInfoField) {
}

// 申请组合通知
func (p *TraderSpi) OnRtnCombAction(pCombAction CThostFtdcCombActionField) {}

// 申请组合录入错误回报
func (p *TraderSpi) OnErrRtnCombActionInsert(pInputCombAction CThostFtdcInputCombActionField, pRspInfo CThostFtdcRspInfoField) {
}

// 请求查询签约银行响应
func (p *TraderSpi) OnRspQryContractBank(pContractBank CThostFtdcContractBankField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询预埋单响应
func (p *TraderSpi) OnRspQryParkedOrder(pParkedOrder CThostFtdcParkedOrderField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询预埋撤单响应
func (p *TraderSpi) OnRspQryParkedOrderAction(pParkedOrderAction CThostFtdcParkedOrderActionField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询交易通知响应
func (p *TraderSpi) OnRspQryTradingNotice(pTradingNotice CThostFtdcTradingNoticeField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询经纪公司交易参数响应
func (p *TraderSpi) OnRspQryBrokerTradingParams(pBrokerTradingParams CThostFtdcBrokerTradingParamsField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询经纪公司交易算法响应
func (p *TraderSpi) OnRspQryBrokerTradingAlgos(pBrokerTradingAlgos CThostFtdcBrokerTradingAlgosField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询监控中心用户令牌
func (p *TraderSpi) OnRspQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken CThostFtdcQueryCFMMCTradingAccountTokenField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 银行发起银行资金转期货通知
func (p *TraderSpi) OnRtnFromBankToFutureByBank(pRspTransfer CThostFtdcRspTransferField) {}

// 银行发起期货资金转银行通知
func (p *TraderSpi) OnRtnFromFutureToBankByBank(pRspTransfer CThostFtdcRspTransferField) {}

// 银行发起冲正银行转期货通知
func (p *TraderSpi) OnRtnRepealFromBankToFutureByBank(pRspRepeal CThostFtdcRspRepealField) {}

// 银行发起冲正期货转银行通知
func (p *TraderSpi) OnRtnRepealFromFutureToBankByBank(pRspRepeal CThostFtdcRspRepealField) {}

// 期货发起银行资金转期货通知
func (p *TraderSpi) OnRtnFromBankToFutureByFuture(pRspTransfer CThostFtdcRspTransferField) {}

// 期货发起期货资金转银行通知
func (p *TraderSpi) OnRtnFromFutureToBankByFuture(pRspTransfer CThostFtdcRspTransferField) {}

// 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
func (p *TraderSpi) OnRtnRepealFromBankToFutureByFutureManual(pRspRepeal CThostFtdcRspRepealField) {}

// 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
func (p *TraderSpi) OnRtnRepealFromFutureToBankByFutureManual(pRspRepeal CThostFtdcRspRepealField) {}

// 期货发起查询银行余额通知
func (p *TraderSpi) OnRtnQueryBankBalanceByFuture(pNotifyQueryAccount CThostFtdcNotifyQueryAccountField) {
}

// 期货发起银行资金转期货错误回报
func (p *TraderSpi) OnErrRtnBankToFutureByFuture(pReqTransfer CThostFtdcReqTransferField, pRspInfo CThostFtdcRspInfoField) {
}

// 期货发起期货资金转银行错误回报
func (p *TraderSpi) OnErrRtnFutureToBankByFuture(pReqTransfer CThostFtdcReqTransferField, pRspInfo CThostFtdcRspInfoField) {
}

// 系统运行时期货端手工发起冲正银行转期货错误回报
func (p *TraderSpi) OnErrRtnRepealBankToFutureByFutureManual(pReqRepeal CThostFtdcReqRepealField, pRspInfo CThostFtdcRspInfoField) {
}

// 系统运行时期货端手工发起冲正期货转银行错误回报
func (p *TraderSpi) OnErrRtnRepealFutureToBankByFutureManual(pReqRepeal CThostFtdcReqRepealField, pRspInfo CThostFtdcRspInfoField) {
}

// 期货发起查询银行余额错误回报
func (p *TraderSpi) OnErrRtnQueryBankBalanceByFuture(pReqQueryAccount CThostFtdcReqQueryAccountField, pRspInfo CThostFtdcRspInfoField) {
}

// 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
func (p *TraderSpi) OnRtnRepealFromBankToFutureByFuture(pRspRepeal CThostFtdcRspRepealField) {}

// 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
func (p *TraderSpi) OnRtnRepealFromFutureToBankByFuture(pRspRepeal CThostFtdcRspRepealField) {}

// 期货发起银行资金转期货应答
func (p *TraderSpi) OnRspFromBankToFutureByFuture(pReqTransfer CThostFtdcReqTransferField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 期货发起期货资金转银行应答
func (p *TraderSpi) OnRspFromFutureToBankByFuture(pReqTransfer CThostFtdcReqTransferField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 期货发起查询银行余额应答
func (p *TraderSpi) OnRspQueryBankAccountMoneyByFuture(pReqQueryAccount CThostFtdcReqQueryAccountField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 银行发起银期开户通知
func (p *TraderSpi) OnRtnOpenAccountByBank(pOpenAccount CThostFtdcOpenAccountField) {}

// 银行发起银期销户通知
func (p *TraderSpi) OnRtnCancelAccountByBank(pCancelAccount CThostFtdcCancelAccountField) {}

// 银行发起变更银行账号通知
func (p *TraderSpi) OnRtnChangeAccountByBank(pChangeAccount CThostFtdcChangeAccountField) {}
