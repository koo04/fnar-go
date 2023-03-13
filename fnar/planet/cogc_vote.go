package planet

type COGCVote struct {
	CompanyName    string `json:"CompanyName"`
	CompanyCode    string `json:"CompanyCode"`
	Influence      int32  `json:"Influence"`
	VoteType       string `json:"VoteType"`
	VoteTimeEpocMs int64  `json:"VoteTimeEpocMs"`
}
