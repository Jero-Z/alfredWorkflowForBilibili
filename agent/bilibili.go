package agent

import (
	"alfred-bilibli-search/alfred"
	"encoding/json"
	"errors"
	"fmt"
	strip "github.com/grokify/html-strip-tags-go"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	SEARCH_HOST               = "api.bilibili.com"
	SEARCH_CURRENT_TYPE_PATCH = "/x/web-interface/search/type"
)

type Client struct {
	QueryType string
	KeyWord   string
	Sort      string
}

var (
	queryTypeMaps           map[string]string
	qyerySortMaps           map[string]map[string]string
	ErrUnsupportedQueryType = errors.New("query type is unsupported")
	ErrUnsupportedSort      = errors.New("sort type is unsupported")
)

func init() {
	queryTypeMaps = make(map[string]string)
	queryTypeMaps = map[string]string{
		"v": "video",         //视频
		"a": "article",       //文章
		"u": "bili_user",     //用户
		"f": "media_bangumi", //番剧
		"m": "media_ft",      //影视
		"l": "live",          //直播
	}
	qyerySortMaps = make(map[string]map[string]string)
	qyerySortMaps = map[string]map[string]string{
		"v": {
			"c":  "click",   //click：点击率
			"p":  "pubdate", //pubdate：发布时间
			"dm": "dm",      //dm：弹幕
			"s":  "stow",    //stow：收藏
		},
	}
}

func (c *Client) SetKeyWord(key string) {
	c.KeyWord = key
}
func (c *Client) SetQueryType(queryType string) error {

	if queryType == "" {
		return nil
	}

	for k, _ := range queryTypeMaps {
		if k == queryType {
			c.QueryType = queryTypeMaps[queryType]
			break
		}
	}
	if c.QueryType == "" {
		return ErrUnsupportedQueryType
	} else {
		return nil
	}

}

func (c *Client) SetQueryWord(queryWord string) error {
	c.KeyWord = queryWord
	return nil
}

func (c *Client) SetQuerySort(querySort string) error {

	if querySort == "" || c.QueryType == "" {
		c.Sort = ""
		return nil
	}

	for k, v := range qyerySortMaps[c.QueryType] {
		if k == querySort {
			c.Sort = v
			break
		}
	}
	return nil
}

/*获取请求url*/
func (c *Client) getQueryUrl() (queryUrl string) {

	u := &url.URL{
		Host:   SEARCH_HOST,
		Scheme: "https",
	}
	u.ForceQuery = true
	q := u.Query()
	u.Path = SEARCH_CURRENT_TYPE_PATCH
	if len(c.QueryType) > 0 {
		q.Set("search_type", c.QueryType)
	} else {
		q.Set("search_type", "video")
	}

	q.Set("page", "1")
	q.Set("order", c.Sort)
	q.Set("keyword", c.KeyWord)
	q.Set("highlight", "1")
	u.RawQuery = q.Encode()
	s := u.String()
	return s
}
func (c *Client) Query(alfredResult *alfred.Result) {
	queryUrl := c.getQueryUrl()

	resp, err := http.Get(queryUrl)
	defer resp.Body.Close()

	if err != nil || resp.StatusCode != 200 {
		log.Fatal(fmt.Sprintf("url:%s请求失败,reasons:%s", queryUrl, resp.Status))
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	res := &searchTypeResult{}
	err = json.Unmarshal(b, res)

	if err != nil {
		fmt.Println("json Unmarshal error:", err)
		return
	}
	for _, v := range res.Data.Result {
		titleText := strip.StripTags(v.Title)
		if len(titleText) > 100 {
			titleText = titleText[:100]
		}
		item := alfred.Item{
			Title:    titleText,
			Valid:    true,
			Subtitle: v.Author,
			Arg:      v.Arcurl,
		}
		alfredResult.Append(item)
	}
}
