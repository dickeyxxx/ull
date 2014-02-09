package ull_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUll(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ull Suite")
}
