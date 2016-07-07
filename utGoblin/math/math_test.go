package math

import (
	"testing"

	"github.com/franela/goblin"
)

func TestMath(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("Test Math", func() {
		g.It("Should add", func() {
			sum := AddInt(23, 78)
			g.Assert(sum).Equal(101)

		})
		g.It("Should sub", func() {
			sub := SubInt(78, 23)
			g.Assert(sub).Equal(55)

		})
	})
}
