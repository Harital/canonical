package shredder

import (
	"crypto/rand"
	"fmt"
	"go/internal/osInterface"
	"io/fs"
	"os"
)

type Shredder struct {
	os osInterface.OsInterface
}

func NewShredder(osInt osInterface.OsInterface) *Shredder {
	return &Shredder{os: osInt}
}

func (s *Shredder) Shred(filePath string) error {
	file, err := s.os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, fs.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close() // do not do before checking the error as the file could not exist

	// Seed the random number generator

	// Generate random data
	dataSize := 1024 // Size of random data in bytes
	randomData := make([]byte, dataSize)

	for i := 0; i < 3; i++ {
		_, generatorErr := rand.Read(randomData)
		if generatorErr != nil {
			return fmt.Errorf(" generating random data: %w", generatorErr)
		}

		// Write random data to the file
		_, writeErr := file.Write(randomData)
		if writeErr != nil {
			return fmt.Errorf(" writing the file %s: %w", filePath, writeErr)
		}
	}

	delErr := s.os.Remove(filePath)
	if delErr != nil {
		return fmt.Errorf(" deleting the file %s: %w", filePath, delErr)
	}

	return nil
}
