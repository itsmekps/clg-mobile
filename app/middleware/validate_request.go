package middleware

import (
	"reflect"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ValidationError returns a structured JSON validation error response
func ValidationError(c *fiber.Ctx, field, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Validation failed",
		"details": []fiber.Map{
			{
				"field": field,
				"error": message,
			},
		},
	})
}

// ValidateRequestDTO dynamically validates request body and parameters based on the type struct
func ValidateRequestDTO(validate *validator.Validate, dto interface{}, customMessages map[string]string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Create a new instance of the DTO for each request
		newDTO := reflect.New(reflect.TypeOf(dto).Elem()).Interface()
		value := reflect.ValueOf(newDTO).Elem()

		for i := 0; i < value.NumField(); i++ {
			field := value.Type().Field(i)
			fieldValue := value.Field(i)

			// Check for param or JSON tags
			paramTag := field.Tag.Get("param")
			jsonTag := field.Tag.Get("json")

			if paramTag != "" {
				// Get the parameter value as a string
				paramValue := c.Params(paramTag)
				if paramValue == "" {
					return ValidationError(c, paramTag, "Parameter is required")
				}

				// Handle conversion based on the expected field type
				switch field.Type {
				case reflect.TypeOf(primitive.ObjectID{}): // Handle MongoDB ObjectID
					objID, err := primitive.ObjectIDFromHex(paramValue)
					if err != nil {
						return ValidationError(c, paramTag, "Parameter must be a valid ID")
					}
					fieldValue.Set(reflect.ValueOf(objID))
					c.Locals(paramTag, objID) // Store the converted ObjectID in Locals

				default:
					// Handle other primitive types (e.g., int, string)
					switch field.Type.Kind() {
					case reflect.Int: // Integer fields
						intValue, err := strconv.Atoi(paramValue)
						if err != nil {
							return ValidationError(c, paramTag, "Parameter must be an integer")
						}
						fieldValue.SetInt(int64(intValue))

					case reflect.String: // String fields
						fieldValue.SetString(paramValue)

					default: // Unsupported types
						return ValidationError(c, paramTag, "Unsupported parameter type")
					}
				}
			}

			if jsonTag != "" && (c.Method() == fiber.MethodPost || c.Method() == fiber.MethodPatch || c.Method() == fiber.MethodPut) {
				// Parse request body for body-based fields
				if err := c.BodyParser(newDTO); err != nil {
					return ValidationError(c, "body", "Invalid request payload")
				}
			}
		}

		// Validate the DTO using struct tags
		if err := validate.Struct(newDTO); err != nil {
			var details []struct {
				Field string `json:"field"`
				Error string `json:"error"`
			}

			for _, err := range err.(validator.ValidationErrors) {

				// Construct the key for the custom message map
				customMsgKey := err.Field() + "." + err.Tag()
				// Get the custom message if available, fallback to default message
				msg := customMessages[customMsgKey]
				if msg == "" {
					msg = "Invalid " + err.Field()
				}

				details = append(details, struct {
					Field string `json:"field"`
					Error string `json:"error"`
				}{
					Field: err.Field(),
					Error: msg,
				})
			}

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "Validation failed",
				"details": details,
			})
		}

		c.Locals("validatedRequest", newDTO)
		return c.Next()
	}
}
