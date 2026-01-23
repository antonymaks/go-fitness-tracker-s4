// Пакет подсчета ежедневной активности:
// кол-во шагов, дистанции в км. и потраченные калории.
package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

// Парсит строку активности,
// возвращает кол-во шагов, время прогулки и возможную ошибку.
func parsePackage(data string) (int, time.Duration, error) {
	dataSlice := strings.Split(data, ",")
	if len(dataSlice) != 2 {
		return 0, 0, fmt.Errorf("incorrect input data format")
	}

	stepsCount, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return 0, 0, err
	}
	if stepsCount <= 0 {
		return 0, 0, fmt.Errorf("amount of steps must be greater, than zero")
	}

	walkDuration, err := time.ParseDuration(dataSlice[1])
	if err != nil {
		return 0, 0, err
	}

	return stepsCount, walkDuration, nil
}

// Парсит строку при помощи parsePackage,
// возращает строку со всеми данными об активности:
// шагии, дистанция и калории.
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ``
	}

	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		outputString := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
			steps, distanceKm, calories)
		return outputString

	}
	return ``
}
