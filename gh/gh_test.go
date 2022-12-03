package gh

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestList(t *testing.T) {
	ctrl := gomock.NewController(t)
	cl := NewMockclientI(ctrl)
	cl.EXPECT().PRDetail().AnyTimes().Return()

	s := &CallClient{
		API: cl,
	}
	s.process()
}
