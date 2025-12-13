package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || height <= 0 || duration <= 0 || weight <= 0 {
		return 0, fmt.Errorf("invalid parameters")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationMinutes) / minInH * walkingCaloriesCoefficient
	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || height <= 0 || duration <= 0 || weight <= 0 {
		return 0, fmt.Errorf("invalid parameters")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationMinutes) / minInH
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	meanSpeed := distance / duration.Hours()
	return meanSpeed
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	lengthStep := height * stepLengthCoefficient
	distance := (lengthStep * float64(steps)) / mInKm
	return distance
}
