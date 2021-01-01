package repo

import (
	"github.com/MakeNowJust/heredoc"
	repoCloneCmd "github.com/cli/cli/pkg/cmd/repo/clone"
	repoCreateCmd "github.com/cli/cli/pkg/cmd/repo/create"
	creditsCmd "github.com/cli/cli/pkg/cmd/repo/credits"
	repoForkCmd "github.com/cli/cli/pkg/cmd/repo/fork"
	repoViewCmd "github.com/cli/cli/pkg/cmd/repo/view"
	repoSecretCmd "github.com/cli/cli/pkg/cmd/repo/secret"
	"github.com/cli/cli/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdRepo(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repo <command>",
		Short: "Create, clone, fork, and view repositories",
		Long:  `Work with GitHub repositories`,
		Example: heredoc.Doc(`
			$ gh repo create
			$ gh repo clone cli/cli
			$ gh repo view --web
			$ gh repo secret
		`),
		Annotations: map[string]string{
			"IsCore": "true",
			"help:arguments": heredoc.Doc(`
				A repository can be supplied as an argument in any of the following formats:
				- "OWNER/REPO"
				- by URL, e.g. "https://github.com/OWNER/REPO"
			`),
		},
	}

	cmd.AddCommand(repoViewCmd.NewCmdView(f, nil))
	cmd.AddCommand(repoForkCmd.NewCmdFork(f, nil))
	cmd.AddCommand(repoCloneCmd.NewCmdClone(f, nil))
	cmd.AddCommand(repoCreateCmd.NewCmdCreate(f, nil))
	cmd.AddCommand(creditsCmd.NewCmdRepoCredits(f, nil))
	cmd.AddCommand(repoSecretCmd.NewCmdSecret(f, nil))

	return cmd
}
