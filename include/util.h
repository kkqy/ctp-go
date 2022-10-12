#ifndef _UTIL_H
typedef struct { 
    char** pArray; 
    int nSize; 
} StringArray;
StringArray *new_string_array(int size);
void set_string_array(StringArray *stringArray, int index, char *s);
void delete_string_array(StringArray *stringArray);
#define _UTIL_H
#endif