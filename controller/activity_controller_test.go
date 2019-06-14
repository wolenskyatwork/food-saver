package controller

import (
	"bytes"
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/wolenskyatwork/food-saver/dao"
	"github.com/wolenskyatwork/food-saver/store"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	Convey("Given an http post to /user/{userId}/activity", t, func() {
		mockStore := store.MockStore{}
		router := NewRouter(mockStore)
		server := httptest.NewServer(router)

		values := dao.Activity{Id: "1", UserId: "1", DateCompleted: "2017-03-14"}
		jsonValue, _ := json.Marshal(values)
		url := server.URL + "/user/1/activity"

		resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
		defer resp.Body.Close()

		Convey("When client receives expected status code", func() {
			So(resp.StatusCode, ShouldEqual, http.StatusCreated)

			Convey("Then the controller called the correct store function with correct inputs", func() {
				// mockStore.AssertCalled(suite.T(), "InsertActivity", 1, 1, "2017-03-14s")
			})
		})
		server.Close()
	})
}

func TestIndex(t *testing.T) {
	Convey("Given http get to /user/{userId}/activity", t, func(){
		knitting := dao.Activity{Name: "knitting", DateCompleted: "2019-05-10"}
		spartan := dao.Activity{Name: "spartan", DateCompleted: "2019-05-10"}
		paleo := dao.Activity{Name: "paleo", DateCompleted: "2019-05-11"}

		mockStore := store.MockStore{}
		mockStore.On("GetActivities").Return([]*dao.Activity{
			&knitting,
			&spartan,
			&paleo,
		}, nil)

		router := NewRouter(mockStore)
		server := httptest.NewServer(router)

		url := server.URL + "/user/1/activity"

		resp, err := http.Get(url)
		if err != nil {
			So(err, ShouldBeNil)
		}

		defer resp.Body.Close()

		So(resp.StatusCode, ShouldEqual, http.StatusOK)

		Convey("When decoding the response", func(){
			activities := []dao.Activity{}
			err := json.NewDecoder(resp.Body).Decode(&activities)

			Convey("Then decode response has correct elements", func(){
				So(err, ShouldBeNil)
				So(activities, ShouldHaveLength, 3)
				So(activities[0], ShouldResemble, dao.Activity{Name: "knitting", DateCompleted: "2019-05-10"})
				So(activities[1], ShouldResemble, dao.Activity{Name: "spartan", DateCompleted: "2019-05-10"})
				So(activities[2], ShouldResemble, dao.Activity{Name: "paleo", DateCompleted: "2019-05-11"})
			})
		})
		server.Close()
	})
}

//
//func TestActivityControllerTestSuite(t *testing.T) {
//	suite.Run(t, new(ActivityControllerTestSuite))
//}
//
//type ActivityControllerTestSuite struct {
//	suite.Suite
//}
//
//var mockStore MockStore
//var router *mux.Router
//var server *httptest.Server
//
//func (suite *ActivityControllerTestSuite) SetupTestSuite() {
//	mockStore = MockStore{}
//	router = NewRouter(mockStore)
//	server = httptest.NewServer(router)
//}
//
//func (suite *ActivityControllerTestSuite) TearDownAllSuite() {
//	server.Close()
//}
