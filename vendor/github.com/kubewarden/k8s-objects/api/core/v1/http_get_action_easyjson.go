// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	intstr "github.com/kubewarden/k8s-objects/apimachinery/pkg/util/intstr"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonE1169c34DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *HTTPGetAction) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "host":
			out.Host = string(in.String())
		case "httpHeaders":
			if in.IsNull() {
				in.Skip()
				out.HTTPHeaders = nil
			} else {
				in.Delim('[')
				if out.HTTPHeaders == nil {
					if !in.IsDelim(']') {
						out.HTTPHeaders = make([]*HTTPHeader, 0, 8)
					} else {
						out.HTTPHeaders = []*HTTPHeader{}
					}
				} else {
					out.HTTPHeaders = (out.HTTPHeaders)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *HTTPHeader
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(HTTPHeader)
						}
						easyjsonE1169c34DecodeGithubComKubewardenK8sObjectsApiCoreV11(in, v1)
					}
					out.HTTPHeaders = append(out.HTTPHeaders, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "path":
			out.Path = string(in.String())
		case "port":
			if in.IsNull() {
				in.Skip()
				out.Port = nil
			} else {
				if out.Port == nil {
					out.Port = new(intstr.IntOrString)
				}
				*out.Port = intstr.IntOrString(in.String())
			}
		case "scheme":
			out.Scheme = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE1169c34EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in HTTPGetAction) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Host != "" {
		const prefix string = ",\"host\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Host))
	}
	{
		const prefix string = ",\"httpHeaders\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.HTTPHeaders == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.HTTPHeaders {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					easyjsonE1169c34EncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *v3)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Path != "" {
		const prefix string = ",\"path\":"
		out.RawString(prefix)
		out.String(string(in.Path))
	}
	{
		const prefix string = ",\"port\":"
		out.RawString(prefix)
		if in.Port == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Port))
		}
	}
	if in.Scheme != "" {
		const prefix string = ",\"scheme\":"
		out.RawString(prefix)
		out.String(string(in.Scheme))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HTTPGetAction) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE1169c34EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HTTPGetAction) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE1169c34EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HTTPGetAction) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE1169c34DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HTTPGetAction) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE1169c34DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
func easyjsonE1169c34DecodeGithubComKubewardenK8sObjectsApiCoreV11(in *jlexer.Lexer, out *HTTPHeader) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
			}
		case "value":
			if in.IsNull() {
				in.Skip()
				out.Value = nil
			} else {
				if out.Value == nil {
					out.Value = new(string)
				}
				*out.Value = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE1169c34EncodeGithubComKubewardenK8sObjectsApiCoreV11(out *jwriter.Writer, in HTTPHeader) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		if in.Value == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Value))
		}
	}
	out.RawByte('}')
}
