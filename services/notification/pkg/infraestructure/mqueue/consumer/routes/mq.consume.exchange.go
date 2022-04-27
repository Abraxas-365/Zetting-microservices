package mqconsumer_routes

import (
	"fmt"
	rabbit "notifications/internal/rabbtit"
	"notifications/pkg/infraestructure/mqueue/consumer/handlers"
)

func ConsumerRoutes(rabbit rabbit.MQueue, mqHandler mqHandler.MQHandler) error {
	fmt.Println("INICIANDO CONSUMER ROUTES")
	//consumer for new work request
	rabbit.StartExangeConsumer("WorkRequest", "workrequest_new", mqHandler.WorkRequest, 1)
	return nil
}
