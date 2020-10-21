package config

// Parser create DB connect path
func Parser(DBUser string, DBPass string, DBProt string, DBName string) string {
	connect := DBUser + ":" + DBPass + "@" + DBProt + "/" + DBName + "?charset=utf8&parseTime=True&loc=Local"
	return connect
}
