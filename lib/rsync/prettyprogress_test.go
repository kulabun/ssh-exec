package rsync

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrettyPrintRsyncOutput_NoFiles(t *testing.T) {
	input := `
sending incremental file list
            0   0%    0.00kB/s    0:00:00 (xfr#0, to-chk=0/22)
`
	want := ``

	writer := new(bytes.Buffer)
	reader := strings.NewReader(input)

	PrettyPrintRsyncOutput(reader, writer)

	got := writer.String()

	println(got)
	if got != want {
		t.Fatalf("PrettyPrintRsyncOutput(%q) = %v, want %v", input, got, want)
	}
}

func TestPrettyPrintRsyncOutput_WithFiles(t *testing.T) {
	input := `
sending incremental file list
file-1
    102,400,000  10%  396.85MB/s    0:00:00 (xfr#13, to-chk=9/22) 
dir-2/file-2
    204,800,000  20%  507.22MB/s    0:00:00 (xfr#14, to-chk=8/22)
dir-3/file-3
    307,200,000  30%  526.87MB/s    0:00:00 (xfr#15, to-chk=7/22)
file-4
    409,600,000  40%  563.63MB/s    0:00:00 (xfr#16, to-chk=6/22)
file-5
    512,000,000  50%  586.84MB/s    0:00:00 (xfr#17, to-chk=5/22)
file-6
    614,400,000  60%  605.28MB/s    0:00:00 (xfr#18, to-chk=4/22)
file-7
    716,800,000  70%  618.05MB/s    0:00:01 (xfr#19, to-chk=3/22)
file-8
    819,200,000  80%  629.51MB/s    0:00:01 (xfr#20, to-chk=2/22)
file-9
    921,600,000  90%  639.18MB/s    0:00:01 (xfr#21, to-chk=1/22)
file-99
  1,024,000,000 100%  645.85MB/s    0:00:01 (xfr#22, to-chk=0/22)
`
	want := "\x1b[2K\r[ 10% ] Uploading file-1\x1b[2K\r - file-1 updated\n\x1b[2K\r[ 20% ] Uploading dir-2/file-2\x1b[2K\r - dir-2/file-2 updated\n\x1b[2K\r[ 30% ] Uploading dir-3/file-3\x1b[2K\r - dir-3/file-3 updated\n\x1b[2K\r[ 40% ] Uploading file-4\x1b[2K\r - file-4 updated\n\x1b[2K\r[ 50% ] Uploading file-5\x1b[2K\r - file-5 updated\n\x1b[2K\r[ 60% ] Uploading file-6\x1b[2K\r - file-6 updated\n\x1b[2K\r[ 70% ] Uploading file-7\x1b[2K\r - file-7 updated\n\x1b[2K\r[ 80% ] Uploading file-8\x1b[2K\r - file-8 updated\n\x1b[2K\r[ 90% ] Uploading file-9\x1b[2K\r - file-9 updated\n\x1b[2K\r[ 100% ] Uploading file-99\x1b[2K\r - file-99 updated\n"

	writer := new(bytes.Buffer)
	reader := strings.NewReader(input)

	PrettyPrintRsyncOutput(reader, writer)

	got := writer.String()
	t.Logf("%q", got)

	if got != want {
		t.Fatalf("PrettyPrintRsyncOutput(%q) = %v, want %v", input, got, want)
	}
}
