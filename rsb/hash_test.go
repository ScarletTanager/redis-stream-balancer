package rsb_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ScarletTanager/redis-stream-balancer/rsb"
)

var _ = Describe("Hash", func() {
	It("Computes the CRC16-XMODEM value correctly", func() {
		Expect(rsb.CRC16([]byte("123456789"))).To(Equal(uint16(0x31C3)))
	})
})
