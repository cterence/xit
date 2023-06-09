package tailout

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/cterence/tailout/internal"
	"github.com/cterence/tailout/tailout/config"
)

func (app *App) Disconnect() error {
	nonInteractive := app.Config.NonInteractive

	var status config.TailscaleStatus

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
