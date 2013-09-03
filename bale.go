package bale

type Logger interface {
	logPrinter
	logFormatter
}

// Log functions that handle arguments in the manner of fmt.Print.
type logPrinter interface {
	// Info logs to the INFO log.
	Info(args ...interface{})

	// Warning logs to the WARNING and INFO logs.
	Warn(args ...interface{})

	// Error logs to the ERROR, WARNING, and INFO logs.
	Error(args ...interface{})

	// Fatal logs to the FATAL, ERROR, WARNING, and INFO logs, and is expected
	// to exit the program.
	Fatal(args ...interface{})
}

// Log functions that handle arguments in the manner of fmt.Printf.
type logFormatter interface {
	// Infof logs to the INFO log.
	Infof(format string, args ...interface{})

	// Warningf logs to the WARNING and INFO logs.
	Warnf(format string, args ...interface{})

	// Errorf logs to the ERROR, WARNING, and INFO logs.
	Errorf(format string, args ...interface{})

	// Fatalf logs to the FATAL, ERROR, WARNING, and INFO logs, and is expected
	// to exit the program.
	Fatalf(format string, args ...interface{})
}

// Log functions that handle arguments in the manner of fmt.Println.
type Logliner interface {
	// Infoln logs to the INFO log.
	Infoln(args ...interface{})

	// Warningln logs to the WARNING and INFO logs.
	Warnln(args ...interface{})

	// Errorln logs to the ERROR, WARNING, and INFO logs.
	Errorln(args ...interface{})

	// Fatalln logs to the FATAL, ERROR, WARNING, and INFO logs, and is expected
	// to exit the program.
	Fatalln(args ...interface{})
}
