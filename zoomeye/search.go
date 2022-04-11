package zoomeye

type Matchers struct {
	Matches []interface{} `json:"matches"`
	facets  string        `json:"facets"`
	total   string        `json:"total"`
}

func (z *ZoomEyeClient) Search(query, facets, subtype string, page int) *Matchers {
	var matches Matchers

	return &matches
}
