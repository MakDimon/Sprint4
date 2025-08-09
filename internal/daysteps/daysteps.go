package daysteps

import (
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	if len(data) <= 6 {
		return 0, 0.0, fmt.Errorf("количество символов меньше ожидаемого") //минимально возможно 0,0h1m -6 символов
	}
	parts := strings.Split(data, ",") //делим строку на шаги и время
	if len(parts) != 2 {
		return 0, 0.0, fmt.Errorf("количество разделителей не соответствует ожидаемому")
	}
	steps, err := strconv.Atoi(parts[0]) // число количество шагов
	if err != nil {
		return 0, 0.0, fmt.Errorf("ошибка при преобразовании количества шагов")
	}
	if steps <= 0 { //проверяем чтоб кол-во шагов было положительным
		return 0, 0.0, fmt.Errorf("ошибка подсчета количества шагов")
	}
	dur, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0.0, fmt.Errorf("ошибка преобразования времени")
	}
	if dur <= 0 {
		return 0, 0.0, fmt.Errorf("значение времени неправильное")
	}
	return steps, dur, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, dur, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	dist := float64(steps) * stepLength / mInKm
	cal, err := spentcalories.WalkingSpentCalories(steps, weight, height, dur)
	if err != nil {
		log.Println(err)
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, dist, cal)
}
