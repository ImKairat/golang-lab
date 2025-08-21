package mathlib

import (
	"errors"
	"math/rand"
)

func GenRandNumber(end int) (int, error) {
	// Generate random number
	if end <= 0 {
		return 0, errors.New("The \"end\" argument must be equal or greeter then 0")
	}
	rand_num := rand.Intn(end)
	return rand_num, nil
}

func GenRandNumberRange(min, max int) (int, error) {
	// Generate random number for range
	if min > max {
		return 0, errors.New("max must be greeter then min")
	}
	if min < 0 {
		return 0, errors.New("min must be equal or greeter then 0")
	}

	result := rand.Intn(max - min) + min
	// Пример: max = 100; min = 20;
	// max - min = 80
	// После генерации числа от 0 до 80, например 79, потом прибавляем к нему min (79 + 20 = 99)
	// Как бы сдвигаем значение от нулю до начадбного значения (min), так всегда получаем рандомное значение в интервале [min, max)

	return result, nil
}
