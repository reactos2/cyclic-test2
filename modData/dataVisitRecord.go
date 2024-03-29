package modData


type SystemKVData struct {
	ID   uint32 `gorm:"primaryKey;serial;column:id;autoIncrement"`
	Name string `gorm:"column:name"`
	ValueInt  int    `gorm:"column:ValueInt"`
	ValueString  string    `gorm:"column:ValueString"`
}