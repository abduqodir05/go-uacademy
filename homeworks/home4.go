package homeworks

func RepeatedCharacter(text string) string {
	for i := 0; i < len(text); i++ {
		if text[i] == text[i+1] {
			return string(text[i])
		}
	}
	return ""
}
