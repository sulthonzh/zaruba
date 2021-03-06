package cmd

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/util"
)

var getEnvKeysByLocationFormat string
var getEnvKeysByLocationCmd = &cobra.Command{
	Use:   "getEnvKeysByLocation <serviceLocation>",
	Short: "Get environment keys by location",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 1 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for getEnvKeysByLocation"))
		}
		envMap, err := util.GetEnvByLocation(args[0])
		if err != nil {
			showErrorAndExit(logger, decoration, err)
		}
		keys := []string{}
		for key, _ := range envMap {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		// output format: json
		if strings.ToLower(getEnvByLocationFormat) == "json" {
			keysJsonB, err := json.Marshal(keys)
			if err != nil {
				showErrorAndExit(logger, decoration, err)
			}
			fmt.Println(string(keysJsonB))
			return
		}
		// output format: not specified
		for _, key := range keys {
			fmt.Println(key)
		}
	},
}

func init() {
	rootCmd.AddCommand(getEnvKeysByLocationCmd)
	getEnvKeysByLocationCmd.Flags().StringVarP(&getEnvKeysByLocationFormat, "format", "f", "", "output format")
}
