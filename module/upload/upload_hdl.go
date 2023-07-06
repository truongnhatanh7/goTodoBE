package upload

import (
	"bufio"
	"fmt"

	"net/http"
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
    fmt.Println(provider)
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

    img, err := provider.SaveFileUploaded(bs,dst)
    if err != nil {
			panic(common.ErrInvalidRequest(err))
    }

		// img := common.Image{
    //   Id:        0,
		// 	Url:       dst,
		// 	Width:     100,
		// 	Height:    100,
		// 	CloudName: "local",
		// 	Extension: "",
		// }

		// img.Fulfill("http://localhost:3000")

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}