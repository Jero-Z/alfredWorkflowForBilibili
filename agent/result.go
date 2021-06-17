package agent

type videoResult struct {
	Type         string `json:"type"`
	Id           int    `json:"id"`
	Author       string `json:"author"`
	Mid          int    `json:"mid"`
	Typeid       string `json:"typeid"`
	Typename     string `json:"typename"`
	Arcurl       string `json:"arcurl"`
	Aid          int    `json:"aid"`
	Bvid         string `json:"bvid"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Arcrank      string `json:"arcrank"`
	Pic          string `json:"pic"`
	Play         int    `json:"play"`
	VideoReview  int    `json:"video_review"`
	Favorites    int    `json:"favorites"`
	Tag          string `json:"tag"`
	Review       int    `json:"review"`
	Pubdate      int    `json:"pubdate"`
	Senddate     int    `json:"senddate"`
	Duration     string `json:"duration"`
	Badgepay     bool   `json:"badgepay"`
	ViewType     string `json:"view_type"`
	IsPay        int    `json:"is_pay"`
	IsUnionVideo int    `json:"is_union_video"`
	RankScore    int    `json:"rank_score"`
}

type typeResult struct {
	Result []videoResult
}

type searchTypeResult struct {
	Code    int
	Message string
	Ttl     int
	Data    typeResult
}
type searchAllResultItem struct {
	ResultType string
	Data       []videoResult
}

type allResult struct {
	Result []searchAllResultItem
}

type searchAllResult struct {
	Code    int
	Message string
	Ttl     int
	Data    allResult
}




