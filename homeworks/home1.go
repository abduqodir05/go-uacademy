package homeworks

import "fmt"

func GetNextDayDate(day, month int) string {
	validate := map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  31,
		10: 30,
		11: 30,
		12: 31,
	}
	if month > 12 {
		return "month can't be more than 12"
	}

	if day > validate[month] {
		return fmt.Sprintf("amount of days for %d month can't be more than %d", month, day)
	}

	if day == validate[month] && month == 12 {
		return fmt.Sprintf("%d, %d", 01, 01)
	}
	
	if day == validate[month] {
		return fmt.Sprintf("%d, %d", 01, month+1)
	}	
	return fmt.Sprintf("%d, %d", day+1, month)
}
