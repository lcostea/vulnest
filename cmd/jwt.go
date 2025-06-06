package cmd

import (
	"github.com/lcostea/vulnest/internal/jwt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(NewJWTCommand())

}

func NewJWTCommand() *cobra.Command {
	var jwtCmd = &cobra.Command{
		Use:   "jwt",
		Short: "Base command for working with JWT tokens",
		Long: `JWT tokens are used for authentication and authorization, they are usually not encrypted. 
	This command will help you try different scenarios with JWT tokens`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	jwtCmd.AddCommand(UpdateJWTToAdmin())
	return jwtCmd
}

func UpdateJWTToAdmin() *cobra.Command {
	var token string //base64 URL encoded JWT token
	var jwtAdminCmd = &cobra.Command{
		Use:   "elevate-to-admin",
		Short: "This command will update a JWT token to have admin privileges",
		Long:  `Updates a JWT token claims to set user admin and role admin without changes to the signature`,
		Run: func(cmd *cobra.Command, args []string) {
			if token == "" {
				log.Fatalf("no JWT token provided")
			}
			updatedToken := jwt.ElevateToAdmin(token)
			log.Infof("JWT token updated to admin: %s", updatedToken)
		},
	}
	jwtAdminCmd.Flags().StringVar(&token, "jwt", "", "Base 64 URL encoded JWT token to update")
	return jwtAdminCmd
}
