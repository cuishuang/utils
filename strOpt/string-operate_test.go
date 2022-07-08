package stringOpt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRuleIdSli(t *testing.T) {

	str := "($124$ && $66$) || $253$"

	rs := GetObjFromTwoSymbol(str)

	expect := []string{"124", "66", "253"}

	for k, item := range rs {
		assert.Equal(t, expect[k], item)
	}

}

func TestReplaceRuneInEvenPosition(t *testing.T) {

	str := "($124$ && $66$) || $253$"
	rs := ReplaceRuneInEvenPosition(str, '$', '#')
	assert.Equal(t, "($124# && $66#) || $253#", rs)

}
