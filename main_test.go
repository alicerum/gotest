package main

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestAdd(t *testing.T) {
	g := gomega.NewWithT(t)

	a := 5
	b := 3
	c := add(a, b)
	g.Expect(c).To(gomega.Equal(a + b))
}
