// models/girl_model.go
package models

type GirlModel struct {
	ID      int
	Name    string     `gorm:"size:16"`
	BoyList []BoyModel `gorm:"foreignKey:GirlID"`
}

type BoyModel struct {
	ID        int
	Name      string `gorm:"size:16"`
	GirlID    int
	GirlModel GirlModel `gorm:"foreignKey:GirlID"`
}
