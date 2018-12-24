package match

import "regexp"

//正则匹配,匹配规则可以一直加
func RegexpMatch(pattern_type string, source string) bool {
	pattern_list := map[string]string{}
	pattern_list["ip"] = "(2(5[0-5]{1}|[0-4]\\d{1})|[0-1]?\\d{1,2})(\\.(2(5[0-5]{1}|[0-4]\\d{1})|[0-1]?\\d{1,2})){3}"
	pattern_list["email"] = "^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\\.[a-zA-Z0-9-]+)*\\.[a-zA-Z0-9]{2,6}$"
	pattern_list["qq"] = "^[1-9]\\d{4,10}$"
	pattern := pattern_list[pattern_type]
	reg := regexp.MustCompile(pattern)
	if res := reg.FindAllString(source, -1); res == nil {
		return false
	} else {
		return true
	}
}
