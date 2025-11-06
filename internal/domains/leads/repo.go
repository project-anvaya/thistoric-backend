package leads

import "gorm.io/gorm"

type Repo struct{ db *gorm.DB }
func NewRepo(db *gorm.DB) *Repo { return &Repo{db: db} }

func (r *Repo) Create(l *Lead) error { return r.db.Create(l).Error }
