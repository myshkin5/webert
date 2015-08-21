package headerhandler_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHeaderhandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HeaderHandler Suite")
}
