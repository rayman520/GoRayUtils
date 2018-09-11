package rayUtils

import (
	"container/list"
	"fmt"
	"os"
	"time"
)

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
	fileOutput *os.File
}

func NewLogger() Logger {
	logger := Logger{}
	logger.loglvl = LOGINFO
	logger.log = list.New()

	return logger
}

func (l Logger) LogStore(str string) {
	l.log.PushBack(str)
}

func (l Logger) LogRemoveLast() {
	if l.log.Back() != nil {
		l.log.Remove(l.log.Back())
	}
}

func (l Logger) LogPrint() {
	for e := l.log.Front(); e != nil; e = e.Next() {
		l.Log(LOGRAW, e.Value.(string))
	}
	l.Log(LOGRAW, "\n\n")
}

func (l Logger) LogFlush() {
	l.Log(LOGINFO, "Log flush called\n\n")
	l.LogPrint()
	l.log = list.New()
}

func (l Logger) Log(loglvl uint8, str string, args ...interface{}) {

	var loglvlStr string
	var strout string
	var lvl uint8

	//TODO Optimize to map
	lvl = loglvl % 6
	if lvl < l.loglvl {
		return
	}
	switch lvl {
	case 0:
		loglvlStr = "\033[38;5;8mDebug"
	case 1:
		loglvlStr = "\033[38;5;2mInfo"
	case 2:
		loglvlStr = "\033[38;5;148mWarn"
	case 3:
		loglvlStr = "\033[38;5;166mError"
	case 4:
		loglvlStr = "\033[38;5;196mFatal"
	}

	if lvl < 5 {
		strout = fmt.Sprintf("HexTranslate - %s - [%s\033[0m]:\t%s\n", time.Now().Format("2006-01-02 15:04:05"), loglvlStr, fmt.Sprintf(str, args...))
	} else {
		strout = fmt.Sprintf("%s", fmt.Sprintf(str, args...))
	}

	if loglvl/6 > 0 {
		l.LogStore(strout)
	} else {
		fmt.Printf("\033[2K\r")
		fmt.Print(strout)
		if l.fileOutput != nil {
			fmt.Fprint(l.fileOutput, strout)
		}
	}
}
