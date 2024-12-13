package database

type Users struct {
	Id       int    `json:"id" gorm:"type:int;not null;unique;primaryKey;auto_increment"`
	Username string `json:"username" gorm:"type:varchar(40);not null"`
	Password string `json:"password" gorm:"type:varchar(100);not null"`
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
	Token    string `json:"token" gorm:"type:varchar(1000)"`
}

type CategoryBook struct {
	Id    int    `json:"id" gorm:"type:int;not null;unique;primaryKey;auto_increment"`
	Name  string `json:"name" gorm:"type:varchar(100);not null"`
	Books []Book `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE" json:"books"`
}

type Book struct {
	Id         int          `json:"id" gorm:"type:int;not null;unique;primaryKey;auto_increment"`
	Title      string       `json:"title" gorm:"type:varchar(100);not null"`
	Exercpt    string       `json:"exercpt" gorm:"type:varchar(100);not null"`
	Creator    string       `json:"creator" gorm:"type:varchar(100);not null"`
	Content    string       `json:"content" gorm:"type:text;not null"`
	CategoryID int          `json:"category_id" gorm:"type:int;not null"`
	Category   CategoryBook `gorm:"foreignKey:CategoryID" json:"category"`
}