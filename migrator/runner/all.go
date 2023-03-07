package runner

import (
	// Import these packages to trigger the registration.
	// Add Postgres -> Postgres migrations at the bottom of the import list.
	_ "github.com/stackrox/rox/migrator/migrations/m_100_to_m_101_cluster_id_netpol_undo_store"
	_ "github.com/stackrox/rox/migrator/migrations/m_101_to_m_102_drop_license_buckets"
	_ "github.com/stackrox/rox/migrator/migrations/m_102_to_m_103_migrate_serial"
	_ "github.com/stackrox/rox/migrator/migrations/m_103_to_m_104_networkpolicy_guidance"
	_ "github.com/stackrox/rox/migrator/migrations/m_104_to_m_105_active_component"
	_ "github.com/stackrox/rox/migrator/migrations/m_105_to_m_106_group_id"
	_ "github.com/stackrox/rox/migrator/migrations/m_106_to_m_107_policy_categories"
	_ "github.com/stackrox/rox/migrator/migrations/m_107_to_m_108_remove_auth_plugin"
	_ "github.com/stackrox/rox/migrator/migrations/m_108_to_m_109_compliance_run_schedules"
	_ "github.com/stackrox/rox/migrator/migrations/m_109_to_m_110_networkpolicy_guidance_2"
	_ "github.com/stackrox/rox/migrator/migrations/m_110_to_m_111_replace_deprecated_resources"
	_ "github.com/stackrox/rox/migrator/migrations/m_111_to_m_112_groups_invalid_values"
	_ "github.com/stackrox/rox/migrator/migrations/m_55_to_m_56_node_scanning_empty"
	_ "github.com/stackrox/rox/migrator/migrations/m_56_to_m_57_compliance_policy_categories"
	_ "github.com/stackrox/rox/migrator/migrations/m_57_to_m_58_update_run_secrets_volume_policy_regex"
	_ "github.com/stackrox/rox/migrator/migrations/m_58_to_m_59_node_scanning_flag_on"
	_ "github.com/stackrox/rox/migrator/migrations/m_59_to_m_60_add_docker_cis_category_to_existing"
	_ "github.com/stackrox/rox/migrator/migrations/m_60_to_m_61_update_network_management_policy_regex"
	_ "github.com/stackrox/rox/migrator/migrations/m_61_to_m_62_multiple_cve_types"
	_ "github.com/stackrox/rox/migrator/migrations/m_62_to_m_63_splunk_source_type"
	_ "github.com/stackrox/rox/migrator/migrations/m_63_to_m_64_exclude_some_openshift_operators_from_policies"
	_ "github.com/stackrox/rox/migrator/migrations/m_64_to_m_65_detect_openshift4_cluster_on_exec_webhooks"
	_ "github.com/stackrox/rox/migrator/migrations/m_65_to_m_66_policy_bug_fixes"
	_ "github.com/stackrox/rox/migrator/migrations/m_66_to_m_67_missing_policy_migrations"
	_ "github.com/stackrox/rox/migrator/migrations/m_67_to_m_68_exclude_pdcsi_from_mount_propagation"
	_ "github.com/stackrox/rox/migrator/migrations/m_68_to_m_69_update_global_access_roles"
	_ "github.com/stackrox/rox/migrator/migrations/m_69_to_m_70_add_xmrig_to_crypto_policy"
	_ "github.com/stackrox/rox/migrator/migrations/m_70_to_m_71_disable_audit_log_collection"
	_ "github.com/stackrox/rox/migrator/migrations/m_71_to_m_72_delete_namespacesac_bucket"
	_ "github.com/stackrox/rox/migrator/migrations/m_72_to_m_73_change_roles_to_sac_v2"
	_ "github.com/stackrox/rox/migrator/migrations/m_73_to_m_74_runtime_policy_event_source"
	_ "github.com/stackrox/rox/migrator/migrations/m_74_to_m_75_severity_policy"
	_ "github.com/stackrox/rox/migrator/migrations/m_75_to_m_76_exclude_compliance_operator_dnf_policy"
	_ "github.com/stackrox/rox/migrator/migrations/m_76_to_m_77_move_roles_to_rocksdb"
	_ "github.com/stackrox/rox/migrator/migrations/m_77_to_m_78_mitre"
	_ "github.com/stackrox/rox/migrator/migrations/m_78_to_m_79_exclude_openshift_sdn_host_pids_policy"
	_ "github.com/stackrox/rox/migrator/migrations/m_79_to_m_80_more_openshift_exclusions"
	_ "github.com/stackrox/rox/migrator/migrations/m_80_to_m_81_rm_demo_policies"
	_ "github.com/stackrox/rox/migrator/migrations/m_81_to_m_82_modify_docker_policies"
	_ "github.com/stackrox/rox/migrator/migrations/m_82_to_m_83_default_pol_flag"
	_ "github.com/stackrox/rox/migrator/migrations/m_83_to_m_84_mitre_fixes"
	_ "github.com/stackrox/rox/migrator/migrations/m_84_to_m_85_exclude_compliance_op_policy"
	_ "github.com/stackrox/rox/migrator/migrations/m_85_to_m_86_apktools_policy"
	_ "github.com/stackrox/rox/migrator/migrations/m_86_to_m_87_microdnf_policy"
	_ "github.com/stackrox/rox/migrator/migrations/m_87_to_m_88_central_secret_policy"
	_ "github.com/stackrox/rox/migrator/migrations/m_88_to_m_89_update_log4shell_policy"
	_ "github.com/stackrox/rox/migrator/migrations/m_89_to_m_90_vuln_state"
	_ "github.com/stackrox/rox/migrator/migrations/m_90_to_m_91_snooze_permissions"
	_ "github.com/stackrox/rox/migrator/migrations/m_91_to_m_92_write_edges_to_graph"
	_ "github.com/stackrox/rox/migrator/migrations/m_92_to_m_93_cleanup_orphaned_rbac_cluster_objs"
	_ "github.com/stackrox/rox/migrator/migrations/m_93_to_m_94_role_accessscopeid"
	_ "github.com/stackrox/rox/migrator/migrations/m_94_to_m_95_cluster_health_status_id"
	_ "github.com/stackrox/rox/migrator/migrations/m_95_to_m_96_alert_scoping_information_at_root"
	_ "github.com/stackrox/rox/migrator/migrations/m_96_to_m_97_modify_default_vulnreportcreator_role"
	_ "github.com/stackrox/rox/migrator/migrations/m_97_to_98_exclude_oauth_sa_kubeadmin_pol"
	_ "github.com/stackrox/rox/migrator/migrations/m_98_to_m_99_process_alert_comments"
	_ "github.com/stackrox/rox/migrator/migrations/m_99_to_m_100_violation_report_branding"

	// The following is the migrations for legacy to Postgres migration. They form a separate sequence.
	// If you have a migration that 1) will release before Postgres Database or 2) does not depend
	// on Postgres Database, you need to increment CurrentDBVersionSeqNum and add the migration to the
	// sequence above (starting with "m_"). Otherwise since your migration depends on the Postgres Database, you need to
	// increment CurrentDBVersionSeqNumWithoutPostgres and add the migration to the sequence below (starting with "n-").
	_ "github.com/stackrox/rox/migrator/migrations/n_01_to_n_02_postgres_clusters"
	_ "github.com/stackrox/rox/migrator/migrations/n_02_to_n_03_postgres_namespaces"
	_ "github.com/stackrox/rox/migrator/migrations/n_03_to_n_04_postgres_deployments"
	_ "github.com/stackrox/rox/migrator/migrations/n_04_to_n_05_postgres_images"
	_ "github.com/stackrox/rox/migrator/migrations/n_05_to_n_06_postgres_active_components"
	_ "github.com/stackrox/rox/migrator/migrations/n_06_to_n_07_postgres_alerts"
	_ "github.com/stackrox/rox/migrator/migrations/n_07_to_n_08_postgres_api_tokens"
	_ "github.com/stackrox/rox/migrator/migrations/n_08_to_n_09_postgres_auth_providers"
	_ "github.com/stackrox/rox/migrator/migrations/n_09_to_n_10_postgres_cluster_cves"
	_ "github.com/stackrox/rox/migrator/migrations/n_10_to_n_11_postgres_cluster_health_statuses"
	_ "github.com/stackrox/rox/migrator/migrations/n_11_to_n_12_postgres_cluster_init_bundles"
	_ "github.com/stackrox/rox/migrator/migrations/n_12_to_n_13_postgres_compliance_domains"
	_ "github.com/stackrox/rox/migrator/migrations/n_13_to_n_14_postgres_compliance_operator_check_results"
	_ "github.com/stackrox/rox/migrator/migrations/n_14_to_n_15_postgres_compliance_operator_profiles"
	_ "github.com/stackrox/rox/migrator/migrations/n_15_to_n_16_postgres_compliance_operator_rules"
	_ "github.com/stackrox/rox/migrator/migrations/n_16_to_n_17_postgres_compliance_operator_scan_setting_bindings"
	_ "github.com/stackrox/rox/migrator/migrations/n_17_to_n_18_postgres_compliance_operator_scans"
	_ "github.com/stackrox/rox/migrator/migrations/n_18_to_n_19_postgres_compliance_run_metadata"
	_ "github.com/stackrox/rox/migrator/migrations/n_19_to_n_20_postgres_compliance_run_results"
	_ "github.com/stackrox/rox/migrator/migrations/n_20_to_n_21_postgres_compliance_strings"
	_ "github.com/stackrox/rox/migrator/migrations/n_21_to_n_22_postgres_configs"
	_ "github.com/stackrox/rox/migrator/migrations/n_22_to_n_23_postgres_external_backups"
	_ "github.com/stackrox/rox/migrator/migrations/n_23_to_n_24_postgres_image_integrations"
	_ "github.com/stackrox/rox/migrator/migrations/n_24_to_n_25_postgres_installation_infos"
	_ "github.com/stackrox/rox/migrator/migrations/n_25_to_n_26_postgres_integration_healths"
	_ "github.com/stackrox/rox/migrator/migrations/n_26_to_n_27_postgres_k8s_roles"
	_ "github.com/stackrox/rox/migrator/migrations/n_27_to_n_28_postgres_log_imbues"
	_ "github.com/stackrox/rox/migrator/migrations/n_28_to_n_29_postgres_network_baselines"
	_ "github.com/stackrox/rox/migrator/migrations/n_29_to_n_30_postgres_network_entities"
	_ "github.com/stackrox/rox/migrator/migrations/n_30_to_n_31_postgres_network_flows"
	_ "github.com/stackrox/rox/migrator/migrations/n_31_to_n_32_postgres_network_graph_configs"
	_ "github.com/stackrox/rox/migrator/migrations/n_32_to_n_33_postgres_networkpolicies"
	_ "github.com/stackrox/rox/migrator/migrations/n_33_to_n_34_postgres_networkpoliciesundodeployments"
	_ "github.com/stackrox/rox/migrator/migrations/n_34_to_n_35_postgres_networkpolicyapplicationundorecords"
	_ "github.com/stackrox/rox/migrator/migrations/n_35_to_n_36_postgres_nodes"
	_ "github.com/stackrox/rox/migrator/migrations/n_36_to_n_37_postgres_notifiers"
	_ "github.com/stackrox/rox/migrator/migrations/n_37_to_n_38_postgres_permission_sets"
	_ "github.com/stackrox/rox/migrator/migrations/n_38_to_n_39_postgres_pods"
	_ "github.com/stackrox/rox/migrator/migrations/n_39_to_n_40_postgres_policies"
	_ "github.com/stackrox/rox/migrator/migrations/n_40_to_n_41_postgres_process_baseline_results"
	_ "github.com/stackrox/rox/migrator/migrations/n_41_to_n_42_postgres_process_baselines"
	_ "github.com/stackrox/rox/migrator/migrations/n_42_to_n_43_postgres_process_indicators"
	_ "github.com/stackrox/rox/migrator/migrations/n_43_to_n_44_postgres_report_configurations"
	_ "github.com/stackrox/rox/migrator/migrations/n_44_to_n_45_postgres_risks"
	_ "github.com/stackrox/rox/migrator/migrations/n_45_to_n_46_postgres_role_bindings"
	_ "github.com/stackrox/rox/migrator/migrations/n_46_to_n_47_postgres_roles"
	_ "github.com/stackrox/rox/migrator/migrations/n_47_to_n_48_postgres_secrets"
	_ "github.com/stackrox/rox/migrator/migrations/n_48_to_n_49_postgres_sensor_upgrade_configs"
	_ "github.com/stackrox/rox/migrator/migrations/n_49_to_n_50_postgres_service_accounts"
	_ "github.com/stackrox/rox/migrator/migrations/n_50_to_n_51_postgres_service_identities"
	_ "github.com/stackrox/rox/migrator/migrations/n_51_to_n_52_postgres_signature_integrations"
	_ "github.com/stackrox/rox/migrator/migrations/n_52_to_n_53_postgres_simple_access_scopes"
	_ "github.com/stackrox/rox/migrator/migrations/n_53_to_n_54_postgres_vulnerability_requests"
	_ "github.com/stackrox/rox/migrator/migrations/n_54_to_n_55_postgres_watched_images"
	_ "github.com/stackrox/rox/migrator/migrations/n_55_to_n_56_postgres_policy_categories"
	_ "github.com/stackrox/rox/migrator/migrations/n_56_to_n_57_postgres_groups"

	// Postgres -> Postgres migrations
	_ "github.com/stackrox/rox/migrator/migrations/m_168_to_m_169_postgres_remove_clustercve_permission"
	_ "github.com/stackrox/rox/migrator/migrations/m_169_to_m_170_collections_sac_resource_migration"
	_ "github.com/stackrox/rox/migrator/migrations/m_170_to_m_171_create_policy_categories_and_edges"
	_ "github.com/stackrox/rox/migrator/migrations/m_171_to_m_172_move_scope_to_collection_in_report_configurations"
	_ "github.com/stackrox/rox/migrator/migrations/m_172_to_m_173_network_flows_partition"
)
