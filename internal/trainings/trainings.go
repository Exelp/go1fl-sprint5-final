package trainings

import (
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("invalid parameters %w", err)
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("invalid parameters %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("invalid data string: %w", err)
	}
	t.Steps = steps
	t.TrainingType = parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("invalid parameters %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("invalid data string: %w", err)

	}

	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var calories float64
	var err error
	switch strings.ToLower(t.TrainingType) {
	case "бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")

	}
	total := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories)
	return total, nil
}
