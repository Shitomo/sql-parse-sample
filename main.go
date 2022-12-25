package main

import (
	"fmt"

	pg_query "github.com/pganalyze/pg_query_go"
)

func main() {
	tree, err := pg_query.ParseToJSON("SELECT 1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", tree)
}
