package main

import (
	"alfred-bilibli-search/agent"
	"alfred-bilibli-search/alfred"
	"os"
)

func main() {
	result := alfred.NewResult()

	queryType, keyWord, sortType := parseArgs(os.Args)

	if len(keyWord) == 0 {
		result.Append(alfred.Item{
			Title:    "请输入搜索内容",
			Valid:    false,
			Subtitle: keyWord,
		})
		result.Run()
	}
	if len(keyWord) > 100 {
		result.Append(alfred.Item{
			Title:    "超出最大检索字符数100",
			Valid:    false,
			Subtitle: keyWord,
		})
		result.Run()
	}
	client := agent.Client{}
	client.KeyWord = keyWord
	if err := client.SetQueryType(queryType); err != nil {
		result.Append(alfred.Item{
			Title:    "暂不请求该请求类型",
			Valid:    false,
			Subtitle: keyWord,
		})
		result.Run()
	}
	if err := client.SetQuerySort(sortType); err != nil {
		result.Append(alfred.Item{
			Title:    "暂不支持该排序类型",
			Valid:    false,
			Subtitle: keyWord,
		})
		result.Run()
	}

	client.Query(result)

	result.Run()
}
