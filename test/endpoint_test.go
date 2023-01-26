package test

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2/utils"
	"io"
	"learn-mock/app"
	"learn-mock/models"
	"log"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

import . "github.com/smartystreets/goconvey/convey"

func TestEndpoint(t *testing.T) {
	var arrayBody = struct {
		Message string           `json:"message"`
		Data    []models.Product `json:"data"`
	}{}
	Convey("Test endpoint", t, func() {
		pool, resource, db, err := SetupDockerTest()
		if err != nil {
			return
		}
		httpApp := app.Setup(&db)
		err = db.AutoMigrate(&models.Product{})
		if err != nil {
			log.Fatal(err)
		}
		Convey("product post request without body should return 400", func() {
			req := httptest.NewRequest("POST", "/api/product", nil)
			res, _ := httpApp.Test(req)
			So(res.StatusCode, ShouldEqual, 200)
		})

		Convey("product post request with body should return 201", func() {
			body := url.Values{}
			body.Add("name", "Some Product")
			req := httptest.NewRequest("POST", "/api/product", strings.NewReader(body.Encode()))
			req.Header.Add("content-type", "application/x-www-form-urlencoded")
			res, _ := httpApp.Test(req)
			So(res.StatusCode, ShouldEqual, 201)
		})

		Convey("get product pagination", func() {
			id := utils.UUID()
			dataInDb := models.Product{Id: id, Name: "soap"}
			tx := db.Model(&models.Product{}).Create(&dataInDb)
			if tx.Error != nil {
				panic(tx.Error)
			}

			Convey("should return 400 when page or limit is not provided", func() {
				req := httptest.NewRequest("GET", "/api/product", nil)
				res, _ := httpApp.Test(req)
				So(res.StatusCode, ShouldEqual, 400)
			})

			Convey("page 1 should return at least 1 data", func() {
				req := httptest.NewRequest("GET", "/api/product?page=1&limit=1", nil)
				res, _ := httpApp.Test(req)
				bytes, err := io.ReadAll(res.Body)
				if err != nil {
					panic(err)
				}
				err = json.Unmarshal(bytes, &arrayBody)
				if err != nil {
					panic(err)
				}
				So(res.StatusCode, ShouldEqual, 200)
				So(len(arrayBody.Data), ShouldNotEqual, 0)
			})

			Convey("page 2 should return 0 data", func() {
				req := httptest.NewRequest("GET", "/api/product?page=2&limit=1", nil)
				res, _ := httpApp.Test(req)
				bytes, err := io.ReadAll(res.Body)
				if err != nil {
					panic(err)
				}
				err = json.Unmarshal(bytes, &arrayBody)
				if err != nil {
					panic(err)
				}
				So(res.StatusCode, ShouldEqual, 200)
				So(len(arrayBody.Data), ShouldEqual, 0)
			})
		})
		Reset(func() {
			if err := pool.Purge(resource); err != nil {
				log.Fatalln(err)
			}
		})
	})
}
