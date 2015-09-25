# Author: Long He
# Andrew ID: longh
# Date: 12/01/2014
# Version: 1.7

# primesSieve takes a list and changes the value in this list
def primeSieve(isComposite):
    biggestPrime = 2 # will hold the biggest prime found so far
    while biggestPrime < len(isComposite):
        # knock out all multiples of biggestPrime
        i = 2 * biggestPrime
        while i < len(isComposite):
            isComposite[i] = True
            i += biggestPrime
        # find the next biggest non-composite number
        biggestPrime += 1
        while biggestPrime < len(isComposite) and isComposite[biggestPrime]:
            biggestPrime += 1

# primes takes the number n and returns a sorted list of the prime numbers <= n 
def primes(n):
    n += 1
    # create a new list
    composites = [False]*n
    # primeSieve() can change the values of composites
    primeSieve(composites)
    primesList = []
    # search from 0 to n (n has increased by 1)
    for i in range (n):
        if not composites[i] and i >= 2:
            # add the found primes into the sorted list
            primesList.append(i) 
    return primesList

# sumOfPrimes takes number k and returns two primes a and b
def sumOfPrimes(k):
    if k == 2:
        return 0, 0
    # get a sorted list of the prime numbers <= k
    primesList = primes(k)
    l = len(primesList)
    # search from the start and end position
    i = 0
    j = l - 1
    while True:
        # make sure a <= b
        if i > j and i != 0:
            # or returns 0, 0 if no such primes exist
            return 0, 0
        # make sure a + b = k
        if primesList[i] + primesList[j] == k:
            return primesList[i], primesList[j]
        elif primesList[i] + primesList[j] < k:
            i += 1 # move to find the target
        else:
            j -= 1 # move to find the target

# allSumOfPrimes takes numbr k and returns a list of all pairs (a, b)
def allSumOfPrimes(k):
    # initialize the output
    all = []
    #  get a sorted list of the prime numbers <= k
    primesList = primes(k)
    l = len(primesList)
    # search from the start and end position
    i = 0
    j = l - 1
    while True:
        # make sure a <= b
        if i > j and i != 0:
            return all
        # make sure a + b == k
        if primesList[i] + primesList[j] == k:
            # add the found prime pairs into the list, if there are no (a, b)
            # pairs, return the empty list []
            all.append((primesList[i], primesList[j]))
            i += 1 # continue to find the next target
        elif primesList[i] + primesList[j] < k:
            i += 1 # move to find the target
        else:
            j -= 1 # move to find the target

# goldbach tests all the even integers <= k to see whether they can be written 
# as the sum of 2 primes
def goldbach(n):
    # initialize the output
    res = []
    res2 = False
    # search z from 2 to n
    for i in range (2, n + 1):
        # make sure z is even and z = a + b (a, b are primes, a <= b)
        if i % 2 == 0 and sumOfPrimes(i) != (0, 0):
            (a, b) = sumOfPrimes(i)
            res.append((i, a, b))
    # return True is every even integer <= k is represented within the list
    if len(res) == n / 2 - 1:
        res2 = True
    return res, res2

# goldbachWidth takes number k and returns a dictionary (map) D such that D[z]
# is the number of ways each even number 2 < z <= k can be written as the sum 
# of two primes
def goldbachWidth(k):
    # initialize output
    D = {}
    # search from 2 to k
    for i in range (2, k + 1):
        # make sure z is even and meets the requirements
        if i % 2 == 0 and sumOfPrimes(i) != (0, 0):
            # save the sum of number of ways each even number 2 < z <= k can be
            # written as the sum of two primes
            D[i] = len(allSumOfPrimes(i))
    return D