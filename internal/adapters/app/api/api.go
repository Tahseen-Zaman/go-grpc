package api

import "github.com/Tahseen-Zaman/go-grpc/internal/ports"

type APIAdapter struct {
	ArithmeticPort ports.ArithmeticPort
	DBAdapter      ports.DBPort
}

func NewAPIAdapter(arithmeticPort ports.ArithmeticPort, dbAdapter ports.DBPort) *APIAdapter {
	return &APIAdapter{
		ArithmeticPort: arithmeticPort,
		DBAdapter:      dbAdapter,
	}
}

func (apia *APIAdapter) GetAddition(a, b int) (int, error) {
	result, err := apia.ArithmeticPort.Addition(a, b)
	if err != nil {
		return 0, err
	}
	if err := apia.DBAdapter.SaveAnswerToHistory(result, "addition"); err != nil {
		return 0, err
	}
	return result, nil
}

func (apia *APIAdapter) GetSubtraction(a, b int) (int, error) {
	result, err := apia.ArithmeticPort.Subtraction(a, b)
	if err != nil {
		return 0, err
	}
	if err := apia.DBAdapter.SaveAnswerToHistory(result, "subtraction"); err != nil {
		return 0, err
	}
	return result, nil
}

func (apia *APIAdapter) GetMultiplication(a, b int) (int, error) {
	result, err := apia.ArithmeticPort.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	if err := apia.DBAdapter.SaveAnswerToHistory(result, "multiplication"); err != nil {
		return 0, err
	}
	return result, nil
}

func (apia *APIAdapter) GetDivision(a, b int) (int, error) {
	result, err := apia.ArithmeticPort.Division(a, b)
	if err != nil {
		return 0, err
	}
	if err := apia.DBAdapter.SaveAnswerToHistory(result, "division"); err != nil {
		return 0, err
	}
	return result, nil
}
