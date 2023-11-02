package cli

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	config       string
	bucket       string
	localfile    string
	cloudfile    string
	version_view bool
)

var RootCmd = &cobra.Command{
	Use:       "oss-admin",
	Short:     "Management for cloud oss",
	Long:      "This is a tool used for various management operations of cloud OSS",
	Example:   "oss-admin [cmd] <options>",
	ValidArgs: []string{"pull", "ls", "push", "du", "bucket"},
	Version:   "1.0.1",
	RunE: func(cmd *cobra.Command, args []string) error {
		if version_view {
			fmt.Println("Version: oss-admin-1.0.1")
			return nil
		}
		return errors.New("no flag to found")
	},
}
