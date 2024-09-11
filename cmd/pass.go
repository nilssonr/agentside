package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "aa",
	Long:  "bb",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := bcrypt.GenerateFromPassword([]byte(args[0]), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(res))
	},
}

func init() {
	rootCmd.AddCommand(passwordCmd)
}
