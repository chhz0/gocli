package gocli

import (
	"github.com/spf13/cobra"
)

type Executer struct {
	exec *cobra.Command
}

func (e *Executer) Execute() error {
	return e.exec.Execute()
}

type Commander interface {
	Name() string
	Commands() []Commander
}

type SimpleCommand struct {
	Use   string
	Short string
	Long  string

	Run func(cmd *cobra.Command, args []string) error
	// withc func(cd *cobra.Command, r *rootCommand) error
	// initc func(cd *Command) error
	Args cobra.PositionalArgs

	Subcommands []*SimpleCommand

	Cmd *cobra.Command
}

func (cd *SimpleCommand) AppendCommands(cmds ...*SimpleCommand) *SimpleCommand {
	cd.Subcommands = append(cd.Subcommands, cmds...)
	return cd
}

func New(scmd *SimpleCommand) *Executer {
	rootCmd := scmd.buildCobra()
	return &Executer{exec: rootCmd}
}

func (cd *SimpleCommand) buildCobra() *cobra.Command {
	if cd.Cmd != nil {
		return cd.Cmd // 如果 cmd 已经初始化，直接返回
	}

	rootCmd := &cobra.Command{
		Use:   cd.Use,
		Short: cd.Short,
		Long:  cd.Long,
		Args:  cd.Args,
		RunE:  cd.Run,
	}

	for _, subcmd := range cd.Subcommands {
		subCmd := subcmd.buildCobra()
		if subCmd != nil {
			rootCmd.AddCommand(subCmd)
		}
	}
	cd.Cmd = rootCmd // 初始化 cmd
	return cd.Cmd
}
