package config

import "os"

type Config struct {
	Port         string
	PostgresDSN  string
	SupabaseJWKS string
}

func Load() Config {
	return Config{
		Port:         getenv("PORT", "8080"),
		PostgresDSN:  must("POSTGRES_DSN"),
		SupabaseJWKS: must("SUPABASE_JWKS_URL"),
	}
}

func must(k string) string {
	v := os.Getenv(k)
	if v == "" { panic("missing env: " + k) }
	return v
}
func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" { return v }
	return def
}
