package handler

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/entity"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/pkg/deliveryutil"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/usecase/todo"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/usecase/todo/model"
	"github.com/spf13/cast"
)

func NewTodoHandler(todoUsecase todo.TodoUsecase) TodoHandler {
	return TodoHandler{
		todoUsecase: todoUsecase,
	}
}

type TodoHandler struct {
	todoUsecase todo.TodoUsecase
}

func (h TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /todo-items Todo CreateTodo
	// create new Todo
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	//   - name: Body
	//     in: body
	//     schema:
	//       "$ref": "#/definitions/TodoDto"
	// responses:
	//   '200':
	//     description: success

	var requestBody model.TodoDto
	if err := deliveryutil.ReadRequestBody(w, r, &requestBody); err != nil {
		deliveryutil.ResponseErrorJson(w, r, err)
		return
	}

	ctx := r.Context()
	result, err := h.todoUsecase.Create(ctx, requestBody)
	if err != nil {
		deliveryutil.ResponseErrorJson(w, r, err)
		return
	}

	deliveryutil.ResponseJson(w, r, result)
}

func (h TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	// swagger:operation PATCH /todo-items/{id} Todo UpdateTodo
	// update Todo
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	//   - name: id
	//     in: path
	//     description: Todo id
	//     type: integer
	//     required: true
	//   - name: Body
	//     in: body
	//     schema:
	//       "$ref": "#/definitions/TodoDto"
	// responses:
	//   '200':
	//     description: success

	var requestBody model.TodoDto
	if err := deliveryutil.ReadRequestBody(w, r, &requestBody); err != nil {
		deliveryutil.ResponseErrorJson(w, r, err)
		return
	}
	requestBody.Id = cast.ToInt64(mux.Vars(r)["id"])

	ctx := r.Context()
	result, err := h.todoUsecase.Update(ctx, requestBody)
	if err != nil {
		deliveryutil.ResponseErrorJson(w, r, err)
		return
	}

	deliveryutil.ResponseJson(w, r, result)
}

func (h TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// swagger:operation DELETE /todo-items/{id} Todo DeleteTodo
	// delete Todo
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   description: Todo id
	//   type: integer
	//   required: true
	// responses:
	//   '200':
	//     description: success

	requestBody := model.TodoDto{
		Id: cast.ToInt64(mux.Vars(r)["id"]),
	}
	ctx := r.Context()
	result, err := h.todoUsecase.Delete(ctx, requestBody)
	if err != nil {
		deliveryutil.ResponseErrorJson(w, r, err)
		return
	}

	deliveryutil.ResponseJson(w, r, result)
}

func (h TodoHandler) GetById(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /todo-items/{id} Todo GetTodoByID
	// get Todo by id
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   description: Todo id
	//   type: integer
	//   required: true
	// responses:
	//   '200':
	//     description: success

	requestBody := model.TodoDto{
		Id: cast.ToInt64(mux.Vars(r)["id"]),
	}
	ctx := r.Context()
	result, err := h.todoUsecase.GetById(ctx, requestBody)
	if err != nil {
		deliveryutil.ResponseErrorJson(w, r, err)
		return
	}

	deliveryutil.ResponseJson(w, r, result)
}

func (h TodoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /todo-items Todo GetAllTodo
	// get all Todo
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	//   - name: is_active
	//     in: query
	//   - name: title
	//     in: query
	//   - name: activity_group_id
	//     in: query
	//   - name: priority
	//     in: query
	// responses:
	//   '200':
	//     description: success

	ctx := r.Context()
	filter := model.TodoDto{}

	if r.URL.Query().Has("title") {
		val := r.URL.Query().Get("title")
		filter.Title = &val
	}
	if r.URL.Query().Has("activity_group_id") {
		val := cast.ToInt64(r.URL.Query().Get("activity_group_id"))
		filter.ActivityGroupId = &val
	}
	if r.URL.Query().Has("is_active") {
		val := cast.ToBool(r.URL.Query().Get("is_active"))
		filter.IsActive = &val
	}
	if val := strings.ReplaceAll(r.URL.Query().Get("priority"), " ", ""); val != "" {
		priority := entity.TodoPriority(val)
		filter.Priority = &priority
	}

	result, err := h.todoUsecase.GetAll(ctx, filter)
	if err != nil {
		deliveryutil.ResponseErrorJson(w, r, err)
		return
	}

	deliveryutil.ResponseJson(w, r, result)
}
