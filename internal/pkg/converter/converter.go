package converter

import "strings"

func ConvertStrArrToPgArr(strArr []string) string {
	return "{" + strings.Join(strArr, ",") + "}"
}
