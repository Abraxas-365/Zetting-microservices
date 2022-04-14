package ports

import "work-request/pkg/core/models"

type WorkRequestMQueue interface {
	Publish(workrequest models.WorkRequest) error
}
