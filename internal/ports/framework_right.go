package ports

type DBPort interface {
	CloseConnection()
	SaveAnswerToHistory(answer int, operation string) error
}
