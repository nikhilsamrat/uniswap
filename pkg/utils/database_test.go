package utils

import (
	"fmt"
	"testing"
)


func TestCreateRocksDB(t *testing.T) {
	db, err := CreateRocksDB("/home/nikhilsamrat/go-buffer/uniswap/database")
	if err != nil {
		fmt.Println("db not created", err)
		return
	}
	db.Close()
}

func TestCreateColumnDB(t *testing.T) {
	db, err := CreateRocksDB("/home/nikhilsamrat/go-buffer/uniswap/testdb")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("created db")
	defer db.Close()
	cfh, err := CreateColumnDB(db, "testColumn")
	defer cfh.Destroy()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("column created")

	db.DropColumnFamily(cfh)
	
}