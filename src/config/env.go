package config

import "os"

func ConfigENV() {
	os.Setenv("PORT", "8081")
	os.Setenv("CONTEXT_PATH", "/feedback-processor-api")
	// DB
	os.Setenv("DB_HOST", "unisun.dynu.com")
	os.Setenv("DB_NAME", "unisuncmsdevdb")
	os.Setenv("DB_USER", "urquhmotrdhwqg")
	os.Setenv("DB_PASS", "efad4bb2169e67ddaa17c21aba5c76efc6a9daa6a06310949eba9a006bf258da")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_SSL", "disable")
	os.Setenv("DB_TIMEZONE", "Asia/Bangkok")
}
