package internal

type MessageType int

const (
	MessageTypeAuth MessageType = iota
	MessageTypeFile
	MessageTypeFileMetaData
	MessageTypeAck
	MessageTypeFileEnd
)

var MessageTypeValue = map[MessageType]string{
	MessageTypeAuth:         "MessageTypeAuth",
	MessageTypeFile:         "MessageTypeFile",
	MessageTypeFileMetaData: "MessageTypeFileMetaData",
	MessageTypeAck:          "MessageTypeAck",
	MessageTypeFileEnd:      "MessageTypeFileEnd",
}

func (ss MessageType) String() string {
	return MessageTypeValue[ss]
}

type Frame struct {
	ProductId        string       `json:"product_id"`
	Token            string       `json:"token"`
	FrameMessageType string       `json:"frame_message_type"`
	Payload          []byte       `json:"payload"`
	FileMetaData     FileMetaData `json:"file_meta_data"`
}

type FileMetaData struct {
	FileName string `json:"file_name"`
	Size     int64  `json:"size"`
}
