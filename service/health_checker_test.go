package service_test

import (
	"context"

	"github.com/ralch/connect/service"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/ralch/connect/service/fake"
)

var _ = Describe("CompositeHealthChecker", func() {
	var checker *service.CompositeHealthChecker

	BeforeEach(func() {
		checker = service.NewCompositeHealthChecker()
		Expect(checker).NotTo(BeNil())
	})

	Describe("Check", func() {
		var fkChecker *FakeHealthChecker

		BeforeEach(func() {
			output := &service.HealthCheckResponse{
				Status: service.HealthStatusServing,
			}

			fkChecker = &FakeHealthChecker{}
			fkChecker.CheckReturns(output, nil)
			// register the fake checker
			checker.Register("ralch.v1.FooService", fkChecker)
		})

		It("executes the service checker successfully", func() {
			input := &service.HealthCheckRequest{
				Service: "ralch.v1.FooService",
			}

			output, err := checker.Check(context.TODO(), input)
			Expect(err).NotTo(HaveOccurred())
			Expect(output).NotTo(BeNil())
			Expect(output.Status).To(Equal(service.HealthStatusServing))
		})

		Context("when the service checker does not exist", func() {
			It("returns an error", func() {
				input := &service.HealthCheckRequest{
					Service: "ralch.v1.BarService",
				}

				output, err := checker.Check(context.TODO(), input)
				Expect(err).To(MatchError("health checker not found for service ralch.v1.BarService"))
				Expect(output).To(BeNil())
			})
		})
	})
})
