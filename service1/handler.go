package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"net/http"
	dProto "service1/proto"
	"strconv"
	"sync"
	"time"
)

type dataLoader struct {
	dProto.UnimplementedDataLoaderServer
	dbConn *sqlx.DB
}

func NewDataLoader(dbConn *sqlx.DB) *dataLoader{
	return &dataLoader{dbConn: dbConn}
}

func (dl *dataLoader) LoadData(ctx context.Context, in *dProto.EmptyParams) (*dProto.EmptyParams, error){

	now := time.Now()
	wg := sync.WaitGroup{}
	response := make([]AutoGenerated, 50)

	for i := 1;i <= 50;i++{
		url := "https://gorest.co.in/public/v1/posts?page="
		method := "GET"

		page := strconv.Itoa(i)

		client := &http.Client {}

		go func(j int) {
			wg.Add(1)
			defer wg.Done()

			req, err := http.NewRequest(method, url+page, nil)
			if err != nil {
				return
			}
			res, err := client.Do(req)
			if err != nil {
				return
			}
			defer res.Body.Close()

			if res.StatusCode != 200 {
				fmt.Println("failed on: ", url+page)
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return
			}

			err = json.Unmarshal(body, &response[j-1])
			if err != nil {
				return
			}

		}(i)

	}

	wg.Wait()

	fmt.Println("since: ", time.Since(now))

	query := `INSERT INTO data (id, user_id, title, body) VALUES (:id, :userID, :title, :body);`
	for i := range response {
		for j := range response[i].Data{
			queryArgs := map[string]interface{}{"id":response[i].Data[j].ID, "userID": response[i].Data[j].UserID, "title":response[i].Data[j].Title, "body":response[i].Data[j].Body}
			_, err := dl.dbConn.NamedExec(query, queryArgs)
			if err != nil{
				fmt.Println("in cycle: ", err.Error())
			}
		}
	}

	return &dProto.EmptyParams{}, nil
}