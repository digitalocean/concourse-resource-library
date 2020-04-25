package log_test

import (
	"log"
	"os"
	"testing"

	rlog "github.com/digitalocean/concourse-resource-library/log"
)

func TestLog(t *testing.T) {
	os.Setenv("LOG_DIRECTORY", "tmp")

	input := rlog.WriteStdin()
	defer rlog.Close()

	// TODO: write stdin
	log.Println(input)

	// TODO: read file contents
	//f, err := os.OpenFile(fmt.Sprintf(rlog.LogFilePattern, dir, time.Now().Format("2006-01-02")), os.O_RD, 0666)
	//if err != nil {
	//	log.Fatalf("error opening file: %v", err)
	//}

	// TODO: assert file contents

}
