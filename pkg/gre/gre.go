package gre

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"

	"airnity.com/router-sidecar/pkg/command"
	"airnity.com/router-sidecar/pkg/config"
	log "github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

type client struct {
	cfg    config.Manager
	logger *log.Logger
}

func (c *client) Synchronize() error {
	tunnels := c.cfg.GetConfig().Tunnels
	funk.ForEach(tunnels, func(tunnel config.Tunnel) {
		fields := log.Fields{
			"tunnel": tunnel.Name,
			"remote": tunnel.Remote,
			"local":  tunnel.Local,
			"state":  tunnel.State,
		}
		if tunnel.State == "absent" {
			if c.tunnelExists(&tunnel) {
				c.logger.WithFields(fields).Info("Deleting tunnel")
				c.deleteTunnel(&tunnel)
			} else {
				c.logger.WithFields(fields).Info("Tunnel already deleted")
			}
		}
		if tunnel.State == "present" {
			if c.tunnelExists(&tunnel) {
				tun, _ := c.getTunnelInfo(tunnel.Name)
				if reflect.DeepEqual(tun, &tunnel) {
					c.logger.WithFields(fields).Info("Tunnel already created, and up to date")
				} else {
					c.logger.WithFields(fields).Info("Tunnel already created, but outdated")
					c.logger.WithFields(fields).Info("Deleting it")
					c.deleteTunnel(tun)
					c.logger.WithFields(fields).Info("Recreating it")
					c.createTunnel(&tunnel)
				}
			} else {
				c.logger.WithFields(fields).Info("Creating tunnel")
				c.createTunnel(&tunnel)
			}
		}
		c.logger.WithFields(fields).Info("Done")
	})
	return nil
}

func (c *client) createTunnel(tunnel *config.Tunnel) error {
	addTunnelCmd := fmt.Sprintf("ip tunnel add %s mode gre remote %s local %s ttl 255", tunnel.Name, tunnel.Remote, tunnel.Local)
	_, err := c.execTunnelCmd(tunnel, addTunnelCmd)
	if err != nil {
		return err
	}
	linkSetCmd := fmt.Sprintf("ip link set %s up", tunnel.Name)
	_, err = c.execTunnelCmd(tunnel, linkSetCmd)
	if err != nil {
		return err
	}
	addAddrCmd := fmt.Sprintf("ip addr add %s dev %s", tunnel.Addr, tunnel.Name)
	_, err = c.execTunnelCmd(tunnel, addAddrCmd)
	if err != nil {
		return err
	}
	rpFilterCmd := fmt.Sprintf("sysctl -w net.ipv4.conf.%s.rp_filter=0", tunnel.Name)
	_, err = c.execTunnelCmd(tunnel, rpFilterCmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) deleteTunnel(tunnel *config.Tunnel) error {
	addAddrCmd := fmt.Sprintf("ip addr del %s dev %s", tunnel.Addr, tunnel.Name)
	_, err := c.execTunnelCmd(tunnel, addAddrCmd)
	if err != nil {
		return err
	}
	linkSetCmd := fmt.Sprintf("ip link set %s down", tunnel.Name)
	_, err = c.execTunnelCmd(tunnel, linkSetCmd)
	if err != nil {
		return err
	}
	addTunnelCmd := fmt.Sprintf("ip tunnel del %s", tunnel.Name)
	_, err = c.execTunnelCmd(tunnel, addTunnelCmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) tunnelExists(tunnel *config.Tunnel) bool {
	addAddrCmd := fmt.Sprintf("ip tunnel show %s", tunnel.Name)
	out, err := c.execTunnelCmd(tunnel, addAddrCmd)
	if err == nil {
		if out == "" {
			return false
		}
		return true
	}
	return false
}

func (c *client) getTunnelInfo(name string) (*config.Tunnel, error) {
	rmt, lcl, err := c.getTunnelAddresses(name)
	if err != nil {
		return nil, err
	}
	addr, err := c.getDevAddress(name)
	if err != nil {
		return nil, err
	}

	tunnel := &config.Tunnel{
		Name:   name,
		Remote: rmt,
		Local:  lcl,
		Addr:   addr,
		State:  "present",
	}

	return tunnel, nil
}

func (c *client) execTunnelCmd(tunnel *config.Tunnel, cmd string) (string, error) {
	fields := log.Fields{
		"tunnel": tunnel.Name,
		"remote": tunnel.Remote,
		"local":  tunnel.Local,
		"state":  tunnel.State,
	}
	out, err := command.ExecCmd(cmd, fields, c.logger)
	if err != nil {
		return out, err
	}
	return out, nil
}

func (c *client) getTunnelAddresses(name string) (string, string, error) {
	remote := ""
	local := ""
	cmd := fmt.Sprintf("ip tunnel show %s", name)
	fields := log.Fields{
		"search": name,
	}
	out, err := command.ExecCmd(cmd, fields, c.logger)
	if err != nil {
		return remote, local, err
	}
	remoteRe := regexp.MustCompile(`remote ([0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3})`)
	localRe := regexp.MustCompile(`local ([0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3})`)
	remoteMatch := remoteRe.FindStringSubmatch(out)
	localMatch := localRe.FindStringSubmatch(out)
	if len(remoteMatch) == 2 {
		remote = remoteMatch[1]
	} else {
		return remote, local, errors.New("could not find remote address")
	}
	if len(localMatch) == 2 {
		local = localMatch[1]
	} else {
		return remote, local, errors.New("could not find local address")
	}
	return remote, local, nil
}

func (c *client) getDevAddress(name string) (string, error) {
	addr := ""
	cmd := fmt.Sprintf("ip a show %s", name)
	fields := log.Fields{
		"search": name,
	}
	out, err := command.ExecCmd(cmd, fields, c.logger)
	if err != nil {
		return addr, err
	}
	re := regexp.MustCompile(`inet ([0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}/[0-9]{2})`)
	match := re.FindStringSubmatch(out)
	if len(match) == 2 {
		addr = match[1]
	} else {
		return addr, errors.New("could not find device address")
	}
	return addr, nil
}
