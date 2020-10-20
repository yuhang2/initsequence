package server

import (
	"fmt"

	"github.com/yuhang2/initsequence/external/aws"
	"github.com/yuhang2/initsequence/server/config"
	"github.com/yuhang2/initsequence/server/handlers"
)

func Serve() {
	appConfig := config.Load()
	fmt.Println("run config.Load(), result: ", appConfig)

	s3.InitializeSomething()

	handlers.UploadPhoto()
}
