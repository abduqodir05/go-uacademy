package main

import (
	"fmt"
	"go-uacademy/homeworks"
)

func main() {
	// home 1
	fmt.Println(homeworks.ReverseArray([]int{1, 2, 3, 4, 5}))

	// home 2
	var chars []byte = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'J', 'I', 'H'}

	err, message, data := homeworks.Drop(chars, 1)

	if err != nil {
		fmt.Println(message)
	} else {
		fmt.Println(string(data)) 
	}

	// home 3

	var users = []homeworks.User{
		{
			Id:        1,
			FirstName: "Shoxrux",
			LastName:  "Safarov",
		},
		{
			Id:        2,
			FirstName: "Abduqodir",
			LastName:  "Musayev",
		},
	}

	var questions = []homeworks.Question{
		{
			Id:      1,
			Content: "Apple qanday meva?",
			Options: [4]homeworks.Option{
				{
					Id:        1,
					Content:   "Olma",
					IsCorrect: true,
				},
				{
					Id:        2,
					Content:   "Banan",
					IsCorrect: false,
				},
				{
					Id:        3,
					Content:   "Nok",
					IsCorrect: false,
				},
				{
					Id:        4,
					Content:   "Anor",
					IsCorrect: false,
				},
			},
		},
		{
			Id:      2,
			Content: "Golang qanday til?",
			Options: [4]homeworks.Option{
				{
					Id:        1,
					Content:   "Mol tili",
					IsCorrect: false,
				},
				{
					Id:        2,
					Content:   "Dasturlash tili",
					IsCorrect: true,
				},
				{
					Id:        3,
					Content:   "Odam tili",
					IsCorrect: false,
				},
				{
					Id:        4,
					Content:   "Til emas",
					IsCorrect: false,
				},
			},
		},
		{
			Id:      3,
			Content: "which is your laptop?",
			Options: [4]homeworks.Option{
				{
					Id:        1,
					Content:   "MacBook",
					IsCorrect: false,
				},
				{
					Id:        2,
					Content:   "Samsung",
					IsCorrect: false,
				},
				{
					Id:        3,
					Content:   "Acer",
					IsCorrect: false,
				},
				{
					Id:        4,
					Content:   "Huawei",
					IsCorrect: true,
				},
			},
		},
		{
			Id:      4,
			Content: "which is your favorite club?",
			Options: [4]homeworks.Option{
				{
					Id:        1,
					Content:   "Barcelona",
					IsCorrect: false,
				},
				{
					Id:        2,
					Content:   "Liverpool",
					IsCorrect: true,
				},
				{
					Id:        3,
					Content:   "Real Madrid",
					IsCorrect: false,
				},
				{
					Id:        4,
					Content:   "Arsenal",
					IsCorrect: false,
				},
			},
		},
		{
			Id:      5,
			Content: "which is your phone?",
			Options: [4]homeworks.Option{
				{
					Id:        1,
					Content:   "RedMi 8",
					IsCorrect: true,
				},
				{
					Id:        2,
					Content:   "Samsung S23 Ultra",
					IsCorrect: false,
				},
				{
					Id:        3,
					Content:   "Iphone 14 Pro Max",
					IsCorrect: false,
				},
				{
					Id:        4,
					Content:   "Huawei P9 Pro",
					IsCorrect: false,
				},
			},
		},
	}

	userMap := homeworks.ConvertToMapUser(users)
	questionMap := homeworks.ConvertToMapQuestion(questions)
	fmt.Println(userMap[1].FirstName + " " + userMap[1].LastName)

	answer := map[uint32]uint8{}
	ball := 0

	answer[1] = 1
	answer[2] = 2
	answer[3] = 4
	answer[4] = 3
	answer[5] = 3

	for questionId, optionId := range answer {

		fmt.Printf("%d. %s\n", questionMap[questionId].Id, questionMap[questionId].Content)

		option := questionMap[questionId].GetOption(optionId)

		fmt.Printf("\t[%c]. %s\n", 64+option.Id, option.Content)
		if option.IsCorrect {
			ball += 10
			fmt.Printf("\t[CORRECT]\n")
		} else {
			ball -= 5
			fmt.Printf("\t[WRONG]\n")
		}
		fmt.Println("Total ball:", ball)
	}
}
