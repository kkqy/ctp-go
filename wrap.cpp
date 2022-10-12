//go:build exclude
#include "include/macro.h"
#include "include/ThostFtdcMdApi.h"
#include "include/ThostFtdcTraderApi.h"
#include "include/macro.h"
#include "include/wrap.h"

CThostFtdcMdApi* CThostFtdcMdApi_CreateFtdcMdApi(const char *pszFlowPath , const bool bIsUsingUdp, const bool bIsMulticast){
    return CThostFtdcMdApi::CreateFtdcMdApi(pszFlowPath,bIsUsingUdp,bIsMulticast);
}
const char * CThostFtdcMdApi_GetApiVersion(){
    return CThostFtdcMdApi::GetApiVersion();
}
CThostFtdcTraderApi *CThostFtdcTraderApi_CreateFtdcTraderApi(const char *pszFlowPath){
    return CThostFtdcTraderApi::CreateFtdcTraderApi(pszFlowPath);
}
const char *CThostFtdcTraderApi_GetApiVersion(){
    return CThostFtdcTraderApi::GetApiVersion();
}