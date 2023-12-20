def mergeAlternateString(word1,word2):
    result = ""
    l = len(word1)
    if len(word2) > l:
        l = len(word2)
    for x in range(0,l):
        if x <  len(word1):
            result += word1[x]
        if x <  len(word2):
            result += word2[x]  
    return result          


print(mergeAlternateString("abc","defgg"))