#ifndef _WRAP_H
#ifdef __cplusplus
extern "C"
{
#endif
    DLLAPI CThostFtdcMdApi *CThostFtdcMdApi_CreateFtdcMdApi(const char *pszFlowPath, bool bIsUsingUdp, bool bIsMulticast);
    DLLAPI const char *CThostFtdcMdApi_GetApiVersion();
    DLLAPI CThostFtdcTraderApi *CThostFtdcTraderApi_CreateFtdcTraderApi(const char *pszFlowPath);
    DLLAPI const char *CThostFtdcTraderApi_GetApiVersion();
#ifdef __cplusplus
}
#endif
#define _WRAP_H
#endif