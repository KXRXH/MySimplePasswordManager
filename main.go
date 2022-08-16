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
		fmt.Println("Cannot get more than one argument at the same time or do not get them at all")
		return
	} else if isAdd && len(args) != 2 {
		fmt.Printf("The following values are not valid: required 2, but got %d\n", len(args))
		return
	} else if isRead && len(args) != 1 {
		fmt.Printf("The following values are not valid: required 1, but got %d\n", len(args))
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
		outTable.AppendRow(table.Row{args[0], database.GetByKey(args[0])})
	} else if isKeys {
		outTable.AppendHeader(table.Row{"#", "Key", "Password"})
		for i, field := range database.GetAllPassword() {
			// Decoding password
			password, err := utils.DecodeString(field.Password)
			if err != nil {
				panic(err)
			}
			outTable.AppendRow(table.Row{i + 1, field.Key, password})
		}
	}
	outTable.Render()
}
