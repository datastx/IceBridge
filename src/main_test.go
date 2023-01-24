package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	// redirect log output to buffer
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	// Save the original command line arguments
	originalArgs := os.Args

	// Set command line arguments for test
	os.Args = []string{"connect", "-p", "postgres://user:password@host/dbname", "-s", "snowflake://user:password@host/dbname", "-d", "mydb", "-c", "config.json"}
	now := time.Now()
	timeStr := now.Format("2006/01/02 15:04:05")
	cli()

	// check if the log output is correct
	expected := fmt.Sprintf("%s PostgreSQL connection string: postgres://user:password@host/dbname\n%s Snowflake connection string: snowflake://user:password@host/dbname\n%s Database name: mydb\n%s Configuration file: config.json\n", timeStr, timeStr, timeStr, timeStr)
	assert.Equal(t, expected, buf.String())
	// restore original command line arguments
	os.Args = originalArgs
}
