package global

type Model struct {
	ID        uint           `json:"ID" gorm:"primarykey"`
	CreatedAt time.Time      `json:"CreatedAt"`
	UpdatedAt time.Time      `json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}