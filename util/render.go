package util

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

/**
后端统一返回 status_code = 200, 如果错误, 则将信息放在 baseResponse
*/

func JsonResponse(w http.ResponseWriter, response interface{}) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		ErrJsonResponse(w, err)
	}
	if _, err := w.Write(jsonResponse); err != nil {
		logrus.Error("write response to client meet err: %+v", err)
	}
}

func ErrJsonResponse(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusOK)
}
