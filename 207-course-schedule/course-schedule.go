func canFinish(numCourses int, prerequisites [][]int) bool {
    // Create a graph from the prerequisites
    graph := make([][]int, numCourses)
    for _, pre := range prerequisites {
        graph[pre[1]] = append(graph[pre[1]], pre[0])
    }

    // visited array to keep track of visited nodes during DFS
    // state: 0 = unvisited, 1 = visiting, 2 = visited
    visited := make([]int, numCourses)
    
    var dfs func(int) bool
    dfs = func(course int) bool {
        if visited[course] == 1 { // cycle detected
            return false
        }
        if visited[course] == 2 { // already visited
            return true
        }

        // Mark the current node as visiting
        visited[course] = 1

        // Visit all the adjacent nodes
        for _, nextCourse := range graph[course] {
            if !dfs(nextCourse) {
                return false
            }
        }

        // Mark the current node as visited
        visited[course] = 2
        return true
    }

    // Check if the graph has a cycle
    for i := 0; i < numCourses; i++ {
        if !dfs(i) {
            return false
        }
    }
    return true
}