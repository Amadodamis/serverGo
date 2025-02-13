package handler

import (
	"net/http"
	"server_go/pkg/transport"
)

func NewUserHTTPServer(ctx context.Context, router *http.ServerMux, endpoints user.EndPoints) {
	router.HandleFunc("/users", UserServer(ctx, service))
}

func UserServeR(ctx context.Context, endpoints user.EndPoints) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		tran := transport.New(w, r, ctx)
		switch r.Method {
		case http.MethodGet:
			tran.Server(
				transport.Endpoint(endpoints.GetAll),
				decodeGetAllUser,
				encodeResponse,
			)
		case http.MethodPost:
			decode := json.NewDecoder(r.Body)
			var req CreateReq
			if err := decode.Decode(&req); err != nil {
				MsgResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			PostUser(ctx, s, w, req)

		default:
			InvalidMethod(w)
		}
	}
}

func decodeGetAllUser(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	status:= http.StatusOK
	w.WriteHeader(status)
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	fmt.Fprintf(w, `{"status": %d, "data":%s}`, status, data)
}

func InvalidMethod(w http.ResponseWriter) {
	status := http.StatusNotFound
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d,"message": "method doesn't exist"}`, status)
}

func MsgResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, status, message)
}

func DataResponse(w http.ResponseWriter, status int, users interface{}) {
	

}
