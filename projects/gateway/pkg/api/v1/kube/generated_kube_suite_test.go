package kube_test

import (
	"testing"

	"github.com/avast/retry-go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/test/helpers"
	"github.com/solo-io/go-utils/testutils/clusterlock"
)

func TestKube(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Generated Kube Types Suite")
}

var locker *clusterlock.TestClusterLocker

var _ = BeforeSuite(func() {
	var err error
	locker, err = clusterlock.NewTestClusterLocker(helpers.MustKubeClient(), clusterlock.Options{})
	Expect(err).NotTo(HaveOccurred())
	Expect(locker.AcquireLock(retry.Attempts(40))).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	locker.ReleaseLock()
})