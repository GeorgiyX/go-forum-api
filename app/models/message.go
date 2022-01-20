package models

type Message struct {
	ErrCode     int    `json:"-"`
	Description string `json:"message"`
	Details     string `json:"message_details"`
}

func (msg *Message) Error() string {
	return msg.Description + " " + msg.Details
}

func (msg *Message) Code() int {
	return msg.ErrCode
}

func (msg *Message) SetDetails(text string) *Message {
	msg.Details = text
	return msg
}
