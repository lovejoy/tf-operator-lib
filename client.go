package tfjob

import (
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"lib/k8s"
)

type KubeflowV1alpha2Interface interface {
	RESTClient() rest.Interface
	TFJobsGetter
}

// KubeflowV1alpha2Client is used to interact with features provided by the kubeflow.org group.
type KubeflowV1alpha2Client struct {
	restClient rest.Interface
}

func (c *KubeflowV1alpha2Client) TFJobs(namespace string) TFJobInterface {
	return newTFJobs(c, namespace)
}

// NewForConfig creates a new KubeflowV1alpha2Client for the given config.
func NewForConfig(c *rest.Config) (*KubeflowV1alpha2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &KubeflowV1alpha2Client{client}, nil
}

// NewForConfigOrDie creates a new KubeflowV1alpha2Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *KubeflowV1alpha2Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new KubeflowV1alpha2Client for the given RESTClient.
func New(c rest.Interface) *KubeflowV1alpha2Client {
	return &KubeflowV1alpha2Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *KubeflowV1alpha2Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

var TfjobClientSetMap map[string]*KubeflowV1alpha2Client

func GetTfjobClientByidc(idc string) (*KubeflowV1alpha2Client, error) {
	if _, ok := TfjobClientSetMap[idc]; ok {
		return TfjobClientSetMap[idc], nil
	}

	config, _ := clientcmd.BuildConfigFromFlags(k8s.GetMasterUrlAndKubeConfig(idc))
	client, err := NewForConfig(config)
	if err == nil {
		if len(TfjobClientSetMap) == 0 {
			TfjobClientSetMap = map[string]*KubeflowV1alpha2Client{}
		}
		TfjobClientSetMap[idc] = client
	}
	return TfjobClientSetMap[idc], err
}
