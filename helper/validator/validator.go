package validator

import "net/url"

//Returns the value as type string and true if its a string
func IsString(value any, expectedStringsInValue []string) (string, bool) {
	if value == nil {
		return "", false
	}
	str, isString := value.(string)
	if !isString {
		return "", false
	}

	if len(expectedStringsInValue) > 0 {
		for _, allowed := range expectedStringsInValue {
			if str == allowed {
				return allowed, true
			}
		}
		return "", false
	}

	return str, isString
}

func IsURL(value any) (*url.URL, bool) {
	if value == nil {
		return nil, false
	}

	str, ok := value.(string)
	if !ok {
		return nil, false
	}

	parsedURL, err := url.Parse(str)
	if err != nil {
		return nil, false
	}

	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return nil, false
	}

	return parsedURL, true
}

func IsArray(value any) ([]any, bool) {
	if value == nil {
		return nil, false
	}

	arr, ok := value.([]any)
	if !ok {
		return nil, false
	}

	if len(arr) == 0 {
		return nil, false
	}

	return arr, true
}

func IsNumber(value any, min int, max int) (int, bool) {
	if value == nil {
		return 0, false
	}

	number, ok := value.(int)
	if !ok {
		return 0, false
	}

	if number < min {
		return 0, false
	}

	if number > max {
		return 0, false
	}

	return number, false
}
