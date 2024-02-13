func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int, numCourses)

	//construct graph
	for i := 0; i < len(prerequisites); i++ {
		graph[prerequisites[i][1]] = append(graph[prerequisites[i][1]], prerequisites[i][0])
	}
	visited := make([]bool, len(graph))
	curNodeOnPath := make([]bool, len(graph))
	
	//check all the path starting from different nodes
	for i := 0; i < len(graph); i++ { 
        if !dfs(graph, i, visited, curNodeOnPath) {
            return false
        }
	}
	return true
}
func dfs(graph [][]int, index int, visited, curNodeOnPath []bool) bool{
    if visited[index] {
        return true
    }
	 if curNodeOnPath[index] {
        curNodeOnPath[index] = false
        return false
    }
    curNodeOnPath[index] = true
	for i := 0; i < len(graph[index]); i++ {
        curNeighbor := graph[index][i]
        if !dfs(graph, curNeighbor, visited, curNodeOnPath) {
            return false
        }
  
	}
    curNodeOnPath[index] = false
    visited[index] = true
	return true
} 
