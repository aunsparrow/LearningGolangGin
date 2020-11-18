package Env

import "os"

// SetDevEnv is Init Env For Dev
func SetDevEnv() {
	os.Setenv("ENV", "dev")
	os.Setenv("PG_USER", "root")
	os.Setenv("PG_HOST", "26257")
	os.Setenv("PG_PORT", "203.151.50.153")
	os.Setenv("PG_DATABASENAME", "qa_irecruit")
	os.Setenv("PG_PASSWORD", "1412")
}
