package tasks

// import (
// 	"fmt"
// )

// type User struct {
// 	id                  int
// 	firstName, LastName string
// }
// type Option struct {
// 	id        int
// 	content   string
// 	isCorrect bool
// }
// type question struct {
// 	id      int
// 	content string
// 	options [4]Option
// }

// type QuestionMap map[int]question
// type UserMap map[int]User

// func ConvertToMapUser(users []User) (userMap UserMap) {
// 	userMap = make(UserMap)
// 	for _, user := range users {
// 		userMap[user.id] = user
// 	}
// 	return
// }

// func ConvertToMapQuestion(questions []question) (questionMap QuestionMap) {
// 	questionMap = make(QuestionMap)
// 	for _, q := range questions {
// 		questionMap[q.id] = q
// 	}
// 	return
// }
// func (q question) getOption(optionId int) Option {
// 	for _, opt := range q.options {
// 		if opt.id == optionId {
// 			return optionId
// 		}
// 	}
// 	return Option{}
// }
// func main() {
// 	var users = []User{
// 		{
// 			id:        1,
// 			firstName: "abduqodir",
// 			LastName:  "musayev",
// 		},
// 		{
// 			id:        2,
// 			firstName: "samandar",
// 			LastName:  "merqoziyev",
// 		},
// 	}
// 	var questions = []question{
// 		{
// 			id:      1,
// 			content: "which is your favourite fruit",
// 			options: [4]Option{
// 				{
// 					id:        1,
// 					content:   "olma",
// 					isCorrect: false,
// 				},
// 				{
// 					id:        2,
// 					content:   "uzum",
// 					isCorrect: false,
// 				},
// 				{
// 					id:        3,
// 					content:   "banan",
// 					isCorrect: true,
// 				},
// 				{
// 					id:        4,
// 					content:   "gilos",
// 					isCorrect: false,
// 				},
// 			},
// 		},
// 		{
// 			id:      2,
// 			content: "what is your brother's name",
// 			options: [4]Option{
// 				{
// 					id:        1,
// 					content:   "john",
// 					isCorrect: false,
// 				},
// 				{
// 					id:        2,
// 					content:   "sylvia",
// 					isCorrect: false,
// 				},
// 				{
// 					id:        3,
// 					content:   "tursunhon",
// 					isCorrect: false,
// 				},
// 				{
// 					id:        4,
// 					content:   "abdulloh",
// 					isCorrect: true,
// 				},
// 			},
// 		},
// 	}

// 	// for _, q := range questions {
// 	// 	fmt.Printf("%d. %s\n", q.id, q.content)
// 	// 	for _, opt := range q.options {
// 	// 		fmt.Printf("\t[%c]. %s\n", 64 + opt.id, opt.content)
// 	// 	}
// 	// }
// 	userMap := ConvertToMapUser(users)
// 	questionMap := ConvertToMapQuestion(questions)
// 	fmt.Println(userMap[1].firstName + " " + userMap[1].LastName)

// 	answer := map[int]int{}
// 	answer[1] = 3
// 	answer[2] = 4

// 	for questionId, optionId := range answer {
// 		fmt.Printf("%d. %s\n", questionMap[questionId].id, questionMap[questionId].content)

// 		option := questionMap[questionId].getOption(optionId)
// 		fmt.Printf("\t[%c]. %s\n", 64+option.id, option.content)
// 		if option.isCorrect {
// 			fmt.Printf("\t[CORRECT]\n")
// 		}else {
// 			fmt.Printf("\t[WRONG]\n")
// 		}
// 	}
// }
