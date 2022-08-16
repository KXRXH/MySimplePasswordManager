package parser

import (
	"errors"
	"flag"

	"github.com/kxrxh/password-manager/constants"
	"github.com/kxrxh/password-manager/utils"
)

func Init(mode *uint8, values *[]string) error {
	var isAdd, isRead, isKeys bool
	flag.BoolVar(&isRead, "get", false, "Get password by key")
	flag.BoolVar(&isKeys, "keys", false, "Get all keys")
	flag.BoolVar(&isAdd, "add", false, "Add password to database")
	flag.Parse()
	*values = flag.Args()
	if (utils.Bool2Int[isAdd] + utils.Bool2Int[isRead] + utils.Bool2Int[isKeys]) != 1 {
		return errors.New("error: unable to use zero or more then one param")
	}
	if isAdd {
		*mode = constants.Add
	} else if isRead {
		*mode = constants.Read
	} else if isKeys {
		*mode = constants.Keys
	}
	return nil
}
