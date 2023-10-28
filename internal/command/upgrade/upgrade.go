package upgrade

import (
	"fmt"
	"github.com/brix-go/brix/config"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var CmdUpgrade = &cobra.Command{
	Use:     "upgrade",
	Short:   "Upgrade the brix command.",
	Long:    "Upgrade the brix command.",
	Example: "brix upgrade",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("go install %s\n", config.BrixCmd)
		cmd := exec.Command("go", "install", config.BrixCmd)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("go install %s error\n", err)
		}
		fmt.Printf("\n🎉 brix upgrade successfully!\n\n")
	},
}
