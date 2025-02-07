package features

//lint:file-ignore U1000 we want to introduce this feature flag unused.

var (
	// csvExport enables CSV export of search results.
	csvExport = registerUnchangeableFeature("Enable CSV export of search results", "ROX_CSV_EXPORT", false)

	// NetworkDetectionBaselineSimulation enables new features related to the baseline simulation part of the network detection experience.
	NetworkDetectionBaselineSimulation = registerUnchangeableFeature("Enable network detection baseline simulation", "ROX_NETWORK_DETECTION_BASELINE_SIMULATION", true)

	// IntegrationsAsConfig enables loading integrations from config
	IntegrationsAsConfig = registerUnchangeableFeature("Enable loading integrations from config", "ROX_INTEGRATIONS_AS_CONFIG", false)

	// QuayRobotAccounts enables Robot accounts as credentials in Quay Image Integration.
	QuayRobotAccounts = registerUnchangeableFeature("Enable Robot accounts in Quay Image Integration", "ROX_QUAY_ROBOT_ACCOUNTS", true)

	// RoxctlNetpolGenerate enables 'roxctl netpol generate' command which integrates with NP-Guard
	RoxctlNetpolGenerate = registerUnchangeableFeature("Enable 'roxctl netpol generate' command", "ROX_ROXCTL_NETPOL_GENERATE", true)

	// RoxSyslogExtraFields enables user to add additional key value pairs in syslog alert notification in cef format.
	RoxSyslogExtraFields = registerUnchangeableFeature("Enable extra fields for syslog integration", "ROX_SYSLOG_EXTRA_FIELDS", true)

	// SourcedAutogeneratedIntegrations enables adding a "source" to autogenerated integrations.
	SourcedAutogeneratedIntegrations = registerUnchangeableFeature("Enable autogenerated integrations with cluster/namespace/secret source", "ROX_SOURCED_AUTOGENERATED_INTEGRATIONS", false)

	// VulnMgmtWorkloadCVEs enables APIs and UI pages for the VM Workload CVE enhancements
	VulnMgmtWorkloadCVEs = registerUnchangeableFeature("Vuln Mgmt Workload CVEs", "ROX_VULN_MGMT_WORKLOAD_CVES", true)

	// PostgresBlobStore enables the creation of the Postgres Blob Store
	PostgresBlobStore = registerUnchangeableFeature("Postgres Blob Store", "ROX_POSTGRES_BLOB_STORE", false)

	// StoreEventHashes stores the hashes of successfully processed objects we receive from Sensor into the database
	StoreEventHashes = registerUnchangeableFeature("Store Event Hashes", "ROX_STORE_EVENT_HASHES", true)

	// PreventSensorRestartOnDisconnect enables a new behavior in Sensor where it avoids restarting when the gRPC connection with Central ends.
	PreventSensorRestartOnDisconnect = registerUnchangeableFeature("Prevent Sensor restart on disconnect", "ROX_PREVENT_SENSOR_RESTART_ON_DISCONNECT", true)

	// SyslogNamespaceLabels enables sending namespace labels as part of the syslog alert notification.
	SyslogNamespaceLabels = registerUnchangeableFeature("Send namespace labels as part of the syslog alert notification", "ROX_SEND_NAMESPACE_LABELS_IN_SYSLOG", true)

	// MoveInitBundlesUI is front-end only move from integrations to clusters route.
	MoveInitBundlesUI = registerUnchangeableFeature("Move init-bundles UI", "ROX_MOVE_INIT_BUNDLES_UI", false)

	// ComplianceEnhancements enables APIs and UI pages for Compliance 2.0
	ComplianceEnhancements = registerUnchangeableFeature("Compliance enhancements", "ROX_COMPLIANCE_ENHANCEMENTS", false)

	// AdministrationEvents enables APIs (including collection) and UI pages for administration events.
	AdministrationEvents = registerFeature("Enable administration events", "ROX_ADMINISTRATION_EVENTS", true)

	// PostgresDatastore defines if PostgresSQL should be used
	PostgresDatastore = registerUnchangeableFeature("Enable Postgres Datastore", "ROX_POSTGRES_DATASTORE", true)

	// ActiveVulnMgmt defines if the active vuln mgmt feature is enabled
	ActiveVulnMgmt = registerFeature("Enable Active Vulnerability Management", "ROX_ACTIVE_VULN_MGMT", false)

	// VulnReportingEnhancements enables APIs and UI pages for VM Reporting enhancements including downloadable reports
	VulnReportingEnhancements = registerFeature("Enable Vulnerability Reporting enhancements", "ROX_VULN_MGMT_REPORTING_ENHANCEMENTS", true)

	// UnifiedCVEDeferral enables APIs and UI pages for unified deferral workflow.
	UnifiedCVEDeferral = registerFeature("Enable new unified Vulnerability deferral workflow", "ROX_VULN_MGMT_UNIFIED_CVE_DEFERRAL", false)

	// SensorReconciliationOnReconnect enables sensors to support reconciliation when reconnecting
	SensorReconciliationOnReconnect = registerFeature("Enable Sensors to support reconciliation on reconnect", "ROX_SENSOR_RECONCILIATION", false)

	// AuthMachineToMachine allows to exchange ID tokens for Central tokens without requiring user interaction.
	AuthMachineToMachine = registerFeature("Enable Auth Machine to Machine functionalities", "ROX_AUTH_MACHINE_TO_MACHINE", false)
)
