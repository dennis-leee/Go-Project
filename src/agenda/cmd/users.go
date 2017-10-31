package cmd

import "github.com/spf13/cobra"

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login",
	Long: ``,
	Run: func(com *cobra.Command, args []string) {
		username, _ := com.Flags().GetString("user")
		checkEmpty("username", username)

		password, _ := com.Flags().GetString("password")
		checkEmpty("password", password)

		cmd.Login(username, password)
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout",
	Long: ``,
	Run: func(com *cobra.Command, args []string) {
		cmd.Logout()
	},
}

var reCmd = &cobra.Command{
	Use:   "register",
	Short: "Register user.",
	Long: `Provide username and password to register, the username is unique`,
	Run: func(com *cobra.Command, args []string) {
		username, _ := com.Flags().GetString("user")
		checkEmpty("username", username)

		password, _ := com.Flags().GetString("password")
		checkEmpty("password", password)

		mail, _ := com.Flags().GetString("mail")
		checkEmpty("mail", mail)

		phone, _ := com.Flags().GetString("phone")
		checkEmpty("phone", phone)

		cmd.Register(username, password, mail, phone)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Long: ``,
	Run: func(com *cobra.Command, args []string) {
		cmd.ShowUsers()
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete your account.",
	Long: `Note: delete cannot goback!`,
	Run: func(com *cobra.Command, args []string) {
		cmd.DeleteUser()
	},
}

func init() {
	reCmd.Flags().StringP("user", "u", "", "Username")
	reCmd.Flags().StringP("password", "p", "", "Help message for username")
	reCmd.Flags().StringP("mail", "m", "", "email.")
	reCmd.Flags().StringP("phone", "t", "", "Phone")

	loginCmd.Flags().StringP("user", "u", "", "Input username")
	loginCmd.Flags().StringP("password", "p", "", "Input password")

	RootCmd.AddCommand(reCmd)
	RootCmd.AddCommand(loginCmd)
	RootCmd.AddCommand(logoutCmd)
	RootCmd.AddCommand(listCmd)
	RootCmd.AddCommand(deleteCmd)
}
