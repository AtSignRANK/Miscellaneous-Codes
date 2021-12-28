import math


def isPrime(n):
    if n <= 1:
        return False
    for i in range(2, int(math.sqrt(n))):
        if n % i == 0:
            return False
    return True


def printPrime():
    n = 2
    while True:
        if isPrime(n):
            print(n)
        n += 1


printPrime()
