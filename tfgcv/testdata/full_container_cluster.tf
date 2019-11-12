resource "google_container_cluster" "full_list_default_1" {
  # TODO: test beta fields.
  # addons_config.cloudrun_config
  # addons_config.cloudrun_config
  # addons_config.cloudrun_config
  # addons_config.istio_config
  # authenticator_groups_config
  # cluster_autoscaling
  # database_encryption
  # enable_binary_authorization
  # enable_intranode_visibility
  # enable_shielded_nodes
  # enable_tpu
  # node_config.workload_metadata_config
  # node_config.sandbox_config
  # pod_security_policy_config
  # release_channel
  # resource_usage_export_config
  # vertical_pod_autoscaling
  # workload_identity_config

  # TODO: need a seperate test case for fields in conflict with others.
  # ip_allocation_policy.cluster_ipv4_cidr_block
  # ip_allocation_policy.create_subnetwork
  # ip_allocation_policy.node_ipv4_cidr_block
  # ip_allocation_policy.services_ipv4_cidr_block
  # ip_allocation_policy.subnetwork_name
  # region
  # zone

  # TODO: problematic fields
  # node_pool

  name = "test-cluster"

  additional_zones = ["test-additional_zones"]
  addons_config {
    horizontal_pod_autoscaling { disabled = true }
    http_load_balancing { disabled = true }
    kubernetes_dashboard { disabled = true }
    network_policy_config { disabled = true }
  }
  cluster_ipv4_cidr         = "10.0.0.42/32"
  default_max_pods_per_node = 42
  description               = "test-description"
  enable_kubernetes_alpha   = true
  enable_legacy_abac        = true
  initial_node_count        = 42
  ip_allocation_policy {
    use_ip_aliases                = true
    cluster_secondary_range_name  = "test-cluster_secondary_range_name"
    services_secondary_range_name = "test-services_secondary_range_name"
  }
  location        = "test-location"
  logging_service = "logging.googleapis.com"
  maintenance_policy {
    daily_maintenance_window {
      start_time = "03:00"
    }
  }
  master_auth {
    username = "test-username"
    password = "test-password"
    client_certificate_config {
      issue_client_certificate = true
    }
  }
  master_authorized_networks_config {
    cidr_blocks {
      cidr_block   = "10.0.0.42/32"
      display_name = "test-display_name1"
    }
    cidr_blocks {
      cidr_block   = "10.0.1.42/32"
      display_name = "test-display_name2"
    }
  }
  min_master_version = "test-min_master_version"
  monitoring_service = "monitoring.googleapis.com"
  network            = "test-network"
  network_policy {
    provider = "CALICO"
    enabled  = true
  }
  node_config {
    disk_size_gb = 42
    disk_type    = "pd-standard"
    guest_accelerator {
      type  = "test-type1"
      count = 1
    }
    guest_accelerator {
      type  = "test-type2"
      count = 1
    }
    image_type = "test-image_type"
    labels = {
      test-label-key = "test-label-value"
    }
    local_ssd_count = 42
    machine_type    = "test-machine_type"
    metadata = {
      test-metadata-key = "test-metadata-value"
    }
    min_cpu_platform = "test-min_cpu_platform"
    oauth_scopes     = ["test-oauth_scopes", "storage-ro"]
    preemptible      = true
    service_account  = "test-service_account"
    tags             = ["test-tags"]
    taint {
      key    = "test-key1"
      value  = "test-value1"
      effect = "NO_SCHEDULE"
    }
    taint {
      key    = "test-key2"
      value  = "test-value2"
      effect = "NO_SCHEDULE"
    }
  }
  node_locations = ["test-node_locations"]
  node_version   = "test-node_version"
  private_cluster_config {
    enable_private_endpoint = true
    enable_private_nodes    = true
    master_ipv4_cidr_block  = "127.0.0.0/28"
  }
  remove_default_node_pool = true
  resource_labels = {
    test-resource_labels-key = "test-resource_labels-value"
  }
  subnetwork = "test-subnetwork"
}