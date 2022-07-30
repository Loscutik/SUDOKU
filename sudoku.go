package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// присваиваем в board наш ввод(из os.Args), разобранный на цифры. Board - это массив 9х9. Мы берем os.Args начиная
	// с первого и до концца [1:], потому как аргументы начинаются с 1го.
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	for _, val := range os.Args[1:] {
		if len(val) != 9 {
			fmt.Println("Error")
			return
		}
	}
	board := parseInput(os.Args[1:])

	// Если backtrack возвращает true, тогда мы выводим, что судоку решено успешно и выводим само судоку, вызывая функцию printobard
	if backtrack(&board) {
		printBoard(&board)
	} else {
		fmt.Printf("Error")
	}
}

// Основаная ф-ция, котоаря решает судоку, используя рекурсию.
func backtrack(board *[9][9]int) bool {
	// будем искать пустую ячейку, вставлять в нее кандидата и проверять не нарушает ли она правила судоку :
	// i идет по строкам
	for i := 0; i < 9; i++ {
		// j идет по столбцам
		for j := 0; j < 9; j++ {
			// Если ячейка не заполнена, будем пытаться найти для нее подходящую цифру
			if board[i][j] == 0 {
				for candidate := 1; candidate <= 9; candidate++ {
					board[i][j] = candidate
					// Проверяем подошел кандитат(цифра) или нет
					if isBoardValid(board) {
						// Если кандитат подошел
						flagBackTrack := backtrack(board)
						// тогда мы вызываем нашу функцию рекурсивно, которая ищет решение уже для новой доски и возвращает true
						// то она нашла правильную цифру
						if flagBackTrack {
							return true
							// если не удалось найти подходящую цифру, то затираем ячейку
						} else { // *
							board[i][j] = 0
						}
						// если текущий кандидат не подошел, стираем значение ячейки
					} else {
						board[i][j] = 0
					}

				}
				// если ни один кандидат не подошел , возвращаем false и возвращаемся на место звездочки *
				return false
			}
		}
	}
	// прошли всю доску и не нашли ни одного пустого места, возвращаем true :
	return true
}

func isBoardValid(board *[9][9]int) bool {
	// check duplicates by row
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			if board[row][col] != 0 {
				counter[board[row][col]]++
				if counter[board[row][col]] > 1 {
					return false
				}
			}
		}

	}

	// check duplicates by column
	for col := 0; col < 9; col++ {
		counter := [10]int{}
		for row := 0; row < 9; row++ {
			if board[row][col] != 0 {
				counter[board[row][col]]++
				if counter[board[row][col]] > 1 {
					return false
				}
			}
		}
	}

	// check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					if board[row][col] != 0 {
						counter[board[row][col]]++
						if counter[board[row][col]] > 1 {
							return false
						}
					}
				}
			}
		}
	}

	return true
}

//
func printBoard(board *[9][9]int) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			fmt.Printf("%d ", board[row][col])
		}

		fmt.Println()

	}
}

func parseInput(input []string) [9][9]int {
	board := [9][9]int{}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			board[row][col], _ = strconv.Atoi(string(input[row][col]))
		}
	}
	return board
}
