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

func main() {
	var strengtName string
	fmt.Println("Введите название силового упражнения: ")
	fmt.Scan(&strengtName)
	var strengtWeight float64
	var counter, strengtReps int
	fmt.Println("Введите количество подходов: ")
	fmt.Scan(&counter)
	sets := []Set{}
	for i := 0; i < counter; i++ {
		fmt.Printf("Введите вес и количество повторений для %d подхода:\n", i+1)
		fmt.Scan(&strengtWeight, &strengtReps)
		currentSet := Set{Weight: strengtWeight, Reps: strengtReps}
		sets = append(sets, currentSet)
	}
	strengt := NewStrengthExercise(strengtName, sets)

	fmt.Println("Введите название кардио упражнения: ")
	var cardioName, cardioType string
	fmt.Scan(&cardioName)
	fmt.Println("Введите тип кардио упражнения (бег/ходьба/велосипед): ")
	fmt.Scan(&cardioType)
	fmt.Println("Введите дистанцию, время выполнения кардио и вес тела: ")
	var cardioDistance, cardioDuration, bodyWeight float64
	fmt.Scan(&cardioDistance, &cardioDuration, &bodyWeight)

	cardio := NewCardioExercise(cardioName, cardioType, cardioDistance, cardioDuration, bodyWeight)

	var bodyweigtName string
	fmt.Println("Введите название упражнения с собственным весом: ")
	fmt.Scan(&bodyweigtName)
	var bodyweigtReps int
	fmt.Println("Введите количество повторений: ")
	fmt.Scan(&bodyweigtReps)

	bw := NewBodyweightExercise(bodyweigtName, bodyweigtReps, bodyWeight)

	exercises := []Exercise{strengt, cardio, bw}

	for _, e := range exercises {
		fmt.Println(e.Summary())
	}
}
