package calculator

import "testing"

func TestTopN(t *testing.T) {
	t.Parallel()

	got, err := Calculate([]string{"10.10", "10.20", "10.30"}, MethodTopN, 2, 0)
	if err != nil {
		t.Fatalf("Calculate() error = %v", err)
	}
	if got != "10.20" {
		t.Fatalf("Calculate() = %q, want %q", got, "10.20")
	}
}

func TestAvgNM(t *testing.T) {
	t.Parallel()

	got, err := Calculate([]string{"10.10", "10.20", "10.30"}, MethodAvgNM, 1, 2)
	if err != nil {
		t.Fatalf("Calculate() error = %v", err)
	}
	if got != "10.15" {
		t.Fatalf("Calculate() = %q, want %q", got, "10.15")
	}
}

func TestCalculateValidation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		values []string
		method Method
		n      uint32
		m      uint32
	}{
		{
			name:   "topn requires n",
			values: []string{"10.10"},
			method: MethodTopN,
			n:      0,
		},
		{
			name:   "avg requires range",
			values: []string{"10.10"},
			method: MethodAvgNM,
			n:      2,
			m:      1,
		},
		{
			name:   "depth too shallow",
			values: []string{"10.10"},
			method: MethodTopN,
			n:      2,
		},
		{
			name:   "invalid price",
			values: []string{"nope"},
			method: MethodTopN,
			n:      1,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if _, err := Calculate(tc.values, tc.method, tc.n, tc.m); err == nil {
				t.Fatal("Calculate() error = nil, want non-nil")
			}
		})
	}
}
