/************************************************************
 * Author:        jackey
 * Date:        2022/11/24
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package controller

import (
	"encoding/json"
	"io"
	"liteIm/pkg/logs"
	"net/http"
)

func getBody(r *http.Request) (res []byte, err error) {
	length := r.ContentLength
	body := make([]byte, length)
	_, err = r.Body.Read(body)
	logs.Info(string(body))
	if err != nil && err != io.EOF {
		logs.Error("getBody", err)
		return nil, err
	}
	return body, nil
}

func writeJson(w http.ResponseWriter, data interface{}) {
	res, _ := json.Marshal(data)
	_, _ = w.Write(res)
}
