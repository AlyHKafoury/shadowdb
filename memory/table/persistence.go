package table

import (
	"log"
	"os"
)

func (table Table) readOrCreate(databaseName string) error {
	fileInfo, err := os.Stat(databaseName)
	if os.IsNotExist(err) {
		if table.file, err = os.Create(databaseName); err != nil {
			return err
		}
	} else {
		if table.file, err = os.OpenFile(databaseName, os.O_RDWR, 0755); err != nil {
			return err
		}

	}
	table.CurrentPage = int8(fileInfo.Size() / 4096)
	table.RowInPage = uint32(fileInfo.Size()-(int64(table.CurrentPage)*4096)) / 291
	return nil
}

func (table Table) readAll() {
	if int8(len(table.readPages)) < table.CurrentPage+1 {
		for i := int8(0); i <= table.CurrentPage; i++ {
			if _, found := table.readPages[i]; !found {
				offset := int64(i) * 4096
				_, err := table.file.ReadAt(table.Pages[i][:], offset)
				if err != nil {
					log.Fatal(err)
				}
				table.readPages[i] = struct{}{}
			}
		}
	}
}

func (table Table) readPage(pageNumber int8) {
	pagePtr := table.Pages[table.CurrentPage][:]
	offset := int64(pageNumber) * 4096
	_, err := table.file.ReadAt(pagePtr, offset)
	if err != nil {
		log.Fatal(err)
	}
	table.readPages[pageNumber] = struct{}{}
}

func (table Table) isCurrentPageLoaded() (found bool) {
	_, found = table.readPages[table.CurrentPage]
	return
}
