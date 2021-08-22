package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRepository_NewLink_error(t *testing.T) {
	_, err := NewLink("", "shortned", time.Now())
	assert.NotNil(t, err, "NewLink didn't return error on empty url")

	_, err = NewLink("url", "", time.Now())
	assert.NotNil(t, err, "NewLink didn't return error on empty shortned")

	_, err = NewLink("url", "shortned", time.Time{})
	assert.NotNil(t, err, "NewLink didn't return error on empty created")
}

func TestRepository_NewLink_valid(t *testing.T) {
	createTime := time.Now().Round(time.Hour)
	l, err := NewLink("url", "shortned", createTime)
	assert.Nil(t, err, "NewLink returned error on valid input")

	assert.Equal(t, "url", l.Url, "wrong Url in Link")
	assert.Equal(t, "shortned", l.Shortned, "wrong Shortned in Link")
	assert.Equal(t, createTime, l.Created, "wrong Created in Link")
}
