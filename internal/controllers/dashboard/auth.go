package dashboard

import (
	"go-boilerplate/internal/dto/request"
	"go-boilerplate/internal/dto/response"
	"go-boilerplate/internal/middleware"
	"go-boilerplate/internal/usecase"
	"go-boilerplate/internal/utils"
	"go-boilerplate/internal/utils/errors"
	"go-boilerplate/pkg/logger"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct {
	AuthUsecase usecase.AuthUseCase
}

func NewAuthController(
	authUC usecase.AuthUseCase,
) *AuthController {
	return &AuthController{AuthUsecase: authUC}
}

func (ctrl *AuthController) LogoutDashboard(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return utils.SetResponseBadRequest(c, "Missing Authorization header", nil)
	}

	tokenString := utils.ExtractBearerToken(authHeader)

	// Hapus token dari Redis
	err := middleware.DeleteTokenFromRedis(c.Context(), tokenString)
	if err != nil {
		logger.Error("Failed to delete token from Redis", err)
		return utils.SetResponseInternalServerError(c, "Failed to logout", err)
	}

	return utils.SetResponseOK(c, "Successfully logged out", nil)
}

func (ctrl *AuthController) ValidateCredentials(c *fiber.Ctx) error {
	var reqLogin request.ReqLogin
	if err := c.BodyParser(&reqLogin); err != nil {
		logger.Error("Failed to parse login request", err)
		return utils.SetResponseBadRequest(c, errors.ErrMessageInvalidRequestData, err)
	}

	err := reqLogin.ValidateRequest()
	if err != nil {
		return utils.SetResponseBadRequest(c, "Invalid username or password", err)
	}

	// Validasi kredensial
	user, err := ctrl.AuthUsecase.LoginDashboard(c.Context(), &reqLogin)
	if err != nil {
		logger.Error("error LoginDashboard", err)
		return utils.SetResponseBadRequest(c, "Login Failed, Invalid username or password", err)
	}

	// Generate temporary token
	temporaryToken, err := middleware.GenerateTemporaryToken(user)
	if err != nil {
		logger.Error("Failed to generate temporary token", err)
		return utils.SetResponseInternalServerError(c, "Failed generate token", err)
	}

	return utils.SetResponseOK(c, "Temporary token generated", response.ResAuth{Token: temporaryToken})
}

func (ctrl *AuthController) GenerateAccessToken(c *fiber.Ctx) error {
	var (
		reqToken request.ReqToken
		err      error
	)
	if err = c.BodyParser(&reqToken); err != nil {
		logger.Error("Failed to parse token request", err)
		return utils.SetResponseBadRequest(c, errors.ErrMessageInvalidRequestData, err)
	}

	err = reqToken.ValidateRequest()
	if err != nil {
		return utils.SetResponseBadRequest(c, errors.ErrMessageInvalidRequestData, err)
	}

	// Validasi temporary token
	user, err := middleware.ValidateTemporaryToken(reqToken.TemporaryToken)
	if err != nil {
		logger.Error("Invalid temporary token", err)
		return utils.SetResponseUnauthorized(c, "Invalid or expired temporary token", err.Error())
	}

	// Generate access token
	user, err = ctrl.AuthUsecase.LoginByUserId(c.Context(), user.ID)
	if err != nil {
		logger.Error("Error GetUserByID", err)
		return utils.SetResponseUnauthorized(c, errors.ErrMessageDataNotFound, err.Error())
	}

	accessToken, err := middleware.GenerateTokenUserDashboard(user)
	if err != nil {
		logger.Error("Failed to generate access token", err)
		return utils.SetResponseInternalServerError(c, "Failed to generate access token", err)
	}

	// Simpan access token di Redis
	claims := jwt.MapClaims{}
	_, _, _ = new(jwt.Parser).ParseUnverified(accessToken, claims)
	exp := time.Unix(int64(claims["exp"].(float64)), 0)

	err = middleware.SaveTokenToRedis(c.Context(), accessToken, exp)
	if err != nil {
		logger.Error("Failed to save access token to Redis", err)
		return utils.SetResponseInternalServerError(c, "Failed to save access token", err)
	}

	return utils.SetResponseOK(c, "Access token generated", response.ResAuth{Token: accessToken})
}
