
import "fmt"
func addBinary(a string, b string) string {
    index := 0
    carry := 0
    result := ""
    for index < len(a) || index < len(b) {

        var sum int
        if index >= len(a) {
            numberB := int(b[len(b)-index-1]-'0')
            sum = numberB + carry
        } else if index >= len(b) {
            numberA := int(a[len(a)-index-1]-'0')
            sum = numberA + carry
        } else {
            numberA := int(a[len(a)-index-1]-'0')
            numberB := int(b[len(b)-index-1]-'0')
            sum = numberA + numberB + carry
        }
        value := sum % 2

        result = strconv.Itoa(value) + result

        carry = sum / 2

        index++
    }

    if carry > 0 {
        result = "1" + result
    }

    return result
}