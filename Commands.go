package rayUtils

import (
	"bufio"
	"container/list"
	"os/exec"
)

func ExecPrintCmd(cmd *exec.Cmd) (*list.List, *list.List) {
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	stdoutLog := list.New()
	stderrLog := list.New()
	cmd.Start()
	go func() {
		logger := NewLogger(cmd.Args[0])
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			m := scanner.Text()
			logger.Log(LOGINFO, m)
			stdoutLog.PushBack(m)
		}
	}()
	go func() {
		logger := NewLogger(cmd.Args[0])
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			m := scanner.Text()
			logger.Log(LOGERROR, m)
			stderrLog.PushBack(m)
		}
	}()
	cmd.Wait()
	return stdoutLog, stderrLog
}

func ExecCmd(cmd *exec.Cmd) {
	cmd.Start()
	cmd.Wait()
}
