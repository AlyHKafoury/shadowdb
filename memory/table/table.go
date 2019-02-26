package table

import (
	"errors"
	"log"
	"unsafe"
)

//Row table row
type Row struct {
	Id       int32
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
}

//New return new table
func New() Table {
	return Table{CurrentPage: 0, RowInPage: 0}
}

//Insert insert row into the latest memory page
func (table *Table) Insert(rowBytes []byte) error {
	if table.RowInPage == uint32(RowsPerPage) && table.CurrentPage == TableMaxPages-1 {
		return errors.New("Table is full")
	}
	pageOffset := table.RowInPage * uint32(RowSize)
	if table.RowInPage > 0 {
		pageOffset += table.RowInPage - 1
	}
	freeMemorySliceInPage := table.Pages[table.CurrentPage][pageOffset:]
	n := copy(freeMemorySliceInPage, rowBytes)
	if n < RowSize-1 {
		return errors.New("Writing to Table Memory Page failed")
	}
	log.Println(pageOffset)
	// log.Printf("%+v\n", table.Pages[table.CurrentPage])
	table.RowInPage++
	if table.RowInPage == uint32(RowsPerPage) {
		if table.CurrentPage == TableMaxPages-1 {
			return errors.New("Table is full")
		}
		table.CurrentPage++
		table.RowInPage = 0
	}
	return nil
}

func (table *Table) ReadRow(page int8, row uint32) []byte {
	dataPage := table.Pages[table.CurrentPage]
	rowStart := row * uint32(RowSize)
	if rowStart > 0 {
		rowStart += row - 1
	}
	rowBytes := dataPage[rowStart : rowStart+uint32(RowSize)+1]
	log.Println(rowStart, rowStart+uint32(RowSize)+1)
	return rowBytes
}
