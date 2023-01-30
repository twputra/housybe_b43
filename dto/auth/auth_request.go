package authdto

type RegisterRequest struct {
	Name      string    `json:"name" gorm:"type: varchar(255)"`
	Username  string    `json:"username" gorm:"type: varchar(255)"`
	Email string `json:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"password" gorm:"type: varchar(255)"`
	Roles     string    `json:"roles" gorm:"type: varchar(255)"`
	Gender string `json:"gender" gorm:"type: varchar(255)"`
	Phone string `json:"phone" gorm:"type: varchar(255)"`
	Address string `json:"address" gorm:"type: varchar(255)"`
	Image      string    `json:"image" gorm:"type: varchar(255)"`
}

type LoginRequest struct {
	Username string `json:"username" gorm:"type: varchar(255)" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
}