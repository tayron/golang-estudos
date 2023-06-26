package mock

import (
	"github.com/golang/mock/gomock"
	"github.com/tayron/golang-estudos/tree/master/testes-unitarios/mocks"
	"testing"
)

func TestFoo(t *testing.T) {
	ctrl := gomock.NewController(t)
	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := mocks.NewMockFoo(ctrl)
	// Asserts: the first and only call to Do() is passed 99 with 101 returned.
	// Anything else will fail.
	m.EXPECT().Do(99).Return(101)

	Bar(m)

}
