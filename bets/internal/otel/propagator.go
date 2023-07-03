//go:generate go run -mod=mod github.com/abice/go-enum@v0.5.6 --nocase

package otel

import (
	"fmt"
	"os"
	"strings"

	xraypropagator "go.opentelemetry.io/contrib/propagators/aws/xray"
	b3propagator "go.opentelemetry.io/contrib/propagators/b3"
	jaegerpropagator "go.opentelemetry.io/contrib/propagators/jaeger"
	otpropagator "go.opentelemetry.io/contrib/propagators/ot"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

// ENUM(tracecontext, baggage, b3, b3multi, jaeger, xray, ottrace)
type Propagator string

func setupPropagators() error {
	env, ok := os.LookupEnv(envPropagators)
	if !ok {
		env = "tracecontext,baggage"
	}

	envs := strings.Split(env, ",")
	propagators := make([]propagation.TextMapPropagator, 0, len(envs))

	for i := range envs {
		propagator, err := ParsePropagator(strings.TrimSpace(envs[i]))
		if err != nil {
			return fmt.Errorf("failed to set the trace propagator: %w", err)
		}

		switch propagator {
		case PropagatorTracecontext:
			propagators = append(propagators, propagation.TraceContext{})
		case PropagatorBaggage:
			propagators = append(propagators, propagation.Baggage{})
		case PropagatorB3:
			propagators = append(propagators, b3propagator.New(b3propagator.WithInjectEncoding(b3propagator.B3SingleHeader)))
		case PropagatorB3multi:
			propagators = append(propagators, b3propagator.New(b3propagator.WithInjectEncoding(b3propagator.B3MultipleHeader)))
		case PropagatorJaeger:
			propagators = append(propagators, jaegerpropagator.Jaeger{})
		case PropagatorXray:
			propagators = append(propagators, xraypropagator.Propagator{})
		case PropagatorOttrace:
			propagators = append(propagators, otpropagator.OT{})
		}
	}

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagators...))

	return nil
}
