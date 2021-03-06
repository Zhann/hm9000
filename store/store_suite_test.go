package store_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry/hm9000/testhelpers/storerunner"
	"os"
	"os/signal"

	"testing"
)

var etcdRunner *storerunner.ETCDClusterRunner

func TestStore(t *testing.T) {
	registerSignalHandler()
	RegisterFailHandler(Fail)

	etcdRunner = storerunner.NewETCDClusterRunner(5001, 1)

	RunSpecs(t, "Store Suite")

	etcdRunner.Stop()
}

var _ = BeforeEach(func() {
	etcdRunner.Stop()
	etcdRunner.Start()
})

func registerSignalHandler() {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)

		select {
		case <-c:
			etcdRunner.Stop()
			os.Exit(0)
		}
	}()
}
