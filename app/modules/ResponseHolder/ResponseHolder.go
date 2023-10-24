package ResponseHolder

type ResponseHolder struct{
	messages [][]byte
	status   int
}

func New() ResponseHolder {
	return ResponseHolder{
		messages: [][]byte{},
		status:   0,
	}
}

func (holder *ResponseHolder) Write(message []byte) (int, error) {
	cup := append(holder.messages, message)

	holder.messages = cup

	return 0, nil
}

func (holder *ResponseHolder) WriteHeader(statusCode int) {
	holder.status = statusCode
}

func (holder ResponseHolder) GetMessages() [][]byte {
	return holder.messages
}

func (holder ResponseHolder) GetStatus() int {
	return holder.status
}
