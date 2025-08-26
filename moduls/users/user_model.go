package users

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey;auto_increment:true;index"`
	Name  string `json:"name" gorm:"varchar(255)"`
	Email string `json:"email" gorm:"varchar(255);unique"`
}


