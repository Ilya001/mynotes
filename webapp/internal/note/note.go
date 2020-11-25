package note

import (
	"fmt"
	"log"

	"github.com/tarantool/go-tarantool"
)

// Note ...
type Note struct {
	config *Config
}

// NotStruct ...
type NotStruct struct {
	ID   uint64 `json:"id"`
	Text string `json:"text"`
}

// New ...
func New(config *Config) *Note {
	return &Note{
		config: config,
	}
}

// GetAllNotes ...
func (note *Note) GetAllNotes(tuplesPerRequest uint32) ([]NotStruct, error) {
	conn, err := tarantool.Connect(note.config.TarantoolAddr, tarantool.Opts{})
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cursor := []interface{}{}

	resp, err := conn.Select("notes", "scanner", 0, tuplesPerRequest, tarantool.IterGt, cursor)
	if err != nil {
		return []NotStruct{}, err
	}

	if resp.Code != tarantool.OkCode {
		return []NotStruct{}, fmt.Errorf("Select failed: %s", resp.Error)
	}

	tuples := resp.Tuples()

	var allNotes []NotStruct
	for _, tuple := range tuples {
		var oneNote NotStruct
		oneNote.ID = tuple[0].(uint64)
		oneNote.Text = tuple[1].(string)
		allNotes = append(allNotes, oneNote)
	}

	return allNotes, nil
}

// GetOneNote ...
func (note *Note) GetOneNote(id int) (NotStruct, error) {
	conn, err := tarantool.Connect(note.config.TarantoolAddr, tarantool.Opts{})
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	resp, err := conn.Select("notes", "primary", 0, 1, tarantool.IterEq, []interface{}{uint64(id)})
	if err != nil {
		return NotStruct{}, err
	}

	if resp.Code != tarantool.OkCode {
		return NotStruct{}, fmt.Errorf("Select failed: %s", resp.Error)
	}

	tuples := resp.Tuples()
	var oneNote NotStruct
	for _, tuple := range tuples {
		oneNote.ID = tuple[0].(uint64)
		oneNote.Text = tuple[1].(string)
	}
	return oneNote, nil
}

// DeleteOneNote ...
func (note *Note) DeleteOneNote(id int) error {
	conn, err := tarantool.Connect(note.config.TarantoolAddr, tarantool.Opts{})
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	resp, err := conn.Delete("notes", "primary", []interface{}{uint64(id)})
	if err != nil {
		return err
	}
	if resp.Code != tarantool.OkCode {
		return fmt.Errorf(resp.Error)
	}
	return nil
}

// CreateNote ...
func (note *Note) CreateNote(not *NotStruct) error {
	conn, err := tarantool.Connect(note.config.TarantoolAddr, tarantool.Opts{})
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	resp, err := conn.Call("auto_increment_text", []interface{}{not.Text})
	if err != nil {
		return err
	}
	if resp.Code != tarantool.OkCode {
		return fmt.Errorf(resp.Error)
	}
	tuples := resp.Tuples()
	for _, tuple := range tuples {
		not.ID = tuple[0].(uint64)
		not.Text = tuple[1].(string)
	}
	return nil
}

// UpdateNote ...
func (note *Note) UpdateNote(not *NotStruct) error {
	conn, err := tarantool.Connect(note.config.TarantoolAddr, tarantool.Opts{})
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	resp, err := conn.Call("update_text", []interface{}{not.ID, not.Text})
	if err != nil {
		return err
	}
	if resp.Code != tarantool.OkCode {
		return fmt.Errorf(resp.Error)
	}
	return nil
}
