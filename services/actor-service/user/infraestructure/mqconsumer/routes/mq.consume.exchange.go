package mqRoutes

import (
	"actor-service/internal/rabbit"
	"actor-service/user/infraestructure/mqconsumer/handlers"
	"fmt"
)

func ConsumerRoutes(rabbit rabbit.MQueue, mqHandler mqHandler.MQHandler) error {
	fmt.Println("INICIANDO CONSUMER ROUTES")
	rabbit.StartExchangeConsumer("User", "created", mqHandler.UserCreated, 1)
	return nil
}
