package api

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"strconv"
	"strings"
)

type Response struct {
	Data Data `json:"data"`
}

type Data struct {
	Result []MetricValue `json:"result"`
}

type MetricValue struct {
	Metric Metric        `json:"metric"`
	Value  []interface{} `json:"value"`
}

type Metric struct {
	Id       string `json:"id"`
	Image    string `json:"image"`
	Instance string `json:"instance"`
	Job      string `json:"job"`
	Name     string `json:"name"`
}

type Final struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func Request(url, username, password string) []Final {

	//http request
	request := gorequest.New()
	res, _, errs := request.Get(url).
		SetBasicAuth(username, password).
		End()

	//error
	if len(errs) != 0 {

	}

	//parsing json
	b, _ := ioutil.ReadAll(res.Body)
	r := Response{
		Data: Data{},
	}
	json.Unmarshal(b, &r)

	//parsing struct
	finals := MetricValueToFinal(r.Data.Result)

	return finals
}

func MetricValueToFinal(metricValues []MetricValue) []Final {

	finals := make([]Final, 0)

	for i := range metricValues {
		final := Final{}
		metricValue := metricValues[i]

		//id
		id := metricValue.Metric.Id
		idArray := strings.Split(id, "/")
		final.Id = idArray[2][0:6]

		//name
		final.Name = metricValue.Metric.Name

		//value
		value := metricValue.Value[1].(string)
		valueFloat, err := strconv.ParseFloat(value, 64)
		if err != nil {

		}

		final.Value = fmt.Sprintf("%.2f", valueFloat)

		finals = append(finals, final)
	}

	return finals
}
