package rayUtils

import (
	"container/list"
	"os"
)

// globals

var rayUtilsGlobals = RayUtilsGlobals{
	Logger: NewLogger("logger"),
}

type RayUtilsGlobals struct {
	BooleanTrueTranslator []string
	BracketOppositesMap   map[byte]byte
	Logger                Logger
}

// logger

const (
	LOGDEBUG uint8 = iota
	LOGINFO
	LOGWARN
	LOGERROR
	LOGFATAL
	LOGRAW
)

type Logger struct {
	loglvl     uint8
	log        *list.List
	name       string
	fileOutput *os.File
}
