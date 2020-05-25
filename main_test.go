package main

import (
	"bytes"
	"strings"
	"testing"

	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

type ResourceList struct {
	Items []resource `yaml:"items"`
}

type resource struct {
	Kind     string   `yaml:"kind"`
	Metadata metadata `yaml:"metadata"`
}

type metadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

func TestFilterReturnsErrorWhenFailToParseFunctionConfig(t *testing.T) {
	readWriter := &kio.ByteReadWriter{
		FunctionConfig: yaml.NewRNode(&yaml.Node{
			Value: ";",
		}),
	}
	f := filter{
		readWriter: readWriter,
	}

	_, err := f.Filter([]*yaml.RNode{})
	if err == nil {
		t.Fatal("expected an error")
	}
}

func TestFilterRemovesResourceMatchingNameNamespaceAndKind(t *testing.T) {
	documents := `
kind: ResourceList
functionConfig:
  data:
    kind: Deployment
    name: matching-name
    namespace: matching-namespace
items:
  - kind: Deployment
    metadata:
      name: matching-name
      namespace: matching-namespace
  - kind: Deployment
    metadata:
      name: different-name
      namespace: matching-namespace
  - kind: Deployment
    metadata:
      name: matching-name
      namespace: different-namespace
  - kind: DaemonSet
    metadata:
      name: matching-name
      namespace: matching-namespace
`

	writer := &bytes.Buffer{}
	err := RunPipeline(strings.NewReader(documents), writer)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	output := string(writer.Bytes())
	expectedMissingDocument := "kind: Deployment\nmetadata:\n  name: matching-name\n  namespace: matching-namespace"
	if strings.Contains(output, expectedMissingDocument) {
		t.Errorf("expected %s to be removed, but found in output %s", expectedMissingDocument, output)
	}

	var resourceList ResourceList
	if err = yaml.Unmarshal(writer.Bytes(), &resourceList); err != nil {
		t.Fatalf("unexpected error while unmarshalling resourceList: %s", err)
	}

	if len(resourceList.Items) != 3 {
		t.Fatalf("expected 3 resources to remain, but had %v", resourceList.Items)
	}

	for _, resource := range resourceList.Items {
		if resource.Kind == "Deployment" && resource.Metadata.Name == "matching-name" && resource.Metadata.Namespace == "matching-namespace" {
			t.Errorf("expected resource to be removed, but it was not")
		}
	}
}
