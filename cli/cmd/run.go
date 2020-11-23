package cmd

import (
	"errors"
	"log"

	"github.com/releaseband/map-switch/analysis"
	"github.com/releaseband/map-switch/services"

	"github.com/spf13/cobra"
)

const (
	cmdMapBySlice = "map_by_slice" //todo: в теги перевести
)

var (
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "generate map to switch case",
		RunE:  run,
	}
	runFlags = struct {
		Path    string
		Command string
	}{}
)

func init() {
	runCmd.PersistentFlags().StringVarP(&runFlags.Path, "filepath", "f",
		"", "path to file where i can find map")

	runCmd.PersistentFlags().StringVarP(&runFlags.Command, "command", "c", cmdMapBySlice,
		"generate command")
}

func validateFlags() error {
	if len(runFlags.Path) == 0 {
		return errors.New("'path' flag doesn't be empty")
	}

	return nil
}

func run(cmd *cobra.Command, args []string) error {
	err := validateFlags()
	if err != nil {
		return err
	}

	recorder := services.NewRecorder()
	switch runFlags.Command {
	case cmdMapBySlice:
		mp := analysis.NewMapParams(
			runFlags.Path,
		)

		err = analysis.GenerateMapByString(recorder, mp)
	default:
		err = errors.New("command not supported")
	}

	return err
}

func Execute() {
	if err := runCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
