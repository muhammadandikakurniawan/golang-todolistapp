package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/pkg/deliveryutil"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/usecase/activity"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/usecase/activity/model"
	"github.com/spf13/cast"
)

func NewActivityHandler(activityUsecase activity.ActivityUsecase) ActivityHandler {
	return ActivityHandler{
		activityUsecase: activityUsecase,
	}
}

type ActivityHandler struct {
	activityUsecase activity.ActivityUsecase
}

func (h ActivityHandler) Create(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /activity-groups Activity CreateActivity
	// create new activity
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	//   - name: Body
	//     in: body
	//     schema:
	//       "$ref": "#/definitions/ActivityDto"
	// responses:
	//   '200':
	//     description: success

	var requestBody model.ActivityDto
	if err := deliveryutil.ReadRequestBody(w, r, &requestBody); err != nil {
		deliveryutil.ResponseErrorJson(w, r, err)
		return
	}

	ctx := r.Context()
	result, err := h.activityUsecase.Create(ctx, requestBody)
	if err != nil {
		deliveryutil.ResponseErrorJson(w, r, err)
		return
	}

	deliveryutil.ResponseJson(w, r, result)
}

func (h ActivityHandler) Update(w http.ResponseWriter, r *http.Request) {
	// swagger:operation PATCH /activity-groups/{id} Activity UpdateActivity
	// update activity
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   description: Activity id
	//   type: integer
	//   required: true
	// - name: Body
	//   in: body
	//   schema:
	//    "$ref": "#/definitions/ActivityDto"
	// responses:
	//   '200':
	//     description: success

	var requestBody model.ActivityDto
	if err := deliveryutil.ReadRequestBody(w, r, &requestBody); err != nil {
		deliveryutil.ResponseErrorJson(w, r, err)
		return
	}
	requestBody.Id = cast.ToInt64(mux.Vars(r)["id"])

	ctx := r.Context()
	result, err := h.activityUsecase.Update(ctx, requestBody)
	if err != nil {
		deliveryutil.ResponseErrorJson(w, r, err)
		return
	}

	deliveryutil.ResponseJson(w, r, result)
}

func (h ActivityHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// swagger:operation DELETE /activity-groups/{id} Activity DeleteActivity
	// delete activity
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   description: Activity id
	//   type: integer
	//   required: true
	// responses:
	//   '200':
	//     description: success

	requestBody := model.ActivityDto{
		Id: cast.ToInt64(mux.Vars(r)["id"]),
	}
	ctx := r.Context()
	result, err := h.activityUsecase.Delete(ctx, requestBody)
	if err != nil {
		deliveryutil.ResponseErrorJson(w, r, err)
		return
	}

	deliveryutil.ResponseJson(w, r, result)
}

func (h ActivityHandler) GetById(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /activity-groups/{id} Activity GetActivityByID
	// get Activity by id
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   description: Activity id
	//   type: integer
	//   required: true
	// responses:
	//   '200':
	//     description: success

	requestBody := model.ActivityDto{
		Id: cast.ToInt64(mux.Vars(r)["id"]),
	}
	ctx := r.Context()
	result, err := h.activityUsecase.GetById(ctx, requestBody)
	if err != nil {
		deliveryutil.ResponseErrorJson(w, r, err)
		return
	}

	deliveryutil.ResponseJson(w, r, result)
}

func (h ActivityHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /activity-groups Activity GetAllActivity
	// get all activity
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	//   - name: title
	//     in: query
	//   - name: email
	//     in: query
	// responses:
	//   '200':
	//     description: success

	ctx := r.Context()

	filter := model.ActivityDto{}

	if r.URL.Query().Has("title") {
		val := r.URL.Query().Get("title")
		filter.Title = &val
	}
	if r.URL.Query().Has("email") {
		val := r.URL.Query().Get("email")
		filter.Email = &val
	}

	result, err := h.activityUsecase.GetAll(ctx, filter)
	if err != nil {
		deliveryutil.ResponseErrorJson(w, r, err)
		return
	}

	deliveryutil.ResponseJson(w, r, result)
}
