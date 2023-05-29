package main

import (
	"bufio"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	asconfig := AsConfig{}
	asconfig.APIVersion = "kubekey.kubesphere.io/v1alpha2"
	asconfig.Metadata.Name = "sample"
	asconfig.Spec.ControlPlaneEndpoint.Domain = "lb.kubesphere.local"
	asconfig.Spec.ControlPlaneEndpoint.Port = 6443
	asconfig.Spec.ControlPlaneEndpoint.Address = ""
	asconfig.Spec.Kubernetes.Version = "v1.21.5"
	asconfig.Spec.Kubernetes.ClusterName = "cluster.local"
	asconfig.Spec.Kubernetes.ContainerManager = "docker"
	asconfig.Spec.Kubernetes.AutoRenewCerts = true
	asconfig.Spec.Etcd.Type = "kubekey"
	asconfig.Spec.Network.Plugin = "calico"
	asconfig.Spec.Network.KubePodsCIDR = "10.233.64.0/18"
	asconfig.Spec.Network.KubeServiceCIDR = "10.233.0.0/18"
	asconfig.Spec.Network.MultusCNI.Enabled = false
	asconfig.Spec.Registry.PrivateRegistry = ""
	asconfig.Spec.Registry.NamespaceOverride = "kubesphere"
	conf, err := yaml.Marshal(&asconfig)
	if err != nil {
		panic(err.Error())
	}

	// fmt.Println(string(conf))

	file, err := os.OpenFile("config.yaml", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}
	write := bufio.NewWriter(file)
	write.Write(conf)
	write.Flush()
	file.Close()

	file2, err := os.OpenFile("config.yaml", os.O_APPEND, 0666)
	V321.Execute(file2, map[string]interface{}{"Tag": "v3.2.1"})
	write.Flush()
	file2.Close()
}
