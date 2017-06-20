package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"golang.org/x/sync/errgroup"

	victorops "github.com/caarlos0/go-victorops"
	"github.com/spf13/cobra"
)

const popup = "https://portal.victorops.com/client/%v/popoutIncident?incidentName=%v"

var cfgFile string
var clientName string
var apiID string
var apiKey string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "victorops-notifier",
	Short: "Listen to alerts and notify you",
	Run: func(cmd *cobra.Command, args []string) {
		var cli = victorops.New("", apiID, apiKey)
		incidents, err := cli.Incidents()
		if err != nil {
			log.Fatalln(err)
		}
		var g errgroup.Group
		for _, incident := range incidents {
			if incident.CurrentPhase == "UNACKED" {
				incident := incident
				g.Go(func() error {
					return notify(incident)
				})
			}
		}
		g.Wait()
	},
}

func notify(incident victorops.Incident) error {
	return exec.Command(
		"terminal-notifier",
		"-title", "VictorOps",
		"-message", incident.EntityState,
		"-subtitle", incident.EntityDisplayName,
		"-open", fmt.Sprintf(popup, clientName, incident.IncidentNumber),
		"-appIcon", "https://help.victorops.com/wp-content/uploads/2017/03/victorops_mark_color.png",
		"-timeout", "10",
	).Run()
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&clientName, "client", "", "VictorOps client name (look at the URL)")
	RootCmd.PersistentFlags().StringVar(&apiID, "id", "", "VictorOps API ID")
	RootCmd.PersistentFlags().StringVar(&apiKey, "key", "", "VictorOps API Key")
}
