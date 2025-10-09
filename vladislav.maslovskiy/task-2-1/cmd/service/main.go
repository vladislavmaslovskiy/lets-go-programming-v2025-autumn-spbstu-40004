package main

import (
	"errors"
	"fmt"
)
var errOperator = errors.New("invalid operation")

func adjustTempurature(lowTempurature int, highTempurature int, askingTempurature int, operation string) (int, int, error) {
	if lowTempurature == -1 && highTempurature == -1 {
		return lowTempurature, highTempurature, nil
	}

	switch operation {
	case ">=":
		if askingTempurature > highTempurature {
			lowTempurature = -1
			highTempurature = -1
		} else if lowTempurature <= askingTempurature && askingTempurature <= highTempurature {
			lowTempurature = askingTempurature
		}
	case "<=":
		if askingTempurature < lowTempurature {
			lowTempurature = -1
			highTempurature = -1
		} else if lowTempurature <= askingTempurature && askingTempurature <= highTempurature {
			highTempurature = askingTempurature
		}
	default:
		return lowTempurature, highTempurature, errOperator
	}

	return lowTempurature, highTempurature, nil
}

func main()
{}
