package module

func ShowMyFile(username string) []MyFile {
	var my []MyFile
	DB.Table("my_files").Where("username=?",username).Find(&my)
	return my
}
func InsertFile(f *MyFile) error {
	return DB.Create(f).Error
}