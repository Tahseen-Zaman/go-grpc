package ports

type DBPort interface {
	CloseConnection()
	SaveAnswerToHistory(answer int32, operation string) error
}
