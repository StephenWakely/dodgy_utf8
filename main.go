package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/coreos/go-systemd/v22/journal"
)

func outputMessage(message string, filename string, journald bool, stdout bool) {
	// Define a slice of bytes containing invalid UTF-8 sequences
	// Example: The byte sequence 0xFF, 0xFE is not valid UTF-8
	invalidUTF8 := []byte{0xFF, 0xFE, 0xFD}

	textBytes := []byte(message)
	textBytes = append([]byte{0xE2, 0x99, 0xA5, 0xE2, 0x99, 0xA5, 0xE2, 0x99, 0xA5, 0xE2, 0x99, 0xA5}, textBytes...)
	textBytes = append(textBytes, invalidUTF8...)

	if filename != "" {
		// Open the file in append mode. If it doesn't exist, create it with 0666 permissions
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("Failed to open file: %v", err)
		}
		defer file.Close()

		_, err = file.Write(textBytes)
		if err != nil {
			log.Fatalf("Failed to write to file: %v", err)
		}

		_, err = file.Write([]byte("\n"))
		if err != nil {
			log.Fatalf("Failed to write to file: %v", err)
		}
	}

	if journald {
		err := journal.Print(journal.PriInfo, string(textBytes))
		if err != nil {
			fmt.Printf("Error sending log to journald: %v\n", err)
		}
	}

	if stdout {
		fmt.Printf("%s\n", string(textBytes))
	}

}

func main() {
	// Get some text from the command args
	text := flag.String("text", "", "The text to append to the file")
	loop := flag.Bool("loop", false, "Loop the text to dump to stdout")
	file := flag.String("filename", "", "Filename to append to")
	journal := flag.Bool("journald", false, "Output to journald")
	stdout := flag.Bool("stdout", false, "Output to stdout")
	flag.Parse()

	if *loop {
		for i := 0; ; i++ {
			outputMessage(fmt.Sprintf("%s %d", *text, i), *file, *journal, *stdout)
			time.Sleep(1 * time.Second)
		}

	} else {
		outputMessage(*text, *file, *journal, *stdout)
	}

}
