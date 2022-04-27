package mqconsumer_routes

import (
	"fmt"
	rabbit "notifications/internal/rabbtit"
	"notifications/pkg/infraestructure/mqueue/consumer/handlers"
)

func ConsumerRoutes(rabbit rabbit.MQueue, mqHandler mqHandler.MQHandler) error {
	fmt.Println("INICIANDO CONSUMER ROUTES")
	//consumer for new work request
	rabbit.StartExangeConsumer("WorkRequest", "workrequest_new", mqHandler.WorkRequestCreated, 1)
	rabbit.StartExangeConsumer("WorkRequest", "workrequest_accepted", mqHandler.WorkRequestAccepted, 1)
	rabbit.StartExangeConsumer("WorkRequest", "workrequest_denied", mqHandler.WorkRequestDenied, 1)
	return nil
}
