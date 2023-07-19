package command

import (
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

func ExecCmd(cmd string, fields log.Fields, logger *log.Logger) (string, error) {
	logger.WithFields(fields).Debug("Exec: ", cmd)
	exec := exec.Command(strings.Fields(cmd)[0], strings.Fields(cmd)[1:]...)
	out, err := exec.Output()
	if err != nil {
		log.Panic(err)
		return string(out), err
	}
	return string(out), nil
}
