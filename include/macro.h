#ifndef _MACRO_H
    #define _NO_CRT_STDIO_INLINE
    #define __USE_MINGW_ANSI_STDIO 1
    #if defined(_MSC_VER) && defined(_DLL)
        #ifdef _LIB
            #define DLLAPI _declspec(dllexport)
        #else
            #define DLLAPI _declspec(dllimport)
        #endif
    #else
        #define DLLAPI
    #endif
#define _MACRO_H
#endif