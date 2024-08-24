package gofp

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Try(t *testing.T) {
	divide := func(a, b int) Try[int] {
		if b == 0 {
			return NewTry(0, errors.New("can not divide by zero"))
		}
		return NewTry(a/b, nil)
	}

	multiplyByTwo := func(a int) Try[int] {
		return NewTry(a*2, nil)
	}

	t.Run("try without failure", func(t *testing.T) {
		o := divide(15, 5).Bind(multiplyByTwo)
		assert.NoError(t, o.Err)
		assert.Equal(t, 6, o.Result)
	})

	t.Run("try with failure", func(t *testing.T) {
		o := divide(15, 0).Bind(multiplyByTwo)
		assert.Error(t, o.Err)
		assert.Equal(t, errors.New("can not divide by zero"), o.Err)
	})
}

func Test_Either(t *testing.T) {
	getValue := func(key string) Either[string] {
		data := map[string]string{"key1": "value1"}
		if v, ok := data[key]; ok {
			return NewEither[string](v, nil)
		}
		return NewEither("", errors.New("key not found"))
	}
	upperCase := func(s string) Either[string] {
		return NewEither(strings.ToUpper(s), nil)
	}

	t.Run("either with success", func(t *testing.T) {
		o := getValue("key1").Bind(upperCase)
		assert.NoError(t, o.Left)
		assert.Equal(t, "VALUE1", o.Right)
	})

	t.Run("either without success", func(t *testing.T) {
		o := getValue("key2").Bind(upperCase)
		assert.Error(t, o.Left)
		assert.Equal(t, errors.New("key not found"), o.Left)
	})
}
