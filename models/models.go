package models

type URL struct {
	ID        string `json:"id" gorm:"primaryKey"`
	TargetURL string `json:"target_url" gorm:"not null"`
	ShortURL  string `json:"short_url" gorm:"unique;not null"`
	Clicked   uint64 `json:"clicked"`
	Random    bool   `json:"random"`
}
