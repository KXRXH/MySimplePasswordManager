package parser

import (
	"flag"
)

func Init(isAdd, isRead, isKeys *bool, values *[]string) {
	flag.BoolVar(isRead, "get", false, "Get password by key")
	flag.BoolVar(isKeys, "keys", false, "Get all keys")
	flag.BoolVar(isAdd, "add", false, "Add password to database")
	flag.Parse()
	*values = flag.Args()
}
