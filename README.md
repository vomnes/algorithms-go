# graph-algorithms-go
Graph algorithms implementation in Golang

## Breadth-First-Search
Breadth-first search (BFS) is an algorithm for traversing or searching tree or graph data structures. It starts at the tree root, and explores all of **the neighbor nodes** at the present depth **prior to moving on to the nodes at the next depth level**. (_Source Wikipedia_)

### Pseudocode
```
1  procedure BFS(G, start_v) is
2      let Q be a queue
3      label start_v as discovered
4      Q.enqueue(start_v)
5      while Q is not empty do
6          v := Q.dequeue()
7          if v is the goal then
8              return v
9          for all edges from v to w in G.adjacentEdges(v) do
10             if w is not labeled as discovered then
11                 label w as discovered
12                 w.parent := v
13                 Q.enqueue(w)
```

## Depth-First-Search
Depth-first search (DFS) is an algorithm for traversing or searching tree or graph data structures. The algorithm starts at the root node and **explores as far as possible along each branch** before backtracking. (_Source Wikipedia_)

### Pseudocode
```
1  procedure DFS-iterative(G, v) is
2      let S be a stack
3      S.push(v)
4      while S is not empty do
5          v = S.pop()
6          if v is not labeled as discovered then
7              label v as discovered
8              for all edges from v to w in G.adjacentEdges(v) do
9                  S.push(w)
```
