// Пакет подсчета ежедневной активности:
// кол-во шагов, дистанции в км. и потраченные калории.
package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
	//spcal "github.com/antonymaks/go-fitness-tracker-s4/internal/spentcalories/spentcalories.go"
	//spcal "spentcalories/spentcalories.go"
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
	//calories := duration // временная затычка до реализации функции подсчета калорий
	//calories := spcal.WalkingStepCalories(steps, weight, height, duration)
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	outputString := fmt.Sprintf("anount of steps: %d\n distance is %0.2f km \n you burn %T cal. \n", steps, distanceKm, calories)

	return outputString
}
