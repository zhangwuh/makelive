package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_convert(t *testing.T) {
	assert.Equal(t, "LCIRETOESIIGEDHN", convert("LEETCODEISHIRING", 3))
	assert.Equal(t, "LDREOEIIECIHNTSG", convert("LEETCODEISHIRING", 4))
	assert.Equal(t, "AB", convert("AB", 1))
	fmt.Println(0 % 1)
}
