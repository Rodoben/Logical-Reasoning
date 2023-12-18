// using a list
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

print("__________________")

// using for loop
def fibonacciLoop(n):
  a, b = 0,1
  for x in range (0,n):
    print(a)
    a,b = b, a+b  
     
fibonacciLoop(10)     