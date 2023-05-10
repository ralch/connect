package service_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi/v5"
	"github.com/ralch/connect/service"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/ralch/connect/service/fake"
)

var _ = Describe("HealthHandler", func() {
	var (
		router  *chi.Mux
		checker *FakeHealthChecker
	)

	BeforeEach(func() {
		router = chi.NewRouter()
		Expect(router).NotTo(BeNil())

		output := &service.HealthCheckResponse{
			Status: service.HealthStatusServing,
		}

		checker = &FakeHealthChecker{}
		checker.CheckReturns(output, nil)

		handler := service.NewHealthHandler(checker)
		handler.Mount(router)
	})

	Describe("Check", func() {
		var (
			w *httptest.ResponseRecorder
			r *http.Request
		)

		BeforeEach(func() {
			w = httptest.NewRecorder()

			input := &FakeHealthCheckRequest{
				Service: "ralch.v1.FooService",
			}

			body := &bytes.Buffer{}
			Expect(json.NewEncoder(body).Encode(input)).To(Succeed())

			r = httptest.NewRequest("POST", "/grpc.health.v1.Health/Check", nil)
			r.Header.Set("Content-Type", "application/json")
			r.Body = io.NopCloser(body)
		})

		It("executes the service checker successfully", func() {
			router.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(200))

			output := &FakeHealthCheckResponse{}
			Expect(json.NewDecoder(w.Body).Decode(output)).To(Succeed())
			Expect(output.Status).To(Equal(""))
		})

		Context("when the service checker execution fails", func() {
			BeforeEach(func() {
				checker.CheckReturns(nil, fmt.Errorf("oh no"))
			})

			It("returns an error", func() {
				router.ServeHTTP(w, r)
				Expect(w.Code).To(Equal(500))

				output := &FakeHealthErrorResponse{}
				Expect(json.NewDecoder(w.Body).Decode(output)).To(Succeed())
				Expect(output.Code).To(Equal("unknown"))
				Expect(output.Message).To(Equal("oh no"))
			})
		})
	})
})
