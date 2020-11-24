package Env

import "os"

// SetDevEnv is Init Env For Dev
func SetDevEnv() {
	os.Setenv("ENV", "dev")
	os.Setenv("PG_USER", "postgres")
	os.Setenv("PG_HOST", "localhost")
	os.Setenv("PG_PORT", "5433")
	os.Setenv("PG_DATABASENAME", "dev_learning")
	os.Setenv("PG_PASSWORD", "abc123")
}
