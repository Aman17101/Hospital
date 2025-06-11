package util

import (
	"flag"
	"os"

	"github.com/Aman17101/Hospital/model"
	"github.com/sirupsen/logrus"
)

var Logger logrus.Logger

func init() {
	Logger = *logrus.New()

	Logger.Out = os.Stdout

	Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,                //add for highlighting
		TimestampFormat: `json:"created_at"`, // Go's reference time
		FullTimestamp:   true,
	})

}

func SetLogger() {
	logLevel := flag.String(model.LogLevel, model.LogLevelInfo, "log-level(debug,info,error,warn)")

	flag.Parse()

	switch *logLevel {
	case model.LogLevelInfo:
		Logger.SetLevel(logrus.InfoLevel)
	case model.LogLevelDebug:
		Logger.SetLevel(logrus.DebugLevel)
	case model.LogLevelError:
		Logger.SetLevel(logrus.ErrorLevel)
	default:
		Logger.SetLevel(logrus.WarnLevel)
	}

}

func Log(logLevel, packagelevel, functionName string, message, parameter interface{}) {
	switch logLevel {
	case model.LogLevelInfo:
		if parameter != nil {
			Logger.Infof("packageLevel %s,functionName %s,message %v,parameter %v", packagelevel, functionName, message, parameter)
		} else {
			Logger.Infof("packageLevel %s,functionName %s,message %v", packagelevel, functionName, message)
		}
		case model.LogLevelError:
		if parameter != nil {
			Logger.Errorf("packageLevel %s,functionName %s,message %v,parameter %v", packagelevel, functionName, message, parameter)
		} else {
			Logger.Errorf("packageLevel %s,functionName %s,message %v", packagelevel, functionName, message)
		}
		case model.LogLevelWarning:
		if parameter != nil {
			Logger.Warnf("packageLevel %s,functionName %s,message %v,parameter %v", packagelevel, functionName, message, parameter)
		} else {
			Logger.Warnf("packageLevel %s,functionName %s,message %v", packagelevel, functionName, message)
		}
		default:
		if parameter != nil {
			Logger.Infof("packageLevel %s,functionName %s,message %v,parameter %v", packagelevel, functionName, message, parameter)
		} else {
			Logger.Infof("packageLevel %s,functionName %s,message %v", packagelevel, functionName, message)
		}
	}
}
