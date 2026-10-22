package main

import (
	"fmt"
)

type Exercise interface {
	Name() string
	Volume() float64
	Summary() string
}

// Подход
type Set struct {
	Weight float64
	Reps   int
}

// Силовое упражнение
type StrengthExercise struct {
	ExerciseName string
	Sets         []Set
}

// Конструктор для силового упражнения
func NewStrengthExercise(name string, set []Set) StrengthExercise {
	return StrengthExercise{
		ExerciseName: name,
		Sets:         set,
	}
}

// Название силового упражнения
func (s StrengthExercise) Name() string {
	return s.ExerciseName
}

// Объем силового упражнения
func (s StrengthExercise) Volume() float64 {
	var sum float64
	for i := 0; i < len(s.Sets); i++ {
		sum += s.Sets[i].Weight * float64(s.Sets[i].Reps)
	}
	return sum
}

// Информация о силовом упражнении
func (s StrengthExercise) Summary() string {
	counter := len(s.Sets)
	return fmt.Sprintf("%s: %d подхода | объём %.1f", s.ExerciseName, counter, s.Volume())
}

// Кардио упражнение
type CardioExercise struct {
	ExerciseName string
	Distanсe     float64
	Duration     float64
	ActivityType string
	BodyWeight   float64
}

// Конструктор кардио упражнения
func NewCardioExercise(name, activityType string, distanse, duration, bodyWeight float64) CardioExercise {
	return CardioExercise{
		ExerciseName: name,
		Distanсe:     distanse,
		Duration:     duration,
		ActivityType: activityType,
		BodyWeight:   bodyWeight,
	}
}

// Название кардио упражнения
func (c CardioExercise) Name() string {
	return c.ExerciseName
}

// объем кардио упражнения
func (c CardioExercise) Volume() float64 {
	if c.Duration == 0 {
		return 0
	}
	var k float64
	switch c.ActivityType {
	case "ходьба":
		k = 0.5
	case "велосипед":
		k = 1.0
	case "бег":
		k = 1.5
	default:
		return 0
	}
	return c.BodyWeight * c.Distanсe * k
}

// информация о кардио упражнении
func (c CardioExercise) Summary() string {
	return fmt.Sprintf("%s: Тип активности: %s | Дистанция: %.2f | Затраченное время: %.0f | Вес: %.2f | Сожжено каллорий:%.1f",
		c.ExerciseName, c.ActivityType, c.Distanсe, c.Duration, c.BodyWeight, c.Volume())
}

// Упражнения с собственным весом
type BodyweightExercise struct {
	ExerciseName string
	Reps         int
	BodyWeight   float64
}

// Конструктор упражнения с собственным весом
func NewBodyweightExercise(name string, reps int, bodyWeight float64) BodyweightExercise {
	return BodyweightExercise{
		ExerciseName: name,
		Reps:         reps,
		BodyWeight:   bodyWeight,
	}
}

// Название упражнения с собственным весом
func (b BodyweightExercise) Name() string {
	return b.ExerciseName
}

// Объем упражнения с собственным весом
func (b BodyweightExercise) Volume() float64 {
	return float64(b.Reps) * b.BodyWeight
}

// Информация об упражнении с собственным весом
func (b BodyweightExercise) Summary() string {
	return fmt.Sprintf("Название упражнения: %s | Вес: %.2f | Количество повторений: %d | Объём: %.2f", b.ExerciseName,
		b.BodyWeight, b.Reps, b.Volume())
}

// Одна тренировка
type Workout struct {
	Date      string
	Exercises []Exercise
}

// Дневник тренировок
type Tracker struct {
	Workouts []Workout
}

func (t Tracker) TotalVolume() float64 {
	var total float64
	for _, w := range t.Workouts {
		for _, e := range w.Exercises {
			total += e.Volume()
		}
	}
	return total
}

func inputStrengthExercise() StrengthExercise {
	fmt.Println("Введите название силового упражнения: ")
	var name string
	fmt.Scan(&name)
	fmt.Println("Введите количество подходов: ")
	var counter int
	fmt.Scan(&counter)
	sets := []Set{}
	var strengtWeight float64
	var strengtReps int
	for i := 0; i < counter; i++ {
		fmt.Printf("Введи вес и количество повторений для %d подхода:", i+1)
		fmt.Scan(&strengtWeight, &strengtReps)
		currentSet := Set{Weight: strengtWeight, Reps: strengtReps}
		sets = append(sets, currentSet)
	}
	return NewStrengthExercise(name, sets)
}

func inputCardioExercise() CardioExercise {
	fmt.Println("Введите название кардио упражнения: ")
	var cardioName, cardioType string
	fmt.Scan(&cardioName)
	fmt.Println("Введите тип кардио упражнения: ")
	fmt.Scan(&cardioType)
	var distanse, duration, bodyweight float64
	fmt.Println("Введите дистанцию (км), время (минуты) и собственный вес (кг): ")
	fmt.Scan(&distanse, &duration, &bodyweight)
	return NewCardioExercise(cardioName, cardioType, distanse, duration, bodyweight)
}

func inputBodyweightExercise() BodyweightExercise {
	fmt.Println("Введите название упражнения с собственным весом: ")
	var name string
	fmt.Scan(&name)
	fmt.Println("Введите количество повторов и собственный вес (кг): ")
	var reps int
	var bodyWeight float64
	fmt.Scan(&reps, &bodyWeight)
	return NewBodyweightExercise(name, reps, bodyWeight)
}

func main() {
	workout := Workout{Date: "2026-10-22"}
	var tracker Tracker
Loop:
	for {
		fmt.Println("1 - силовое упражнение")
		fmt.Println("2 - кардио упражнение")
		fmt.Println("3 - упражнение с собственным весом")
		fmt.Println("0 - выход")
		fmt.Println("Выберите одно из предложенного: ")
		var choice int8
		fmt.Scan(&choice)
		switch choice {
		case 1:
			workout.Exercises = append(workout.Exercises, inputStrengthExercise())
		case 2:
			workout.Exercises = append(workout.Exercises, inputCardioExercise())
		case 3:
			workout.Exercises = append(workout.Exercises, inputBodyweightExercise())
		case 0:
			tracker.Workouts = append(tracker.Workouts, workout)
			break Loop
		default:
			fmt.Println("Вы ввели неверное число ")
		}

	}
	fmt.Printf("Общий объем за тренировку: %.1f\n", tracker.TotalVolume())
}
