
def palindrome(s):
    s = s.lower()
    for i in range (len(s)//2):
        if s[i]!= s[len(s)-1-i]:
            return False 
    return True
print(palindrome("naman"))    
print(palindrome("A man a plan a canal Panama")) 