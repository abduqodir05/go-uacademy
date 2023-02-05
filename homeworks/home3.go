package homeworks

type User struct {
	Id uint32
	FirstName, LastName string
}

type Option struct {
	Id uint8
	Content string
	IsCorrect bool
}

type Question struct {
	Id uint32
	Content string
	Options [4]Option
}

type QuestionMap map[uint32]Question
type UserMap map[uint32]User

func ConvertToMapUser(users []User) (userMap UserMap) {
	userMap = UserMap{}

	for _, user := range users {
		userMap[user.Id] = user
	}

	return
}

func ConvertToMapQuestion(questions []Question) (questionMap QuestionMap) {
	questionMap = QuestionMap{}

	for _, question := range questions {
		questionMap[question.Id] = question
	}

	return
}

func (q Question) GetOption(optionId uint8) Option {

	for _, option := range q.Options {
		if option.Id == optionId {
			return option
		}
	}

	return Option{}
}
