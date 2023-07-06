package uploadprovider

import (
	"context"

	"github.com/truongnhatanh7/goTodoBE/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}