package models

type Profile struct {
	ID           int    `gorm:"primaryKey"`
	Name         string `gorm:"size:100"`
	Summary      string `gorm:"type:text"`
	Skills       string `gorm:"type:text"`
	ResumeLink   string `gorm:"size:255"`
	LinkedinLink string `gorm:"size:255"`
	GithubLink   string `gorm:"size:255"`
	ProfileImage string `gorm:"size:255"`
	CreatedAt    string `gorm:"autoCreateTime"`
	UpdatedAt    string `gorm:"autoUpdateTime"`
}

type ProfileRepositoryImpl interface {
	GetProfile(id string) (Profile, error)
}
