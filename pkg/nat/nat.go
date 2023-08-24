package nat

import (
	"fmt"

	"airnity.com/network-manager/pkg/command"
	"airnity.com/network-manager/pkg/config"
	log "github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

type client struct {
	cfg    config.Manager
	logger *log.Logger
}

func (c *client) Synchronize() error {
	sourceNat := c.cfg.GetConfig().NatRules.SourceNat
	for _, rule := range sourceNat {
		err := c.createSourceNatRule(&rule)
		if err != nil {
			return err
		}
	}

	destNat := c.cfg.GetConfig().NatRules.DestNat
	funk.ForEach(destNat, func(rule config.NatRule) {
		c.createDestNatRule(&rule)
	})

	return nil
}

func (c *client) createSourceNatRule(rule *config.NatRule) error {
	optionalArgs := generateOptionalArgs(rule)
	addNatCmd := fmt.Sprintf("iptables -t nat -A POSTROUTING -d %s -o %s %s-j SNAT --to %s", rule.Destination, rule.Interface, optionalArgs, rule.TranslatedIP)
	_, err := c.execNatCmd(rule, addNatCmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) createDestNatRule(rule *config.NatRule) error {
	optionalArgs := generateOptionalArgs(rule)
	addNatCmd := fmt.Sprintf("iptables -t nat -A PREROUTING -d %s -i %s %s-j DNAT --to-destination %s", rule.Destination, rule.Interface, optionalArgs, rule.TranslatedIP)
	_, err := c.execNatCmd(rule, addNatCmd)
	if err != nil {
		return err
	}
	return nil
}

func generateOptionalArgs(rule *config.NatRule) string {
	var optionalArgs string
	if rule.Source != "" {
		optionalArgs = optionalArgs + fmt.Sprintf("-s %s ", rule.Source)
	}
	if rule.Proto != "" {
		optionalArgs = optionalArgs + fmt.Sprintf("-p %s ", rule.Proto)
	}
	if rule.Port != 0 {
		optionalArgs = optionalArgs + fmt.Sprintf("--dport %d ", rule.Port)
	}
	return optionalArgs
}

func (c *client) execNatCmd(rule *config.NatRule, cmd string) (string, error) {
	fields := log.Fields{
		"type":        "destNat",
		"source":      rule.Source,
		"dest":        rule.Destination,
		"translateIP": rule.TranslatedIP,
		"interface":   rule.Interface,
		"port":        rule.Port,
		"proto":       rule.Proto,
	}
	out, err := command.ExecCmd(cmd, fields, c.logger)
	if err != nil {
		return out, err
	}
	return out, nil
}
