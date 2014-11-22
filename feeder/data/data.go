package data

type Entry struct {
	ID         int
	UserID     int
	Server     string
	Method     string
	Controller string
	Action     string
	SubjectID  int
	Event      string
}
