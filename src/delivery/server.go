package delivery

import (
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/container"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/delivery/http"
)

func Run(dependencyContainer *container.Container) {
	http.Run(dependencyContainer)
}
