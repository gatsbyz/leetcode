func findOrder(numCourses int, prerequisites [][]int) []int {
    search := make([][]int, numCourses) // watch out

    for i := range prerequisites {
        search[prerequisites[i][1]] = append(search[prerequisites[i][1]], prerequisites[i][0])
    }

    stack := []int{}
    visited := make([]int, numCourses)

    var dfs func(course int) bool
    // 0 not visited 1 in progress 2 completed
    dfs = func(course int) bool {
        if visited[course] == 1 {
            return false
        } else if visited[course] == 2 {
            return true
        }

        visited[course] = 1

        for i := range search[course] {
            if !dfs(search[course][i]) {
                return false
            }
        }

        visited[course] = 2
        stack = append(stack, course)
        return true
    }

    for i := 0 ; i < numCourses ; i++ {
        if !dfs(i) {
            return []int{}
        }
    }

    for i,j := 0,len(stack)-1 ; i < j ; i, j = i+1, j-1 {
        stack[j], stack[i] = stack[i], stack[j]
    }

    return stack
}