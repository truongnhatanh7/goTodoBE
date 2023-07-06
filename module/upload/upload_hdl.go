package upload

import (
	"bufio"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/truongnhatanh7/goTodoBE/common"
	"github.com/truongnhatanh7/goTodoBE/component/uploadprovider"
	"gorm.io/gorm"
)

// NOTE: this handler is very simple and please do not use it in practice
// Instead, I recommend you should check "Upload Image to AWS S3 and CDN with CloudFront"

func Upload(db *gorm.DB, provider *uploadprovider.S3Provider) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		dst := fmt.Sprintf("static/%d.%s", time.Now().UTC().UnixNano(), fileHeader.Filename)

		file, err := fileHeader.Open()
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		defer file.Close()

		var bs []byte
		_, err = bufio.NewReader(file).Read(bs)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		img, err := provider.SaveFileUploaded(bs, dst)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		w, h, err := getImageDimension(bufio.NewReader(file))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

    var CDN_URL string = os.Getenv("CDN_URL")
		img.Width = w
		img.Height = h
    img.Fulfill(CDN_URL)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}

func getImageDimension(reader io.Reader) (int, int, error) {
	image, _, err := image.DecodeConfig(reader)
	if err != nil {
		return 0, 0, err
	}
	return image.Width, image.Height, nil
}
