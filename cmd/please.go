package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/logger"
	"github.com/state-alchemists/zaruba/runner"
)

var pleaseEnv []string
var pleaseKwargs []string
var pleaseFile string

// pleaseCmd represents the please command
var pleaseCmd = &cobra.Command{
	Use:   "please",
	Short: "Ask Zaruba to do something for you",
	Long:  "💀 Ask Zaruba to do something for you",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.NewConfig(pleaseFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		// process globalEnv
		for _, env := range pleaseEnv {
			conf.AddGlobalEnv(env)
		}
		// process kwargs from flag
		for _, kwarg := range pleaseKwargs {
			if err = conf.AddKwargs(kwarg); err != nil {
				fmt.Println(err)
				return
			}
		}
		//  distinguis taskNames and additional kwargs
		taskNames := []string{}
		for _, arg := range args {
			if strings.Contains(arg, "=") {
				conf.AddKwargs(arg)
				continue
			}
			taskNames = append(taskNames, arg)
		}
		// init
		if err = conf.Init(); err != nil {
			fmt.Println(err)
			return
		}
		// show list of available tasks if no task provided
		if len(taskNames) == 0 {
			taskIndentation := strings.Repeat(" ", 6)
			taskFieldIndentation := taskIndentation + strings.Repeat(" ", 5)
			d := logger.NewDecoration()
			publishedTask := conf.GetPublishedTask()
			logger.Printf("%sPlease what?%s\n", d.Bold, d.Normal)
			logger.Printf("Here are some possible tasks you can execute:\n")
			for _, taskName := range conf.SortedTaskNames {
				if task, exist := publishedTask[taskName]; exist {
					fmt.Printf("%s%s %szaruba please %s%s%s\n", taskIndentation, task.Icon, d.Important, d.Bold, task.Name, d.Normal)
					fmt.Printf("%s%s%sDECLARED ON:%s%s %s%s\n", taskFieldIndentation, d.Important, d.Dim, d.Normal, d.Dim, task.FileLocation, d.Normal)
					showTaskDescription(task, taskFieldIndentation)
				}
			}
			return
		}
		// run
		r := runner.NewRunner(conf, taskNames)
		if err := r.Run(); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pleaseCmd)
	// get current working directory
	dir, err := os.Getwd()
	if err != nil {
		dir = "."
	}
	// define defaultPleaseFile
	defaultPleaseFile := filepath.Join(dir, "main.zaruba.yaml")
	if _, err := os.Stat(defaultPleaseFile); os.IsNotExist(err) {
		defaultPleaseFile = "${ZARUBA_HOME}/scripts/core.zaruba.yaml"
	}
	// define defaultPleaseKwargs
	defaultPleaseKwargs := []string{}
	defaultKwargsFile := filepath.Join(dir, "default.kwargs.yaml")
	if _, err := os.Stat(defaultKwargsFile); !os.IsNotExist(err) {
		defaultPleaseKwargs = append(defaultPleaseKwargs, defaultKwargsFile)
	}
	// register flags
	pleaseCmd.Flags().StringVarP(&pleaseFile, "file", "f", defaultPleaseFile, "custom file")
	pleaseCmd.Flags().StringArrayVarP(&pleaseEnv, "environment", "e", []string{}, "environment file or pairs (e.g: '-e environment.env' or '-e key=val')")
	pleaseCmd.Flags().StringArrayVarP(&pleaseKwargs, "kwargs", "k", defaultPleaseKwargs, "yaml file or pairs (e.g: '-k value.yaml' or '-k key=val')")
}

func showTaskDescription(task *config.Task, fieldIndentation string) {
	if task.Description != "" {
		d := logger.NewDecoration()
		description := strings.TrimSpace(task.Description)
		rows := strings.Split(description, "\n")
		for index, row := range rows {
			if index == 0 {
				row = fmt.Sprintf("%sDESCRIPTION:%s %s%s", d.Important, d.Normal, d.Dim, row)
			}
			fmt.Printf("%s%s%s%s\n", fieldIndentation, d.Dim, row, d.Normal)
		}
	}
}
