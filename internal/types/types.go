package types

func ValidateData(data map[string]interface{}, fields map[string]string) bool {
	for key, value := range data {
		expectedType, ok := fields[key]
		if !ok {
			return false // Field not in template
		}
		if !IsTypeMatch(value, expectedType) {
			return false // Field type mismatch
		}
	}
	return true
}

func IsTypeMatch(value interface{}, expectedType string) bool {
	switch expectedType {
	case "string":
		_, ok := value.(string)
		return ok
	case "boolean":
		_, ok := value.(bool)
		return ok
	default:
		return false
	}
}
