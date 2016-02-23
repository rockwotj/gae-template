package app

type Test struct {
	Timestamp int64  `json:"timestamp"`
	Message   string `json:"message"`
}

type TestResults []Test

type Error struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}
