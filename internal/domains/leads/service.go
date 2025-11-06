package leads

type Service struct{ r *Repo }
func NewService(r *Repo) *Service { return &Service{r:r} }

func (s *Service) Create(l *Lead) error { return s.r.Create(l) }
