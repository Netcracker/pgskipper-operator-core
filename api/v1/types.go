// Copyright 2024-2025 NetCracker Technology Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +kubebuilder:object:generate=true
// +groupName=qubership.org

package types

import (
	v1 "k8s.io/api/core/v1"
)

type BackupDaemon struct {
	Resources              *v1.ResourceRequirements `json:"resources,omitempty"`
	DockerImage            string                   `json:"image,omitempty"`
	Affinity               v1.Affinity              `json:"affinity,omitempty"`
	Storage                Storage                  `json:"storage,omitempty"`
	PgHost                 string                   `json:"pgHost,omitempty"`
	EvictionPolicy         string                   `json:"evictionPolicy,omitempty"`
	BackupSchedule         string                   `json:"backupSchedule,omitempty"`
	GranularEviction       string                   `json:"granularEviction,omitempty"`
	JobFlag                string                   `json:"jobFlag,omitempty"`
	ConnectTimeout         string                   `json:"connectTimeout,omitempty"`
	GranularBackupSchedule string                   `json:"granularBackupSchedule,omitempty"`
	DatabasesToSchedule    string                   `json:"databasesToSchedule,omitempty"`
	WalArchiving           bool                     `json:"walArchiving,omitempty"`
	AllowPrefix            bool                     `json:"allowPrefix,omitempty"`
	ExcludedExtensions     string                   `json:"excludedExtensions,omitempty"`
	UseEvictionPolicyFirst string                   `json:"useEvictionPolicyFirst,omitempty"`
	EvictionBinaryPolicy   string                   `json:"evictionBinaryPolicy,omitempty"`
	ArchiveEvictionPolicy  string                   `json:"archiveEvictionPolicy,omitempty"`
	SecurityContext        v1.PodSecurityContext    `json:"securityContext,omitempty"`
	PriorityClassName      string                   `json:"priorityClassName,omitempty"`
	S3Storage              *S3Storage               `json:"s3Storage,omitempty"`
	PodLabels              map[string]string        `json:"podLabels,omitempty"`
	ExternalPv             *ExternalPv              `json:"externalPv,omitempty"`
	SslMode                string                   `json:"sslMode,omitempty"`
}

type MetricCollector struct {
	Resources             *v1.ResourceRequirements `json:"resources,omitempty"`
	DockerImage           string                   `json:"image,omitempty"`
	Affinity              v1.Affinity              `json:"affinity,omitempty"`
	InfluxDbHost          string                   `json:"influxDbHost,omitempty"`
	InfluxDatabase        string                   `json:"influxDatabase,omitempty"`
	MetricsProfile        string                   `json:"metricsProfile,omitempty"`
	CollectionInterval    int                      `json:"collectionInterval,omitempty"`
	SecurityContext       v1.PodSecurityContext    `json:"securityContext,omitempty"`
	TelegrafPluginTimeout int                      `json:"telegrafPluginTimeout,omitempty"`
	DevMetricsTimeout     int                      `json:"devMetricsTimeout,omitempty"`
	DevMetricsInterval    int                      `json:"devMetricsInterval,omitempty"`
	PriorityClassName     string                   `json:"priorityClassName,omitempty"`
	OcExecTimeout         int                      `json:"ocExecTimeout,omitempty"`
	PodLabels             map[string]string        `json:"podLabels,omitempty"`
	SslMode               string                   `json:"sslMode,omitempty"`
}

// Vault DbEngine configuration
type DbEngine struct {
	Enabled               bool   `json:"enabled,omitempty"`
	Name                  string `json:"name,omitempty"`
	MaxOpenConnections    int    `json:"maxOpenConnections,omitempty"`
	MaxIdleConnections    int    `json:"maxIdleConnections,omitempty"`
	MaxConnectionLifetime string `json:"maxConnectionLifetime,omitempty"`
}

// Storage Describes Storage that will be used by patroni
type Storage struct {
	// +kubebuilder:validation:Pattern=`^[0-9]+(m|Ki|Mi|Gi|Ti|Pi|Ei|k|M|G|T|P|E)$`
	Size         string   `json:"size,omitempty"`
	Type         string   `json:"type,omitempty"`
	StorageClass string   `json:"storageClass,omitempty"`
	Volumes      []string `json:"volumes,omitempty"`
	Nodes        []string `json:"nodes,omitempty"`
	Selectors    []string `json:"selectors,omitempty"`
	AccessModes  []string `json:"accessModes,omitempty"`
}

type VaultRegistration struct {
	DockerImage string   `json:"dockerImage,omitempty"`
	Enabled     bool     `json:"enabled,omitempty"`
	Path        string   `json:"path,omitempty"`
	Url         string   `json:"url,omitempty"`
	Role        string   `json:"role,omitempty"`
	Method      string   `json:"method,omitempty"`
	Token       string   `json:"token,omitempty"`
	DbEngine    DbEngine `json:"dbEngine,omitempty"`
}

type ConsulRegistration struct {
	CheckInterval   string            `json:"checkInterval,omitempty"`
	CheckTimeout    string            `json:"checkTimeout,omitempty"`
	DeregisterAfter string            `json:"deregisterAfter,omitempty"`
	Host            string            `json:"host,omitempty"`
	ServiceName     string            `json:"serviceName,omitempty"`
	Meta            map[string]string `json:"meta,omitempty"`
	Tags            []string          `json:"tags,omitempty"`
	LeaderMeta      map[string]string `json:"leaderMeta,omitempty"`
	LeaderTags      []string          `json:"leaderTags,omitempty"`
}

type CloudSql struct {
	Project        string `json:"project,omitempty"`
	Instance       string `json:"instance,omitempty"`
	AuthSecretName string `json:"authSecretName,omitempty"`
}

type S3Storage struct {
	Url             string `json:"url,omitempty"`
	AccessKeyId     string `json:"accessKeyId,omitempty"`
	SecretAccessKey string `json:"secretAccessKey,omitempty"`
	Bucket          string `json:"bucket,omitempty"`
	Prefix          string `json:"prefix,omitempty"`
	UntrustedCert   bool   `json:"untrustedCert,omitempty"`
	Region          string `json:"region,omitempty"`
}

type ExternalPv struct {
	Name         string `json:"name,omitempty"`
	Capacity     string `json:"capacity,omitempty"`
	StorageClass string `json:"storageClass,omitempty"`
}
