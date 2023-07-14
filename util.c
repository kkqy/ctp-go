#include "include/macro.h"
#include <stdlib.h>
#include <string.h>
#include "include/util.h"

StringArray *NewStringArray(int size)
{
    char **pArray = (char **)calloc(sizeof(char *), size);
    StringArray *stringArray = (StringArray *)malloc(sizeof(StringArray));
    stringArray->pArray = pArray;
    stringArray->nSize = size;
    return stringArray;
}

void SetStringArray(StringArray *stringArray, int index, char *s)
{
    if (stringArray->pArray[index] != NULL)
    {
        free(stringArray->pArray[index]);
    }
    unsigned int len = strlen(s);
    stringArray->pArray[index] = (char *)malloc(len + 1);
    strcpy(stringArray->pArray[index], s);
    stringArray->pArray[index][len] = '\0';
}

void DeleteStringArray(StringArray *stringArray)
{
    if (stringArray != NULL)
    {
        int i;
        for (i = 0; i < stringArray->nSize; i++)
        {
            if (stringArray->pArray[i] != NULL)
            {
                free(stringArray->pArray[i]);
            }
        }
        free(stringArray);
    }
}