package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	//	lenStep                    = 0.65 // средняя длина шага.	//не используется в данном модуле
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	if len(data) < 10 {
		return 0, "", 0.0, fmt.Errorf("количество символов меньше ожидаемого") //минимально возможно 10 символов 1,Бег,1h1m
	}
	parts := strings.Split(data, ",") //делим строку на шаги, активность и время
	if len(parts) != 3 {
		return 0, "", 0.0, fmt.Errorf("количество разделителей не соответствует ожидаемому")
	}
	steps, err := strconv.Atoi(parts[0]) // число количество шагов
	if err != nil {
		return 0, "", 0.0, fmt.Errorf("ошибка при преобразовании количества шагов")
	}
	if steps <= 0 { //шагов должно быть больше 0
		return 0, "", 0.0, fmt.Errorf("ошибка подсчета количества шагов")
	}
	dur, err := time.ParseDuration(parts[2]) //преобразуем время
	if err != nil {
		return 0, "", 0.0, fmt.Errorf("ошибка преобразования времени")
	}
	if dur <= 0 {
		return 0, "", 0.0, fmt.Errorf("неверно задано время")
	}
	return steps, parts[1], dur, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	return stepLengthCoefficient * height * float64(steps) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 { //проверка на больше 0
		return 0.0
	}
	return distance(steps, height) / duration.Hours() // Вычисляем среднюю скорость
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activity, dur, err := parseTraining(data)
	if err != nil {
		log.Println(err)
	}
	var cal float64
	switch activity { //в соответствии с видом тренировки, считаем калории
	case "Ходьба":
		cal, err = WalkingSpentCalories(steps, weight, height, dur)
	case "Бег":
		cal, err = RunningSpentCalories(steps, weight, height, dur)
	default:
		err = fmt.Errorf("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}
	averSpeed := meanSpeed(steps, height, dur) //получаем среднюю скорость
	dist := distance(steps, height)            //получаем дистанцию
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activity, dur.Hours(), dist, averSpeed, cal), nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0.0, fmt.Errorf("ошибка в количестве шагов")
	}
	if weight <= 0 {
		return 0.0, fmt.Errorf("ошибка в весе пользователя")
	}
	if height <= 0 {
		return 0.0, fmt.Errorf("ошибка в росте пользователя")
	}
	if duration <= 0 {
		return 0.0, fmt.Errorf("ошибка в длительности занятия")
	}
	averSpeed := meanSpeed(steps, height, duration)
	return weight * averSpeed * duration.Minutes() / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	rsc, err := RunningSpentCalories(steps, weight, height, duration)
	return walkingCaloriesCoefficient * rsc, err
}
