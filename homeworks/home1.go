package homeworks

func DecodeMessage(key string, message string) string {
	keySorted := []byte{}
	myMap := map[byte]int{}
	for i, _ := range key {
		if key[i] != ' ' && myMap[key[i]] == 0 {
			keySorted = append(keySorted, key[i])
		}
		myMap[key[i]]++
	}

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	newMap := map[byte]byte{}

	j := 0
	for i := 0; i < len(keySorted); i++ {
		if keySorted[i] != ' ' {
			j++
		}
	}
	return ""
}
