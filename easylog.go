package easylogger

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger(config string) error {
	if config == "" {
		return errors.New("No config found")
	} else {
		er := godotenv.Load("log.env")
		if er != nil {
			return er
		} else {
			logfile := os.Getenv("LOG_LOCATION")
			if logfile == "" {
				logfile = "./app.log"
			}
			logsize := os.Getenv("LOG_MAX_SIZE")
			lsize, err := strconv.Atoi(logsize)
			if err != nil {
				// default log sizze 100 MB before rotary
				lsize = 100
			}
			logage := os.Getenv("LOG_AGE")
			lage, er := strconv.Atoi(logage)
			if er != nil {
				lage = 28
			}
			lloger := &lumberjack.Logger{
				Filename: logfile,
				MaxSize:  lsize,
				MaxAge:   lage,
				Compress: true,
			}
			// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

			writerConsole := zerolog.ConsoleWriter{Out: os.Stdout}

			fileLogger := zerolog.New(lloger).With().Timestamp().Logger()
			multilogger := zerolog.MultiLevelWriter(fileLogger, writerConsole)
			log.Logger = zerolog.New(multilogger)
			return nil
		}
	}
}
