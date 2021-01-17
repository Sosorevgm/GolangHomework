package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestGetPrimeCount(t *testing.T) {

	/*
			Пытался созздать анонимную структуру с вложенной анонимной структурой,
			но не смог разрбраться в итоге

		var cases = []struct {
			name  string
			input int64
			want  struct {
				mResult string
				mError  error
			}
		}{
			{
				name:  "Case positive number 5",
				input: 5,
				want: {
					mResult: "Найдено простых чисел: 2",
					mError:  nil,
				},
			},
			{
				name:  "Case positive number 10",
				input: 10,
				want: {
					mResult: "Найдено простых чисел: 4",
					mError:  nil,
				},
			},
			{
				name:  "Case negative number",
				input: -5,
				want: {
					mResult: "",
					mError:  errors.New("число должно быть положительным"),
				},
			},
			{
				name:  "Case zero",
				input: 0,
				want: {
					mResult: "",
					mError:  errors.New("число должно быть положительным"),
				},
			},
		}
	*/

	/*
		Тест не проходит в кейсах, где ошибка не равна nil. Не могу понять почему.
		Пробовал создавать свою ошибку и сравнивать, но это тоже не сработало.
	*/

	var cases = []struct {
		name       string
		input      int64
		wantResult string
		wantError  error
	}{
		{
			name:       "Case positive number 5",
			input:      5,
			wantResult: "Найдено простых чисел: 2",
			wantError:  nil,
		},
		{
			name:       "Case positive number 10",
			input:      10,
			wantResult: "Найдено простых чисел: 4",
			wantError:  nil,
		},
		{
			name:       "Case negative number",
			input:      -5,
			wantResult: "",
			wantError:  errors.New("число должно быть положительным"),
		},
		{
			name:       "Case zero",
			input:      0,
			wantResult: "",
			wantError:  errors.New("число должно быть положительным"),
		},
	}

	for _, tc := range cases {
		getResult, getError := getPrimeCount(tc.input)
		if (getResult != tc.wantResult) || (getError != tc.wantError) {
			t.Errorf("%s excpected result: %v, excpected error: %v , got result: %v, got error: %v",
				tc.name, tc.wantResult, tc.wantError, getResult, getError)
		}
	}
}

func ExampleGetPrimeCount() {
	getResult, getError := getPrimeCount(5)
	fmt.Printf("%v,%v", getResult, getError)
	// Output: Найдено простых чисел: 2,<nil>
}

func BenchmarkGetPrimeCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getPrimeCount(5)
	}
}
