package fixture_test

import (
	"bytes"
	"ganiulis/fixture"
	"io"
	"log"
	"os"
	"sync"
	"testing"
)

func TestRunFixture(t *testing.T) {
	days := "31"
	os.Args[1] = days

	actual := captureOutput(func() {
		fixture.RunFixture()
	})

	testdata := "../testdata/fixture.txt"
	expected, _ := os.ReadFile(testdata)

	if actual != string(expected) {
		t.Error("Fixtures have failed.")
	}
}

func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()

	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	stderr := os.Stderr

	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr

		log.SetOutput(os.Stderr)
	}()

	os.Stdout = writer
	os.Stderr = writer

	log.SetOutput(writer)

	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()

	wg.Wait()

	f()

	writer.Close()

	return <-out
}
