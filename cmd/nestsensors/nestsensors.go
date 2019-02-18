package main

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	Quiet    bool
	Verbose  bool
	User     string
	Password string
)

var rootCmd = &cobra.Command{
	Use:   "nestsensors",
	Short: "Program for getting at Nest sensor data not exposed in the offial API",
	Long:  "Program for getting at Nest sensor data not exposed in the offial API.  Works via scraping the home.nest.com frontend JSON.",
	PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
		initLogging(Quiet, Verbose)

		if !cmd.Flags().Changed("user") {
			return errors.New("no user specified")
		}
		if !cmd.Flags().Changed("password") {
			return errors.New("no password specified")
		}
		return nil
	},
}

// TODO: Maybe rename to nesthome or something else better.

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Quiet, "quiet", "q", false, "Activate quiet log output")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Activate verbose log output")

	rootCmd.PersistentFlags().StringVarP(&User, "user", "u", "", "Nest username to authenticate with")
	rootCmd.PersistentFlags().StringVarP(&Password, "password", "p", "", "Nest account password")

	rootCmd.AddCommand(
		sensorsCmd,
		loginCmd,
	)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initLogging(quiet bool, verbose bool) {
	if quiet && verbose {
		log.Fatal("Invalid logging flag combination: cannot turn on both quiet and verbose modes")
	}

	level := log.InfoLevel
	if verbose {
		level = log.DebugLevel
	}
	if quiet {
		level = log.ErrorLevel
	}
	log.SetLevel(level)
}
