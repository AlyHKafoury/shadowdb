package table

import (
	"errors"
	"fmt"
	"os"
	"unsafe"
)

//Row table row
type Row struct {
	ID       int32
	Username [32]byte
	Email    [255]byte
}

const (
	//RowSize size of each row
	RowSize = int(unsafe.Sizeof(*new(Row))) - 1
	//PageSize size of internal storage page
	PageSize = 4096
	//TableMaxPages number of pages per table
	TableMaxPages = 100
	//RowsPerPage number of rows per page
	RowsPerPage = PageSize / RowSize
	//TableSize actual table size in bytes
	TableSize = PageSize * TableMaxPages
	//RowsPerTable is the max rows per table
	RowsPerTable = RowsPerPage * TableMaxPages
)

//Page is a Memory Page
type Page [PageSize]byte

//Table is a Memory Table
type Table struct {
	Pages       [TableMaxPages]Page
	CurrentPage int8
	RowInPage   uint32
	file        *os.File
	readPages   map[int8]struct{}
}

//New return new table
func New(databaseName string) Table {
	newTable := Table{CurrentPage: 0, RowInPage: 0, readPages: make(map[int8]struct{})}
	newTable.readOrCreate(databaseName)
	return newTable
}

//Insert insert row into the latest memory page
func (table *Table) Insert(rowBytes []byte) error {
	if table.RowInPage == uint32(RowsPerPage) && table.CurrentPage == TableMaxPages-1 {
		fmt.Println("Table is full")
		return errors.New("Table is full")
	}
	pageOffset := table.RowInPage * uint32(RowSize)
	freeMemorySliceInPage := table.Pages[table.CurrentPage][pageOffset:]
	n := copy(freeMemorySliceInPage, rowBytes)
	if n < RowSize-1 {
		return errors.New("Writing to Table Memory Page failed")
	}
	// log.Println(pageOffset)
	// log.Printf("%+v\n", table.Pages[table.CurrentPage])
	table.RowInPage++
	if table.RowInPage == uint32(RowsPerPage) {
		if table.CurrentPage == TableMaxPages-1 {
			return nil
		}
		table.CurrentPage++
		table.RowInPage = 0
	}
	return nil
}

func (table *Table) ReadRow(page int8, row uint32) []byte {
	dataPage := table.Pages[table.CurrentPage]
	rowStart := row * uint32(RowSize)
	rowBytes := dataPage[rowStart : rowStart+uint32(RowSize)+1]
	// log.Println(rowStart, rowStart+uint32(RowSize)+1)
	return rowBytes
}
