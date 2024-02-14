func canFinish(numCourses int, prerequisites [][]int) bool {
    // Build the graph using an adjacency list.
    // graph[i] contains a list of courses that course i depends on.
    adjList := make([][]int, numCourses)
    for _, prereq := range prerequisites {
        adjList[prereq[0]] = append(adjList[prereq[0]], prereq[1])
    }

    // visited[i] == 0 means course i has not been visited,
    // visited[i] == 1 means course i is being visited (in the current DFS path),
    // visited[i] == 2 means course i has been fully visited (all descendants explored).
    visited := make([]int, numCourses)

    // Define the DFS function to explore the graph.
    var dfs func(int) bool
    dfs = func(course int) bool {
        if visited[course] == 1 {
            // We are visiting a course that is already being visited,
            // which indicates a cycle.
            return false
        }
        if visited[course] == 2 {
            // This course has been fully explored previously,
            // no need to explore again. No cycle detected from this course.
            return true
        }

        // Mark the course as being visited.
        visited[course] = 1

        // Visit all the prerequisites of the current course.
        for _, prereq := range adjList[course] {
            if !dfs(prereq) {
                // If a cycle is detected in any prerequisite, return false.
                return false
            }
        }

        // Mark the course as fully visited.
        visited[course] = 2
        return true
    }

    // Perform DFS from each course to detect cycles.
    for i := 0; i < numCourses; i++ {
        if !dfs(i) {
            // If a cycle is detected from any course, return false.
            return false
        }
    }

    // If no cycles are detected, it is possible to finish all courses.
    return true
}
