package utils

func canN[T any](matrix [][]T, i int, j int) bool {
	return canDoStepsN(matrix, i, j, 1)
}

func canDoStepsN[T any](matrix [][]T, i int, j int, steps int) bool {
	return i-steps >= 0
}

func N[T any](matrix [][]T, i int, j int) T {
	return stepsN(matrix, i, j, 1)
}

func stepsN[T any](matrix [][]T, i int, j int, steps int) T {
	return matrix[i-steps][j]
}

func canS[T any](matrix [][]T, i int, j int) bool {
	return canDoStepsS(matrix, i, j, 1)
}

func canDoStepsS[T any](matrix [][]T, i int, j int, steps int) bool {
	return i+steps < len(matrix)
}

func S[T any](matrix [][]T, i int, j int) T {
	return stepsS(matrix, i, j, 1)
}

func stepsS[T any](matrix [][]T, i int, j int, steps int) T {
	return matrix[i+steps][j]
}

func canW[T any](matrix [][]T, i int, j int) bool {
	return canDoStepsW(matrix, i, j, 1)
}

func canDoStepsW[T any](matrix [][]T, i int, j int, steps int) bool {
	return j-steps >= 0
}

func W[T any](matrix [][]T, i int, j int) T {
	return stepsW(matrix, i, j, 1)
}

func stepsW[T any](matrix [][]T, i int, j int, steps int) T {
	return matrix[i][j-steps]
}

func canE[T any](matrix [][]T, i int, j int) bool {
	return canDoStepsE(matrix, i, j, 1)
}

func canDoStepsE[T any](matrix [][]T, i int, j int, steps int) bool {
	return len(matrix) > 0 && j+steps < len(matrix[0])
}

func E[T any](matrix [][]T, i int, j int) T {
	return stepsE(matrix, i, j, 1)
}

func stepsE[T any](matrix [][]T, i int, j int, steps int) T {
	return matrix[i][j+steps]
}
