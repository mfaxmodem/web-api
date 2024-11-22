package common

import (
	"log"
	"regexp"
)

const MobileNumberPattern string = `09(1[0-9]|2[0-9]|3[0-9]|9[0-9])[0-9]{7}$)`

func MobileNumberValidator(mobileNUmber string) bool {
	res, err := regexp.MatchString(MobileNumberPattern, mobileNUmber)
	if err != nil {
		log.Print(err.Error())
	}
	return res
}
