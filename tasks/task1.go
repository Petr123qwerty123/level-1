package tasks

import "fmt"

type Human struct {
	FirstName string
	LastName  string
}

func (h *Human) GetFullName() string {
	return fmt.Sprintf("%s %s", h.FirstName, h.LastName)
}

// встраиваем Human в Action
type Action struct {
	Human
}

func Task1() {
	action := Action{
		Human{
			"Ivan",
			"Ivanov",
		},
	}

	// теперь можно использовать метод Human у объекта Action
	fmt.Println(action.GetFullName())
}
