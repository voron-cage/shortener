package shortener

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"shortener/storage"
)

type Service struct {
	tntStorage storage.TNTStorage
	logger     *zap.Logger
}

func NewService(tntCfg *storage.TarantoolConfig) *Service {
	return &Service{
		tntStorage: storage.NewTarantoolStorage(tntCfg),
		logger:     zap.L(),
	}
}

func (s *Service) CreateShortLink(c *gin.Context) {
	userData, ok := c.Get("user")
	if !ok {
		c.Status(http.StatusUnauthorized)
	}
	user := userData.(*storage.UserTnt)
	fmt.Println(user)
	url := URLRequest{}
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	linkID, err := s.tntStorage.GetLinkSeqIDNext()
	if err != nil {
		s.logger.Error("failed to get sequence from links tarantool space", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}
	tntLink := storage.LinkTnt{
		ID:          linkID,
		UserID:      user.ID,
		OriginalURL: url.URL,
		ShortURL:    getShortURLFromInt(user.ID) + getShortURLFromInt(linkID),
	}

	err = s.tntStorage.InsertLink(&tntLink)
	if err != nil {
		s.logger.Error("failed to insert link into tarantool space", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}
	uri := buildAbsoluteURI(c.Request.TLS != nil, c.Request.Host, tntLink.ShortURL)
	c.JSON(http.StatusOK, map[string]string{"short_url": uri})
}

func (s *Service) RedirectLink(c *gin.Context) {
	shortURL, ok := c.Params.Get("shortURL")
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}
	tntLink, err := s.tntStorage.GetByShortURL(shortURL)
	if err != nil {
		if err == storage.NotFoundError {
			c.Status(http.StatusNotFound)
			return
		}
		s.logger.Error("failed to fetch by short_url from tarantool", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Redirect(http.StatusMovedPermanently, tntLink.OriginalURL)
}
