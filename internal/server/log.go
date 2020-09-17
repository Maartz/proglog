package server

import (
	"fmt"
	"sync"
)

type Log struct {
	mu      sync.Mutex
	records []Record
}

//NewLog returns a reference to a new Log struct
func NewLog() *Log {
	return &Log{}
}

//Append append a record to the log
func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

//Read read a record with the given index, it use that index to look up the record in the slice
func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if offset >= uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}

type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

//ErrOffsetNotFound returns an error saying that the offset doesn’t exist
var ErrOffsetNotFound = fmt.Errorf("offset not found")
