package configuration

import (
	"os"
	"strconv"
)

func GetEnvOrDefaultString(name, defaultValue string) string {
	value := os.Getenv("APP_" + name)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetEnvOrDefaultInt(name string, defaultValue int) int {
	valueStr := os.Getenv("APP_" + name)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func GetEnvOrDefaultBool(name string, defaultValue bool) bool {
	valueStr := os.Getenv("APP_" + name)
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultValue
}

// func getEnvOrDefaultFloat(name string, defaultValue float64) float64 {
// 	valueStr := os.Getenv("APP_" + name)
// 	if value, err := strconv.ParseFloat(valueStr, 64); err == nil {
// 		return value
// 	}
// 	return defaultValue
// }
