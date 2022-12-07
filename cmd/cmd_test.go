package cmd

import (
	"bytes"
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	t.Run("good", func(t *testing.T) {
		flag.CommandLine.Set("n", "1")

		buffer := bytes.Buffer{}
		c := New(&buffer)
		err := c.Run()

		if err != nil {
			t.Error(err)
		}
	})

	t.Run("not exist -n option", func(t *testing.T) {
		flag.CommandLine.Set("n", "-1")

		buffer := bytes.Buffer{}
		c := New(&buffer)
		err := c.Run()

		if err == nil {
			t.Error("expect: fail")
		}
		assert.Equal(t, err, PRNumNotExistError)
	})
}
