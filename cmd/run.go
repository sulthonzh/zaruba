package cmd

import (
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/logger"
	"github.com/state-alchemists/zaruba/modules/runner"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run [project-dir]",
	Short: "Run project",
	Long:  `Zaruba will run all defined services`,
	Run: func(cmd *cobra.Command, args []string) {
		// get projectDir
		projectDir, err := filepath.Abs(".")
		if err != nil {
			logger.Fatal(err)
		}
		p, err := config.NewProjectConfig(projectDir)
		if err != nil {
			logger.Fatal(err)
		}
		// invoke action
		stopChan := make(chan bool)
		errChan := make(chan error)
		executedChan := make(chan bool)
		go runner.Run(projectDir, p, args, stopChan, executedChan, errChan)
		<-executedChan
		// listen to kill signal
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			stopChan <- true
		}()
		// wait for errChan
		err = <-errChan
		if err != nil {
			logger.Fatal(err)
		}
	},
}
