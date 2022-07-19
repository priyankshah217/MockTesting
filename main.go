package main

import (
	"MockTesting/http"
	"MockTesting/service"
	"fmt"
)

func main() {
	//testServer := httptest.NewServer(h.HandlerFunc(func(res h.ResponseWriter, req *h.Request) {
	//	res.WriteHeader(h.StatusNotFound)
	//	res.Write([]byte("No Record Found"))
	//}))
	//defer func() { testServer.Close() }()
	//fmt.Println(testServer.URL)
	//req, err := h.NewRequest(h.MethodPost, testServer.URL, nil)
	//resp, err := h.DefaultClient.Do(req)
	//fmt.Println(resp.StatusCode)
	//bytes, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return
	//}
	//fmt.Println(string(bytes))
	//if err != nil {
	//	return
	//}
	//httpClient := http.NewHttp()
	//response, err := httpClient.Get("https://jsonplaceholder.typicode.com/posts")
	//if err != nil {
	//	return
	//}
	//fmt.Println(response[0].Title)
	//jsonPlaceHolder := service.NewJsonPlaceHolder()
	//posts, err := jsonPlaceHolder.GetPosts("https://jsonplaceholder.typicode.com/posts")
	//if err != nil {
	//	return
	//}
	//fmt.Println(posts[0].Title)
	customClient := http.NewCustomClient()
	jsonPlaceHolder := service.NewJsonPlaceHolder(customClient)
	getPosts, err := jsonPlaceHolder.GetPosts("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return
	}
	fmt.Println(getPosts[0].Title)
}
