package errs

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrap(t *testing.T) {

	baseErr := errors.New("base err")
	wrappedErr := fmt.Errorf("wrapped error {%w}", baseErr)
	t.Logf("wrapped err: %+v", wrappedErr)
	t.Logf("base err: %+v", baseErr)
	unwrappedErr := errors.Unwrap(wrappedErr)
	t.Logf("unwrapped err: %+v", unwrappedErr)
	assert.Equal(t, baseErr, unwrappedErr)

	sonErr := New(-10000, "son")
	fatherErr := fmt.Errorf("father: %w", sonErr)

	var err *Error
	errors.As(fatherErr, &err)
	t.Logf("err: %+v", err)

}
