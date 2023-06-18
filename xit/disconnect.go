package xit

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/cterence/xit/internal"
	"github.com/spf13/viper"
)

func (app *App) Disconnect() error {
	nonInteractive := viper.GetBool("non_interactive")

	var status internal.TailscaleStatus

	out, err := exec.Command("tailscale", "debug", "prefs").CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to get tailscale preferences: %w", err)
	}

	json.Unmarshal(out, &status)

	if status.ExitNodeID == "" {
		return fmt.Errorf("not connected to an exit node")
	}

	err = internal.RunTailscaleUpCommand("tailscale up --exit-node=", nonInteractive)
	if err != nil {
		return fmt.Errorf("failed to run tailscale up command: %w", err)
	}

	fmt.Println("Disconnected.")
	return nil
}