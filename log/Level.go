package log

import (
	"strconv"
	"strings"
)

type Level int

const (
	// LevelDefault means the log entry has no assigned severity level.
	LevelDefault = Level(0)
	// LevelDebug means debug or trace information.
	LevelDebug = Level(100)
	// LevelInfo means routine information, such as ongoing status or performance.
	LevelInfo = Level(200)
	// LevelNotice means normal but significant events, such as start up, shut down, or configuration.
	LevelNotice = Level(300)
	// LevelWarning means events that might cause problems.
	LevelWarning = Level(400)
	// LevelError means events that are likely to cause problems.
	LevelError = Level(500)
	// LevelCritical means events that cause more severe problems or brief outages.
	LevelCritical = Level(600)
	// LevelAlert means a person must take an action immediately.
	LevelAlert = Level(700)
	// LevelEmergency means one or more systems are unusable.
	LevelEmergency = Level(800)
)

var levelName = map[Level]string{
	LevelDefault:   "Default",
	LevelDebug:     "Debug",
	LevelInfo:      "Info",
	LevelNotice:    "Notice",
	LevelWarning:   "Warning",
	LevelError:     "Error",
	LevelCritical:  "Critical",
	LevelAlert:     "Alert",
	LevelEmergency: "Emergency",
}

var nameToLevel = map[string]Level{
	"default":   LevelDefault,
	"debug":     LevelDebug,
	"info":      LevelInfo,
	"notice":    LevelNotice,
	"warning":   LevelWarning,
	"error":     LevelError,
	"critical":  LevelCritical,
	"alert":     LevelAlert,
	"emergency": LevelEmergency,
}

// String converts a severity level to a string.
func (level Level) String() string {
	// same as proto.EnumName
	s, ok := levelName[level]
	if ok {
		return s
	}

	return strconv.Itoa(int(level))
}

func ParseLevel(s string) Level {
	sl := strings.ToLower(s)
	lvl, ok := nameToLevel[sl]
	if ok {
		return lvl
	}
	return LevelDefault
}
