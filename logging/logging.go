package logging

import (
	"fmt"
	"io"
	"os"

	"github.com/eddwinpaz/customer-svc/customer/entity"
	mylog "github.com/sirupsen/logrus"
)

// InitializeLogging asdas
func InitializeLogging(logFileName string) {

	var file, err = os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("%s: %s", entity.ErrOpenFileLog.Error(), err.Error())
	}

	mulWrt := io.MultiWriter(os.Stdout, file)
	mylog.SetOutput(mulWrt)

	mylog.SetFormatter(&mylog.JSONFormatter{})
	// // mylog.SetFormatter(&log.TextFormatter{})
	mylog.SetFormatter(&mylog.TextFormatter{
		ForceColors:   true, // Seems like automatic color detection doesn't work on windows terminals
		FullTimestamp: true,
	})

}
