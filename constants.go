/*
modification history
--------------------
2018/7/26, by lovejoy, create
*/

package tfjob

const (
	// EnvKubeflowNamespace is ENV for kubeflow namespace specified by user.
	EnvKubeflowNamespace = "KUBEFLOW_NAMESPACE"

	// DefaultPortName is name of the port used to communicate between PS and
	// workers.
	DefaultPortName = "tfjob-port"
	// DefaultContainerName is the name of the TFJob container.
	DefaultContainerName = "tensorflow"
	// DefaultPort is default value of the port.
	DefaultPort = 2222
	// DefaultRestartPolicy is default RestartPolicy for TFReplicaSpec.
	DefaultRestartPolicy = RestartPolicyNever
)
