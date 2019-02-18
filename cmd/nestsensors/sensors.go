package main

import (
	"github.com/spf13/cobra"

	"jaytaylor.com/nestsensors"
)

var sensorsCmd = &cobra.Command{
	Use: "sensors",
	RunE: func(cmd *cobra.Command, _ []string) error {
		s := nestsensors.New(User, Password)
		if err := s.Login(); err != nil {
			return err
		}
		if err := s.Raw(); err != nil {
			return err
		}
		return nil
	},
}
