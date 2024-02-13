func decodeMessage(key string, message string) string {
    decoded := []byte{}
    dict := map[byte]byte{
		' ': ' ',
	}
    index := 0
    for i := 0 ; i < len(key) ; i++ {
        if dict[key[i]] == 0 {
            dict[key[i]] = byte(index + 97)
            index++
        }
    }
    for i := 0 ; i < len(message) ; i++ {
        decoded = append(decoded, dict[message[i]])
    }
    return string(decoded)
}