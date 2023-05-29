package main

import (
	"text/template"

	"github.com/lithammer/dedent"
)

type AsConfig struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

type Metadata struct {
	Name string `yaml:"name"`
}

type Hosts struct {
	Name            string `yaml:"name"`
	Address         string `yaml:"address"`
	InternalAddress string `yaml:"internalAddress"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
}

type RoleGroups struct {
	Etcd         []string `yaml:"etcd"`
	ControlPlane []string `yaml:"control-plane"`
	Worker       []string `yaml:"worker"`
	Registry     string   `yaml:"registry"`
}

type ControlPlaneEndpoint struct {
	Domain  string `yaml:"domain"`
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type Kubernetes struct {
	Version          string `yaml:"version"`
	ClusterName      string `yaml:"clusterName"`
	AutoRenewCerts   bool   `yaml:"autoRenewCerts"`
	ContainerManager string `yaml:"containerManager"`
}

type Etcd struct {
	Type string `yaml:"type"`
}

type MultusCNI struct {
	Enabled bool `yaml:"enabled"`
}

type Network struct {
	Plugin          string    `yaml:"plugin"`
	KubePodsCIDR    string    `yaml:"kubePodsCIDR"`
	KubeServiceCIDR string    `yaml:"kubeServiceCIDR"`
	MultusCNI       MultusCNI `yaml:"multusCNI"`
}

type Registry struct {
	PrivateRegistry    string        `yaml:"privateRegistry"`
	NamespaceOverride  string        `yaml:"namespaceOverride"`
	RegistryMirrors    []interface{} `yaml:"registryMirrors"`
	InsecureRegistries []interface{} `yaml:"insecureRegistries"`
}

type Spec struct {
	Hosts                []Hosts              `yaml:"hosts"`
	RoleGroups           RoleGroups           `yaml:"roleGroups"`
	ControlPlaneEndpoint ControlPlaneEndpoint `yaml:"controlPlaneEndpoint"`
	Kubernetes           Kubernetes           `yaml:"kubernetes"`
	Etcd                 Etcd                 `yaml:"etcd"`
	Network              Network              `yaml:"network"`
	Registry             Registry             `yaml:"registry"`
	Addons               []interface{}        `yaml:"addons"`
}

var V321 = template.Must(template.New("v3.2.1").Parse(
	dedent.Dedent(`
---
apiVersion: installer.kubesphere.io/v1alpha1
kind: ClusterConfiguration
metadata:
  name: ks-installer
  namespace: kubesphere-system
  labels:
    version: {{ .Tag }}
spec:
  persistence:
    storageClass: ""
  authentication:
    jwtSecret: ""
  local_registry: ""
  namespace_override: ""
  # dev_tag: ""
  etcd:
    monitoring: false
    endpointIps: localhost
    port: 2379
    tlsEnable: true
  common:
    core:
      console:
        enableMultiLogin: true
        port: 30880
        type: NodePort
    # apiserver:
    #  resources: {}
    # controllerManager:
    #  resources: {}
    redis:
      enabled: false
      volumeSize: 2Gi
    openldap:
      enabled: false
      volumeSize: 2Gi
    minio:
      volumeSize: 20Gi
    monitoring:
      # type: external
      endpoint: http://prometheus-operated.kubesphere-monitoring-system.svc:9090
      GPUMonitoring:
        enabled: false
    gpu:
      kinds:         
      - resourceName: "nvidia.com/gpu"
        resourceType: "GPU"
        default: true
    es:
      # master:
      #   volumeSize: 4Gi
      #   replicas: 1
      #   resources: {}
      # data:
      #   volumeSize: 20Gi
      #   replicas: 1
      #   resources: {}
      logMaxAge: 7
      elkPrefix: logstash
      basicAuth:
        enabled: false
        username: ""
        password: ""
      externalElasticsearchHost: ""
      externalElasticsearchPort: ""
  alerting:
    enabled: false
    # thanosruler:
    #   replicas: 1
    #   resources: {}
  auditing:
    enabled: false
    # operator:
    #   resources: {}
    # webhook:
    #   resources: {}
  devops:
    enabled: false
    jenkinsMemoryLim: 2Gi
    jenkinsMemoryReq: 1500Mi
    jenkinsVolumeSize: 8Gi
    jenkinsJavaOpts_Xms: 512m
    jenkinsJavaOpts_Xmx: 512m
    jenkinsJavaOpts_MaxRAM: 2g
  events:
    enabled: false
    # operator:
    #   resources: {}
    # exporter:
    #   resources: {}
    # ruler:
    #   enabled: true
    #   replicas: 2
    #   resources: {}
  logging:
    enabled: false
    containerruntime: docker
    logsidecar:
      enabled: true
      replicas: 2
      # resources: {}
  metrics_server:
    enabled: false
  monitoring:
    storageClass: ""
    # kube_rbac_proxy:
    #   resources: {}
    # kube_state_metrics:
    #   resources: {}
    # prometheus:
    #   replicas: 1
    #   volumeSize: 20Gi
    #   resources: {}
    #   operator:
    #     resources: {}
    #   adapter:
    #     resources: {}
    # node_exporter:
    #   resources: {}
    # alertmanager:
    #   replicas: 1
    #   resources: {}
    # notification_manager:
    #   resources: {}
    #   operator:
    #     resources: {}
    #   proxy:
    #     resources: {}
    gpu:
      nvidia_dcgm_exporter:
        enabled: false
        # resources: {}
  multicluster:
    clusterRole: none 
  network:
    networkpolicy:
      enabled: false
    ippool:
      type: none
    topology:
      type: none
  openpitrix:
    store:
      enabled: false
  servicemesh:
    enabled: false
  kubeedge:
    enabled: false   
    cloudCore:
      nodeSelector: {"node-role.kubernetes.io/worker": ""}
      tolerations: []
      cloudhubPort: "10000"
      cloudhubQuicPort: "10001"
      cloudhubHttpsPort: "10002"
      cloudstreamPort: "10003"
      tunnelPort: "10004"
      cloudHub:
        advertiseAddress:
          - ""
        nodeLimit: "100"
      service:
        cloudhubNodePort: "30000"
        cloudhubQuicNodePort: "30001"
        cloudhubHttpsNodePort: "30002"
        cloudstreamNodePort: "30003"
        tunnelNodePort: "30004"
    edgeWatcher:
      nodeSelector: {"node-role.kubernetes.io/worker": ""}
      tolerations: []
      edgeWatcherAgent:
        nodeSelector: {"node-role.kubernetes.io/worker": ""}
        tolerations: []
`)))
