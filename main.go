package main

import (
	"fmt"
	"github.com/linxGnu/grocksdb"
)

func main() {
	bbto := grocksdb.NewDefaultBlockBasedTableOptions()
	bbto.SetBlockCache(grocksdb.NewLRUCache(3 << 30))

	opts := grocksdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbto)
	opts.SetCreateIfMissing(true)

	cfs, err := grocksdb.ListColumnFamilies(opts, "/home/nikhilsamrat/go-buffer/uniswap/database")
	if err != nil {
		fmt.Println("could not get family names", err)
	}

	fmt.Println(cfs)

	db, _, err := grocksdb.OpenDbColumnFamilies(opts, "/home/nikhilsamrat/go-buffer/uniswap/database", cfs, []*grocksdb.Options{opts, opts})
	if err != nil {
		fmt.Println("could not open the database", err)
	}
	
	ro := grocksdb.NewDefaultReadOptions()
	wo := grocksdb.NewDefaultWriteOptions()

	cfOpts := grocksdb.NewDefaultOptions()
	// cfOpts.SetBlockBasedTableFactory(bbto)
	// cfOpts.SetCreateIfMissingColumnFamilies(true)
	defer cfOpts.Destroy()
	
	_, err = db.CreateColumnFamily(cfOpts, "testColumn")
	if err != nil {
		fmt.Println("could not create column", err)
	}
	fmt.Println("column created")

	// if ro and wo are not used again, be sure to Close them.
	err = db.Put(wo, []byte("foo"), []byte("bar"))
	value, err := db.Get(ro, []byte("foo"))
	fmt.Println(string(value.Data()))
	defer value.Free()
	err = db.Delete(wo, []byte("foo"))
}
