package cmd

import (
	"github.com/lcostea/vulnest/internal/app"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(NewS3Command())

}

func NewS3Command() *cobra.Command {
	var s3Cmd = &cobra.Command{
		Use:   "s3",
		Short: "Base command for finding misconfigurations in AWS S3 buckets",
		Long: `AWS S3 buckets are used for a big variety of use-cases and are often misconfigured. 
	This command will help you find those misconfigurations`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	s3Cmd.AddCommand(NewScanCommand())
	return s3Cmd
}

func NewScanCommand() *cobra.Command {
	var bucket string
	var extendedOutput bool
	var scanCmd = &cobra.Command{
		Use:   "scan",
		Short: "Scan a list of S3 buckets for permissions misconfigurations",
		Long: `Scan a specific S3 bucket for permissions misconfigurations:
		list
		read
		write
		download`,
		Run: func(cmd *cobra.Command, args []string) {
			app.ScanBucket(bucket, extendedOutput)
		},
	}
	scanCmd.Flags().StringVar(&bucket, "bucket", "", "Multiple buckets separated by comma to scan")
	scanCmd.Flags().BoolVar(&extendedOutput, "extended-output", false, "Once you find a misconfiguration, print more details on the bucket")
	return scanCmd
}
