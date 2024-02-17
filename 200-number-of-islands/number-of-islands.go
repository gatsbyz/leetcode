func numIslands(grid [][]byte) int {
    count := 0
    
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == '1' {
                count++
                
                // dfs(grid, i, j)
				bfs(grid, i, j)
            }
        }
    }
    
    return count
}

func dfs(grid [][]byte, r, c int) {
    if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) || grid[r][c] != '1' {
        return
    }
    
    grid[r][c] = '0'
    dfs(grid, r+1, c)
    dfs(grid, r-1, c)
    dfs(grid, r, c+1)
    dfs(grid, r, c-1)
}

func bfs(grid [][]byte, r, c int) {
    q := [][2]int{}
    q = append(q, [2]int{r,c})
    grid[r][c] = 0
    
    offsets := [4][2]int{{-1,0}, {1,0}, {0,1}, {0,-1}}
    for len(q) > 0 {
        curr := q[0]
        q = q[1:]
        
        for _, offset := range offsets {
            x := curr[0] + offset[0]
            y := curr[1] + offset[1]
            
            if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x]) && grid[x][y] == '1' {
                q = append(q, [2]int{x,y})
                grid[x][y] = 0
            }
        }
    }
}