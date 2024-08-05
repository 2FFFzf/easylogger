package easylogger

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

// default value : name : app.log, max file size :  200MB, max file age : 28 day
func InitLogger(config string) error {
	if config == "" {
		lloger := &lumberjack.Logger{
			Filename: "app.log",
			MaxSize:  200,
			MaxAge:   28,
			Compress: true,
		}
		writerConsole := zerolog.ConsoleWriter{Out: os.Stdout}
		multilogger := zerolog.MultiLevelWriter(lloger, writerConsole)
		log.Logger = zerolog.New(multilogger)
		zerolog.TimeFieldFormat = time.RFC822
		log.Logger = log.With().Timestamp().Logger()

		log.Info().Msg("Rotary Logger")
		return nil
	} else {
		er := godotenv.Load("log.env")
		if er != nil {
			//if failed to open use default value
			lloger := &lumberjack.Logger{
				Filename: "app.log",
				MaxSize:  200,
				MaxAge:   28,
				Compress: true,
			}
			writerConsole := zerolog.ConsoleWriter{Out: os.Stdout}
			multilogger := zerolog.MultiLevelWriter(lloger, writerConsole)
			log.Logger = zerolog.New(multilogger)
			zerolog.TimeFieldFormat = time.RFC822
			log.Logger = log.With().Timestamp().Logger()
			log.Error().Err(er).Stack().Caller().Msg("hala")
			return nil
		} else {
			logfile := os.Getenv("LOG_LOCATION")
			if logfile == "" {
				logfile = "./app.log"
			}
			logsize := os.Getenv("LOG_MAX_SIZE")
			lsize, err := strconv.Atoi(logsize)
			if err != nil {
				// default log sizze 200 MB before rotary
				lsize = 200
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

			multilogger := zerolog.MultiLevelWriter(lloger, writerConsole)
			log.Logger = zerolog.New(multilogger)
			zerolog.TimeFieldFormat = time.RFC822
			log.Logger = log.With().Timestamp().Logger()
			log.Info().Msg("Rotary Logger")
			return nil
		}
	}
}
