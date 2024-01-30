# easylogger
---
> implementation 
> - [godotenv](http://github.com/joho/godotenv)
> - [zerolog](https://github.com/rs/zerolog)
> - [lumberjack v2](http://gopkg.in/natefinch/lumberjack.v2)
> for easy rotary logging

| log.env file located with same folder of your compiled app

> log.env file :
> LOG_LOCATION : absolute path log filename
> LOG_MAX_SIZE : maximum log size (in MB) default 100MB
> LOG_AGE : maximum age of file before deleted, default 28 days