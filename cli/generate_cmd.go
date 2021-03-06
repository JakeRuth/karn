package cli

import (
	"io"
	"os"

	parse "github.com/GrantSeltzer/karn/parse"
	"github.com/spf13/cobra"
)

type GenerateOptions struct {
	declarationDirectory string
	seccomp              bool
	apparmor             bool
	outputDirectory      string
}

func NewGenerateCmd(out io.Writer) *cobra.Command {

	genOpts := GenerateOptions{}

	generateCmd := &cobra.Command{
		Use:   "generate [<DECLARATION>,...]",
		Short: "generate seccomp and apparmor profiles from a karn profile",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: verify arguments
			return genOpts.Run(out, args)
		},
	}

	homedir := os.Getenv("HOME")

	g := generateCmd.PersistentFlags()
	g.StringVarP(&genOpts.declarationDirectory, "declarations", "d", homedir+"/.karn/declarations", "directory of declaration definitions")
	g.BoolVar(&genOpts.seccomp, "seccomp", false, "output seccomp profile")
	g.BoolVar(&genOpts.apparmor, "apparmor", false, "output apparmor profile")

	return generateCmd
}

func (genOpts *GenerateOptions) Run(out io.Writer, args []string) error {

	if genOpts.seccomp {
		err := parse.WriteSeccompProfile(out, args, genOpts.declarationDirectory)
		if err != nil {
			return err
		}
	}

	if genOpts.apparmor {
		err := parse.WriteAppArmorProfile(out, args, genOpts.declarationDirectory)
		if err != nil {
			return err
		}
	}

	return nil
}
