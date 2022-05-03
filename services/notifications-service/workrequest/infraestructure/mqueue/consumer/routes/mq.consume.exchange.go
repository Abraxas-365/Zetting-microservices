package mqRoutes

import (
	"fmt"
	rabbit "notifications/internal/rabbit"
	"notifications/workrequest/infraestructure/mqueue/consumer/handlers"
)

func ConsumerRoutes(rabbit rabbit.MQueue, mqHandler mqHandler.MQHandler) error {
	fmt.Println("INICIANDO CONSUMER ROUTES")
	//consumer for new work request
	rabbit.StartExchangeConsumer("WorkRequest", "new", mqHandler.WorkRequestCreated, 1)
	rabbit.StartExchangeConsumer("WorkRequest", "accepted", mqHandler.WorkRequestAccepted, 1)
	rabbit.StartExchangeConsumer("WorkRequest", "denied", mqHandler.WorkRequestDenied, 1)
	return nil
}
