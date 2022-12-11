package rsync

import (
	"fmt"
	"regexp"
	"strconv"
)

type Progress struct {
	ProcessedSize string
	Speed         string
	ChunksRemain  int
	ChunksTotal   int
	PercentDone   int
	ElapsedTime   string
}

// parse rsync otput with regexp to Progress
func parseRsyncProgress(line string) (*Progress, error) {
	progress := &Progress{}

	//      1,024,000,000 100%  953.67GB/s    0:00:00 (xfr#22, to-chk=0/23)
	matcher := regexp.MustCompile(` +([\d,\.]+) *(\d+)% +([\d,\.]+\S{2}/s) *([0-9:]+) +\(xfr#\d+, to-chk=(\d+)/(\d+)\)`)
	if matcher.MatchString(line) {
		progress.ProcessedSize = matcher.FindStringSubmatch(line)[1]
		progress.PercentDone, _ = strconv.Atoi(matcher.FindStringSubmatch(line)[2])
		progress.Speed = matcher.FindStringSubmatch(line)[3]
		progress.ElapsedTime = matcher.FindStringSubmatch(line)[4]
		progress.ChunksRemain, _ = strconv.Atoi(matcher.FindStringSubmatch(line)[5])
		progress.ChunksTotal, _ = strconv.Atoi(matcher.FindStringSubmatch(line)[6])
		return progress, nil
	}

	return nil, fmt.Errorf("Failed to parse rsync progress output: `%s`", line)
}
