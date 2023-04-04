package must

import (
	"testing"

	"github.com/Revolution1/rutil/ptr"
	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {
	assert.EqualValues(t, "1", Default("", "1"))
	assert.EqualValues(t, []string{"1"}, Default(nil, []string{"1"}))
	assert.EqualValues(t, ptr.Of("1"), Default(nil, ptr.Of("1")))
	assert.EqualValues(t, []string{"1"}, Default([]string{}, []string{"1"}))
}
