package utils

import (
	"log"
	"strconv"
)

func Convert_string_float(string_parameter string) float64 {
	string_f, err := strconv.ParseFloat(string_parameter, 64)
	if err != nil {
		log.Fatalf("Unable to convert %s to float value", string_parameter)
	}
	return string_f
}
