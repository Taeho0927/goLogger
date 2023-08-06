package logger

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

func Warn(l LogContent) {
	logFile, err := Init()
	if err != nil {
		panic(err)
	}

	defer func() {
		logFile.Close()
		p := recover()
		if p != nil {
			fmt.Println("LOG FILE ERROR", p)
		}
		return
	}()

	white := color.New(color.FgWhite)
	bold := white.Add(color.BgYellow)

	level := "[WARN]"
	message := fmt.Sprintf("{project:%s, handler:%s, function:%s, message:%s, line:%d, time:%s}", l.Project, l.Handler, l.Function, l.Message, l.CodeLine, l.Time)

	log.SetFlags(log.Ldate | log.Ltime)
	log.SetPrefix(fmt.Sprintf("%s", bold.Sprint(level)))

	logFile.WriteString(fmt.Sprintf("%s %s\n", level, message))
	log.Println(message)
}
