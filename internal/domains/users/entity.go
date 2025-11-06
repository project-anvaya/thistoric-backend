package users

type User struct {
	ID    string `json:"id"`    // Supabase auth UID
	Email string `json:"email"` // from JWT
}
