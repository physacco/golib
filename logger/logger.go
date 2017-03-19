package logger

import (
  "io"
  "os"
  "log"
)

// verbosity levels
const (
    FATAL = iota
    ERROR
    WARNING
    INFO
    DEBUG
    VERBOSE
)

var verbosity int = VERBOSE

func GetLevel() int {
    return verbosity
}

func SetLevel(level int) {
    verbosity = level
}

func SetOutput(writer io.Writer) {
    log.SetOutput(writer)
}

func SetOutputFile(name string) (err error) {
    file, err := os.OpenFile(name,
        os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        return
    }

    log.SetOutput(file)
    return
}

func Fatal(format string, v ...interface{}) {
    if verbosity >= FATAL {
        log.Fatalf("FATAL " + format, v...)
    }
}

func Error(format string, v ...interface{}) {
    if verbosity >= ERROR {
        log.Printf("ERROR " + format, v...)
    }
}

func Warning(format string, v ...interface{}) {
    if verbosity >= WARNING {
        log.Printf("WARN  " + format, v...)
    }
}

func Info(format string, v ...interface{}) {
    if verbosity >= INFO {
        log.Printf("INFO  " + format, v...)
    }
}

func Debug(format string, v ...interface{}) {
    if verbosity >= DEBUG {
        log.Printf("DEBUG " + format, v...)
    }
}

func Verbose(format string, v ...interface{}) {
    if verbosity >= VERBOSE {
        log.Printf("VERBO " + format, v...)
    }
}
