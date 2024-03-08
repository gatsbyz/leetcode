
import "strconv"
func validWordAbbreviation(word string, abbr string) bool {
    wordIndex := 0
    abbrIndex := 0

    for wordIndex < len(word) && abbrIndex < len(abbr) {
        if abbr[abbrIndex] == word[wordIndex] {
            wordIndex++
            abbrIndex++
            continue
        }
        if isLetter(abbr[abbrIndex]) {
            return false
        }
        if abbr[abbrIndex] == '0' {
            return false
        }
        numStr := []byte{}
        for abbrIndex < len(abbr) {
            if !isLetter(abbr[abbrIndex]) {
                numStr = append(numStr, abbr[abbrIndex])
            } else {
                break
            }
            abbrIndex++
        }
        num, _ := strconv.Atoi(string(numStr))
        wordIndex += num
    }

    if wordIndex == len(word) && abbrIndex == len(abbr) {
        return true
    }
    return false
}

func isLetter(s byte) bool {
    if 'a' <= s && s <= 'z' {
        return true
    }
    return false
}