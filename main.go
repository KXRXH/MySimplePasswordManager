package main

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/kxrxh/password-manager/constants"
	"github.com/kxrxh/password-manager/database"
	"github.com/kxrxh/password-manager/parser"
	"github.com/kxrxh/password-manager/theme"
	"github.com/kxrxh/password-manager/utils"
)

func main() {
	var (
		mode uint8
		args []string
	)
	err := parser.Init(&mode, &args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	config := parser.ParseJson(utils.GetExecPath() + "/pm.json")
	database.InitDb(config.DbPath)
	// Building a table for data
	outTable := table.NewWriter()
	outTable.SetStyle(theme.Theme[config.Theme])
	outTable.SetOutputMirror(os.Stdout)

	switch mode {
	case constants.Add: // -add <key> <value>
		if len(args) != 2 {
			fmt.Printf("error: expected 2 arguments, got %d\n", len(args))
		} else {
			database.AddToDb(args[0], args[1])
		}
		return
	case constants.Read: // -read <key>
		if len(args) != 1 {
			fmt.Printf("error: expected 1 arguments, got %d\n", len(args))
			return
		}
		outTable.AppendHeader(table.Row{"Key", "Password"})
		password, err := database.GetByKey(args[0])
		if err != nil {
			fmt.Printf("error: unable to find key with the given values: %v\n", args[0])
			return
		}
		outTable.AppendRow(table.Row{args[0], password})
	case constants.Keys: // -keys
		outTable.AppendHeader(table.Row{"#", "Key", "Password"})
		for i, field := range database.GetAllPassword() {
			// Decoding password
			password, err := utils.DecodeString(field.Password)
			if err != nil {
				fmt.Printf("error: unable to decode password: %v\n", err)
				return
			}
			outTable.AppendRow(table.Row{i + 1, field.Key, password})
		}
	}
	outTable.Render()
}
