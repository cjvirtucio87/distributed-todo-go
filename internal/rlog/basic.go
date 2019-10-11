package rlog

import (
	"cjvirtucio87/distributed-todo-go/internal/dto"
	"errors"
	"fmt"
)

type BasicLog struct {
	backend []dto.Entry
}

func (l *BasicLog) AddEntries(e dto.EntryInfo) {
	l.backend = append(l.backend[:e.NextIndex], e.Entries...)
}

func (l *BasicLog) Count() int {
	return len(l.backend)
}

func (l *BasicLog) Entry(idx int) (dto.Entry, error) {
	e := l.backend[idx]

	if &e == nil {
		return e, errors.New(
			fmt.Sprintf(
				"could not retrieve entry for index %d",
				idx,
			),
		)
	}

	return e, nil
}

func (l *BasicLog) Entries(start, end int) ([]dto.Entry, bool) {
	if start < 0 || end > l.Count() {
		return []dto.Entry{}, false
	}

	return l.backend[start:end], true
}
