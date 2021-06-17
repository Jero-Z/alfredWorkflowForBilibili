package main

import (
	"fmt"
	"regexp"
	"strings"
)

/*

 */
var (
	queryPattern = regexp.MustCompile(`t=\w`)
	sortPattern  = regexp.MustCompile(`s=\w`)
)

func parseArgs(args []string) (queryType string, keyWord string, sort string) {
	// 0的位置为路径  参数个数少于2个则无输入
	if len(args) < 2 {
		return
	}
	hasQueryType := false
	hasSort := false

	queryString := ""
	queryArgs := args[1:]
	positionCount := len(queryArgs) - 1

	for i, v := range queryArgs {
		str := fmt.Sprintf("%s ", v)
		if i == positionCount {
			str = v
		}
		queryString += str
	}
	queryTypeStr := ""
	if v := queryPattern.FindAllStringSubmatch(queryString, -1); len(v) > 0 { // 匹配成功则代表指定了搜索类型
		hasQueryType = true
		queryTypeStr = v[0][0]
		queryType = queryTypeStr[2:]
	}
	querySortStr := ""
	if sortMatch := sortPattern.FindAllStringSubmatch(queryString, -1); len(sortMatch) > 0 {
		hasSort = true
		querySortStr = sortMatch[0][0]
		sort = querySortStr[2:]
	}
	keyWord = queryString

	if hasQueryType { //有类型限制
		keyWord = strings.Replace(keyWord, queryTypeStr, "", -1)
	}

	if hasSort {
		keyWord = strings.Replace(keyWord, querySortStr, "", -1)
	}
	keyWord = strings.TrimSpace(keyWord)
	return
}
