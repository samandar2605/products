package v1

import (
	"errors"

	"github.com/samandar2605/products/config"
	"github.com/samandar2605/products/storage"
)

type handlerV1 struct {
	cfg     *config.Config
	storage storage.StorageI
}

var (
	ErrForbidden = errors.New("forbidden")
)

type HandlerV1Options struct {
	Cfg     *config.Config
	Storage storage.StorageI
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		cfg:     options.Cfg,
		storage: options.Storage,
	}
}

// func errorResponse(err error) *models.ErrorResponse {
// 	return &models.ErrorResponse{
// 		Error: err.Error(),
// 	}
// }
