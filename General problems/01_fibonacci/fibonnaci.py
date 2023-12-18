
def fibonacci(n):
  fib = []
  if n>=1:
    fib.append(0)
  if n>=2:
    fib.append(1)
  for x in range (2, n):
    fib.append(fib[x-1]+fib[x-2])
  return fib      
  

print(fibonacci(10))