package idcard

import (
	"fmt"
	"regexp"
	"strconv"
)

var id18 = regexp.MustCompile(`^\d{6}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`)

// CheckID18 18位身份证正则表达式验证
func CheckID18(idCard string) bool {
	return id18.MatchString(idCard)
}

var id15 = regexp.MustCompile(`^\d{6}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}$`)

// CheckID15 15位身份证正则表达式验证
func CheckID15(idCard string) bool {
	return id15.MatchString(idCard)
}

var W = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}

// Check 身份证有效性认证
func Check(idCard string) bool {
	sum := 0
	for idx, val := range W {
		i, err := strconv.ParseInt(string(idCard[idx]), 10, 32)
		if err != nil {
			fmt.Println(err)
			return false
		}
		sum += int(i) * val
	}
	// 校验位是X，则表示10
	if idCard[17] == 'X' || idCard[17] == 'x' {
		sum += 10
	} else {
		i, err := strconv.ParseInt(string(idCard[17]), 10, 32)
		if err != nil {
			return false
		}
		sum += int(i)
	}
	// 如果除11模1，则校验通过
	return sum%11 == 1
}
