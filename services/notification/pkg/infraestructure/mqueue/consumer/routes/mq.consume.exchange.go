package mqconsumer_routes

import (
	"fmt"
	rabbit "notifications/internal/rabbtit"
	"notifications/pkg/infraestructure/mqueue/consumer/handlers"
)

func ConsumerRoutes(rabbit rabbit.MQueue, mqHandler mqHandler.MQHandler) error {
	fmt.Println("INICIANDO CONSUMER ROUTES")
	rabbit.StartExangeConsumer("WorkRequest", "new_workrequest", mqHandler.WorkRequest, 1)
	return nil
}
