package protocol

type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
} // let's keep interafce{} because why not?
// also that's compatibility!
