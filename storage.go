package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Storage interface {
	LoadTasks() ([]Task, error)
	SaveTasks(tasks []Task) error
}

type FileStorage struct {
	filename string
}

func NewFileStorage(filename string) *FileStorage {
	return &FileStorage{filename: filename}
}

func (fs *FileStorage) LoadTasks() ([]Task, error) {
	if _, err := os.Stat(fs.filename); os.IsNotExist(err) {
		return []Task{}, nil
	}

	data, err := ioutil.ReadFile(fs.filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return tasks, nil
}

func (fs *FileStorage) SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling tasks to JSON: %v", err)
	}

	if err := ioutil.WriteFile(fs.filename, data, 0644); err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}
