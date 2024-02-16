/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type key struct {
    node *TreeNode
    col int // the column of the node
}

func verticalOrder(root *TreeNode) [][]int {
    queue := []key{ key{root, 0}}
    traversal := [][]int{}
    colMap := map[int][]int{}
    minCol := 0
    // 0: 4
    //

    for len(queue) > 0 {
        n := len(queue)

        for i:=0 ; i < n ; i++ {
            current := queue[0]
            queue = queue[1:]

            if current.node == nil {
                continue
            }

            if minCol > current.col {
                minCol = current.col
            }

            colMap[current.col] = append(colMap[current.col], current.node.Val)

            left := key{ current.node.Left, current.col - 1}
            right := key{ current.node.Right, current.col + 1}

            queue = append(queue, left, right)
        }
    }

    fmt.Println(minCol, colMap, len(colMap))
    for i:=minCol ; i < minCol + len(colMap) ; i++ {
        traversal = append(traversal, colMap[i])
    }

    return traversal

}