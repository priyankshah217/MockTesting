package service

import (
	"MockTesting/http"
	"MockTesting/model"
	"encoding/json"
)

type JsonPlaceHolder struct {
	h http.Http
}

func NewJsonPlaceHolder(h http.Http) *JsonPlaceHolder {
	return &JsonPlaceHolder{h: h}
}

func (j JsonPlaceHolder) GetPosts(url string) ([]*model.PostResponse, error) {
	response, err := j.h.Get(url)
	if err != nil {
		return nil, err
	}
	var postResponseData []*model.PostResponse
	err = json.NewDecoder(response.Body).Decode(&postResponseData)
	if err != nil {
		return nil, err
	}
	return postResponseData, nil
}
