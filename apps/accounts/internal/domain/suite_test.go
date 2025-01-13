package domain_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/onsi/ginkgo/v2/reporters"
	"github.com/onsi/gomega/gexec"
	"os"
)

func TestDomain(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Accounts Domain Suite", []Reporter{junitReporter})
}

var _ = BeforeEach(func() {
	session, err := gexec.Start(exec.Command("echo", "Starting test..."), GinkgoWriter, GinkgoWriter)
	Expect(err).ShouldNot(HaveOccurred())
	Eventually(session).Should(gexec.Exit(0))
})
