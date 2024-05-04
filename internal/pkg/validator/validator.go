package validator

import "net/url"

func IsString(s interface{}) bool {
	_, ok := s.(string)
	return ok
}

func IsNumber(n interface{}) bool {
	_, ok := n.(int)
	return ok
}

func isValidUrl(strUrl string) bool {
	if strUrl == "" {
		return true // Allow empty strings, as this will be handled by other validations
	}

	_, err := url.ParseRequestURI(strUrl)
	return err == nil
}

func IsValidImageUrls(arr []string) bool {
	if len(arr) == 0 {
		return false
	}

	for _, item := range arr {
		if item == "" {
			return false
		}
		if !isValidUrl(item) {
			return false
		}
	}

	return true
}
