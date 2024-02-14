func findOrder(numCourses int, prerequisites [][]int) []int {
    // Adjacency list to represent the graph
    graph := make([][]int, numCourses)
    for _, pre := range prerequisites {
        graph[pre[1]] = append(graph[pre[1]], pre[0])
    }

    // Visited array to track the state of each node (0 = unvisited, 1 = visiting, 2 = visited)
    visited := make([]int, numCourses)
    // Stack to store the topological order in reverse
    stack := []int{}

    // DFS function to visit nodes
    var dfs func(course int) bool
    dfs = func(course int) bool {
        if visited[course] == 1 { // Cycle detected
            return false
        }
        if visited[course] == 2 { // Already visited
            return true
        }

        // Mark as visiting
        visited[course] = 1
        for _, nextCourse := range graph[course] {
            if !dfs(nextCourse) {
                return false
            }
        }
        // Mark as visited
        visited[course] = 2
        // Add to stack (reverse topological order)
        stack = append(stack, course)
        return true
    }

    // Perform DFS for each course
    for i := 0; i < numCourses; i++ {
        if !dfs(i) {
            return []int{} // Cycle detected, return empty array
        }
    }

    // Reverse the stack to get the correct topological order
    for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
        stack[i], stack[j] = stack[j], stack[i]
    }
    return stack
}