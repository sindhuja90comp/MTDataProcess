package main // Important: Belongs to the main package now

import (
    "os"
)

// ResultSaver struct to hold file writer.
type ResultSaver struct {
    fileWriter *os.File
    filename   string
}

// NewResultSaver creates a new ResultSaver.
func NewResultSaver(filename string) (*ResultSaver, error) {
    file, err := os.Create(filename)
    if err != nil {
        return nil, err
    }
    return &ResultSaver{fileWriter: file, filename: filename}, nil
}

// SaveResult saves a result to the file.
func (rs *ResultSaver) SaveResult(result string) error {
    _, err := rs.fileWriter.WriteString(result + "\n")
    if err != nil {
        return err
    }
    return nil
}

// Close closes the file writer.
func (rs *ResultSaver) Close() error {
    if rs.fileWriter != nil {
        return rs.fileWriter.Close()
    }
    return nil
}