package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var layout = "2006/01/02 03:04:05"
var client *http.Client
var testServer *httptest.Server
var initTestPortfolios = [2]Portfolio{
	{
		ID: 0, Title: "title_" + time.Now().Format(layout), UID: "dummy-token", Author: "author",
		Abstruct: "", Readme: "readme", Source: "source", Link: "link", CreatedAt: time.Now(), UpdatedAt: time.Now(),
	},
	{
		ID: 0, Title: "puttitle_" + time.Now().Format(layout), UID: "dummy-token", Author: "putauthor",
		Abstruct: "", Readme: "putreadme", Source: "putsource", Link: "putlink", CreatedAt: time.Now(), UpdatedAt: time.Now(),
	},
}

var searchedPortfolios = []Portfolio{
	{
		ID: 0, Title: "searchedtitle_" + time.Now().Format(layout), UID: "dummy-token", Author: "searchedAuthor",
		Abstruct: "", Readme: "searched readme", Source: "source", Link: "link", CreatedAt: time.Now(), UpdatedAt: time.Now(),
	},
	{
		ID: 0, Title: "searchedtitle2_" + time.Now().Format(layout), UID: "dummy-token", Author: "searchedAuthor",
		Abstruct: "", Readme: "searched readme2", Source: "source", Link: "link", CreatedAt: time.Now(), UpdatedAt: time.Now(),
	},
	{
		ID: 0, Title: "searchedtitle3_" + time.Now().Format(layout), UID: "dummy-token", Author: "searchedAuthor2",
		Abstruct: "", Readme: "searched readme3", Source: "source", Link: "link", CreatedAt: time.Now(), UpdatedAt: time.Now(),
	},
	{
		ID: 0, Title: "hogehogetitle_" + time.Now().Format(layout), UID: "dummy-token", Author: "searchedAuthor2",
		Abstruct: "", Readme: "hogehoge readme", Source: "source", Link: "link", CreatedAt: time.Now(), UpdatedAt: time.Now(),
	},
}

// intにはsearchedPortfoliosの何番目が取得できるかを２進数でセットする
var testCase = map[string]map[string]int{
	"title":  map[string]int{"searchedtitle": 7},    //0111
	"author": map[string]int{"searchedAuthor2": 12}, //1100
	"readme": map[string]int{"hogehoge": 8},         //1000
}

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	dir, _ := os.Getwd()
	filePath := strings.Replace(dir, "server", "db/migrations", 1)
	Migration("file:" + filePath)
	client = new(http.Client)
	midd := &MockMiddleware{}
	engine := setup(midd)
	testServer = httptest.NewServer(engine)
	defer testServer.Close()

	for _, p := range searchedPortfolios {
		postedJSON, _ := json.Marshal(p)
		initReq, _ := http.NewRequest("POST", testServer.URL+"/api/portfolios", bytes.NewBuffer([]byte(postedJSON)))
		initReq.Header.Set("Authorization", "Bearer "+p.UID)
		client.Do(initReq)
	}

	code := m.Run()

	req, _ := http.NewRequest("GET", testServer.URL+"/api/portfolios", nil)

	resp, _ := client.Do(req)
	respBody, _ := ioutil.ReadAll(resp.Body)
	var portfolios []Portfolio
	json.Unmarshal(respBody, &portfolios)
	for _, p := range portfolios {
		if p.UID == initTestPortfolios[0].UID {
			req, _ = http.NewRequest("DELETE", testServer.URL+"/api/portfolios/"+strconv.FormatInt(p.ID, 10), nil)
			req.Header.Set("Authorization", "Bearer "+initTestPortfolios[0].UID)
			client.Do(req)
		}
	}

	os.Exit(code)
}

func TestPortfolio(t *testing.T) {
	var createdID int64

	t.Run("get portfolio list", func(t *testing.T) {
		req, _ := http.NewRequest("GET", testServer.URL+"/api/portfolios", nil)

		resp, _ := client.Do(req)
		respBody, _ := ioutil.ReadAll(resp.Body)
		var portfolios []Portfolio
		json.Unmarshal(respBody, &portfolios)
		for _, p := range searchedPortfolios {
			foundPortfolio := findPortfolio(p, portfolios)
			assert.NotNil(t, foundPortfolio)
		}
	})

	t.Run("get portfolio list with query", func(t *testing.T) {
		for k, v := range testCase {
			for query, expectedNO := range v {
				req, _ := http.NewRequest("GET", testServer.URL+"/api/portfolios?"+k+"="+query, nil)
				resp, _ := client.Do(req)
				respBody, _ := ioutil.ReadAll(resp.Body)
				var portfolios []Portfolio
				json.Unmarshal(respBody, &portfolios)
				for i, p := range searchedPortfolios {
					foundPortfolio := findPortfolio(p, portfolios)
					if (expectedNO>>i)&1 == 1 {
						assert.NotNil(t, foundPortfolio)
					} else {
						assert.Nil(t, foundPortfolio)
					}
				}
			}
		}
	})

	t.Run("post portfolio", func(t *testing.T) {
		postedJSON, _ := json.Marshal(initTestPortfolios[0])
		initReq, _ := http.NewRequest("POST", testServer.URL+"/api/portfolios", bytes.NewBuffer([]byte(postedJSON)))
		initReq.Header.Set("Authorization", "Bearer "+initTestPortfolios[0].UID)
		resp, _ := client.Do(initReq)
		respBody, _ := ioutil.ReadAll(resp.Body)
		var portfolio Portfolio
		json.Unmarshal(respBody, &portfolio)
		assert.True(t, portfolio.equal(initTestPortfolios[0]))
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		createdID = portfolio.ID
	})

	t.Run("get portfolio /:id", func(t *testing.T) {
		req, _ := http.NewRequest("GET", testServer.URL+"/api/portfolios/"+strconv.FormatInt(createdID, 10), nil)

		resp, _ := client.Do(req)
		respBody, _ := ioutil.ReadAll(resp.Body)
		var portfolio Portfolio
		json.Unmarshal(respBody, &portfolio)
		assert.True(t, portfolio.equal(initTestPortfolios[0]))
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("put portfolio /:id", func(t *testing.T) {
		putJSON, _ := json.Marshal(initTestPortfolios[1])
		req, _ := http.NewRequest("PUT", testServer.URL+"/api/portfolios/"+strconv.FormatInt(createdID, 10),
			bytes.NewBuffer([]byte(putJSON)))
		req.Header.Set("Authorization", "Bearer "+initTestPortfolios[0].UID)

		resp, _ := client.Do(req)
		respBody, _ := ioutil.ReadAll(resp.Body)
		var portfolio Portfolio
		json.Unmarshal(respBody, &portfolio)
		assert.True(t, portfolio.equal(initTestPortfolios[1]))
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("verify error", func(t *testing.T) {
		postedJSON, _ := json.Marshal(initTestPortfolios[0])
		req, _ := http.NewRequest("POST", testServer.URL+"/api/portfolios", bytes.NewBuffer([]byte(postedJSON)))
		resp, _ := client.Do(req)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

		postedJSON, _ = json.Marshal(initTestPortfolios[0])
		req, _ = http.NewRequest("POST", testServer.URL+"/api/portfolios", bytes.NewBuffer([]byte(postedJSON)))
		req.Header.Set("Authorization", "Bearer ")
		resp, _ = client.Do(req)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("put verify mismatch", func(t *testing.T) {
		putJSON, _ := json.Marshal(initTestPortfolios[1])
		req, _ := http.NewRequest("PUT", testServer.URL+"/api/portfolios/"+strconv.FormatInt(createdID, 10),
			bytes.NewBuffer([]byte(putJSON)))
		req.Header.Set("Authorization", "Bearer other_token")

		resp, _ := client.Do(req)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("delete portfolio /:id", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", testServer.URL+"/api/portfolios/"+strconv.FormatInt(createdID, 10), nil)
		req.Header.Set("Authorization", "Bearer "+initTestPortfolios[0].UID)
		resp, _ := client.Do(req)
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)

		req, _ = http.NewRequest("GET", testServer.URL+"/api/portfolios/"+strconv.FormatInt(createdID, 10), nil)
		resp, _ = client.Do(req)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}

func findPortfolio(p Portfolio, array []Portfolio) *Portfolio {
	for _, portfolio := range array {
		if portfolio.equal(p) {
			return &portfolio
		}
	}
	return nil
}

func (p *Portfolio) equal(expected Portfolio) bool {
	return p.Title == expected.Title &&
		p.UID == expected.UID &&
		p.Author == expected.Author &&
		p.Abstruct == expected.Abstruct &&
		p.Readme == expected.Readme &&
		p.Source == expected.Source &&
		p.Link == expected.Link
}
