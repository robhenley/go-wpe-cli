package ssh_keys

import (
	"github.com/spf13/cobra"
)

// sshKeysCmd represents the sshKeysCmd command
var SSHKeysCmd = &cobra.Command{
	Use:   "ssh-keys",
	Short: "Command for operations on SSH keys",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	SSHKeysCmd.AddCommand(SSHKeysListCmd)
	SSHKeysCmd.AddCommand(SSHKeysCreateCmd)
	SSHKeysCmd.AddCommand(SSHKeysDeleteCmd)
}
