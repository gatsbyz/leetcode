func findThePrefixCommonArray(A []int, B []int) []int {
        resA := make(map[int]bool, 0)
        resB := make(map[int]bool, 0)
        c := 0
        res := make([]int, 0)
        for i, n := range A{
           resA[n] = true
           resB[B[i]] = true
           if resB[n]{
               c += 1
           }
           if resA[B[i]]{
               c += 1
           }

           if n == B[i]{
               c -= 1
            } 
           res = append(res, c)      
        }
        return res
}