package cmd

import (
	"github.com/abiosoft/colima/config"
	"github.com/abiosoft/colima/runtime/container"
	"github.com/abiosoft/colima/runtime/container/docker"
	"github.com/spf13/cobra"
	"strings"
)

// startCmd represents the start command
// TODO detect the default container runtime
// TODO replace $HOME env var.
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Colima",
	Long: `Start Colima with the specified container runtime (and kubernetes if --with-kubernetes is passed).

Kubernetes requires at least 2 CPUs and 2.3GiB memory.

For verbose output, tail the log file "$HOME/Library/Caches/colima/out.log".
`,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(app.Start())
	},
}

const (
	defaultCPU     = 2
	defaultMemory  = 4
	defaultDisk    = 60
	defaultSSHPort = 41122
)

var appConfig config.Config

func init() {
	runtimes := strings.Join(container.Names(), ", ")

	rootCmd.AddCommand(startCmd)
	startCmd.Flags().BoolVarP(&appConfig.Kubernetes, "with-kubernetes", "k", false, "start VM with Kubernetes")
	startCmd.Flags().StringVarP(&appConfig.Runtime, "runtime", "r", docker.Name, "container runtime, one of ["+runtimes+"]")
	startCmd.Flags().IntVarP(&appConfig.VM.CPU, "cpu", "c", defaultCPU, "number of CPUs")
	startCmd.Flags().IntVarP(&appConfig.VM.Memory, "memory", "m", defaultMemory, "memory in GiB")
	startCmd.Flags().IntVarP(&appConfig.VM.Disk, "disk", "d", defaultDisk, "disk size in GiB")
	startCmd.Flags().IPSliceVarP(&appConfig.VM.DNS, "dns", "n", nil, "DNS servers for the VM")

	// internal
	startCmd.Flags().IntVar(&appConfig.VM.SSHPort, "ssh-port", defaultSSHPort, "SSH port for the VM")
	startCmd.Flags().MarkHidden("ssh-port")

	// not sure of the usefulness of env vars for now considering that interactions will be with the containers, not the VM.
	// leaving it undocumented until there is a need.
	startCmd.Flags().StringToStringVarP(&appConfig.VM.Env, "env", "e", nil, "environment variables for the VM")
	startCmd.Flags().MarkHidden("env")
}
