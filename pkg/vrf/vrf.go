package vrf

import (
	"fmt"
	"os"

	"airnity.com/router-sidecar/pkg/command"
	"airnity.com/router-sidecar/pkg/config"
	log "github.com/sirupsen/logrus"
)

type client struct {
	cfg    config.Manager
	logger *log.Logger
}

func (c *client) Synchronize() error {
	vrfs := c.cfg.GetConfig().VRFs
	for _, vrf := range vrfs {
		err := c.createVRF(&vrf)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *client) createVRF(vrf *config.VRF) error {
	//create table
	err := appendToFile("/etc/iproute2/rt_tables", fmt.Sprintf("%d %s", vrf.TableID, vrf.Name))
	if err != nil {
		return err
	}
	//add vrf with table
	addVrfCmd := fmt.Sprintf("ip link add dev %s type vrf table %d", vrf.Name, vrf.TableID)
	_, err = c.execVRFCmd(vrf, addVrfCmd)
	if err != nil {
		return err
	}
	//up vrf
	upVrfCmd := fmt.Sprintf("ip link set dev %s up", vrf.Name)
	_, err = c.execVRFCmd(vrf, upVrfCmd)
	if err != nil {
		return err
	}
	return nil
}

func appendToFile(filename, content string) error {
	// Open file with O_APPEND mode (or create if doesn't exist)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Append content followed by a newline
	if _, err := file.WriteString(content + "\n"); err != nil {
		return err
	}

	return nil
}

func (c *client) execVRFCmd(vrf *config.VRF, cmd string) (string, error) {
	fields := log.Fields{
		"name":    vrf.Name,
		"tableID": vrf.TableID,
	}
	out, err := command.ExecCmd(cmd, fields, c.logger)
	if err != nil {
		return out, err
	}
	return out, nil
}
