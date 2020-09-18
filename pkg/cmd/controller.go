package hub

import (
	"github.com/spf13/cobra"

	"github.com/openshift/library-go/pkg/controller/controllercmd"

	"github.com/qiujian16/acm-submariner/pkg/hub"
	"github.com/qiujian16/acm-submariner/pkg/version"
)

func NewController() *cobra.Command {
	cmd := controllercmd.
		NewControllerCommandConfig("submariner-controller", version.Get(), hub.RunControllerManager).
		NewCommand()
	cmd.Use = "controller"
	cmd.Short = "Start the ACM Submariner Controller"

	return cmd
}
