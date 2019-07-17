satisfied = False


# n choose 2 is n!/(2!(n-2)!

k = 2
while not satisfied:
    s = 0
    for i in range(0, k):
        for j in range(i + 1, k):
            s += 1 / 365
    if s >= 1:
        print(k)
        break
    else:
        k += 1
