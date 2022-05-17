package tests

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestGophkeeper(t *testing.T) {
	suite.Run(t, new(TestGophkeeperAuth))
	suite.Run(t, new(TestEndToEnd))
}
