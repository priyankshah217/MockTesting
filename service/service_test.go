package service

import (
	"MockTesting/http"
	mockHttp "MockTesting/mock"
	"MockTesting/model"
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	h "net/http"
	"net/http/httptest"
	"testing"
)

const URL = "https://jsonplaceholder.typicode.com/posts"

var mockResponses = []model.PostResponse{
	{
		0,
		0,
		"Test Response",
		"Test Body",
	}, {
		1,
		1,
		"Test Response 1",
		"Test Body 1",
	},
}

type MockHttpClient struct{}

func (m MockHttpClient) Get(_ string) (*h.Response, error) {
	b, _ := json.Marshal(mockResponses)
	return &h.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBuffer(b)),
	}, nil
}

type MockHttpClientTestify struct {
	mock.Mock
}

func (m *MockHttpClientTestify) Get(url string) (*h.Response, error) {
	arguments := m.Called(url)
	return arguments.Get(0).(*h.Response), arguments.Error(1)
}

type CustomTestSuite struct {
	suite.Suite
}

func TestCustomTestSuites(t *testing.T) {
	suite.Run(t, new(CustomTestSuite))
}

// Mock using gomock/mockgen
func (t *CustomTestSuite) TestJsonPlaceHolder_GetPosts_Using_Gmock() {
	controller := gomock.NewController(t.T())
	defer controller.Finish()
	newMockHttp := mockHttp.NewMockHttp(controller)
	b, err := json.Marshal(mockResponses)
	if err != nil {
		return
	}
	newMockHttp.EXPECT().Get(gomock.Any()).Return(&h.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBuffer(b)),
	}, nil)
	jsonPlaceHolder := NewJsonPlaceHolder(newMockHttp)
	getPosts, err := jsonPlaceHolder.GetPosts(URL)
	t.Require().Nil(err, "error should not be nil")
	t.Require().Len(getPosts, 2, "length of posts should be 2")
	t.Require().Equal(getPosts[0].Id, 0, "id should be 0")
}

// Mock using httpTestServer
func (t *CustomTestSuite) TestJsonPlaceHolder_GetPosts_Using_HttpTest() {
	testServer := httptest.NewServer(h.HandlerFunc(func(res h.ResponseWriter, req *h.Request) {
		byteResponse, err := json.Marshal(mockResponses)
		if err != nil {
			return
		}
		res.WriteHeader(h.StatusOK)
		_, _ = res.Write(byteResponse)
	}))
	defer func() { testServer.Close() }()
	customClient := http.NewCustomClient()
	jsonPlaceHolder := NewJsonPlaceHolder(customClient)
	getPosts, err := jsonPlaceHolder.GetPosts(testServer.URL)
	t.Require().Nil(err, "error should not be nil")
	t.Require().Len(getPosts, 2, "length of posts should be 2")
	t.Require().Equal(getPosts[0].Id, 0, "id should be 0")
}

// Mock using Testify
func (t *CustomTestSuite) TestJsonPlaceHolder_GetPosts_Using_Testify() {
	m := &MockHttpClientTestify{}
	b, _ := json.Marshal(mockResponses)
	m.On("Get", URL).Return(&h.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBuffer(b)),
	}, nil)
	jsonPlaceHolder := NewJsonPlaceHolder(m)
	getPosts, err := jsonPlaceHolder.GetPosts(URL)
	t.Require().Nil(err, "error should not be nil")
	t.Require().Len(getPosts, 2, "length of posts should be 2")
	t.Require().Equal(getPosts[0].Id, 0, "id should be 0")
}

// Mock using custom client
func (t *CustomTestSuite) TestJsonPlaceHolder_GetPosts_Using_CustomClient() {
	m := &MockHttpClient{}
	jsonPlaceHolder := NewJsonPlaceHolder(m)
	getPosts, err := jsonPlaceHolder.GetPosts(URL)
	t.Require().Nil(err, "error should not be nil")
	t.Require().Len(getPosts, 2, "length of posts should be 2")
	t.Require().Equal(getPosts[0].Id, 0, "id should be 0")
}
