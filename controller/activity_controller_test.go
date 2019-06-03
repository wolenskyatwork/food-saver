package controller

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/wolenskyatwork/food-saver/dao"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestActivityController(t *testing.T) {
	mockStore := new(MockStore)
	router := NewRouter(mockStore)
	server := httptest.NewServer(router)

	Convey("Given http response", t, func(){
		knitting := dao.Activity{Name: "knitting", DateCompleted: "2019-05-10"}
		spartan := dao.Activity{Name: "spartan", DateCompleted: "2019-05-10"}
		paleo := dao.Activity{Name: "paleo", DateCompleted: "2019-05-11"}

		mockStore.On("GetActivities").Return([]*dao.Activity{
			&knitting,
			&spartan,
			&paleo,
		}, nil).Once()

		url := server.URL + "/activity"
		resp, _ := http.Get(url)
		defer resp.Body.Close()

		//this So is ok, it doesn't follow given when then, but it is just a check in a preparation step
		So(resp.StatusCode, ShouldEqual, http.StatusOK)

		Convey("When decoding the response", func(){
			activities := []dao.Activity{}
			err := json.NewDecoder(resp.Body).Decode(&activities)

			Convey("Then decode response has correct elements", func(){
				So(err, ShouldBeNil)
				So(activities, ShouldHaveLength, 3)
				So(activities[0], ShouldResemble, dao.Activity{Name: "knitting", DateCompleted: "2019-05-10"})
			})
		})

		server.Close()
	})
}

