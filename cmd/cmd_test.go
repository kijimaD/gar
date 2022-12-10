package cmd

import (
	"bytes"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	t.Run("valid PR number arg", func(t *testing.T) {
		buffer := bytes.Buffer{}

		ctrl := gomock.NewController(t)
		cl := NewMockRunner(ctrl)
		cl.EXPECT().Run().Times(1).Return(0, "yes", nil)

		c := New(&buffer, cl)
		err := c.Execute([]string{"any", "1"})

		if err != nil {
			t.Error(err)
		}
	})
	t.Run("not exist PR number arg", func(t *testing.T) {
		buffer := bytes.Buffer{}

		ctrl := gomock.NewController(t)
		cl := NewMockRunner(ctrl)
		cl.EXPECT().Run().Times(0).Return(0, "yes", nil)

		c := New(&buffer, cl)
		err := c.Execute([]string{"any"})

		if err == nil {
			t.Error("expect: fail")
		}
		assert.Equal(t, PRNumNotExistError, err)
	})
	t.Run("invalid PR number arg", func(t *testing.T) {
		buffer := bytes.Buffer{}

		ctrl := gomock.NewController(t)
		cl := NewMockRunner(ctrl)
		cl.EXPECT().Run().Times(0).Return(0, "yes", nil)

		c := New(&buffer, cl)
		err := c.Execute([]string{"any", "あ"})

		if err == nil {
			t.Error("expect: fail")
		}
	})
}
