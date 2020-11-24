package Migration

type User struct {
	UserId    string `gorm:"column:UserId;type:uuid;primary_key;default:uuid_generate_v4()"`
	FirstName string `gorm:"column:FirstName;type:varchar(100)"`
	LastName  string `gorm:"column:LastName;type:varchar(100)"`
	Age       int    `gorm:"column:Age"`
	Active    int    `gorm:"column:Active"`
	CreatedAt string `gorm:"column:CreatedAt;type:timestamp"`
}
