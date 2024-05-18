package types

import "testing"

func TestIsTypeMatch(t *testing.T) {
	tests := []struct {
		value        interface{}
		expectedType string
		expected     bool
	}{
		{"hello", "string", true},
		{true, "boolean", true},
		{123, "string", false},
		{"hello", "boolean", false},
		{true, "string", false},
	}

	for _, test := range tests {
		result := IsTypeMatch(test.value, test.expectedType)
		if result != test.expected {
			t.Errorf("isTypeMatch(%v, %s) = %v; want %v", test.value, test.expectedType, result, test.expected)
		}
	}
}

func TestValidateData(t *testing.T) {
	fields := map[string]string{
		"Name":        "string",
		"HealthCheck": "boolean",
	}

	tests := []struct {
		data     map[string]interface{}
		expected bool
	}{
		{map[string]interface{}{"Name": "Service", "HealthCheck": true}, true},
		{map[string]interface{}{"Name": "Service", "HealthCheck": "true"}, false}, // Incorrect type
		{map[string]interface{}{"Name": "Service"}, true},                         // Missing optional field
		{map[string]interface{}{"Name": 123, "HealthCheck": true}, false},         // Incorrect type
	}

	for _, test := range tests {
		result := ValidateData(test.data, fields)
		if result != test.expected {
			t.Errorf("validateData(%v) = %v; want %v", test.data, result, test.expected)
		}
	}
}
