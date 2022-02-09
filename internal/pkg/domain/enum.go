package domain

type ResolutionStrategy string

const (
	APPLY_LOCAL  ResolutionStrategy = "LOCAL"
	APPLY_REMOTE ResolutionStrategy = "REMOTE"
)
