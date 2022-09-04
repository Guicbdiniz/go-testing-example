package calculator

import "testing"

func TestIntCalculator(t *testing.T) {
	calculator := CreateIntCalculator(0, 0)
	expected := 0
	got := 0
	a := 0
	b := 0

	t.Run("Sum 5 5", func(t *testing.T) {
		t.Parallel()

		a = 5
		b = 5
		expected = 10
		got = calculator.SetValues(a, b).Add().Result()

		if expected != got {
			t.Fatalf("Sum %d and %d did not work as expected. Expected: %d, got: %d", a, b, expected, got)
		}
	})

	t.Run("Sum -20 10", func(t *testing.T) {
		t.Parallel()

		a = -20
		b = 10
		expected = -10
		got = calculator.SetValues(a, b).Add().Result()

		if expected != got {
			t.Fatalf("Sum %d and %d did not work as expected. Expected: %d, got: %d", a, b, expected, got)
		}
	})

	t.Run("Subtract 798 564", func(t *testing.T) {
		t.Parallel()

		a = 798
		b = 564
		expected = 234
		got = calculator.SetValues(a, b).Subtract().Result()

		if expected != got {
			t.Fatalf("Subtract %d and %d did not work as expected. Expected: %d, got: %d", a, b, expected, got)
		}
	})

	t.Run("Subtract -56 10", func(t *testing.T) {
		t.Parallel()

		a = -56
		b = 10
		expected = -66
		got = calculator.SetValues(a, b).Subtract().Result()

		if expected != got {
			t.Fatalf("Subtract %d and %d did not work as expected. Expected: %d, got: %d", a, b, expected, got)
		}
	})

	t.Run("Multiply 0 10", func(t *testing.T) {
		t.Parallel()

		a = 0
		b = 10
		expected = 0
		got = calculator.SetValues(a, b).Multiply().Result()

		if expected != got {
			t.Fatalf("Multiply %d and %d did not work as expected. Expected: %d, got: %d", a, b, expected, got)
		}
	})

	t.Run("Multiply -79 15", func(t *testing.T) {
		t.Parallel()

		a = -79
		b = 15
		expected = -1185
		got = calculator.SetValues(a, b).Multiply().Result()

		if expected != got {
			t.Fatalf("Multiply %d and %d did not work as expected. Expected: %d, got: %d", a, b, expected, got)
		}
	})

	t.Run("Divide 45 13", func(t *testing.T) {
		t.Parallel()

		a = 45
		b = 13
		expected = 3
		got = calculator.SetValues(a, b).Divide().Result()

		if expected != got {
			t.Fatalf("Divide %d and %d did not work as expected. Expected: %d, got: %d", a, b, expected, got)
		}
	})

	t.Run("Divide 14 -2", func(t *testing.T) {
		t.Parallel()

		a = 14
		b = -2
		expected = -7
		got = calculator.SetValues(a, b).Divide().Result()

		if expected != got {
			t.Fatalf("Divide %d and %d did not work as expected. Expected: %d, got: %d", a, b, expected, got)
		}
	})
}
