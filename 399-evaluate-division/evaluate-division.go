type Graph map[string]map[string]float64

// dfs is a helper function for depth-first search through the graph to find the result of division.
// It tries to find a path from 'start' to 'end' and multiplies the weights of edges along the path.
func (g Graph) dfs(start, end string, visited map[string]bool) (float64, bool) {
	if _, ok := g[start]; !ok {
		return -1, false // start node doesn't exist
	}
	if ratio, ok := g[start][end]; ok {
		return ratio, true // direct edge exists
	}
	visited[start] = true
	for node, value := range g[start] {
		if visited[node] { // Skip visited nodes to prevent cycles
			continue
		}
		if result, found := g.dfs(node, end, visited); found {
			return result * value, true // Multiply edge weights along the path
		}
	}
	return -1, false
}

// calcEquation calculates the results of division queries based on given equations and values.
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(Graph)
	for i, equation := range equations {
		a, b := equation[0], equation[1]
		if graph[a] == nil {
			graph[a] = make(map[string]float64)
		}
		if graph[b] == nil {
			graph[b] = make(map[string]float64)
		}
		graph[a][b] = values[i]        // Edge a -> b
		graph[b][a] = 1 / values[i]    // Edge b -> a (inverse relation)
	}

	results := make([]float64, len(queries))
	for i, query := range queries {
		start, end := query[0], query[1]
		if result, found := graph.dfs(start, end, make(map[string]bool)); found {
			results[i] = result
		} else {
			results[i] = -1 // Query result cannot be determined
		}
	}
	return results
}
