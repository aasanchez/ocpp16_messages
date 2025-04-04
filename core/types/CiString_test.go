package types

import (
	"math/rand"
	"strings"
	"testing"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// generateStr creates a random alphanumeric string of the given length.
func generateStr(length int) string {
	var sb strings.Builder
	sb.Grow(length)

	for range make([]struct{}, length) {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func TestCiString20_IsValid(t *testing.T) {
	valid := CiString20(generateStr(20)) // Generates a 20-character string
	if !valid.IsValid() {
		t.Errorf("Expected string %q to be valid", valid)
	}

	invalid := CiString20(generateStr(21)) // Generates a 21-character string
	if invalid.IsValid() {
		t.Errorf("Expected string %q to be invalid", invalid)
	}
}

func TestCiString20_MarshalText(t *testing.T) {
	string20 := generateStr(20)
	s := CiString20(string20) // Generates a 20-character string
	data, err := s.MarshalText()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if string(data) != string20 {
		t.Errorf("Expected '%s', got: %s", string20, data)
	}

	s = CiString20(generateStr(21)) // Generates a 21-character string
	_, err = s.MarshalText()
	if err == nil {
		t.Error("Expected error for string exceeding 20 characters")
	}
}

func TestCiString20_UnmarshalText(t *testing.T) {
	string20 := generateStr(20)
	var s CiString20
	err := s.UnmarshalText([]byte(string20)) // Generates a 20-character string
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if s.String() != string20 { // Compare using the String() method
		t.Errorf("Expected '%s', got: %q", string20, s.String())
	}

	err = s.UnmarshalText([]byte(generateStr(21))) // Generates a 21-character string
	if err == nil {
		t.Error("Expected error for string exceeding 20 characters")
	}
}

func TestCiString25_IsValid(t *testing.T) {
	valid := CiString25(generateStr(25)) // Generates a 25-character string
	if !valid.IsValid() {
		t.Errorf("Expected string %q to be valid", valid)
	}

	invalid := CiString25(generateStr(26)) // Generates a 26-character string
	if invalid.IsValid() {
		t.Errorf("Expected string %q to be invalid", invalid)
	}
}

func TestCiString25_MarshalText(t *testing.T) {
	string25 := generateStr(25)
	s := CiString25(string25) // Generates a 25-character string
	data, err := s.MarshalText()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if string(data) != string25 {
		t.Errorf("Expected '%s', got: %s", string25, data)
	}

	s = CiString25(generateStr(26)) // Generates a 26-character string
	_, err = s.MarshalText()
	if err == nil {
		t.Error("Expected error for string exceeding 25 characters")
	}
}

func TestCiString25_UnmarshalText(t *testing.T) {
	string25 := generateStr(25)
	var s CiString25
	err := s.UnmarshalText([]byte(string25)) // Generates a 25-character string
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if s.String() != string25 { // Compare using the String() method
		t.Errorf("Expected '%s', got: %q", string25, s.String())
	}

	err = s.UnmarshalText([]byte(generateStr(26))) // Generates a 26-character string
	if err == nil {
		t.Error("Expected error for string exceeding 25 characters")
	}
}

func TestCiString50_IsValid(t *testing.T) {
	valid := CiString50(generateStr(50)) // Generates a 50-character string
	if !valid.IsValid() {
		t.Errorf("Expected string %q to be valid", valid)
	}

	invalid := CiString50(generateStr(51)) // Generates a 51-character string
	if invalid.IsValid() {
		t.Errorf("Expected string %q to be invalid", invalid)
	}
}

func TestCiString50_MarshalText(t *testing.T) {
	string50 := generateStr(50)
	s := CiString50(string50) // Generates a 50-character string
	data, err := s.MarshalText()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if string(data) != string50 {
		t.Errorf("Expected '%s', got: %s", string50, data)
	}

	s = CiString50(generateStr(51)) // Generates a 51-character string
	_, err = s.MarshalText()
	if err == nil {
		t.Error("Expected error for string exceeding 50 characters")
	}
}

func TestCiString50_UnmarshalText(t *testing.T) {
	string50 := generateStr(50)
	var s CiString50
	err := s.UnmarshalText([]byte(string50)) // Generates a 50-character string
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if s.String() != string50 { // Compare using the String() method
		t.Errorf("Expected '%s', got: %q", string50, s.String())
	}

	err = s.UnmarshalText([]byte(generateStr(51))) // Generates a 51-character string
	if err == nil {
		t.Error("Expected error for string exceeding 50 characters")
	}
}

func TestCiString255_IsValid(t *testing.T) {
	valid := CiString255(generateStr(255)) // Generates a 255-character string
	if !valid.IsValid() {
		t.Errorf("Expected string %q to be valid", valid)
	}

	invalid := CiString255(generateStr(256)) // Generates a 256-character string
	if invalid.IsValid() {
		t.Errorf("Expected string %q to be invalid", invalid)
	}
}

func TestCiString255_MarshalText(t *testing.T) {
	string255 := generateStr(255)
	s := CiString255(string255) // Generates a 255-character string
	data, err := s.MarshalText()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if string(data) != string255 {
		t.Errorf("Expected '%s', got: %s", string255, data)
	}

	s = CiString255(generateStr(256)) // Generates a 256-character string
	_, err = s.MarshalText()
	if err == nil {
		t.Error("Expected error for string exceeding 255 characters")
	}
}

func TestCiString255_UnmarshalText(t *testing.T) {
	string255 := generateStr(255)
	var s CiString255
	err := s.UnmarshalText([]byte(string255)) // Generates a 255-character string
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if s.String() != string255 { // Compare using the String() method
		t.Errorf("Expected '%s', got: %q", string255, s.String())
	}

	err = s.UnmarshalText([]byte(generateStr(256))) // Generates a 256-character string
	if err == nil {
		t.Error("Expected error for string exceeding 255 characters")
	}
}

func TestCiString500_IsValid(t *testing.T) {
	valid := CiString500(generateStr(500)) // Generates a 500-character string
	if !valid.IsValid() {
		t.Errorf("Expected string %q to be valid", valid)
	}

	invalid := CiString500(generateStr(501)) // Generates a 501-character string
	if invalid.IsValid() {
		t.Errorf("Expected string %q to be invalid", invalid)
	}
}

func TestCiString500_MarshalText(t *testing.T) {
	string500 := generateStr(500)
	s := CiString500(string500) // Generates a 500-character string
	data, err := s.MarshalText()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if string(data) != string500 {
		t.Errorf("Expected '%s', got: %s", string500, data)
	}

	s = CiString500(generateStr(501)) // Generates a 501-character string
	_, err = s.MarshalText()
	if err == nil {
		t.Error("Expected error for string exceeding 500 characters")
	}
}

func TestCiString500_UnmarshalText(t *testing.T) {
	string500 := generateStr(500)
	var s CiString500
	err := s.UnmarshalText([]byte(string500)) // Generates a 500-character string
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if s.String() != string500 { // Compare using the String() method
		t.Errorf("Expected '%s', got: %q", string500, s.String())
	}

	err = s.UnmarshalText([]byte(generateStr(501))) // Generates a 501-character string
	if err == nil {
		t.Error("Expected error for string exceeding 500 characters")
	}
}
