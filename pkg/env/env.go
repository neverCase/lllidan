package env

import (
	"fmt"
	"k8s.io/klog/v2"
	"os"
	"os/exec"
	"strings"
)

const (
	NodeName          = "NODE_NAME"
	HostIP            = "HOST_IP"
	PodName           = "POD_NAME"
	PodNamespace      = "POD_NAMESPACE"
	PodIP             = "POD_IP"
	PodServiceAccount = "POD_SERVICE_ACCOUNT"
)

// getHostName gets the hostname of the host machine if the container is started by docker run --net=host
func getHostName() (string, error) {
	cmd := exec.Command("/bin/hostname")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	hostname := strings.TrimSpace(string(out))
	if hostname == "" {
		return "", fmt.Errorf("no hostname get from cmd '/bin/hostname' in the container, please check")
	}
	return hostname, nil
}

// GetHostName gets the hostname of host machine
func GetHostName() (string, error) {
	hostName := os.Getenv("HOST_NAME")
	if hostName != "" {
		return hostName, nil
	}
	klog.Info("get HOST_NAME from env failed, is env.(\"HOST_NAME\") already set? Will use hostname instead")
	return getHostName()
}

// GetPodIp returns the ip which has been allocated to the pod in the k8s cluster
func GetPodIP() (string, error) {
	podIp := os.Getenv(PodIP)
	if podIp != "" {
		return podIp, nil
	}
	return "", fmt.Errorf("no podIp get from '%s' in the container", PodIP)
}
