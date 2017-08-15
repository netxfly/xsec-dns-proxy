package models_test

import (
	"testing"
	"xsec-dns-server/models"
)

func TestNewDbEngine(t *testing.T) {
	t.Log(models.NewDbEngine())
}
