package homeworks

import (
	"fmt"
	"strings"
)

func filterI(stringForSearch string) []string {
	names := []string{"asadulloh", "asadbek", "abduqodir", "asliddin"}

	filterdNames := make([]string, len(names))

	for i := 0; i < len(names); i++ {
		if strings.Contains(names[i], stringForSearch) {
			filterdNames = append(filterdNames, names[i])
		}
	}

	return filterdNames

}
func main() {
	name := "as"
	fmt.Println(filterI(name))
}
