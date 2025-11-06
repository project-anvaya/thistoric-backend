package vendors

import (
	"strings"
	"gorm.io/gorm"
)

type Repo struct{ db *gorm.DB }
func NewRepo(db *gorm.DB) *Repo { return &Repo{db: db} }

func (r *Repo) List(q, category, location string, limit, offset int) ([]Vendor, error) {
	var out []Vendor
	tx := r.db.Model(&Vendor{})
	if category != "" { tx = tx.Where("LOWER(category) = ?", strings.ToLower(category)) }
	if location != "" { tx = tx.Where("LOWER(location) = ?", strings.ToLower(location)) }
	if q != "" {
		q = "%" + strings.ToLower(strings.TrimSpace(q)) + "%"
		tx = tx.Where("LOWER(name) LIKE ? OR LOWER(description) LIKE ?", q, q)
	}
	err := tx.Order("id DESC").Limit(limit).Offset(offset).Find(&out).Error
	return out, err
}

func (r *Repo) Get(id uint) (*Vendor, error) {
	var v Vendor
	if err := r.db.First(&v, id).Error; err != nil { return nil, err }
	return &v, nil
}
