package middle_ware

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type logWrite struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (lw logWrite) Write(b []byte) (int, error) {
	lw.body.Write(b)
	return lw.ResponseWriter.Write(b)
}

func LogMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		reqBody, _ := ioutil.ReadAll(request.Body)
		logrus.Infof("path: %s, method: %s, remote_addr: %s, body: %s",
			request.URL, request.Method, request.RemoteAddr, reqBody)
		request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		lw := &logWrite{writer, &bytes.Buffer{}}
		h.ServeHTTP(lw, request)
		logrus.Infof("path: %s, method: %s, remote_addr: %s, resp: %s",
			request.URL, request.Method, request.RemoteAddr, lw.body.String())
	})
}
