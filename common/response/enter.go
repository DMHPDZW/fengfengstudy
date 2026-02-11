package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Response(r *http.Request, w http.ResponseWriter, res any, err error) {
	//根据不同的错误码，返回不同的错误信息
	if err != nil {
		body := Body{
			Code: 400,
			Data: nil,
			Msg:  "错误",
		}
		httpx.WriteJson(w, http.StatusOK, body)
		return
	}
	body := Body{
		Code: 200,
		Data: res,
		Msg:  "成功",
	}
	httpx.WriteJson(w, http.StatusOK, body)

}
