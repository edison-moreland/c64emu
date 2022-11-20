package trace

import (
	"fmt"
	"strings"
)

type TraceEvent struct {
	System string
	Action string
	Info   []TraceInfo
}

func (te *TraceEvent) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%-10s --- %-15s", te.System, te.Action))

	for _, ti := range te.Info {
		sb.WriteString(fmt.Sprintf(" %s=%s", ti.Key, ti.Value))
	}

	return sb.String()
}

type TraceInfo struct {
	Key   string
	Value string
}

func String(key, value string) TraceInfo {
	return TraceInfo{Key: key, Value: value}
}

func Uint16(key string, value uint16) TraceInfo {
	return TraceInfo{Key: key, Value: fmt.Sprintf("$%04X", value)}
}

func Byte(key string, value byte) TraceInfo {
	return TraceInfo{Key: key, Value: fmt.Sprintf("$%02X", value)}
}

type Tracer struct {
	out chan TraceEvent
}

func NewTracer() *Tracer {
	return &Tracer{
		out: make(chan TraceEvent),
	}
}

func (t *Tracer) Close() error {
	close(t.out)
	return nil
}

func (t *Tracer) System(system string) *SystemTracer {
	return &SystemTracer{
		p:      t,
		system: system,
	}
}

func (t *Tracer) Out() <-chan TraceEvent {
	return t.out
}

type SystemTracer struct {
	p      *Tracer
	system string
}

func (st *SystemTracer) Trace(action string, info ...TraceInfo) {
	st.p.out <- TraceEvent{
		System: st.system,
		Action: action,
		Info:   info,
	}
}

func (st *SystemTracer) System(system string) *SystemTracer {
	return &SystemTracer{
		system: fmt.Sprintf("%s.%s", st.system, system),
		p:      st.p,
	}
}
