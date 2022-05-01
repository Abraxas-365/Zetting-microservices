package main

import (
	"fmt"
	rabbit "notifications/internal/rabbit"
	notificationApp "notifications/notification/application"
	notificationRepo "notifications/notification/infraestructure/repository"
	restNotification "notifications/notification/infraestructure/rest/handlers"
	"notifications/notification/infraestructure/rest/routes"
	projectApp "notifications/project/application"
	projectMQHandler "notifications/project/infraestructure/mqueue/consumer/handlers"
	projectMQRoutes "notifications/project/infraestructure/mqueue/consumer/routes"
	projectRepo "notifications/project/infraestructure/repository"
	userApp "notifications/user/application"
	userMQHandler "notifications/user/infraestructure/mqueue/consumer/handlers"
	userMQRoutes "notifications/user/infraestructure/mqueue/consumer/routes"
	userRepo "notifications/user/infraestructure/repository"
	workRequestApp "notifications/workrequest/application"
	workrequestMQHandler "notifications/workrequest/infraestructure/mqueue/consumer/handlers"
	workrequestMQRoutes "notifications/workrequest/infraestructure/mqueue/consumer/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	//Infraestructure
	mongoUri := os.Getenv("MONGODB_URI")
	mqUri := os.Getenv("MQ_URI")
	mq, err := rabbit.NewMQueueConection(mqUri)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	//Repositories
	NotificationRepo, _ := notificationRepo.NewMongoRepository(mongoUri, "Zetting", 10, "Notifications")
	UserRepo, _ := userRepo.NewMongoRepository(mongoUri, "Zetting", 10, "NotificationUsers")
	ProjectRepo, _ := projectRepo.NewMongoRepository(mongoUri, "Zetting", 10, "NotificationProjects")

	//applications
	NotificationApp := notificationApp.NewNotificationApplication(NotificationRepo)
	UserApp := userApp.NewUserApplication(UserRepo)
	ProjectApp := projectApp.NewProjectApplication(ProjectRepo)
	WorkRequestApp := workRequestApp.NewWorkRequestApplication(NotificationApp)

	//MQ Handlers
	ProjectMQHandler := projectMQHandler.NewMQHandler(ProjectApp)
	UserMQHandler := userMQHandler.NewMQHandler(UserApp)
	WorkRequestMQHandler := workrequestMQHandler.NewMQHandler(WorkRequestApp)

	// MQ Routes
	projectMQRoutes.ConsumerRoutes(mq, ProjectMQHandler)
	userMQRoutes.ConsumerRoutes(mq, UserMQHandler)
	workrequestMQRoutes.ConsumerRoutes(mq, WorkRequestMQHandler)

	//Rest
	app := fiber.New()
	app.Use(logger.New())
	RestNotificationhandler := restNotification.NewNotificationHandler(NotificationApp)
	routes.NotificationRoute(app, RestNotificationhandler) //User routes
	fmt.Println("inicando en puerto 3002")
	app.Listen(":3002")

}
