package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vne",
	Short: "Vulnest is a tool to check for misconfigurations fast and easy",
	Long: `A fast and easy tool to check for security misconfigurations especially
                starting with AWS S3 buckets`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

// func init() {
//   cobra.OnInitialize(initConfig)
//   rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
//   rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
//   rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
//   rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
//   rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
// }

// func initConfig() {
// 	cli.SetLogFormat(cmdutil.LogFormat)
// 	cli.SetLogLevel(cmdutil.LogLevel)
// }

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
