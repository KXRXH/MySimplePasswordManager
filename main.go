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
	var isAdd, isRead, isKeys bool
	var values []string
	parser.Init(&isAdd, &isRead, &isKeys, &values)
	boolSum := utils.Bool2Int[isAdd] + utils.Bool2Int[isRead] + utils.Bool2Int[isKeys]
	if boolSum > 1 || boolSum == 0 {
		fmt.Println("Cannot get more than one argument at the same time")
		return
	} else if isAdd && len(values) != 2 {
		fmt.Printf("The following values are not valid: required 2, but got %d\n", len(values))
		return
	} else if isRead && len(values) != 1 {
		fmt.Printf("The following values are not valid: required 1, but got %d\n", len(values))
		return
	}
	config := parser.ParseJson(utils.GetPwd() + "/pm.json")
	database.InitDb(config.DbPath)
	t := table.NewWriter()
	t.SetStyle(theme.Theme[config.Theme])
	t.SetOutputMirror(os.Stdout)
	if isAdd {
		database.AddToDb(values[0], values[1])
	} else if isRead {
		t.AppendHeader(table.Row{"Key", "Password"})
		t.AppendRow(table.Row{values[0], database.GetByKey(values[0])})
	} else if isKeys {
		t.AppendHeader(table.Row{"#", "Key", "Password"})
		for i, field := range database.GetAllPassword() {
			pswrd, _ := utils.DecodeString(field.Password)
			t.AppendRow(table.Row{i + 1, field.Key, pswrd})
		}
	}
	t.Render()
}
