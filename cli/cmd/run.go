package cmd

import (
	"errors"

	"github.com/releaseband/map-switch/generator"
	"github.com/spf13/cobra"
)

var (
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "generate map to switch",
		RunE:  run,
	}
	Path string
)

func init() {
	runCmd.PersistentFlags().StringVarP(&Path, "path", "f", "", "src file path")
}

func validateFlags() error {
	if len(Path) == 0 {
		return errors.New("'path' flag doesn't be empty")
	}

	return nil
}

func run(_ *cobra.Command, _ []string) error {
	if err := validateFlags(); err != nil {
		return err
	}

	return generator.Run(Path)
}

func Execute() error {
	return runCmd.Execute()
}
