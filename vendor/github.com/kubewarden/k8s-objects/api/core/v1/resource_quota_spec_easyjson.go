// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	resource "github.com/kubewarden/k8s-objects/apimachinery/pkg/api/resource"
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

func easyjsonFc0a5813DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *ResourceQuotaSpec) {
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
		case "hard":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Hard = make(map[string]resource.Quantity)
				} else {
					out.Hard = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 resource.Quantity
					v1 = resource.Quantity(in.String())
					(out.Hard)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "scopeSelector":
			if in.IsNull() {
				in.Skip()
				out.ScopeSelector = nil
			} else {
				if out.ScopeSelector == nil {
					out.ScopeSelector = new(ScopeSelector)
				}
				easyjsonFc0a5813DecodeGithubComKubewardenK8sObjectsApiCoreV11(in, out.ScopeSelector)
			}
		case "scopes":
			if in.IsNull() {
				in.Skip()
				out.Scopes = nil
			} else {
				in.Delim('[')
				if out.Scopes == nil {
					if !in.IsDelim(']') {
						out.Scopes = make([]string, 0, 4)
					} else {
						out.Scopes = []string{}
					}
				} else {
					out.Scopes = (out.Scopes)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.Scopes = append(out.Scopes, v2)
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
func easyjsonFc0a5813EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in ResourceQuotaSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Hard) != 0 {
		const prefix string = ",\"hard\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('{')
			v3First := true
			for v3Name, v3Value := range in.Hard {
				if v3First {
					v3First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v3Name))
				out.RawByte(':')
				out.String(string(v3Value))
			}
			out.RawByte('}')
		}
	}
	if in.ScopeSelector != nil {
		const prefix string = ",\"scopeSelector\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonFc0a5813EncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *in.ScopeSelector)
	}
	{
		const prefix string = ",\"scopes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Scopes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v4, v5 := range in.Scopes {
				if v4 > 0 {
					out.RawByte(',')
				}
				out.String(string(v5))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResourceQuotaSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFc0a5813EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResourceQuotaSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFc0a5813EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResourceQuotaSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFc0a5813DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResourceQuotaSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFc0a5813DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
func easyjsonFc0a5813DecodeGithubComKubewardenK8sObjectsApiCoreV11(in *jlexer.Lexer, out *ScopeSelector) {
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
		case "matchExpressions":
			if in.IsNull() {
				in.Skip()
				out.MatchExpressions = nil
			} else {
				in.Delim('[')
				if out.MatchExpressions == nil {
					if !in.IsDelim(']') {
						out.MatchExpressions = make([]*ScopedResourceSelectorRequirement, 0, 8)
					} else {
						out.MatchExpressions = []*ScopedResourceSelectorRequirement{}
					}
				} else {
					out.MatchExpressions = (out.MatchExpressions)[:0]
				}
				for !in.IsDelim(']') {
					var v6 *ScopedResourceSelectorRequirement
					if in.IsNull() {
						in.Skip()
						v6 = nil
					} else {
						if v6 == nil {
							v6 = new(ScopedResourceSelectorRequirement)
						}
						easyjsonFc0a5813DecodeGithubComKubewardenK8sObjectsApiCoreV12(in, v6)
					}
					out.MatchExpressions = append(out.MatchExpressions, v6)
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
func easyjsonFc0a5813EncodeGithubComKubewardenK8sObjectsApiCoreV11(out *jwriter.Writer, in ScopeSelector) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"matchExpressions\":"
		out.RawString(prefix[1:])
		if in.MatchExpressions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v7, v8 := range in.MatchExpressions {
				if v7 > 0 {
					out.RawByte(',')
				}
				if v8 == nil {
					out.RawString("null")
				} else {
					easyjsonFc0a5813EncodeGithubComKubewardenK8sObjectsApiCoreV12(out, *v8)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonFc0a5813DecodeGithubComKubewardenK8sObjectsApiCoreV12(in *jlexer.Lexer, out *ScopedResourceSelectorRequirement) {
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
		case "operator":
			if in.IsNull() {
				in.Skip()
				out.Operator = nil
			} else {
				if out.Operator == nil {
					out.Operator = new(string)
				}
				*out.Operator = string(in.String())
			}
		case "scopeName":
			if in.IsNull() {
				in.Skip()
				out.ScopeName = nil
			} else {
				if out.ScopeName == nil {
					out.ScopeName = new(string)
				}
				*out.ScopeName = string(in.String())
			}
		case "values":
			if in.IsNull() {
				in.Skip()
				out.Values = nil
			} else {
				in.Delim('[')
				if out.Values == nil {
					if !in.IsDelim(']') {
						out.Values = make([]string, 0, 4)
					} else {
						out.Values = []string{}
					}
				} else {
					out.Values = (out.Values)[:0]
				}
				for !in.IsDelim(']') {
					var v9 string
					v9 = string(in.String())
					out.Values = append(out.Values, v9)
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
func easyjsonFc0a5813EncodeGithubComKubewardenK8sObjectsApiCoreV12(out *jwriter.Writer, in ScopedResourceSelectorRequirement) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"operator\":"
		out.RawString(prefix[1:])
		if in.Operator == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Operator))
		}
	}
	{
		const prefix string = ",\"scopeName\":"
		out.RawString(prefix)
		if in.ScopeName == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.ScopeName))
		}
	}
	{
		const prefix string = ",\"values\":"
		out.RawString(prefix)
		if in.Values == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v10, v11 := range in.Values {
				if v10 > 0 {
					out.RawByte(',')
				}
				out.String(string(v11))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
