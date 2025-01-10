package messages

// LoginRequestMessages contains custom messages for LoginRequest validation
var Login = map[string]string{
	"Email.required":    "Email is required and cannot be empty",
	"Email.email":       "Email must be a valid email address",
	"Password.required": "Password is required",
	"Password.min":      "Password must be at least 6 characters long",
}

var GetUser = map[string]string{
	"ID.required": "The user ID is required and cannot be empty",
}
