// Copyright 2013 Cameron Walters. All Right Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package bale provides wrappers around go logging libraries, allowing users of the
// library to easily swap between logging implementations without changing client code.
package bale

// Logger defines the minimum set of methods implemented by all logging implementations.
type Logger interface {
	logPrinter
	logFormatter
}

// Log functions that handle arguments like fmt.Print.
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

// Log functions that handle arguments like fmt.Printf.
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

// Log functions that handle arguments like fmt.Println.
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
