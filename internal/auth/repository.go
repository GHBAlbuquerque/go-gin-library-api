package auth

// ClientRepository queries the InMemoryStore to validate the existence of a client.
type ClientRepository interface {
	Validate(clientID, clientSecret string) bool
}

// InMemoryClientRepo stores credentials (id and secret) in memory.
type InMemoryClientRepo struct {
	credentials map[string]string // id & secret
}

func NewInMemoryClientRepo(seed map[string]string) *InMemoryClientRepo {
	return &InMemoryClientRepo{credentials: seed}
}

// Validate returns if requesting client (clientID and clientSecret) exist in the database.
func (r *InMemoryClientRepo) Validate(clientID, clientSecret string) bool {
	secret, exists := r.credentials[clientID]
	if !exists {
		return false
	}

	if secret != clientSecret {
		return false
	}

	return true
}
