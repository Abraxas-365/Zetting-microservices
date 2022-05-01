package mqRoutes

import (
	"fmt"
	"projects/internal/rabbit"
	"projects/user/infraestructure/mqueue/consumer/handlers"
)

func ConsumerRoutes(rabbit rabbit.MQueue, mqHandler mqHandler.MQHandler) error {
	fmt.Println("INICIANDO CONSUMER ROUTES")
	rabbit.StartExchangeConsumer("User", "created", mqHandler.UserCreated, 1)
	return nil
}
