// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"entity"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register a new user",
	Long: 'input some informations to register a new account',
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("user")
		checkEmpty("username", username)
		password, _ := cmd.Flags().GetString("password")
		checkEmpty("password", password)
		email, _ := cmd.Flags().GetString("email")
		checkEmpty("email", email)
		phone, _ := cmd.Flags().GetString("phone")
		checkEmpty("phone", phone)
		//创建用户实体，返回是否成功
		//fmt.Printf(User.create())
	},
}

func init() {
	RootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringP("user", "u", "", "login account")
	registerCmd.Flags().StringP("password", "p", "", "login password")
	registerCmd.Flags().StringP("email", "e", "", "user's email")
	registerCmd.Flags().StringP("phone", "m", "", "user's phone number")
}
