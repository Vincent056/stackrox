package service

import (
	"context"

	"github.com/stackrox/rox/central/deployment/datastore"
	multiplierStore "github.com/stackrox/rox/central/multiplier/store"
	processIndicatorDataStore "github.com/stackrox/rox/central/processindicator/datastore"
	riskManager "github.com/stackrox/rox/central/risk/manager"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/pkg/grpc"
	"github.com/stackrox/rox/pkg/logging"
)

var (
	log = logging.LoggerForModule()
)

// Service provides the interface to the microservice that serves alert data.
type Service interface {
	grpc.APIService

	AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error)

	v1.DeploymentServiceServer
}

// New returns a new Service instance using the given DataStore.
func New(datastore datastore.DataStore, processIndicators processIndicatorDataStore.DataStore, multipliers multiplierStore.Store, manager riskManager.Manager) Service {
	return &serviceImpl{
		datastore:         datastore,
		processIndicators: processIndicators,
		multipliers:       multipliers,
		manager:           manager,
	}
}
