package helpers

import (
	"bytes"
	"encoding/binary"
	"shadowdb/memory/table"
)

func RowToBytes(row table.Row) ([]byte, error) {
	buffer := bytes.NewBuffer(make([]byte, 0))
	if err := binary.Write(buffer, binary.LittleEndian, row); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func BytesToRow(rowBytes []byte) (table.Row, error) {
	var row table.Row
	buffer := bytes.NewBuffer(rowBytes)
	if err := binary.Read(buffer, binary.LittleEndian, &row); err != nil {
		return row, err
	}
	return row, nil
}
