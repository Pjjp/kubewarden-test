// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
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

func easyjsonE5b093eaDecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *EndpointSubset) {
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
		case "addresses":
			if in.IsNull() {
				in.Skip()
				out.Addresses = nil
			} else {
				in.Delim('[')
				if out.Addresses == nil {
					if !in.IsDelim(']') {
						out.Addresses = make([]*EndpointAddress, 0, 8)
					} else {
						out.Addresses = []*EndpointAddress{}
					}
				} else {
					out.Addresses = (out.Addresses)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *EndpointAddress
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(EndpointAddress)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Addresses = append(out.Addresses, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "notReadyAddresses":
			if in.IsNull() {
				in.Skip()
				out.NotReadyAddresses = nil
			} else {
				in.Delim('[')
				if out.NotReadyAddresses == nil {
					if !in.IsDelim(']') {
						out.NotReadyAddresses = make([]*EndpointAddress, 0, 8)
					} else {
						out.NotReadyAddresses = []*EndpointAddress{}
					}
				} else {
					out.NotReadyAddresses = (out.NotReadyAddresses)[:0]
				}
				for !in.IsDelim(']') {
					var v2 *EndpointAddress
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(EndpointAddress)
						}
						(*v2).UnmarshalEasyJSON(in)
					}
					out.NotReadyAddresses = append(out.NotReadyAddresses, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "ports":
			if in.IsNull() {
				in.Skip()
				out.Ports = nil
			} else {
				in.Delim('[')
				if out.Ports == nil {
					if !in.IsDelim(']') {
						out.Ports = make([]*EndpointPort, 0, 8)
					} else {
						out.Ports = []*EndpointPort{}
					}
				} else {
					out.Ports = (out.Ports)[:0]
				}
				for !in.IsDelim(']') {
					var v3 *EndpointPort
					if in.IsNull() {
						in.Skip()
						v3 = nil
					} else {
						if v3 == nil {
							v3 = new(EndpointPort)
						}
						(*v3).UnmarshalEasyJSON(in)
					}
					out.Ports = append(out.Ports, v3)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonE5b093eaEncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in EndpointSubset) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"addresses\":"
		out.RawString(prefix[1:])
		if in.Addresses == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v4, v5 := range in.Addresses {
				if v4 > 0 {
					out.RawByte(',')
				}
				if v5 == nil {
					out.RawString("null")
				} else {
					(*v5).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"notReadyAddresses\":"
		out.RawString(prefix)
		if in.NotReadyAddresses == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.NotReadyAddresses {
				if v6 > 0 {
					out.RawByte(',')
				}
				if v7 == nil {
					out.RawString("null")
				} else {
					(*v7).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"ports\":"
		out.RawString(prefix)
		if in.Ports == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Ports {
				if v8 > 0 {
					out.RawByte(',')
				}
				if v9 == nil {
					out.RawString("null")
				} else {
					(*v9).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EndpointSubset) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE5b093eaEncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EndpointSubset) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE5b093eaEncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EndpointSubset) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE5b093eaDecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EndpointSubset) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE5b093eaDecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
