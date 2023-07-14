//go:build windows
#include "include/macro.h"
#include "include/ThostFtdcMdApi.h"
#include "include/ThostFtdcTraderApi.h"
#include "include/wrap.h"

CThostFtdcMdApi *CThostFtdcMdApi::CreateFtdcMdApi(const char *pszFlowPath, const bool bIsUsingUdp, const bool bIsMulticast)
{
    return CThostFtdcMdApi_CreateFtdcMdApi(pszFlowPath, bIsUsingUdp, bIsMulticast);
}
const char *CThostFtdcMdApi::GetApiVersion()
{
    return CThostFtdcMdApi_GetApiVersion();
}
CThostFtdcTraderApi *CThostFtdcTraderApi::CreateFtdcTraderApi(const char *pszFlowPath)
{
    return CThostFtdcTraderApi_CreateFtdcTraderApi(pszFlowPath);
}
const char *CThostFtdcTraderApi::GetApiVersion()
{
    return CThostFtdcTraderApi_GetApiVersion();
}