from unionfind import UnionFind


def one_two_away():
    two_poss = {}
    for i in range(24):
        for j in range(24):
            if j != i:
                two_poss[(1 << i) + (1 << j)] = 1
    return [1 << i for i in range(24)], two_poss.keys()


def cluster(nodes):
    uf = UnionFind()
    one_diff, two_diff = one_two_away()

    for v in nodes:
        uf.add(v)

    for vindex, v in enumerate(nodes):

        # squash all the nodes with distance 1 away into me
        od = [(v ^ i) for i in one_diff]
        td = [(v ^ i) for i in two_diff]
        for d in od + td:
            if d in nodes and not uf.connected(v, d):
                uf.union(v, d)
                print("smashed", v, d)

    print(len(list(uf.components())))


if __name__ == "__main__":
    nodes = {}
    with open("course_3_p2q2.txt", "r") as f:
        # 200000 24
        for lindex, line in enumerate(f):
            i = line.split()
            key = int("".join(i), base=2)
            nodes[key] = 1

    cluster(nodes)
