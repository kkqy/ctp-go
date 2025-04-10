#ifndef _UTIL_H
typedef struct { 
    char** pArray; 
    int nSize; 
} StringArray;
StringArray *NewStringArray(int size);
void SetStringArray(StringArray *stringArray, int index, char *s);
void DeleteStringArray(StringArray *stringArray);
#define _UTIL_H
#endif