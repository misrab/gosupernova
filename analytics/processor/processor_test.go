package processor

import (
	"testing"
	//"fmt"
)

func TestInferType(t *testing.T) {
	if (InferType(" ") != "empty") { t.Errorf("Failed for type empty") }

	if (InferType(" 234 ") != "float64") { t.Errorf("Failed for type float64") }

	if (InferType(" fds 3343") != "string") { t.Errorf("Failed for type string") }

	//if (InferType(" fds 3343") != "undefined") { t.Errorf("Failed for type undefined") }

	//if (InferType(" fds 3343") != "undefined") { t.Errorf("Failed for type undefined") }
}