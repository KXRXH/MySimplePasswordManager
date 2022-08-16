package main

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/kxrxh/password-manager/database"
	"github.com/kxrxh/password-manager/parser"
	"github.com/kxrxh/password-manager/theme"
	"github.com/kxrxh/password-manager/utils"
)

func main() {
	var (
		isAdd, isRead, isKeys bool
		args                  []string
	)
	parser.Init(&isAdd, &isRead, &isKeys, &args)
	// Check if params and args are correct
	boolSum := utils.Bool2Int[isAdd] + utils.Bool2Int[isRead] + utils.Bool2Int[isKeys]
	if boolSum > 1 || boolSum == 0 {
		fmt.Println("error: cannot get more than one argument or do not get them at all")
		fmt.Println("avaliable arguments:\n\t-add <key> <value>\n\t-read <key>\n\t-keys")
		return
	} else if isAdd && len(args) != 2 {
		fmt.Printf("error: the following values are not valid: required 2, but got %d\n", len(args))
		return
	} else if isRead && len(args) != 1 {
		fmt.Printf("error: the following values are not valid: required 1, but got %d\n", len(args))
		return
	}
	if isAdd {
		database.AddToDb(args[0], args[1])
		return
	}
	config := parser.ParseJson(utils.GetExecPath() + "/pm.json")
	database.InitDb(config.DbPath)
	// Building a table for data
	outTable := table.NewWriter()
	outTable.SetStyle(theme.Theme[config.Theme])
	outTable.SetOutputMirror(os.Stdout)

	if isRead {
		outTable.AppendHeader(table.Row{"Key", "Password"})
		password, err := database.GetByKey(args[0])
		if err != nil {
			fmt.Printf("error: unable to find key with the given values: %v\n", args[0])
			return
		}
		outTable.AppendRow(table.Row{args[0], password})
	} else if isKeys {
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
