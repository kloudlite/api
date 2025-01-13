package domain_test

import (
	"testing"
	"github.com/onsi/ginkgo/v2/reporters"
	"github.com/onsi/gomega/gexec"
	"os"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var junitReporter *reporters.JUnitReporter
var session *gexec.Session

func TestDomain(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter = reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Domain Suite", []Reporter{junitReporter})
}

var _ = BeforeEach(func() {
	var err error
	session, err = gexec.Start(exec.Command("echo", "Starting test..."), GinkgoWriter, GinkgoWriter)
	Expect(err).ShouldNot(HaveOccurred())
	Eventually(session).Should(gexec.Exit(0))
})
