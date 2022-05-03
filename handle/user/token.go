package user

import (
	"time"

	"github.com/ez-deploy/ezdeploy/handle/db"
	"github.com/ez-deploy/ezdeploy/models"
	"github.com/thanhpk/randstr"
	"github.com/wuhuizuo/sqlm"
)

const (
	// keep token alive for a week.
	defaultSessionTokenExpireTime = time.Hour * 24 * 7

	defaultTokenLength = 128
)

// TokenOperationImpl impl restapi.UserOperation and restapi.Authable interface.
type TokenOperationImpl struct {
	Tables *db.Tables
}

func (t *TokenOperationImpl) deleteTokenByID(id int64) error {
	return t.Tables.Token.Delete(sqlm.SelectorFilter{"id": id})
}

func (t *TokenOperationImpl) getTokenByValue(value string) (*models.Token, error) {
	selector := sqlm.SelectorFilter{"value": value}

	retToken := &models.Token{}
	if err := t.Tables.Token.Get(selector, retToken); err != nil {
		return nil, err
	}

	return retToken, nil
}

func (t *TokenOperationImpl) createToken(userID int64, tokenType string) (*models.Token, error) {
	newToken := &models.Token{
		UserID:    userID,
		Type:      tokenType,
		Value:     randstr.Base64(defaultTokenLength),
		CreateAt:  time.Now().Unix(),
		UpdateAt:  time.Now().Unix(),
		ExpiredAt: time.Now().Add(defaultSessionTokenExpireTime).Unix(),
	}

	id, err := t.Tables.Token.Insert(newToken)
	if err != nil {
		return nil, err
	}

	newToken.ID = id

	return newToken, nil
}
