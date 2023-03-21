package deliveryutil

import (
	"encoding/json"
	"fmt"
	"net/http"

	pkgError "github.com/muhammadandikakurniawan/golang-todolistapp/src/pkg/error"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/pkg/model"
)

func Log(data interface{}, r *http.Request) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("error : %v\n", err)
		return
	}
	fmt.Printf("response %s : %s\n", r.URL.Path, string(dataBytes))
}

func ResponseJson[T any](w http.ResponseWriter, r *http.Request, model model.BaseResponseModel[T]) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(model.HttpStatusCode)
	Log(model, r)
	json.NewEncoder(w).Encode(model)
}

func ResponseErrorJson(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	model := model.BaseResponseModel[struct{}]{
		Message:        "internal server error",
		Status:         "Error",
		ErrorMessage:   err.Error(),
		HttpStatusCode: pkgError.INTERNAL_SERVER_ERROR.ToHttpStatus(),
		StatusCode:     string(pkgError.INTERNAL_SERVER_ERROR),
	}
	Log(model, r)
	json.NewEncoder(w).Encode(model)
}

func ReadRequestBody(w http.ResponseWriter, r *http.Request, dst interface{}) (err error) {
	err = json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		ResponseErrorJson(w, r, err)
		return
	}
	return
}
