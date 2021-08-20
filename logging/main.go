package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	//"log"
)

//Config for logging
type Config struct {
	// Enable console logging
	ConsoleLoggingEnabled bool

	// FileLoggingEnabled makes the framework log to a file
	// the fields below can be skipped if this value is false!
	FileLoggingEnabled bool
	// Filename is the name of the logfile which will be placed inside the directory
	Filename string
	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int
	// MaxBackups the max number of rolled files to keep
	MaxBackups int
	// MaxAge the max age in days to keep a logfile
	MaxAge int
	// Compress denotes whether to compress backup files or not
	Compress bool
}

//Configure configures a zero log instance
func (config *Config) Configure() zerolog.Logger {
	var writers []io.Writer
	if config.ConsoleLoggingEnabled {
		writers = append(writers, os.Stdout)
	}
	if config.FileLoggingEnabled {
		writers = append(writers, config.NewRollingFile())
	}
	logger := zerolog.New(io.MultiWriter(writers...)).With().Timestamp().Logger()
	return logger
}

//NewRollingFile creates a new rolling log file with lumberjack
func (config *Config) NewRollingFile() io.Writer {
	return &lumberjack.Logger{
		Filename:   config.Filename,
		MaxBackups: config.MaxBackups, // files
		MaxSize:    config.MaxSize,    // megabytes
		MaxAge:     config.MaxAge,     // days
		Compress:   config.Compress,
	}
}

func main() {
	_, err := os.Stat("logs")
	if os.IsNotExist(err) {
		fmt.Println("Directory logs does not exist. Proceeding with directory creation")
		os.Mkdir("logs", 0755)
	}
	_, err = os.Stat("logs/test.log")
	if os.IsNotExist(err) {
		fmt.Println("test log file does not exist")
	}
	fileDescriptor, err := os.OpenFile("logs/test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error while creating file : ", err.Error())
		return
	}
	logConfig := &Config{
		ConsoleLoggingEnabled: true,
		FileLoggingEnabled:    true,
		MaxSize:               1,
		MaxAge:                30,
		Compress:              false,
		Filename:              fileDescriptor.Name(),
		MaxBackups:            3,
	}
	logger := logConfig.Configure()
	logger.Info().Msg("Started generating test logs")
	payload := `Software Engineer with over three and half years of total experience and two years of expertise in developing microservices-based applications with Golang. Experience in working with popular cloud-native technologies. Proficient communication skills and ability to perform well in a team.`
	index := 1
	for {
		logger.Info().Int("index", index).Msg(payload)
		index++
		time.Sleep(time.Millisecond)
	}
}
