package rsync

import (
  "testing"
)

func TestParseRsyncProgress(t *testing.T) {
  input := "      1,024,000,000 100%  953.67GB/s    0:00:00 (xfr#22, to-chk=0/23)"
  want := Progress{
    ProcessedSize: "1,024,000,000",
    PercentDone:   100,
    Speed:         "953.67GB/s",
    ElapsedTime:   "0:00:00",
    ChunksRemain:  0,
    ChunksTotal:   23,
  }

  got, err := parseRsyncProgress(input)

  if err != nil {
    t.Fatalf("parseRsyncProgress(%s) returned error: %s", input, err)
  }

  if *got != want {
    t.Fatalf("parseRsyncProgress(%q) = %v, want %v", input, got, want)
  }
}
