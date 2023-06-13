package response

import (
	"bytes"
	contract2 "github.com/ArtisanCloud/PowerLibs/v3/logger/contract"
	"log"
	"net/http"
	"net/http/httputil"
)

func LogResponse(logger contract2.LoggerInterface, response *http.Response) {
	var output bytes.Buffer
	output.Write([]byte("\r\n------------------\r\n"))
	output.Write([]byte("response content:\r\n"))
	dumpRes, _ := httputil.DumpResponse(response, true)
	output.Write(dumpRes)
	log.Println(output.String())
}
