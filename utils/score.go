package utils

func CalculateScore(correctAnswer int, totalQuestion int) int {
	return (correctAnswer * 100) / totalQuestion
}
