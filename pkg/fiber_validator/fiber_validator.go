package fiber_validator

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en2 "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
)

func TranslateError(err error, validate *validator.Validate) map[string]string {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = en2.RegisterDefaultTranslations(validate, trans)
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)

	errs := make(map[string]string)

	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs[e.Field()] = translatedErr.Error()
	}
	return errs
}

func Validate(model interface{}, c *fiber.Ctx) (bool, error) {

	if err := c.BodyParser(model); err != nil {
		return false, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}

	validation := validator.New()

	err := validation.Struct(model)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
		}
		errs := TranslateError(err, validation)
		return false, c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
			"fields":  errs,
		})
	}

	return true, nil
}
