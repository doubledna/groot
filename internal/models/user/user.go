package user

type User struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Username string `gorm:"not null;unique" json:"username"`
	Password string `gorm:"not null" json:"password"`
}
