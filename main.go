package main

import (
	"fmt"
	"os"

	ds "github.com/ipfs/go-datastore"
	dsq "github.com/ipfs/go-datastore/query"
	badger "github.com/ipfs/go-ds-badger"
	badger2 "github.com/ipfs/go-ds-badger2"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("must pass two arguments, old path and new path")
		os.Exit(1)
	}

	old, err := badger.NewDatastore(os.Args[1], nil)
	if err != nil {
		fmt.Println("opening old datastore: ", err)
		os.Exit(1)
	}
	defer old.Close()

	nds, err := badger2.NewDatastore(os.Args[2], nil)
	if err != nil {
		fmt.Println("opening new datastore: ", err)
		os.Exit(1)
	}
	defer nds.Close()

	res, err := old.Query(dsq.Query{})
	if err != nil {
		fmt.Println("querying: ", err)
		os.Exit(1)
	}

	for e := range res.Next() {
		if err := nds.Put(ds.NewKey(e.Key), e.Value); err != nil {
			fmt.Println("failed to write value: ", err)
			os.Exit(1)
		}
	}

	fmt.Println("done")
}
