package mqRoutes

import (
	"fmt"
	"projects/internal/rabbit"
	"projects/workrequest/infraestructure/mqueue/consumer/handlers"
)

func ConsumerRoutes(rabbit rabbit.MQueue, mqHandler mqHandler.MQHandler) error {
	fmt.Println("INICIANDO CONSUMER ROUTES")
	rabbit.StartExchangeConsumer("WorkRequest", "accepted", mqHandler.WorkRequestAccepted, 1)
	return nil
}
