/*
 * Databricks
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package models

type ClustersCreateRequest struct {
	NumWorkers int32 `json:"num_workers,omitempty"`

	Autoscale *ClustersAutoScale `json:"autoscale,omitempty"`

	ClusterName string `json:"cluster_name,omitempty"`

	SparkVersion string `json:"spark_version"`

	SparkConf map[string]string `json:"spark_conf,omitempty"`

	AwsAttributes *ClustersAwsAttributes `json:"aws_attributes,omitempty"`

	NodeTypeId string `json:"node_type_id"`

	DriverNodeTypeId string `json:"driver_node_type_id,omitempty"`

	SshPublicKeys []string `json:"ssh_public_keys,omitempty"`

	CustomTags map[string]string `json:"custom_tags,omitempty"`

	ClusterLogConf *ClustersClusterLogConf `json:"cluster_log_conf,omitempty"`

	SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`

	AutoterminationMinutes int32 `json:"autotermination_minutes,omitempty"`

	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
}
