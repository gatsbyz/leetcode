
import "fmt"
func findSpecialInteger(arr []int) int {
    limit := len(arr) / 4
    var count int
    current := arr[0]
    for _, element := range arr {
        fmt.Println(element, count)
        if current == element {
            count++
            if count > limit {
                return element
            }
        } else {
            count=1
            current = element
        }
    }
    return -1
}
