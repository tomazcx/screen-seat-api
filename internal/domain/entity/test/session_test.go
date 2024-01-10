package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tomazcx/screen-seat-api/internal/domain/entity"
)

func TestCreateNewSession(t *testing.T) {
	testDate := time.Now()
	testRoom := "A1"
	isSubtitled := false
	movie := entity.Movie{}	

	session := entity.NewSession(movie, testDate, testRoom, isSubtitled)

	assert.NotNil(t, session.ID)
	assert.Equal(t, session.Room, testRoom)
	assert.Equal(t, session.DateTime, testDate)
	assert.Equal(t, session.IsSubtitled, isSubtitled)
	assert.Equal(t, session.Movie, movie)
}	

