package calculator

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
)

// Method identifies the calculation strategy requested by the client.
type Method string

const (
	MethodTopN  Method = "topN"
	MethodAvgNM Method = "avgNM"
)

// Calculate computes a price from the provided depth values.
func Calculate(values []string, method Method, n, m uint32) (string, error) {
	if len(values) == 0 {
		return "", errors.New("empty order book")
	}

	switch method {
	case MethodTopN:
		if n == 0 {
			return "", errors.New("n must be greater than zero")
		}
		if int(n) > len(values) {
			return "", fmt.Errorf("requested depth %d exceeds available values %d", n, len(values))
		}
		if _, _, err := parseDecimal(values[n-1]); err != nil {
			return "", err
		}
		return values[n-1], nil
	case MethodAvgNM:
		if n == 0 || m == 0 {
			return "", errors.New("n and m must be greater than zero")
		}
		if n > m {
			return "", errors.New("n must be less than or equal to m")
		}
		if int(m) > len(values) {
			return "", fmt.Errorf("requested depth %d exceeds available values %d", m, len(values))
		}

		total := new(big.Rat)
		scale := 0
		for _, value := range values[n-1 : m] {
			rat, valueScale, err := parseDecimal(value)
			if err != nil {
				return "", err
			}
			total.Add(total, rat)
			if valueScale > scale {
				scale = valueScale
			}
		}

		count := new(big.Rat).SetInt64(int64(m - n + 1))
		total.Quo(total, count)
		return formatDecimal(total, scale+8), nil
	default:
		return "", fmt.Errorf("unsupported method %q", method)
	}
}

func parseDecimal(value string) (*big.Rat, int, error) {
	scale := 0
	if parts := strings.Split(value, "."); len(parts) == 2 {
		scale = len(parts[1])
	}

	rat := new(big.Rat)
	if _, ok := rat.SetString(value); !ok {
		return nil, 0, fmt.Errorf("invalid decimal value %q", value)
	}
	return rat, scale, nil
}

func formatDecimal(value *big.Rat, precision int) string {
	formatted := value.FloatString(precision)
	formatted = strings.TrimRight(formatted, "0")
	formatted = strings.TrimRight(formatted, ".")
	if formatted == "" {
		return "0"
	}
	return formatted
}
