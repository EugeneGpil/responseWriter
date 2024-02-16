package should_set_and_get_header

import (
	"testing"

	"github.com/EugeneGpil/responseWriter"
	"github.com/EugeneGpil/tester"
)

const headerKey = "Content-Type"
const headerValue = "application/json"

func Test_should_set_and_get_header(t *testing.T) {
	tester.SetTester(t)

	writer := responseWriter.New()

	writer.Header().Add(headerKey, headerValue)

	result := writer.Header().Get(headerKey)

	tester.AssertSame(headerValue, result)
}
