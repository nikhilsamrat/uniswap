package utils

import (
	"github.com/linxGnu/grocksdb"
)


type Database struct {
	pools *grocksdb.ColumnFamilyHandle
	arbitrages *grocksdb.ColumnFamilyHandle
	lastBlockNumber uint64
}

func CreateRocksDB(path string) (*grocksdb.DB, error) {
	bbto := grocksdb.NewDefaultBlockBasedTableOptions()
	bbto.SetBlockCache(grocksdb.NewLRUCache(3 << 30))
	defer bbto.Destroy()

	opts := grocksdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbto)
	opts.SetCreateIfMissing(true)
	defer opts.Destroy()

	return grocksdb.OpenDb(opts, path)
}

func OpenRocksDB(path string) (*grocksdb.DB, []*grocksdb.ColumnFamilyHandle, error) {
	bbto := grocksdb.NewDefaultBlockBasedTableOptions()
	bbto.SetBlockCache(grocksdb.NewLRUCache(3 << 30))
	defer bbto.Destroy()

	opts := grocksdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbto)
	defer opts.Destroy()

	cfs, err := grocksdb.ListColumnFamilies(opts, path)
	if err != nil {
		return nil, nil, err
	}
	db, cfhs, err := grocksdb.OpenDbColumnFamilies(opts, path, cfs, []*grocksdb.Options{opts, opts})
	return db, cfhs, nil
}

func CreateColumnDB(db *grocksdb.DB, name string) (*grocksdb.ColumnFamilyHandle, error) {
	bbto := grocksdb.NewDefaultBlockBasedTableOptions()
	bbto.SetBlockCache(grocksdb.NewLRUCache(3 << 30))

	cfOpts := grocksdb.NewDefaultOptions()
	cfOpts.SetBlockBasedTableFactory(bbto)
	defer cfOpts.Destroy()
	
	return db.CreateColumnFamily(cfOpts, name)
}
