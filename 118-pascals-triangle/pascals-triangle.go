func generate(numRows int) [][]int {
    pascal := make([][]int, numRows)
    for i:=1;i<=numRows;i++ {
        line := make([]int, i)
        for j:=0;j<i;j++ {
            if j==0 || j==i-1 {
                line[j] = 1
            } else {
                line[j] = pascal[i-2][j-1] + pascal[i-2][j]
            }

        }
        pascal[i-1] = line
    }
    return pascal
}