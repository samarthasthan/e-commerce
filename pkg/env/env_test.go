package env

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	os.Setenv("TEST_KEY_1", "test_value")
	value := GetEnv("TEST_KEY_1", "")
	if value != "test_value" {
		t.Errorf("GetEnv() = %v; want %v", value, "test_value")
	}

	key := "TEST_KEY_2"
	defaultValue := "default"
	value = GetEnv(key, defaultValue)
	if value != defaultValue {
		t.Errorf("GetEnv() = %v; want %v", value, defaultValue)
	}
}
