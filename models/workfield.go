package models

type WorkFieldItem int

const (
	Mobile WorkFieldItem = iota
	Web
	Backend
	DevOps
	Manager
	Owner
)

func (w WorkFieldItem) String() string {
	return [...]string{"Mobile", "Web", "Backend", "DevOps", "Manager", "Owner"}[w]
}
