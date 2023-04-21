package rsb_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRsb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rsb Suite")
}
