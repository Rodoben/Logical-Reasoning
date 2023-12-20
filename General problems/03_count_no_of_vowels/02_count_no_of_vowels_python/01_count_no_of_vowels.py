

def countVowels(s):
   vowels = dict()
   for x in s:
      x = x.lower()
      if x in "aeiou":
        vowels[x] = vowels.get(x,0)+1
   return vowels    

print(countVowels("aasseecciieoO"))   
print(countVowels(["a","a","e","e","i"]))   
print(countVowels(("a","a","e","e","i")))   