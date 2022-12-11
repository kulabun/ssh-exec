package rsync

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
)

func PrettyPrintRsyncOutput(reader io.Reader, writer io.Writer) {
	var lastFile string
	var fileInProgress string

	br := bufio.NewReader(reader)
	bw := bufio.NewWriter(writer)

	fileMatcher := regexp.MustCompile(`^(\S.*)$`)

	for {
		line, err := readLine(br)
		if err != nil {
			if err == io.EOF {
				if fileInProgress != "" {
					writeLine(bw, fmt.Sprintf("\033[2K\r - %s updated\n", fileInProgress))
				}
			} else {
				log.Printf("Failed to read rsync output: %s\n", err)
			}
			return
		}

		if line == "" || line == "sending incremental file list" {
			continue
		}

		if line != "" && fileMatcher.MatchString(line) {
			lastFile = fileMatcher.FindString(line)
		} else {
			progress, err := parseRsyncProgress(line)
			if err != nil {
				// ignore lines that doesn't match expected output
				continue
			}
			if fileInProgress != "" && fileInProgress != lastFile {
				writeLine(bw, fmt.Sprintf("\033[2K\r - %s updated\n", fileInProgress))
			}
			fileInProgress = lastFile
			if fileInProgress != "" {
				writeLine(bw, fmt.Sprintf("\033[2K\r[ %d%% ] Uploading %s", progress.PercentDone, fileInProgress))
			}
		}
	}
}

func writeLine(writer *bufio.Writer, line string) error {
	_, err := writer.WriteString(line)
	if err != nil {
		return err
	}
	return writer.Flush()
}

func readLine(reader *bufio.Reader) (string, error) {
	var buffer []byte
	for {
		bytes, isPrefix, err := reader.ReadLine()
		if err != nil {
			return "", err
		}

		buffer = append(buffer, bytes...)
		if !isPrefix {
			break
		}
	}
	return string(buffer), nil
}
