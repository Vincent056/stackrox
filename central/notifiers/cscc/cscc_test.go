package cscc

import (
	"context"
	"testing"

	"github.com/gogo/protobuf/types"
	clusterMocks "github.com/stackrox/rox/central/cluster/datastore/mocks"
	"github.com/stackrox/rox/central/notifiers/cscc/findings"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/cryptoutils/cryptocodec"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

const (
	projectKey = "FJ"
)

func TestWithFakeCSCC(t *testing.T) {

	sourceID := "organizations/0000000000/sources/0000000000"

	s := &storage.Notifier{
		Name:         "FakeSCC",
		UiEndpoint:   "https://central.stackrox",
		Type:         "scc",
		LabelDefault: projectKey,
		Config: &storage.Notifier_Cscc{
			Cscc: &storage.CSCC{
				ServiceAccount: "test_service_account",
				SourceId:       sourceID,
			},
		},
	}

	cluster := &storage.Cluster{
		Id:   "test_id",
		Name: "test_cluster",
		Status: &storage.ClusterStatus{
			ProviderMetadata: &storage.ProviderMetadata{
				Zone: "test_zone",
				Provider: &storage.ProviderMetadata_Google{
					Google: &storage.GoogleProviderMetadata{
						Project:     "test_project",
						ClusterName: "test_cluster",
					},
				},
			},
		},
	}

	mockCtrl := gomock.NewController(t)
	clusterStore := clusterMocks.NewMockDataStore(mockCtrl)
	clusterStore.EXPECT().GetCluster(gomock.Any(), "test_id").Return(cluster, true, nil)
	scc, _ := newCSCC(s, cryptocodec.Singleton(), "stackrox")

	alertID := "myAlertID"
	severity := findings.High

	testAlert := &storage.Alert{
		Id: alertID,
		Policy: &storage.Policy{
			Id:             "myPolicyID",
			Name:           "myPolicy",
			Description:    "Fake policy",
			PolicySections: []*storage.PolicySection{},
			Severity:       storage.Severity_HIGH_SEVERITY,
		},
		Entity: &storage.Alert_Deployment_{Deployment: &storage.Alert_Deployment{
			Name:      "myDeployment",
			Id:        "myDeploymentID",
			ClusterId: "test_id",
		}},
		Time: types.TimestampNow(),
	}
	findingID := ""
	var finding *findings.Finding
	var err error
	findingID, finding, err = scc.initFinding(context.Background(), testAlert, clusterStore)
	assert.NoError(t, err)
	assert.Equal(t, "myAlertID", findingID)
	assert.NotEmpty(t, finding)
	assert.Equal(t, severity, finding.Severity)
	assert.Equal(t, sourceID, finding.Parent)
	assert.Contains(t, finding.URL, alertID)

}
