package storage

type InMemoryStorage struct {
	data map[string]string
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		data: make(map[string]string),
	}
}

func (s *InMemoryStorage) Get(shortURL string) (string, bool) {
	originalURL, exists := s.data[shortURL]
	return originalURL, exists
}

func (s *InMemoryStorage) Save(shortURL, originalURL string) {
	s.data[shortURL] = originalURL
}
