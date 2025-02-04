package romantests

import (
	"testing"
	"github.com/Hindmalti/kata-roman-numerals/main"
)
func TestTreat(t *testing.T){
	result := main.Treat(1)
	println(result)
}