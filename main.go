package main

import (
    "fmt"
    "math"
)

// calculateBMI вычисляет индекс массы тела и классифицирует его
// weight — вес в килограммах (дробное число)
// height — рост в метрах (дробное число)
// Возвращает: BMI (округлённый до 1 знака), категорию и ошибку
//
// Правила валидации:
// - weight должен быть > 0 и < 1000
// - height должен быть > 0 и < 3
// - Формула: BMI = weight / (height * height)
// - Категории:
//   * BMI < 18.5 → "underweight"
//   * 18.5 <= BMI < 25 → "normal"
//   * 25 <= BMI < 30 → "overweight"
//   * BMI >= 30 → "obese"
// - Результат BMI округлён до 1 знака после запятой

func calculateBMI(weight, height float64) (bmi float64, category string, err error) {
 if weight <= 0 {
     return 0, "", fmt.Errorf("weight must be greater than 0 (got %.1f)", weight)
 }
 if weight >= 1000 {
     return 0, "", fmt.Errorf("weight must be less than 1000 (got %.1f)", weight)
 }
 if height <= 0 {
     return 0, "", fmt.Errorf("height must be greater than 0 (got %.2f)", height)
 }
 if height >= 3 {
     return 0, "", fmt.Errorf("height must be less than 3 (got %.2f)", height)
 }

 bmi = weight / (height * height)
 bmi = math.Round(bmi*10) / 10

 switch {
    case bmi < 18.5:
    category = "underweight"
    case bmi < 25:
    category = "normal"
    case bmi < 30:
    category = "overweight"
    default:
    category = "obese"
 }

 return bmi, category, nil
}