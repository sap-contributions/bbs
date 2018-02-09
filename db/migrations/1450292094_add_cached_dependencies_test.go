package migrations_test

import (
	"code.cloudfoundry.org/bbs/db/migrations"
	"code.cloudfoundry.org/bbs/migration"
	"code.cloudfoundry.org/lager/lagertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Add Cache Dependencies Migration", func() {
	var (
		migration migration.Migration
		logger    *lagertest.TestLogger
	)

	BeforeEach(func() {
		logger = lagertest.NewTestLogger("test")
		migration = migrations.NewAddCachedDependencies()
	})

	It("appends itself to the migration list", func() {
		Expect(migrations.AllMigrations()).To(ContainElement(migration))
	})

	Describe("Version", func() {
		It("returns the timestamp from which it was created", func() {
			Expect(migration.Version()).To(BeEquivalentTo(1450292094))
		})
	})

	Describe("Up", func() {
		It("returns nil", func() {
			migrationErr := migration.Up(logger)
			Expect(migrationErr).NotTo(HaveOccurred())
		})
	})
})
