package utils

import (
	"fmt"
	"log/slog"
	"os"
)

/**
*
* BRIEF: This file will create the Logger that will be used
*		 across the server
* AUTHOR: Elias Lopes
 */

// this method will take the configs raw string and return the slog level
func getHandlerLvl(lvl LogLevel) slog.Level {
	// get the log level
	switch lvl {
	case "INFO":
		return slog.LevelInfo
	case "DEBUG":
		return slog.LevelDebug
	case "WARNING":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}

}

// this method opens/creates log files and opens them for writing
func init_log_handlers(logger *slog.Logger, logFiles []LogFileHandler) error {
	// multiple output handlers
	if len(logFiles) > 0 {
		// slice of handlers
		var handlers []slog.Handler

		// create handler for each
		for _, handlerObj := range logFiles {

			// open file for readng and writing
			file, err := os.OpenFile(
				handlerObj.FilePath,
				os.O_CREATE|os.O_WRONLY|os.O_APPEND,
				// farthest non-zero: Owner, Group, Others
				0662, // octal permissions
			)

			// handle IO error
			if err != nil {
				return fmt.Errorf("[ERROR] When attemping to open file")
			}

			defer file.Close()

			// create output handler
			handler := slog.NewTextHandler(file, &slog.HandlerOptions{
				Level: getHandlerLvl(handlerObj.LogLvl),
			})

			// append handler
			handlers = append(handlers, handler)
		}

		// add stdout too
		handlers = append(handlers, slog.NewTextHandler(os.Stdout, nil))

		// FIXME: ADD HANDLERS TO THE LOGGER
		// logger = slog.New(&MultiHandler{handlers})

	} else {
		// single output handler (stdout)
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

	}

	return nil
}

// this method to create a logger object
func NewLogger(cfg LoggerConfig) (*slog.Logger, error) {
	// declare logger object
	var logger slog.Logger

	// set up output handlers
	err := init_log_handlers(&logger, cfg.FileOutputHandler)

	// handler error
	if err != nil {
		fmt.Printf("[ERROR] When setting up output handlers: %v\n", err)
		return &logger, err
	}

	// set up log level

	return &logger, nil
}
