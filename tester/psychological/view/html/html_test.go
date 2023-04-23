package html

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type HtmlTestSuite struct {
	suite.Suite
}

func (s *HtmlTestSuite) TestResult() {
	p := ResultParams{
		Title:   "설문 결과",
		Message: fmt.Sprintf("당신은 %s 타입입니다!", "착한"),
	}

	output, err := Result(p)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), *output)
	assert.Contains(s.T(), *output, "<!DOCTYPE HTML>")
	assert.Contains(s.T(), *output, "당신은 착한 타입입니다!")
}

func TestHtmlTestSuite(t *testing.T) {
	suite.Run(t, new(HtmlTestSuite))
}
