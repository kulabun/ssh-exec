package commands

import "github.com/spf13/cobra"

var (
  rootCmd = &cobra.Command{
    Use: "ssh-exec",
    Short: "ssh-exec is a tool for command execution over ssh",
    Long: "ssh-exec is a tool for command execution over ssh",
    SilenceUsage: true,
    SilenceErrors: false,
  }
)

type flags struct {
  profile string
}

func Execute() error {
  return rootCmd.Execute()
}

func init() {
}
