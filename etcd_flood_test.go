package main

import (
	"time"

	"github.com/onsi/etcd-flood/flood"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("When rolling etcd", func() {
	var node0, node1, node2 *gexec.Session

	BeforeEach(func() {
		node0 = StartNode(VERSION, 3, 0, DataDir(0, true), "-snapshot-count=1000")
		node1 = StartNode(VERSION, 3, 1, DataDir(1, true), "-snapshot-count=1000")
		node2 = StartNode(VERSION, 3, 2, DataDir(2, true), "-snapshot-count=1000")

		etcdFlood = flood.NewFlood(STORE_SIZE, WRITERS, HEAVY_READERS, LIGHT_READERS, WATCHERS, Machines(0, 1, 2))
		etcdFlood.Flood()
		time.Sleep(10 * time.Second)
	})

	Context("the happy path", func() {
		It("should work when nothing messes with it", func() {
			flood.YellowBanner("checking...")
			Ω(KeysOnNode(0)).Should(Equal(STORE_SIZE))
			Ω(KeysOnNode(1)).Should(Equal(STORE_SIZE))
			Ω(KeysOnNode(2)).Should(Equal(STORE_SIZE))
		})
	})

	It("should work when the first node is shut down", func() {
		flood.YellowBanner("shutting down node 0")
		node0.Interrupt().Wait()
		time.Sleep(5 * time.Second)

		flood.YellowBanner("checking...")
		Ω(KeysOnNode(1)).Should(Equal(STORE_SIZE))
		Ω(KeysOnNode(2)).Should(Equal(STORE_SIZE))
	})

	It("should work when the first node comes back", func() {
		flood.YellowBanner("shutting down node 0")
		node0.Interrupt().Wait()
		time.Sleep(5 * time.Second)

		flood.YellowBanner("restarting node 0...")
		node0 = StartNode(VERSION, 3, 0, DataDir(0, false), "-snapshot-count=1000")
		time.Sleep(5 * time.Second)

		flood.YellowBanner("checking...")
		Ω(KeysOnNode(0)).Should(Equal(STORE_SIZE))
		Ω(KeysOnNode(1)).Should(Equal(STORE_SIZE))
		Ω(KeysOnNode(2)).Should(Equal(STORE_SIZE))
	})
})
