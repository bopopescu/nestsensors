package main

import (
	"github.com/spf13/cobra"

	"jaytaylor.com/nestsensors"
)

var loginCmd = &cobra.Command{
	Use: "login",
	RunE: func(cmd *cobra.Command, _ []string) error {
		s := nestsensors.New(User, Password)
		if err := s.Login(); err != nil {
			return err
		}
		return nil
	},
}
