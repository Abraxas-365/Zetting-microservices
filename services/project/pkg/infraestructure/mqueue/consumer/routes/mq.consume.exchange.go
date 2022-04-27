package mqconsumer_routes

import (
	"fmt"
	"projects/internal/rabbit"
	"projects/pkg/infraestructure/mqueue/consumer/handlers"
)

func ConsumerRoutes(rabbit rabbit.MQueue, mqHandler mqHandler.MQHandler) error {
	fmt.Println("INICIANDO CONSUMER ROUTES")
	rabbit.StartExangeConsumer("WorkRequest", "anser_workrequest", mqHandler.AnsweredWorkRequest, 1)
	return nil
}
