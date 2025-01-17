package events

import (
	"encoding/json"
	"github.com/zwjlink/okex"
)

type (
	Basic struct {
		ID    string         `json:"id,omitempty"`
		Event string         `json:"event"`
		Code  int            `json:"code,omitempty,string"`
		Msg   string         `json:"msg,omitempty"`
		Op    okex.Operation `json:"op,omitempty"`
		Arg   *Argument      `json:"arg,omitempty"`
		Args  []*Argument    `json:"args,omitempty"`
		Data  []*Argument    `json:"data,omitempty"`
	}
	Argument struct {
		Arg        map[string]interface{}
		UntypedArg []interface{}
	}
	Success struct {
		Code int            `json:"code,omitempty,string"`
		Msg  string         `json:"msg,omitempty"`
		ID   string         `json:"id,omitempty"`
		Op   okex.Operation `json:"op,omitempty"`
		Data []*Argument    `json:"data,omitempty"`
	}
	Error struct {
		Event string         `json:"event,omitempty"`
		Msg   string         `json:"msg,omitempty"`
		Op    string         `json:"op,omitempty"`
		Code  okex.JSONInt64 `json:"code"`
		Args  []*Argument    `json:"args,omitempty"`
		Arg   *Argument      `json:"arg,omitempty"`
		Data  []*Argument    `json:"data,omitempty"`
		ID    string         `json:"id,omitempty"`
	}
	Login struct {
		Event string `json:"event"`
		Code  string `json:"code"`
		Msg   string `json:"msg"`
	}
	Subscribe struct {
		Event string    `json:"event"`
		Arg   *Argument `json:"arg"`
	}
	Unsubscribe struct {
		Event string    `json:"event"`
		Arg   *Argument `json:"arg"`
	}
)

func (a *Argument) Get(k string) (interface{}, bool) {
	v, ok := a.Arg[k]
	return v, ok
}

func (a *Argument) UnmarshalJSON(buf []byte) error {
	a.Arg = make(map[string]interface{})
	if json.Unmarshal(buf, &a.Arg) != nil {
		return json.Unmarshal(buf, &a.UntypedArg)
	}

	return nil
}
