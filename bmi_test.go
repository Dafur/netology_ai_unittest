package main

import (
	"math"
	"testing"
)

func TestCalculateBMI(t *testing.T) {
	tests := []struct {
		name     string
		weight   float64
		height   float64
		wantBMI  float64
		wantCat  string
		wantErr  bool
	}{
		// Категории BMI
		{"underweight (low border)", 45.0, 1.80, 13.9, "underweight", false},
		{"underweight (mid)", 50.0, 1.85, 14.6, "underweight", false},
		{"normal (low border)", 55.0, 1.70, 19.0, "normal", false},          // ~18.99 → округляется до 19.0
		{"normal (mid)", 70.0, 1.75, 22.9, "normal", false},
		{"normal (high border)", 84.0, 1.73, 28.0, "overweight", false},     // 27.96 → 28.0 (уже overweight)
		{"overweight (low border)", 85.0, 1.73, 28.5, "overweight", false},
		{"overweight (mid)", 90.0, 1.70, 31.1, "obese", false},               // 31.14 → 31.1 (уже obese)
		{"obese (low border)", 100.0, 1.70, 34.6, "obese", false},
		{"obese (high value)", 120.0, 1.60, 46.9, "obese", false},

		// Edge cases: граничные значения валидации (должны давать ошибку)
		{"weight == 0", 0.0, 1.70, 0, "", true},
		{"weight < 0", -5.0, 1.70, 0, "", true},
		{"weight == 1000", 1000.0, 1.70, 0, "", true},
		{"height == 0", 70.0, 0.0, 0, "", true},
		{"height < 0", 70.0, -0.5, 0, "", true},
		{"height == 3", 70.0, 3.0, 0, "", true},

		// Edge cases: пограничные BMI-значения (важно из-за округления)
		{"BMI exactly 18.5 (before rounding)", 55.425, 1.73, 18.5, "normal", false}, // 18.495 → 18.5
		{"BMI just below 18.5", 55.0, 1.73, 18.4, "underweight", false},             // 18.43 → 18.4
		{"BMI exactly 25 (before rounding)", 74.375, 1.73, 25.0, "normal", false},   // 24.995 → 25.0
		{"BMI just below 25", 74.0, 1.73, 24.9, "normal", false},                    // 24.87 → 24.9
		{"BMI exactly 30 (before rounding)", 89.325, 1.73, 30.0, "obese", false},   // 29.995 → 30.0
		{"BMI just below 30", 89.0, 1.73, 29.9, "overweight", false},                // 29.88 → 29.9

		// Крайние допустимые значения
		{"min weight, min height", 0.1, 0.01, 1000.0, "obese", false},
		{"max weight allowed (999.9)", 999.9, 1.0, 999.9, "obese", false},
		{"max height allowed (2.99)", 80.0, 2.99, 8.9, "underweight", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange: входные данные уже заданы в tt

			// Act
			bmi, category, err := calculateBMI(tt.weight, tt.height)

			// Assert: проверка ошибки
			if (err != nil) != tt.wantErr {
				if tt.wantErr {
					t.Errorf("expected error for case %q, got nil", tt.name)
				} else {
					t.Errorf("unexpected error for case %q: %v", tt.name, err)
				}
				return
			}

			// Если ошибка не ожидается — проверяем результаты
			if !tt.wantErr {
				if math.Abs(bmi-tt.wantBMI) > 1e-9 {
					t.Errorf("BMI mismatch for case %q: want %.1f, got %.1f", tt.name, tt.wantBMI, bmi)
				}

				if category != tt.wantCat {
					t.Errorf("category mismatch for case %q: want %q, got %q", tt.name, tt.wantCat, category)
				}
			}
		})
	}
}
