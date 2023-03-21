package http

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/muhammadandikakurniawan/golang-todolistapp/cmd/app/docs"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/container"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/delivery/http/handler"
	"github.com/rs/cors"
	"github.com/spf13/cast"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Run(dependencyContainer *container.Container) (err error) {
	router := mux.NewRouter()
	appChan := make(chan os.Signal, 1)
	signal.Notify(appChan, os.Interrupt, syscall.SIGTERM)

	port := cast.ToString(dependencyContainer.AppConfig.HttpPort)
	docs.Setup()
	docs.SwaggerInfo.Host = ":" + port
	router.PathPrefix("/api-docs").Handler(httpSwagger.WrapHandler)

	cors := cors.New(cors.Options{
		AllowCredentials: true,
	})

	// ================ START SETUP HANDLERS ================
	activityHandler := handler.NewActivityHandler(dependencyContainer.ActivityUsecase)
	todoHandler := handler.NewTodoHandler(dependencyContainer.TodoUsecase)
	// ================ END SETUP HANDLERS ================

	activityRoute := router.PathPrefix("/activity-groups").Subrouter()
	activityRoute.HandleFunc("", activityHandler.Create).Methods(http.MethodPost)
	activityRoute.HandleFunc("/{id}", activityHandler.Update).Methods(http.MethodPatch)
	activityRoute.HandleFunc("/{id}", activityHandler.Delete).Methods(http.MethodDelete)
	activityRoute.HandleFunc("/{id}", activityHandler.GetById).Methods(http.MethodGet)
	activityRoute.HandleFunc("", activityHandler.GetAll).Methods(http.MethodGet)

	todoRoute := router.PathPrefix("/todo-items").Subrouter()
	todoRoute.HandleFunc("", todoHandler.Create).Methods(http.MethodPost)
	todoRoute.HandleFunc("/{id}", todoHandler.Update).Methods(http.MethodPatch)
	todoRoute.HandleFunc("/{id}", todoHandler.Delete).Methods(http.MethodDelete)
	todoRoute.HandleFunc("/{id}", todoHandler.GetById).Methods(http.MethodGet)
	todoRoute.HandleFunc("", todoHandler.GetAll).Methods(http.MethodGet)

	handler := cors.Handler(router)

	go func() {
		log.Printf("HTTP running on port %s.", port)
		log.Println(http.ListenAndServe(fmt.Sprintf(":%s", port), handler))
	}()

	<-appChan
	log.Println("HTTP SERVER CLOSED")
	return
}
