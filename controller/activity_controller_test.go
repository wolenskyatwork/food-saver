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
		mockStore := &store.MockStore{}

		mockStore.On("CreateActivity", dao.Activity{Id: "1", UserId: "1", DateCompleted: "2017-03-14"}).Return(nil)
		router := NewRouter(mockStore)
		server := httptest.NewServer(router)

		values := dao.Activity{Id: "1", UserId: "1", DateCompleted: "2017-03-14"}
		jsonValue, _ := json.Marshal(values)
		url := server.URL + "/user/1/activity"

		resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))

		Convey("When client receives expected status code", func() {
			So(resp.StatusCode, ShouldEqual, http.StatusCreated)

			// TODO hook up to full db? This is kind of not testing everything
			Convey("Then the controller called the correct store function with correct inputs", func() {
				mockStore.AssertCalled(t, "CreateActivity", values)
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
		mockStore.On("GetActivities", "1").Return([]*dao.Activity{
			&knitting,
			&spartan,
			&paleo,
		}, nil)

		router := NewRouter(&mockStore)
		server := httptest.NewServer(router)

		url := server.URL + "/user/1/activity"

		resp, err := http.Get(url)
		if err != nil {
			So(err, ShouldBeNil)
		}

		defer resp.Body.Close()

		So(resp.StatusCode, ShouldEqual, http.StatusOK)

		Convey("When decoding the response", func(){
			activities := make(map[string][]*dao.Activity)
			err := json.NewDecoder(resp.Body).Decode(&activities)

			Convey("Then decode response has correct elements", func(){
				So(err, ShouldBeNil)
				So(activities, ShouldHaveLength, 2)
				So(activities["2019-05-10"], ShouldResemble, []*dao.Activity{
					&knitting,
					&spartan,
				})
				So(activities["2019-05-11"], ShouldResemble, []*dao.Activity{
					&paleo,
				})
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
