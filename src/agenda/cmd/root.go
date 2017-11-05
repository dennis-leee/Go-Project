package cmd

import (
		"fmt"
		"io/ioutil"
		"log"
		"os"
	homedir 
		"github.com/spf13/cobra"
		"github.com/spf13/viper"
		"github.com/mitchellh/go-homedir"
)

var RootCmd = &cobra.Command{
	Use:   "Agenda",
	Short: "A CLI meeting manager",
	Long: `Agenda supports different operation on meetings including register, create meeting, query and so on.
			It's a cooperation homework assignment for service computing.`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
