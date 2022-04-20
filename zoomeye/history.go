package zoomeye

var (
	historyPath = "/both/search?history=true&ip=%s"
)

type HistoryRecord struct {
	Count int           `json:"count"`
	Data  []interface{} `json:"data"`
}

func (z *ZoomEyeClient) BothHistoty(ipaddr string) (hr *HistoryRecord, err error) {
	return
}
