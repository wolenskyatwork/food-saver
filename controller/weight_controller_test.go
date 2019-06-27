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

func TestWeightIndex(t *testing.T) {
	Convey("Given http get to /user/{userId}/weight", t, func(){
		w0 := dao.Weight{
			Id: "1",
			UserId: "1",
			Weight: 147.3,
			WeightDate: "2019-06-26",
		}
		w1 := dao.Weight{
			Id: "2",
			UserId: "1",
			Weight: 146.3,
			WeightDate: "2019-06-21",
		}
		w2 := dao.Weight{
			Id: "3",
			UserId: "1",
			Weight: 144.3,
			WeightDate: "2019-06-18",
		}

		mockStore := store.MockStore{}
		mockStore.On("GetWeights", "1").Return([]*dao.Weight{
			&w0,
			&w1,
			&w2,
		}, nil)

		router := NewRouter(&mockStore)
		server := httptest.NewServer(router)

		url := server.URL + "/user/1/weight"

		resp, err := http.Get(url)

		if err != nil {
			So(err, ShouldBeNil)
		}

		defer resp.Body.Close()

		So(resp.StatusCode, ShouldEqual, http.StatusOK)

		Convey("When decoding the response", func(){
			var weights []*dao.Weight
			err := json.NewDecoder(resp.Body).Decode(&weights)
			So(err, ShouldBeNil)

			Convey("Then decode response has correct elements", func(){
				So(weights, ShouldHaveLength, 3)
				So(weights[0], ShouldResemble, &w0)
				So(weights[1], ShouldResemble, &w1)
				So(weights[2], ShouldResemble, &w2)
			})
		})
		server.Close()
	})
}


func TestWeightCreate(t *testing.T) {
	Convey("Given an http post to /user/{userId}/weight", t, func() {
		mockStore := &store.MockStore{}

		mockStore.On("CreateWeight", dao.Weight{Id: "1", UserId: "1", WeightDate: "2017-03-14", Weight: 142.7}).Return(nil)
		router := NewRouter(mockStore)
		server := httptest.NewServer(router)

		values := dao.Weight{Id: "1", UserId: "1", WeightDate: "2017-03-14", Weight: 142.7}
		jsonValue, _ := json.Marshal(values)
		url := server.URL + "/user/1/weight"

		resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))

		Convey("When client receives expected status code", func() {
			So(resp.StatusCode, ShouldEqual, http.StatusCreated)

			// TODO hook up to full db? This is kind of not testing everything
			Convey("Then the controller called the correct store function with correct inputs", func() {
				mockStore.AssertCalled(t, "CreateWeight", values)
			})
		})
		server.Close()
	})
}
