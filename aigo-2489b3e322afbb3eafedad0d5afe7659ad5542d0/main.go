package main

import (
	"fmt"
	"net/http"
	"os"

	"gitlab.com/axisnet-modernization/backend-squad/axisnet-package/aigo/app/repository"
	"gitlab.com/axisnet-modernization/backend-squad/axisnet-package/aigo/app/routes"

	"gitlab.com/axisnet-modernization/backend-squad/axisnet-package/aigo/app/controllers"

	"gitlab.com/axisnet-modernization/backend-squad/axisnet-package/aigo/app/usecase"

	"gitlab.com/axisnet-modernization/library/adapter"

	_ "github.com/joho/godotenv/autoload"

	"github.com/sirupsen/logrus"
	"gitlab.com/axisnet-modernization/library/utils"
)

const ServiceName = "Aigo"

func init() {
	utils.LoadConfig(ServiceName)
	adapter.LoadComponent()
	adapter.LoadMongo()
}

func main() {
	dbapps := adapter.DB()
	repo := repository.NewRepo(dbapps)
	usecase := usecase.NewUC(repo)
	ctrl := controllers.NewControllers(usecase)

	route := routes.NewRoute(ctrl)
	router := route.Router()

	logrus.Infof("[SERVER] starting server in port :%v", os.Getenv("SERVER_PORT"))
	logrus.Fatalln(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("SERVER_PORT")), handler(router)))
	logrus.Exit(0)
}

func handler(router http.Handler) http.Handler {
	wrapper := utils.Wrapper(router)
	return wrapper
}
