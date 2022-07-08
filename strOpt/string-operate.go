package stringOpt

import "strings"

// 将 ( $124$ && $66$) || $253$ 这样的字符串，解析得到$符之间的数字，不使用正则。
// GetObjFromTwoSymbol  "($124$ && $66$) || $253$"  ---> ["124", "66", "253"]
func GetObjFromTwoSymbol(rawStr string) (obj []string) {

	sharpStr := ReplaceRuneInEvenPosition(rawStr, '$', '#') //  "($124# && $66#) || $253#"

	strSli := strings.SplitN(sharpStr, "#", -1) // ["($124", "&& $66", ") || $253"]

	for _, v := range strSli {

		if strings.Contains(v, "$") {
			pos := strings.Index(v, "$")
			obj = append(obj, v[pos+1:])
		}

	}
	return
}

/*
将第偶数个指定字符换成另一个字符
ReplaceRuneInEvenPosition 将第偶数个指定字符换成另一个字符
	"($124$&& $66#$) || $253$" ---> "($124# && $66#) || $253#"
*/
func ReplaceRuneInEvenPosition(str string, oldRune, newRune uint8) string {
	var newStr []uint8
	j := 0
	for i := 0; i < len(str); i++ {
		c := str[i]

		if str[i] == oldRune {
			j++
			if j%2 == 0 {
				c = newRune
			}
		}
		newStr = append(newStr, c)
	}
	return string(newStr)
}
