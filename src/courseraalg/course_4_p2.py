import itertools
import math
import copy

INF = 99999999999999


def get_input(f):
    nodes = []
    D = []

    lines = open(f).readlines()
    print(len(lines))
    N = len(lines)
    for l in lines:
        items = l.split()
        nodes.append((float(items[0]), float(items[1])))

    for i in range(N):
        row = []
        for j in range(N):
            if i != j:
                dist = math.pow(nodes[i][0] - nodes[j][0], 2) + math.pow(nodes[i][1] - nodes[j][1], 2)
                row.append(math.sqrt(dist))
            else:
                row.append(0)
        D.append(row)

    return D, N


def run_soln(D, N):

    A = {}
    A[(0,), 0] = 0

    for m in range(1, N):
        print(m)
        A_new = {}
        # to form a tuple that contains 0, gen 1 that doesnt of m-1 and add it
        # m would have been +1 here since we want inclusive. so we have m + 1 - 1 = m
        for s in itertools.combinations(range(1, N), m):
            S = (0,) + s
            for j in S:
                if j != 0:
                    minimum = INF
                    for k in S:  # min k in S, k != j
                        if k != j:
                            # tuple(x for x in tupleX if condition)
                            minus_j = tuple(x for x in S if x != j)
                            val = -1
                            if (minus_j, k) not in A and k == 0:
                                val = INF
                            else:
                                val = A[minus_j, k] + D[k][j]
                            if val < minimum:
                                minimum = val
                    A_new[S, j] = minimum
        A = copy.copy(A_new)

    minimum = INF
    for j in range(1, N):
        val = A[tuple(range(N)), j] + D[j][0]
        if val < minimum:
            minimum = val
    return minimum


"""
https://www.coursera.org/learn/algorithms-npcomplete/discussions/weeks/2/threads/FNo8tpFPEeeH2hLapWklpg

The crazy breakup into two graphs solution here came from the above post

Basically, what I think is happening is that we have two graphs far apart from each other,
so they must be connected by the two shortest edges connecting the two graphs
so you find a tour in graph 1,
find a tour in graph 2
then you delete the same edge in both graphs where you "graft" the two together (delete it twice because if you only did it from one graph you would end up with a tour with a weird cut across it
"""

D, N = get_input("course_4_p2_sec1.txt")
sol1 = run_soln(D, N)
print(sol1)

D2, N2 = get_input("course_4_p2_sec2.txt")
sol2 = run_soln(D2, N2)
print(sol2)

print(sol1 + sol2 - 2 * D[11][12])
