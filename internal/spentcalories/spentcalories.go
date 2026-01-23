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
	steps, activityType, duration, err := parseTraining(data)
	if err != nil {
		fmt.Println("training parse error", err)
		return "", err
	}

	if steps <= 0 {
		err := fmt.Errorf("amount of steps must be positive: %d", steps)
		fmt.Println(err)
		return "", err
	}

	activityType = strings.ToLower(strings.TrimSpace(activityType))

	var distanceKm, speedKmh, calories float64

	distanceMeters := float64(steps) * lenStep
	distanceKm = distanceMeters / mInKm

	durationHours := duration.Hours()

	activityName := ""
	switch activityType {
	case "ходьба", "walking", "ходьба,":
		// Устанавливаем русское название для вывода
		activityName = "Ходьба"

		if durationHours > 0 {
			speedKmh = distanceKm / durationHours
		}

		calories, err = WalkingSpentCalories(steps, weight, height, duration)

	case "бег", "running", "run", "бег,":
		activityName = "Бег"

		if durationHours > 0 {
			speedKmh = distanceKm / durationHours
		}

		calories, err = RunningSpentCalories(steps, weight, height, duration)

	default:
		err := fmt.Errorf("unnown training type: %s", activityType)
		fmt.Println(err)
		return "", err
	}

	// Формируем строку результата
	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		activityName,
		durationHours,
		distanceKm,
		speedKmh,
		calories,
	)

	return result, nil
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
