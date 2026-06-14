package protocol

type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
} // let's keep interafce{} because why not?
// also that's compatibility!

type FileOffer struct {
	Name   string `json:"name"`
	Size   int64  `json:"size"`
	SHA256 string `json:"sha256"`
}

type TextTrans struct {
	Text string `json:"text"`
}

type CBTrans struct {
	Text string `json:"text"`
}
