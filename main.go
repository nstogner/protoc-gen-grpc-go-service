package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/nstogner/kit/log"
)

func main() {
	encodeResponse(
		generateResponse(
			parseRequest(
				decodeRequest(os.Stdin),
			),
		),
		os.Stdout,
	)
}

// decodeRequest unmarshals the protobuf request.
func decodeRequest(r io.Reader) *plugin.CodeGeneratorRequest {
	var req plugin.CodeGeneratorRequest
	input, err := ioutil.ReadAll(r)
	if err != nil {
		log.WithErr(err).Fatal("unable to read stdin")
	}
	if err := proto.Unmarshal(input, &req); err != nil {
		log.WithErr(err).Fatal("unable to marshal stdin as protobuf")
	}
	return &req
}

// parseRequest wrangles the request to fit needs of the template.
func parseRequest(req *plugin.CodeGeneratorRequest) params {
	pf := req.GetProtoFile()[0]
	ps := params{
		ServiceDescriptorProto: *pf.GetService()[0],
		PackageName:            pf.GetPackage(),
		ProtoName:              pf.GetName(),
	}

	for _, mtd := range ps.ServiceDescriptorProto.GetMethod() {
		m := method{
			MethodDescriptorProto: *mtd,
			serviceName:           ps.ServiceDescriptorProto.GetName(),
		}
		ps.Methods = append(ps.Methods, m)
	}
	return ps
}

// generateResponse executes the template.
func generateResponse(ps params) *plugin.CodeGeneratorResponse {
	w := &bytes.Buffer{}
	if err := tmpl.Execute(w, ps); err != nil {
		log.WithErr(err).Fatal("unable to execute template")
	}

	fmted, err := format.Source([]byte(w.String()))
	if err != nil {
		log.WithErr(err).Fatal("unable to go-fmt output")
	}

	fileName := "service.go"
	fileContent := string(fmted)

	return &plugin.CodeGeneratorResponse{
		File: []*plugin.CodeGeneratorResponse_File{
			{
				Name:    &fileName,
				Content: &fileContent,
			},
		},
	}
}

// encodeResponse marshals the protobuf response.
func encodeResponse(resp *plugin.CodeGeneratorResponse, w io.Writer) {
	outBytes, err := proto.Marshal(resp)
	if err != nil {
		log.WithErr(err).Fatal("unable to marshal response to protobuf")
	}

	if _, err := w.Write(outBytes); err != nil {
		log.WithErr(err).Fatal("unable to write protobuf to stdout")
	}
}

// params is the data provided to the template.
type params struct {
	descriptor.ServiceDescriptorProto
	ProtoName   string
	PackageName string
	Methods     []method
}

type method struct {
	descriptor.MethodDescriptorProto
	serviceName string
}

// The following methods are used by the template.
func (m method) TrimmedInput() string {
	return strings.TrimPrefix(m.GetInputType(), ".")
}
func (m method) TrimmedOutput() string {
	return strings.TrimPrefix(m.GetOutputType(), ".")
}
func (m method) StreamName() string {
	return fmt.Sprintf("%s_%sServer", m.serviceName, m.GetName())
}

var tmpl = template.Must(template.New("server").Parse(`
// Code generated by protoc-gen-service
// source: {{.ProtoName}}
// DO NOT EDIT!

package main

type service struct{}

{{ range .Methods }}
	{{ if .GetClientStreaming }}
		{{ if .GetServerStreaming }}
// {{.Name}} streams outputs and listens to a stream of inputs.
func (s *service) {{.Name}}(stream {{$.PackageName}}.{{.StreamName}}) error {
	for {
		input, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// TODO: Do something with input
		_ = input

		// TODO: Stream some meaningful output
		if err := stream.Send(&{{.TrimmedOutput}}{}); err != nil {
			return err
		}
	}

	return nil
}
		{{ else }}
// {{.Name}} sends a single output for a streamed input.
func (s *service) {{.Name}}(stream {{$.PackageName}}.{{.StreamName}}) error {
	for {
		input, err := stream.Recv()
		if err == io.EOF {
			// TODO: Send some meaningful output
			return stream.SendAndClose(&{{.TrimmedOutput}}{})
		}
		if err != nil {
			return err
		}

		// TODO: Do something with the input message
		_ = input
	}

	return nil
}
		{{ end }}
	{{ else }}
		{{ if .GetServerStreaming }}
// {{.Name}} streams output for a single input.
func (s *service) {{.Name}}(input *{{.TrimmedInput}}, stream {{$.PackageName}}.{{.StreamName}}) error {
	// TODO: Do something with the input
	_ = input

	// TODO: Stream some meaningful output
	for i := 0; i < 10; i++ {
		if err := stream.Send(&{{.TrimmedOutput}}{}); err != nil {
			return err
		}
	}

	return nil
}
		{{ else }}
// {{.Name}} sends a single output for a single input.
func (s *service) {{.Name}}(ctx context.Context, input *{{.TrimmedInput}}) (*{{.TrimmedOutput}}, error) {
	// TODO: Do something with the input
	_ = input

	// TODO: Send some meaningful output
	return &{{.TrimmedOutput}}{}, nil
}
		{{ end }}
	{{ end }}

{{ end }}
`))
