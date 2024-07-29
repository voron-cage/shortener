package auth

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"shortener/storage"
)

type Service struct {
	tntStorage storage.TNTStorage
	secretKey  string
	logger     *zap.Logger
}

func NewService(secretKey string, tntCfg *storage.TarantoolConfig) *Service {
	return &Service{
		tntStorage: storage.NewTarantoolStorage(tntCfg),
		secretKey:  secretKey,
		logger:     zap.L(),
	}
}

func (s *Service) Register(c *gin.Context) {
	regReq := RegisterUserRequest{}
	if err := c.ShouldBindJSON(&regReq); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if err := regReq.Validate(); err != nil {
		c.JSON(err.Status, err)
		return
	}
	user, err := s.tntStorage.GetUser(regReq.Email)
	if err != nil && err != storage.NotFoundError {
		s.logger.Error("failed to fetch user from tarantool", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}
	if user != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"email": "user already exist"})
		return
	}

	userID, err := s.tntStorage.GetAuthUserIDSeqNext()
	if err != nil {
		s.logger.Error("failed to fetch sequence from auth_user space", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}
	tntUser := storage.UserTnt{
		ID:           userID,
		Email:        regReq.Email,
		Username:     regReq.Username,
		HashPassword: getHashPassword(regReq.Password, s.secretKey),
	}
	if err := s.tntStorage.InsertUser(&tntUser); err != nil {
		s.logger.Error("failed to insert user into tarantool", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.Status(201)
}

func (s *Service) Login(c *gin.Context) {
	logReq := LoginUserRequest{}
	if err := c.ShouldBindJSON(&logReq); err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	user, err := s.tntStorage.GetUser(logReq.Email)
	if err != nil {
		if err == storage.NotFoundError {
			c.JSON(http.StatusOK, map[string]string{"error": "wrong email or password"})
			return
		}
		s.logger.Error("failed to fetch user from tarantool", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	if compareHashPassword(logReq.Password, s.secretKey, user.HashPassword) {
		token, err := obtainJWTToken(user.Email, s.secretKey)
		if err != nil {
			s.logger.Error("failed to obtain jwt token", zap.Error(err))
		}
		c.JSON(http.StatusOK, map[string]string{"access": token})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"error": "wrong email or password"})
}
