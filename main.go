package main

import (
	"fmt"
	"io"
	"os"

	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func main() {

	if err := RunPipeline(os.Stdin, os.Stdout); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

func RunPipeline(reader io.Reader, writer io.Writer) error {
	readWriter := &kio.ByteReadWriter{
		Reader: reader,
		Writer: writer,
	}
	pipeline := kio.Pipeline{
		Inputs:  []kio.Reader{readWriter},
		Filters: []kio.Filter{filter{readWriter: readWriter}},
		Outputs: []kio.Writer{readWriter},
	}

	return pipeline.Execute()
}

type filter struct {
	readWriter *kio.ByteReadWriter
}

type functionConfig struct {
	Data struct {
		Kind      string `yaml:"kind"`
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `yaml:"data"`
}

func (f filter) Filter(in []*yaml.RNode) ([]*yaml.RNode, error) {
	out := []*yaml.RNode{}

	marshalledFunctionConfig, err := f.readWriter.FunctionConfig.String()
	if err != nil {
		return out, err
	}

	var config functionConfig
	if err := yaml.Unmarshal([]byte(marshalledFunctionConfig), &config); err != nil {
		return out, err
	}

	for _, resource := range in {
		meta, err := resource.GetMeta()
		if err != nil {
			return out, err
		}

		if meta.Kind != config.Data.Kind || meta.Name != config.Data.Name || meta.Namespace != config.Data.Namespace {
			out = append(out, resource)
		}
	}

	return out, nil
}
