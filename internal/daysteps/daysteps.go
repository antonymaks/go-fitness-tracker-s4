// Пакет подсчета ежедневной активности:
// кол-во шагов, дистанции в км. и потраченные калории.
package daysteps

import (
	"errors"
	"fmt"
	"log"
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

var err error

// Парсит строку активности,
// возвращает кол-во шагов, время прогулки и возможную ошибку.
func parsePackage(data string) (int, time.Duration, error) {
	dataSlice := strings.Split(data, ",")
	if len(dataSlice) == 2 {
		stepsCount, err := strconv.Atoi(dataSlice[0])
		if err != nil {
			err = errors.New("fail to conver string to integer")
			log.Println(err)
			return 0, 0, err

		}
		if stepsCount <= 0 {
			err = errors.New("amount of steps must be greater, than zero")
			log.Println(err)
			return 0, 0, err
		}

		walkDuration, err := time.ParseDuration(dataSlice[1])
		if err != nil {
			err = errors.New("fail to parse time")
			log.Println(err)
			return 0, 0, err
		}
		if walkDuration <= 0 {
			err = errors.New("duration of training must be greater, than zero")
			log.Println(err)
			return 0, 0, err
		}

		return stepsCount, walkDuration, nil
	}

	err = errors.New("incorrect data format")
	log.Println(err)
	return 0, 0, err
}

// Парсит строку при помощи parsePackage,
// возращает строку со всеми данными об активности:
// шагии, дистанция и калории.
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		fmt.Println("error in DayActionInfo: ", err)
		return ``
	}

	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println(err)
		fmt.Println("error occured: ", err)
	}

	outputString := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distanceKm, calories)
	return outputString

}
