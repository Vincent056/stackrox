// Code originally generated by pg-bindings generator.

//go:build sql_integration

package n39ton40

import (
	"context"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	legacy "github.com/stackrox/rox/migrator/migrations/n_39_to_n_40_postgres_policies/legacy"
	pgStore "github.com/stackrox/rox/migrator/migrations/n_39_to_n_40_postgres_policies/postgres"
	pghelper "github.com/stackrox/rox/migrator/migrations/postgreshelper"
	"github.com/stackrox/rox/pkg/bolthelper"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/suite"
	bolt "go.etcd.io/bbolt"
)

func TestMigration(t *testing.T) {
	suite.Run(t, new(postgresMigrationSuite))
}

type postgresMigrationSuite struct {
	suite.Suite
	ctx context.Context

	legacyDB   *bolt.DB
	postgresDB *pghelper.TestPostgres
}

var _ suite.TearDownTestSuite = (*postgresMigrationSuite)(nil)

func (s *postgresMigrationSuite) SetupTest() {
	var err error
	s.legacyDB, err = bolthelper.NewTemp(s.T().Name() + ".db")
	s.NoError(err)

	s.Require().NoError(err)

	s.ctx = sac.WithAllAccess(context.Background())
	s.postgresDB = pghelper.ForT(s.T(), false)
}

func (s *postgresMigrationSuite) TearDownTest() {
	testutils.TearDownDB(s.legacyDB)
	s.postgresDB.Teardown(s.T())
}

func (s *postgresMigrationSuite) TestPolicyMigration() {
	newStore := pgStore.New(s.postgresDB.DB)
	legacyStore := legacy.New(s.legacyDB)

	// Prepare data and write to legacy DB
	var policys []*storage.Policy
	for i := 0; i < 200; i++ {
		policy := &storage.Policy{}
		s.NoError(testutils.FullInit(policy, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		policys = append(policys, policy)
		s.NoError(legacyStore.Upsert(s.ctx, policy))
	}

	// Move
	s.NoError(move(s.ctx, s.postgresDB.GetGormDB(), s.postgresDB.DB, legacyStore))

	// Verify
	count, err := newStore.Count(s.ctx)
	s.NoError(err)
	s.Equal(len(policys), count)
	for _, policy := range policys {
		fetched, exists, err := newStore.Get(s.ctx, policy.GetId())
		s.NoError(err)
		s.True(exists)
		s.Equal(policy, fetched)
	}
}
