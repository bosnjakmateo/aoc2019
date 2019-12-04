package utils

import (
	"io/ioutil"
	"strings"
)

func FileInt(filePath string) (values []int) {
	return ToIntArray(FileString(filePath))
}

func FileString(filePath string) (values []string) {
	bytes, _ := ioutil.ReadFile(filePath)

	str := string(bytes)

	return strings.Split(str, ",")
}
