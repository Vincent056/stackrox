// Code generated by MockGen. DO NOT EDIT.
// Source: data.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/data.go -source data.go
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	framework "github.com/stackrox/rox/central/compliance/framework"
	v1 "github.com/stackrox/rox/generated/api/v1"
	compliance "github.com/stackrox/rox/generated/internalapi/compliance"
	storage "github.com/stackrox/rox/generated/storage"
	set "github.com/stackrox/rox/pkg/set"
	gomock "go.uber.org/mock/gomock"
)

// MockImageMatcher is a mock of ImageMatcher interface.
type MockImageMatcher struct {
	ctrl     *gomock.Controller
	recorder *MockImageMatcherMockRecorder
}

// MockImageMatcherMockRecorder is the mock recorder for MockImageMatcher.
type MockImageMatcherMockRecorder struct {
	mock *MockImageMatcher
}

// NewMockImageMatcher creates a new mock instance.
func NewMockImageMatcher(ctrl *gomock.Controller) *MockImageMatcher {
	mock := &MockImageMatcher{ctrl: ctrl}
	mock.recorder = &MockImageMatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageMatcher) EXPECT() *MockImageMatcherMockRecorder {
	return m.recorder
}

// Match mocks base method.
func (m *MockImageMatcher) Match(image *storage.ImageName) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Match", image)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Match indicates an expected call of Match.
func (mr *MockImageMatcherMockRecorder) Match(image any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Match", reflect.TypeOf((*MockImageMatcher)(nil).Match), image)
}

// MockComplianceDataRepository is a mock of ComplianceDataRepository interface.
type MockComplianceDataRepository struct {
	ctrl     *gomock.Controller
	recorder *MockComplianceDataRepositoryMockRecorder
}

// MockComplianceDataRepositoryMockRecorder is the mock recorder for MockComplianceDataRepository.
type MockComplianceDataRepositoryMockRecorder struct {
	mock *MockComplianceDataRepository
}

// NewMockComplianceDataRepository creates a new mock instance.
func NewMockComplianceDataRepository(ctrl *gomock.Controller) *MockComplianceDataRepository {
	mock := &MockComplianceDataRepository{ctrl: ctrl}
	mock.recorder = &MockComplianceDataRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComplianceDataRepository) EXPECT() *MockComplianceDataRepositoryMockRecorder {
	return m.recorder
}

// CISDockerTriggered mocks base method.
func (m *MockComplianceDataRepository) CISDockerTriggered() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CISDockerTriggered")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CISDockerTriggered indicates an expected call of CISDockerTriggered.
func (mr *MockComplianceDataRepositoryMockRecorder) CISDockerTriggered() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CISDockerTriggered", reflect.TypeOf((*MockComplianceDataRepository)(nil).CISDockerTriggered))
}

// CISKubernetesTriggered mocks base method.
func (m *MockComplianceDataRepository) CISKubernetesTriggered() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CISKubernetesTriggered")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CISKubernetesTriggered indicates an expected call of CISKubernetesTriggered.
func (mr *MockComplianceDataRepositoryMockRecorder) CISKubernetesTriggered() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CISKubernetesTriggered", reflect.TypeOf((*MockComplianceDataRepository)(nil).CISKubernetesTriggered))
}

// Cluster mocks base method.
func (m *MockComplianceDataRepository) Cluster() *storage.Cluster {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster")
	ret0, _ := ret[0].(*storage.Cluster)
	return ret0
}

// Cluster indicates an expected call of Cluster.
func (mr *MockComplianceDataRepositoryMockRecorder) Cluster() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockComplianceDataRepository)(nil).Cluster))
}

// ComplianceOperatorResults mocks base method.
func (m *MockComplianceDataRepository) ComplianceOperatorResults() map[string][]*storage.ComplianceOperatorCheckResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComplianceOperatorResults")
	ret0, _ := ret[0].(map[string][]*storage.ComplianceOperatorCheckResult)
	return ret0
}

// ComplianceOperatorResults indicates an expected call of ComplianceOperatorResults.
func (mr *MockComplianceDataRepositoryMockRecorder) ComplianceOperatorResults() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComplianceOperatorResults", reflect.TypeOf((*MockComplianceDataRepository)(nil).ComplianceOperatorResults))
}

// Deployments mocks base method.
func (m *MockComplianceDataRepository) Deployments() map[string]*storage.Deployment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deployments")
	ret0, _ := ret[0].(map[string]*storage.Deployment)
	return ret0
}

// Deployments indicates an expected call of Deployments.
func (mr *MockComplianceDataRepositoryMockRecorder) Deployments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deployments", reflect.TypeOf((*MockComplianceDataRepository)(nil).Deployments))
}

// HasProcessIndicators mocks base method.
func (m *MockComplianceDataRepository) HasProcessIndicators() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasProcessIndicators")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasProcessIndicators indicates an expected call of HasProcessIndicators.
func (mr *MockComplianceDataRepositoryMockRecorder) HasProcessIndicators() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasProcessIndicators", reflect.TypeOf((*MockComplianceDataRepository)(nil).HasProcessIndicators))
}

// HostScraped mocks base method.
func (m *MockComplianceDataRepository) HostScraped(node *storage.Node) *compliance.ComplianceReturn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HostScraped", node)
	ret0, _ := ret[0].(*compliance.ComplianceReturn)
	return ret0
}

// HostScraped indicates an expected call of HostScraped.
func (mr *MockComplianceDataRepositoryMockRecorder) HostScraped(node any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HostScraped", reflect.TypeOf((*MockComplianceDataRepository)(nil).HostScraped), node)
}

// ImageIntegrations mocks base method.
func (m *MockComplianceDataRepository) ImageIntegrations() []*storage.ImageIntegration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageIntegrations")
	ret0, _ := ret[0].([]*storage.ImageIntegration)
	return ret0
}

// ImageIntegrations indicates an expected call of ImageIntegrations.
func (mr *MockComplianceDataRepositoryMockRecorder) ImageIntegrations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageIntegrations", reflect.TypeOf((*MockComplianceDataRepository)(nil).ImageIntegrations))
}

// Images mocks base method.
func (m *MockComplianceDataRepository) Images() []*storage.ListImage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Images")
	ret0, _ := ret[0].([]*storage.ListImage)
	return ret0
}

// Images indicates an expected call of Images.
func (mr *MockComplianceDataRepositoryMockRecorder) Images() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Images", reflect.TypeOf((*MockComplianceDataRepository)(nil).Images))
}

// K8sRoleBindings mocks base method.
func (m *MockComplianceDataRepository) K8sRoleBindings() []*storage.K8SRoleBinding {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "K8sRoleBindings")
	ret0, _ := ret[0].([]*storage.K8SRoleBinding)
	return ret0
}

// K8sRoleBindings indicates an expected call of K8sRoleBindings.
func (mr *MockComplianceDataRepositoryMockRecorder) K8sRoleBindings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "K8sRoleBindings", reflect.TypeOf((*MockComplianceDataRepository)(nil).K8sRoleBindings))
}

// K8sRoles mocks base method.
func (m *MockComplianceDataRepository) K8sRoles() []*storage.K8SRole {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "K8sRoles")
	ret0, _ := ret[0].([]*storage.K8SRole)
	return ret0
}

// K8sRoles indicates an expected call of K8sRoles.
func (mr *MockComplianceDataRepositoryMockRecorder) K8sRoles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "K8sRoles", reflect.TypeOf((*MockComplianceDataRepository)(nil).K8sRoles))
}

// NetworkFlows mocks base method.
func (m *MockComplianceDataRepository) NetworkFlows() []*storage.NetworkFlow {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkFlows")
	ret0, _ := ret[0].([]*storage.NetworkFlow)
	return ret0
}

// NetworkFlows indicates an expected call of NetworkFlows.
func (mr *MockComplianceDataRepositoryMockRecorder) NetworkFlows() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkFlows", reflect.TypeOf((*MockComplianceDataRepository)(nil).NetworkFlows))
}

// NetworkGraph mocks base method.
func (m *MockComplianceDataRepository) NetworkGraph() *v1.NetworkGraph {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkGraph")
	ret0, _ := ret[0].(*v1.NetworkGraph)
	return ret0
}

// NetworkGraph indicates an expected call of NetworkGraph.
func (mr *MockComplianceDataRepositoryMockRecorder) NetworkGraph() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkGraph", reflect.TypeOf((*MockComplianceDataRepository)(nil).NetworkGraph))
}

// NetworkPolicies mocks base method.
func (m *MockComplianceDataRepository) NetworkPolicies() map[string]*storage.NetworkPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkPolicies")
	ret0, _ := ret[0].(map[string]*storage.NetworkPolicy)
	return ret0
}

// NetworkPolicies indicates an expected call of NetworkPolicies.
func (mr *MockComplianceDataRepositoryMockRecorder) NetworkPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkPolicies", reflect.TypeOf((*MockComplianceDataRepository)(nil).NetworkPolicies))
}

// NodeResults mocks base method.
func (m *MockComplianceDataRepository) NodeResults() map[string]map[string]*compliance.ComplianceStandardResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeResults")
	ret0, _ := ret[0].(map[string]map[string]*compliance.ComplianceStandardResult)
	return ret0
}

// NodeResults indicates an expected call of NodeResults.
func (mr *MockComplianceDataRepositoryMockRecorder) NodeResults() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeResults", reflect.TypeOf((*MockComplianceDataRepository)(nil).NodeResults))
}

// Nodes mocks base method.
func (m *MockComplianceDataRepository) Nodes() map[string]*storage.Node {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Nodes")
	ret0, _ := ret[0].(map[string]*storage.Node)
	return ret0
}

// Nodes indicates an expected call of Nodes.
func (mr *MockComplianceDataRepositoryMockRecorder) Nodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nodes", reflect.TypeOf((*MockComplianceDataRepository)(nil).Nodes))
}

// Notifiers mocks base method.
func (m *MockComplianceDataRepository) Notifiers() []*storage.Notifier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Notifiers")
	ret0, _ := ret[0].([]*storage.Notifier)
	return ret0
}

// Notifiers indicates an expected call of Notifiers.
func (mr *MockComplianceDataRepositoryMockRecorder) Notifiers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notifiers", reflect.TypeOf((*MockComplianceDataRepository)(nil).Notifiers))
}

// Policies mocks base method.
func (m *MockComplianceDataRepository) Policies() map[string]*storage.Policy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Policies")
	ret0, _ := ret[0].(map[string]*storage.Policy)
	return ret0
}

// Policies indicates an expected call of Policies.
func (mr *MockComplianceDataRepositoryMockRecorder) Policies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Policies", reflect.TypeOf((*MockComplianceDataRepository)(nil).Policies))
}

// PolicyCategories mocks base method.
func (m *MockComplianceDataRepository) PolicyCategories() map[string]set.StringSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PolicyCategories")
	ret0, _ := ret[0].(map[string]set.StringSet)
	return ret0
}

// PolicyCategories indicates an expected call of PolicyCategories.
func (mr *MockComplianceDataRepositoryMockRecorder) PolicyCategories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PolicyCategories", reflect.TypeOf((*MockComplianceDataRepository)(nil).PolicyCategories))
}

// RegistryIntegrations mocks base method.
func (m *MockComplianceDataRepository) RegistryIntegrations() []framework.ImageMatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegistryIntegrations")
	ret0, _ := ret[0].([]framework.ImageMatcher)
	return ret0
}

// RegistryIntegrations indicates an expected call of RegistryIntegrations.
func (mr *MockComplianceDataRepositoryMockRecorder) RegistryIntegrations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegistryIntegrations", reflect.TypeOf((*MockComplianceDataRepository)(nil).RegistryIntegrations))
}

// SSHProcessIndicators mocks base method.
func (m *MockComplianceDataRepository) SSHProcessIndicators() []*storage.ProcessIndicator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SSHProcessIndicators")
	ret0, _ := ret[0].([]*storage.ProcessIndicator)
	return ret0
}

// SSHProcessIndicators indicates an expected call of SSHProcessIndicators.
func (mr *MockComplianceDataRepositoryMockRecorder) SSHProcessIndicators() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SSHProcessIndicators", reflect.TypeOf((*MockComplianceDataRepository)(nil).SSHProcessIndicators))
}

// ScannerIntegrations mocks base method.
func (m *MockComplianceDataRepository) ScannerIntegrations() []framework.ImageMatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScannerIntegrations")
	ret0, _ := ret[0].([]framework.ImageMatcher)
	return ret0
}

// ScannerIntegrations indicates an expected call of ScannerIntegrations.
func (mr *MockComplianceDataRepositoryMockRecorder) ScannerIntegrations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScannerIntegrations", reflect.TypeOf((*MockComplianceDataRepository)(nil).ScannerIntegrations))
}

// UnresolvedAlerts mocks base method.
func (m *MockComplianceDataRepository) UnresolvedAlerts() []*storage.ListAlert {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnresolvedAlerts")
	ret0, _ := ret[0].([]*storage.ListAlert)
	return ret0
}

// UnresolvedAlerts indicates an expected call of UnresolvedAlerts.
func (mr *MockComplianceDataRepositoryMockRecorder) UnresolvedAlerts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnresolvedAlerts", reflect.TypeOf((*MockComplianceDataRepository)(nil).UnresolvedAlerts))
}
