package gh

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	cl := NewMockclientI(ctrl)
	cl.EXPECT().List().AnyTimes().Return()

	s := &CallClient{
		API: cl,
	}
	s.run()
}
