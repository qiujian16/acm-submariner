package main

import (
	goflag "flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	submarinercmd "github.com/qiujian16/acm-submariner/pkg/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	utilflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/logs"

	"github.com/qiujian16/acm-submariner/pkg/version"
)

// The registration binary contains both the hub-side controllers for the
// registration API and the spoke agent.

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	pflag.CommandLine.SetNormalizeFunc(utilflag.WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)

	logs.InitLogs()
	defer logs.FlushLogs()

	command := newSubmarinerControllerCommand()
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func newSubmarinerControllerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "registration",
		Short: "Spoke Cluster Registration",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}

	if v := version.Get().String(); len(v) == 0 {
		cmd.Version = "<unknown>"
	} else {
		cmd.Version = v
	}

	cmd.AddCommand(submarinercmd.NewController())

	return cmd
}
