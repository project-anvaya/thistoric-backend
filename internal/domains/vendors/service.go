package vendors

type Service struct{ r *Repo }
func NewService(r *Repo) *Service { return &Service{r:r} }

func (s *Service) List(q, cat, loc string, limit, offset int) ([]Vendor, error) {
	return s.r.List(q, cat, loc, limit, offset)
}
func (s *Service) Get(id uint) (*Vendor, error) { return s.r.Get(id) }
