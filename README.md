# log
This package allows developers to easily create structure logs.  

## Installation
This package was written in the Go programming language.  Once you have Go installed, you can install this package by issuing the following command.
```
go get github.com/forestgiant/log
```

Once the command above has finished, you should be able to import the package for use in your own code as follows.
```
import "github.com/forestgiant/log"
```

## Log Levels
This package currently supports logging at the following levels.  These levels are automatically associated with the `level` key in the log output.
- Alert - alert
- Critical - critical
- Debug - debug
- Emergency - emergency
- Error - error
- Info - info
- Notice - notice
- Warning - warning

These can be used in the following manner:
```
logger := log.Logger{}
logger.Info("Info log message.", "key", "value")
```

## Log Formats
This package provides `Formatters` capable of structuring logs in the following formats.
- JSON - default
- logfmt

By default, messages are formatted using the JSON Formatter.  If you require another format, it is easy to provide a formatter to your logger as follows.
```
logger := log.Logger{Formatter: logger.LogfmtFormatter{}}
```

## Log Output
By default, this package will log to `os.Stdout`.  If you would prefer, output can be written to an alternate `io.Writer`.  The following example shows how you would set up a logger to write to a bytes buffer.
```
buffer := &bytes.Buffer{}
logger := log.Logger{Writer: buffer}
```

## Context
Each logger has a series of key-value pairs that make up it's context.  If values exist in the logger's context, these values will be present in the log output.  A key-value pair can be added to a logger's context as follows:
```
logger := log.Logger{}.With("key", "value")
```

## Value Generators
Unlike most values that are added to a logger's context, a ValueGenerator is a function that will be evaluated at the time the log entry is formatted.  This is useful in scenarios where you may want to provide some runtime context to your logs, such as line numbers or timestamps.

## Inspiration
This package was largely based on [Go kit's log package](https://github.com/go-kit/kit/tree/master/log).  Through reading the source, as well as a number of conversations they pointed out in their GitHub issues, I feel I was able to get a much better grasp on the topic of structured logging in general.  I'd like to extend my thanks to that team for having made such a great resource!
