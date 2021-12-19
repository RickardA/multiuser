package app

func NewClient(rwyRepo runwayRepository, conflcitRepo conflictRepository) Client {
	return Client{
		runwayRepo:   rwyRepo,
		conflictRepo: conflcitRepo,
	}
}

type Client struct {
	runwayRepo   runwayRepository
	conflictRepo conflictRepository
}

type runwayRepository interface {
}

type conflictRepository interface {
}
