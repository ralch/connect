package health_test

import (
	"context"

	"github.com/ralch/connect-go/controller/health"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/ralch/connect-go/controller/health/fake"
)

var _ = Describe("ServiceChecker", func() {
	var checker *health.ServiceChecker

	BeforeEach(func() {
		checker = health.NewServiceChecker()
		Expect(checker).NotTo(BeNil())
	})

	Describe("Check", func() {
		var fkChecker *FakeChecker

		BeforeEach(func() {
			output := &health.CheckResponse{
				Status: health.StatusServing,
			}

			fkChecker = &FakeChecker{}
			fkChecker.CheckReturns(output, nil)
			// register the fake checker
			checker.Register("ralch.v1.FooService", fkChecker)
		})

		It("executes the service checker successfully", func() {
			input := &health.CheckRequest{
				Service: "ralch.v1.FooService",
			}

			output, err := checker.Check(context.TODO(), input)
			Expect(err).NotTo(HaveOccurred())
			Expect(output).NotTo(BeNil())
			Expect(output.Status).To(Equal(health.StatusServing))
		})

		Context("when the service checker does not exist", func() {
			It("returns an error", func() {
				input := &health.CheckRequest{
					Service: "ralch.v1.BarService",
				}

				output, err := checker.Check(context.TODO(), input)
				Expect(err).To(MatchError("health checker not found for service ralch.v1.BarService"))
				Expect(output).To(BeNil())
			})
		})
	})
})
