package audit

import "github.com/denisgrosek/changelock/internal/identity"

func EnsureDecisionHash(event Event) Event {
	if event.DecisionHash != "" {
		return event
	}

	event.DecisionHash = identity.DecisionHash(identity.DecisionInput{
		PolicyBundleHash: event.PolicyBundleHash,
		ImageDigest:      event.Digest,
		RequestID:        event.RequestID,
		Decision:         event.Decision,
		Component:        event.Component,
		Repo:             event.Repo,
		Environment:      event.Environment,
	})
	return event
}
