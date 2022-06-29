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

func easyjson7c8bc0beDecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *PodStatus) {
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
		case "conditions":
			if in.IsNull() {
				in.Skip()
				out.Conditions = nil
			} else {
				in.Delim('[')
				if out.Conditions == nil {
					if !in.IsDelim(']') {
						out.Conditions = make([]*PodCondition, 0, 8)
					} else {
						out.Conditions = []*PodCondition{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *PodCondition
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(PodCondition)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Conditions = append(out.Conditions, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "containerStatuses":
			if in.IsNull() {
				in.Skip()
				out.ContainerStatuses = nil
			} else {
				in.Delim('[')
				if out.ContainerStatuses == nil {
					if !in.IsDelim(']') {
						out.ContainerStatuses = make([]*ContainerStatus, 0, 8)
					} else {
						out.ContainerStatuses = []*ContainerStatus{}
					}
				} else {
					out.ContainerStatuses = (out.ContainerStatuses)[:0]
				}
				for !in.IsDelim(']') {
					var v2 *ContainerStatus
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(ContainerStatus)
						}
						(*v2).UnmarshalEasyJSON(in)
					}
					out.ContainerStatuses = append(out.ContainerStatuses, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "ephemeralContainerStatuses":
			if in.IsNull() {
				in.Skip()
				out.EphemeralContainerStatuses = nil
			} else {
				in.Delim('[')
				if out.EphemeralContainerStatuses == nil {
					if !in.IsDelim(']') {
						out.EphemeralContainerStatuses = make([]*ContainerStatus, 0, 8)
					} else {
						out.EphemeralContainerStatuses = []*ContainerStatus{}
					}
				} else {
					out.EphemeralContainerStatuses = (out.EphemeralContainerStatuses)[:0]
				}
				for !in.IsDelim(']') {
					var v3 *ContainerStatus
					if in.IsNull() {
						in.Skip()
						v3 = nil
					} else {
						if v3 == nil {
							v3 = new(ContainerStatus)
						}
						(*v3).UnmarshalEasyJSON(in)
					}
					out.EphemeralContainerStatuses = append(out.EphemeralContainerStatuses, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "hostIP":
			out.HostIP = string(in.String())
		case "initContainerStatuses":
			if in.IsNull() {
				in.Skip()
				out.InitContainerStatuses = nil
			} else {
				in.Delim('[')
				if out.InitContainerStatuses == nil {
					if !in.IsDelim(']') {
						out.InitContainerStatuses = make([]*ContainerStatus, 0, 8)
					} else {
						out.InitContainerStatuses = []*ContainerStatus{}
					}
				} else {
					out.InitContainerStatuses = (out.InitContainerStatuses)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *ContainerStatus
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(ContainerStatus)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.InitContainerStatuses = append(out.InitContainerStatuses, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "message":
			out.Message = string(in.String())
		case "nominatedNodeName":
			out.NominatedNodeName = string(in.String())
		case "phase":
			out.Phase = string(in.String())
		case "podIP":
			out.PodIP = string(in.String())
		case "podIPs":
			if in.IsNull() {
				in.Skip()
				out.PodIPs = nil
			} else {
				in.Delim('[')
				if out.PodIPs == nil {
					if !in.IsDelim(']') {
						out.PodIPs = make([]*PodIP, 0, 8)
					} else {
						out.PodIPs = []*PodIP{}
					}
				} else {
					out.PodIPs = (out.PodIPs)[:0]
				}
				for !in.IsDelim(']') {
					var v5 *PodIP
					if in.IsNull() {
						in.Skip()
						v5 = nil
					} else {
						if v5 == nil {
							v5 = new(PodIP)
						}
						(*v5).UnmarshalEasyJSON(in)
					}
					out.PodIPs = append(out.PodIPs, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "qosClass":
			out.QosClass = string(in.String())
		case "reason":
			out.Reason = string(in.String())
		case "startTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.StartTime).UnmarshalJSON(data))
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
func easyjson7c8bc0beEncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in PodStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"conditions\":"
		out.RawString(prefix[1:])
		if in.Conditions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.Conditions {
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
		const prefix string = ",\"containerStatuses\":"
		out.RawString(prefix)
		if in.ContainerStatuses == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.ContainerStatuses {
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
	{
		const prefix string = ",\"ephemeralContainerStatuses\":"
		out.RawString(prefix)
		if in.EphemeralContainerStatuses == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v10, v11 := range in.EphemeralContainerStatuses {
				if v10 > 0 {
					out.RawByte(',')
				}
				if v11 == nil {
					out.RawString("null")
				} else {
					(*v11).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.HostIP != "" {
		const prefix string = ",\"hostIP\":"
		out.RawString(prefix)
		out.String(string(in.HostIP))
	}
	{
		const prefix string = ",\"initContainerStatuses\":"
		out.RawString(prefix)
		if in.InitContainerStatuses == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v12, v13 := range in.InitContainerStatuses {
				if v12 > 0 {
					out.RawByte(',')
				}
				if v13 == nil {
					out.RawString("null")
				} else {
					(*v13).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Message != "" {
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	if in.NominatedNodeName != "" {
		const prefix string = ",\"nominatedNodeName\":"
		out.RawString(prefix)
		out.String(string(in.NominatedNodeName))
	}
	if in.Phase != "" {
		const prefix string = ",\"phase\":"
		out.RawString(prefix)
		out.String(string(in.Phase))
	}
	if in.PodIP != "" {
		const prefix string = ",\"podIP\":"
		out.RawString(prefix)
		out.String(string(in.PodIP))
	}
	{
		const prefix string = ",\"podIPs\":"
		out.RawString(prefix)
		if in.PodIPs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.PodIPs {
				if v14 > 0 {
					out.RawByte(',')
				}
				if v15 == nil {
					out.RawString("null")
				} else {
					(*v15).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.QosClass != "" {
		const prefix string = ",\"qosClass\":"
		out.RawString(prefix)
		out.String(string(in.QosClass))
	}
	if in.Reason != "" {
		const prefix string = ",\"reason\":"
		out.RawString(prefix)
		out.String(string(in.Reason))
	}
	if true {
		const prefix string = ",\"startTime\":"
		out.RawString(prefix)
		out.Raw((in.StartTime).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PodStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7c8bc0beEncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PodStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7c8bc0beEncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PodStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7c8bc0beDecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PodStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7c8bc0beDecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
