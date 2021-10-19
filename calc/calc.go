package calc

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func convertString(val string) (float64, bool, string) {

	if strings.Contains(val, ",") {
		val = strings.Replace(val, ",", ".", -1)
	} else if !strings.Contains(val, ".") {
		val = val + ".0"
	}

	res, err := strconv.ParseFloat(val, 64)

	if err != nil {
		return 0, false, "Cannot convert >>" + val + "<<"
	}

	return res, true, ""
}

func floatToStr(val float64) string {
	return strconv.FormatFloat(val, 'f', -1, 64)
}

func validateParam(a string) (val float64, isErr bool, errMsg string) {
	val, isOk, errMsg := convertString(a)
	if !isOk {
		return 0, true, errMsg
	}

	return val, false, ""
}

func validateParams(a string, b string) (val1 float64, val2 float64, isErr bool, errMsg string) {
	val1, isErr, errMsg = validateParam(a)
	if isErr {
		return 0, 0, true, errMsg
	}

	val2, isErr, errMsg = validateParam(b)
	if isErr {
		return 0, 0, true, errMsg
	}

	return val1, val2, false, ""
}

func Add(a string, b string) string {

	defer func() string {
		if err := recover(); err != nil {
			return fmt.Sprintf("Cannot add %v to %v", a, b)
		}
		return ""
	}()

	val1, val2, isErr, errMsg := validateParams(a, b)
	if isErr {
		return errMsg
	}

	return fmt.Sprintf("%v + %v =\n%v\n", a, b, floatToStr(val1+val2))
}

func Subtract(a string, b string) string {

	defer func() string {
		if err := recover(); err != nil {
			return fmt.Sprintf("Cannot subtract %v from %v", b, a)
		}
		return ""
	}()

	val1, val2, isErr, errMsg := validateParams(a, b)
	if isErr {
		return errMsg
	}

	return fmt.Sprintf("%v - %v =\n%v\n", a, b, floatToStr(val1-val2))
}

func Multiply(a string, b string) string {

	val1, val2, isErr, errMsg := validateParams(a, b)
	if isErr {
		return errMsg
	}

	return fmt.Sprintf("%v × %v =\n%v\n", a, b, floatToStr(val1*val2))
}

func Divide(a string, b string) string {

	defer func() string {
		if err := recover(); err != nil {
			return fmt.Sprintf("Cannot divide %v by %v", a, b)
		}
		return ""
	}()

	val1, val2, isErr, errMsg := validateParams(a, b)
	if isErr {
		return errMsg
	}

	if val1 == 0 && val2 == 0 {
		return "0 divided by 0 is undefined."
	}

	if val2 == 0 {
		return "Cannot divide by zero!"
	}

	res := val1 / val2

	if math.IsNaN(res) {
		return fmt.Sprintf("%v ÷ %v is not a number", a, b)
	}

	return fmt.Sprintf("%v ÷ %v =\n%v", a, b, floatToStr(res))
}

func Power(a string, b string) string {

	defer func() string {
		if err := recover(); err != nil {
			return fmt.Sprintf("Cannot calculate %v to the power of %v", a, b)
		}
		return ""
	}()

	val1, val2, isErr, errMsg := validateParams(a, b)
	if isErr {
		return errMsg
	}

	if val1 == 0 && val2 == 0 {
		return "0 to the power of 0 is undefined."
	}

	res := math.Pow(val1, val2)

	if math.IsNaN(res) {
		return fmt.Sprintf("%v ^ %v is not a number", a, b)
	}

	return fmt.Sprintf("%v ^ %v =\n%v\n", a, b, math.Pow(val1, val2))
}

func Pi() string {
	return "π =\n" + floatToStr(math.Pi)
}

func TwoPi() string {
	return "2π =\n" + floatToStr(2*math.Pi)
}

func PiSqrd() string {
	return "√π =" + floatToStr(math.SqrtPi)
}

func E() string {
	return "e =\n" + floatToStr(math.Exp(1))
}
