package should_write_response_status

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/responseWriter"
	"github.com/EugeneGpil/tester"
)

func Test_should_write_response_status(t *testing.T) {
	tester.SetTester(t)

	writer := responseWriter.New()

	writer.WriteHeader(http.StatusOK)

	tester.AssertSame(http.StatusOK, writer.GetStatus())
}
