from unionfind import UnionFind


def find_minimal(uf, edges):
    # now we have to search for the lowest cost edge that connects them
    m = 9999999999999
    for i in range(1, 501):
        for j in range(1, 501):
            if i != j and not uf.connected(i, j):
                for eindex in start_at[i]:
                    if edges[eindex][1] == i and edges[eindex][2] == j and edges[eindex][3] < m:
                        m = edges[eindex][3]
                        print(m)


def cluster(edges, start_at):
    uf = UnionFind()

    sortede = sorted(edges, key=lambda x: x[3])

    for k in range(1, 501):
        uf.add(k)

    for _, u, v, w in sortede:

        if len(list(uf.components())) == 4:
            break
        elif not uf.connected(u, v):
            uf.union(u, v)

    find_minimal(uf, edges)


if __name__ == "__main__":
    edges = []
    start_at = {}
    with open("course_3_p2q1.txt", "r") as f:
        for lindex, line in enumerate(f):
            i = line.split()
            start = int(i[0])
            end = int(i[1])
            edges.append((lindex, start, end, int(i[2])))
            if start not in start_at:
                start_at[start] = []
            if end not in start_at:
                start_at[end] = []
            # we need these later when we need to quickly scan through edjes that connect clusters
            start_at[start].append(lindex)
            start_at[end].append(lindex)

    cluster(edges, start_at)
