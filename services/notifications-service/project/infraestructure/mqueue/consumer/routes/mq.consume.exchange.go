package mqRoutes

import (
	"fmt"
	"notifications/internal/rabbit"
	"notifications/project/infraestructure/mqueue/consumer/handlers"
)

func ConsumerRoutes(rabbit rabbit.MQueue, mqHandler mqHandler.MQHandler) error {
	fmt.Println("INICIANDO CONSUMER ROUTES")
	rabbit.StartExchangeConsumer("Project", "created", mqHandler.ProjectCreated, 1)
	return nil
}
