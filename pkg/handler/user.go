package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"server_go/internal/user"
	"server_go/pkg/transport"
)

func NewUserHTTPServer(ctx context.Context, router *http.ServeMux, endpoints user.EndPoints) {
	router.HandleFunc("/users", UserServer(ctx, endpoints))
}

func UserServer(ctx context.Context, endpoints user.EndPoints) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		tran := transport.New(w, r, ctx)
		switch r.Method {
		case http.MethodGet:
			tran.Server(
				transport.Endpoint(endpoints.GetAll),
				decodeGetAllUser,
				encodeResponse,
				encodeError)
				return
		/*case http.MethodPost:
			decode := json.NewDecoder(r.Body)
			var req CreateReq
			if err := decode.Decode(&req); err != nil {
				MsgResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			PostUser(ctx, s, w, req)

		default:
			InvalidMethod(w)*/
		}
		InvalidMethod(w)
	}
}

func decodeGetAllUser(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(ctx context.Context,w http.ResponseWriter,resp interface{})error{
	data,err:=json.Marshal(resp)
	if err!=nil{
	return err
	}
	status:= http.StatusOK
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/-json; charset=utf-8")
	fmt.Fprintf(w,`{"status" %d, "data":%s}`, status,data)
	return nil
	}

func encodeError(_ context.Context,err error, w http.ResponseWriter){
	status:= http.StatusInternalServerError
	w.WriteHeader(status)
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	fmt.Fprintf(w, `{"status": %d, "message":%s}`, status, err.Error())
}


func InvalidMethod(w http.ResponseWriter) {
	status := http.StatusNotFound
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d,"message": "method doesn't exist"}`, status)
}
