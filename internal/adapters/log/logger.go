package logz

import "log"

type Logger struct{}

func New() *Logger { return &Logger{} }
func (l *Logger) Infof(f string, v ...any)  { log.Printf("[INFO] "+f, v...) }
func (l *Logger) Errorf(f string, v ...any) { log.Printf("[ERR]  "+f, v...) }
