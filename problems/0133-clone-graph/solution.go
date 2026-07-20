package clonegraph

/*
Constraints:
	- Undirected: Yes → edges are bidirectional → ALWAYS has cycles → must track visited
	- Has cycle: Yes
	- Node self-loop: No
	- Number of nodes: 0 - 100 → CAN be empty (node == nil) → guard and return nil
	- Node.Val: unique, in range [1, n] → can key the map by Val or by pointer
	- No repeated edges (no multi-edges) → each neighbor appears once, link directly
	- Disconnected: No → one node is enough to reach everything, no need to scan all nodes

Expectation:
	- Deep copy: NEW nodes, neighbors are also NEW, sharing no node with the original.
	- New node: same Val, different address. Same "shape" (connection structure).
*/

/*
Example: <Skip>
*/

/*
Brute: DFS clone each node on visit. Undirected → going A→B then B→A clones A
       AGAIN (nothing remembers A was cloned) → infinite loop.
       → This is exactly why we need a map to remember cloned nodes.
Time: INF, Space: INF
*/

/*
Pattern: DFS/BFS (input is a graph). Map[*Node]*Node = original → clone:
		acts as "visited" (key exists = already cloned) AND lets us look up
		the clone to link neighbors. (Islands only needed yes/no visited.)
*/

/*
Target:
	Time:  O(V + E) → each node visited once (V); each edge traversed when linking (E)
	Space: O(V) → map (V pairs) + recursion/queue depth (≤ V). Output graph NOT counted.
*/

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visitedNodes := map[int]*Node{node.Val: {Val: node.Val}}
	stack := []*Node{node}

	for len(stack) > 0 {
		rawNode := stack[0]
		stack = stack[1:]

		for _, rawNeighborNode := range rawNode.Neighbors {
			newNode, _ := visitedNodes[rawNode.Val]

			if newNeighborNode, ok := visitedNodes[rawNeighborNode.Val]; ok {
				newNode.Neighbors = append(newNode.Neighbors, newNeighborNode)
			} else {
				newNeighborNode := &Node{Val: rawNeighborNode.Val, Neighbors: make([]*Node, 0)}
				newNode.Neighbors = append(newNode.Neighbors, newNeighborNode)
				visitedNodes[newNeighborNode.Val] = newNeighborNode
				stack = append(stack, rawNeighborNode)
			}
		}
	}

	return visitedNodes[node.Val]
}
