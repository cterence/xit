package tailout

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/cterence/tailout/tailout/config"
	"github.com/cterence/tailout/tailout/tailscale"
)

func (app *App) Status() error {
	c := tailscale.NewClient(&app.Config.Tailscale)

	nodes, err := c.GetActiveNodes()
	if err != nil {
		return err
	}

	out, err := exec.Command("tailscale", "debug", "prefs").CombinedOutput()

	if err != nil {
		return fmt.Errorf("failed to get tailscale preferences: %w", err)
	}

	var status config.TailscaleStatus
	var currentNode config.Node

	json.Unmarshal(out, &status)

	if status.ExitNodeID != "" {
		currentNode, err = c.GetNode(status.ExitNodeID)
		if err != nil {
			return fmt.Errorf("failed to get node: %w", err)
		}
	}

	if len(nodes) == 0 {
		fmt.Println("No active node created by tailout found.")
	} else {
		fmt.Println("Active nodes created by tailout:")
		for _, node := range nodes {
			if currentNode.Hostname == node.Hostname {
				fmt.Println("-", node.Hostname, "[Connected]")
			} else {
				fmt.Println("-", node.Hostname)
			}
		}
	}

	// Query for the public IP address of this Node
	out, err = exec.Command("curl", "ifconfig.me").Output()

	if err != nil {
		return fmt.Errorf("failed to get public IP: %w", err)
	}

	fmt.Println("Public IP:", string(out))
	return nil
}
