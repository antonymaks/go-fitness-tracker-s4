package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	dataSlice := strings.Split(data, ",")
	if len(dataSlice) != 3 {
		return 0, ``, 0, fmt.Errorf("incorrect input data format")
	}

	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return 0, ``, 0, err
	}
	if steps <= 0 {
		return 0, ``, 0, fmt.Errorf("amount of steps must be greater, than zero")
	}

	activityType := dataSlice[1]

	duration, err := time.ParseDuration(dataSlice[2])
	if err != nil {
		return 0, ``, 0, err
	}

	return steps, activityType, duration, nil
}

func distance(steps int, height float64) float64 {
	lenghtOfStep := height * stepLengthCoefficient
	distanceMeter := lenghtOfStep * float64(steps)
	distanceKm := distanceMeter / mInKm

	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	meanSpeed := distance(steps, height) / duration.Hours()

	return meanSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	return ``, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("incorrect input data format")
	}

	meanSpeed := meanSpeed(steps, height, duration)
	calories := (weight * meanSpeed * duration.Minutes()) / minInH

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("incorrect input data format")
	}

	meanSpeed := meanSpeed(steps, height, duration)
	calories := ((weight * meanSpeed * duration.Minutes()) / minInH) * walkingCaloriesCoefficient

	return calories, nil
}
