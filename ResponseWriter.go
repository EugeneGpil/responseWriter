package responseWriter

import (
	"net/http"

	"github.com/EugeneGpil/responseWriter/app/modules/ResponseHolder"
)

type ResponseWriter struct {
	holder *ResponseHolder.ResponseHolder
}

func New() ResponseWriter {
	holder := ResponseHolder.New()

	return ResponseWriter{
		holder: &holder,
	}
}

// Implementation of net/http.responseWriter interface
func (writer ResponseWriter) Header() http.Header {
	return http.Header{}
}

// Implementation of net/http.responseWriter interface
func (writer ResponseWriter) Write(message []byte) (int, error) {
	writer.holder.Write(message)

	return 0, nil
}

// Implementation of net/http.responseWriter interface
func (writer ResponseWriter) WriteHeader(statusCode int) {
	writer.holder.WriteHeader(statusCode)
}

func (writer ResponseWriter) GetMessages() [][]byte {
	return writer.holder.GetMessages()
}

func (writer ResponseWriter) GetStatus() int {
	return writer.holder.GetStatus()
}
