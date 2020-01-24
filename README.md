# Graph Algorithms
Graph algorithms implementation in Golang

## Breadth-first Search
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

## Depth-first Search
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

## A* Search Algorithm
A* (pronounced "A-star") is a graph traversal and path search algorithm, which is often used in computer science due to its completeness, optimality, and optimal efficiency. In practical travel-routing systems, it is generally outperformed by algorithms which can pre-process the graph to attain better performance, as well as memory-bounded approaches; however, A* is still the best solution in many cases.

At each iteration of its main loop, **A Star needs to determine which of its paths to extend**. It does so based **on the cost of the path and an estimate of the cost required to extend the path all the way to the goal**.  
Specifically, A* selects the path that minimizes :  

> f(n) = g(n) + h(n)

*g(n)* is the cost of the path from the start node to n
*h(n)* is a heuristic function that estimates the cost of the cheapest path from n to the goal for example using a [Euclidean distance formula](https://en.wikipedia.org/wiki/Euclidean_distance)
(_Source Wikipedia_)

### Pseudocode
```
function reconstruct_path(cameFrom, current)
    total_path := {current}
    while current in cameFrom.Keys:
        total_path.prepend(current)
        current := cameFrom[current]
    return total_path

// A* finds a path from start to goal.
// h is the heuristic function. h(n) estimates the cost to reach goal from node n.
function A_Star(start, goal, h)
    // The set of discovered nodes that may need to be (re-)expanded.
    // Initially, only the start node is known.
    openSet := {start}

    // For node n, cameFrom[n] is the node immediately preceding it on the cheapest path from start to n currently known.
    cameFrom := an empty map

    // For node n, gScore[n] is the cost of the cheapest path from start to n currently known.
    gScore := map with default value of Infinity
    gScore[start] := 0

    // For node n, fScore[n] := gScore[n] + h(n).
    fScore := map with default value of Infinity
    fScore[start] := h(start)

    while openSet is not empty
        current := the node in openSet having the lowest fScore[] value
        if current = goal
            return reconstruct_path(cameFrom, current)

        openSet.Remove(current)
        for each neighbor of current
            // d(current,neighbor) is the weight of the edge from current to neighbor
            // tentative_gScore is the distance from start to the neighbor through current
            tentative_gScore := gScore[current] + d(current, neighbor)
            if tentative_gScore < gScore[neighbor]
                // This path to neighbor is better than any previous one. Record it!
                cameFrom[neighbor] := current
                gScore[neighbor] := tentative_gScore
                fScore[neighbor] := gScore[neighbor] + h(neighbor)
                if neighbor not in openSet
                    openSet.add(neighbor)

    // Open set is empty but goal was never reached
    return failure
```
