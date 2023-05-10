package service_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suite")
}
