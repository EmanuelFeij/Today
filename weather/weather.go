package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	MapRegions regionIds
	MapWeather weatherMap
)

func init() {
	err1, err2 := ReadRegions("weather/regions.json"), getWeatherDescription("weather/description.json")
	if err1 != nil || err2 != nil {
		log.Panic("Resources couldn't be read")
	}
}

type region struct {
	IdRegiao      int    `json:"idRegiao"`
	IdAreaAviso   string `json:"idAreaAviso"`
	IdConcelho    int    `json:"idConcelho"`
	GlobalIdLocal int    `json:"GlobalIdLocal"`
	Latitude      string `json:"latitude"`
	IdDistrito    int    `json:"idDistrito"`
	Local         string `json:"local"`
	Longitude     string `json:"longitude"`
}

type regionIds map[string]region

func ReadRegions(path string) error {
	MapRegions = make(regionIds)
	file, err := ioutil.ReadFile("weather/regions.json")
	if err != nil {
		return err
	}

	content := make([]region, 0)
	err = json.Unmarshal(file, &content)
	if err != nil {
		log.Fatalf("%v", err)
	}
	for _, r := range content {
		MapRegions[strings.ToLower(r.Local)] = r
	}

	return nil
}

type weatherDescription struct {
	DescIdWeatherTypeEN string `json:"descIdWeatherTypeEN"`
	DescIdWeatherTypePT string `json:"descIdWeatherTypePT"`
	IdWeatherType       int    `json:"idWeatherType"`
}

type weatherMap map[int]weatherDescription

func getWeatherDescription(path string) error {
	MapWeather = make(weatherMap)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	content := make([]weatherDescription, 0)
	json.Unmarshal(file, &content)

	for _, wDes := range content {
		MapWeather[wDes.IdWeatherType] = wDes
	}
	return nil
}

type weatherData struct {
	PrecipitaProb  string `json:"precipitaProb"`
	TMin           string `json:"tMin"`
	TMax           string `json:"tMax"`
	PredWindDir    string `json:"predWindDir"`
	IdWeatherType  int    `json:"idWeatherType"`
	ClassWindSpeed int    `json:"classWindSpeed"`
	Longitude      string `json:"longitude"`
	ForecastDate   string `json:"forecastDate"`
	Latitude       string `json:"latitude"`
}

type ipmaResponse struct {
	Owner         string `json:"owner"`
	Country       string `json:"country"`
	Data          []weatherData
	GlobalIdLocal int    `json:"globalIdLocal"`
	DataUpdate    string `json:"dataUpdate"`
}

func GetWeather(city string, nextDays bool) (string, error) {

	region, ok := MapRegions[strings.ToLower(city)]
	city = strings.Title(strings.ToLower(city))
	if !ok {
		return "", fmt.Errorf("city %v not found", city)
	}
	cityID := region.GlobalIdLocal
	apiRequest := fmt.Sprintf("https://api.ipma.pt/open-data/forecast/meteorology/cities/daily/%v.json", cityID)
	resp, err := http.Get(apiRequest)
	if err != nil {
		return "", fmt.Errorf("error getting api response: %v", err)
	}
	defer resp.Body.Close()
	var ipmaRes ipmaResponse
	err = json.NewDecoder(resp.Body).Decode(&ipmaRes)
	if err != nil {
		return "", fmt.Errorf("error decoding api response: %v", err)
	}
	if !nextDays {
		return PrettyPrintIpmaResponseToday(ipmaRes, city), nil
	}
	return PrettyPrintIpmaResponseSeveralDays(ipmaRes, city), nil
}

func PrettyPrintIpmaResponseSeveralDays(ir ipmaResponse, city string) string {
	data := ir.Data
	prettierData := make([]string, 0, len(data))
	for _, day := range data {
		prettierData = append(
			prettierData,
			fmt.Sprintf("City: %v\nDate: %v\nTMax: %v\nTMin: %v\nPrecipitaProblably: %v\nDescription: %v\n",
				city,
				day.ForecastDate,
				day.TMax,
				day.TMin,
				day.PrecipitaProb,
				MapWeather[day.IdWeatherType].DescIdWeatherTypeEN,
			),
		)
	}

	return strings.Join(prettierData, "\n")
}

func PrettyPrintIpmaResponseToday(ir ipmaResponse, city string) string {
	data := ir.Data[0]

	return fmt.Sprintf("City: %v\nDate: %v\nTMax: %v\nTMin: %v\nPrecipitaProblably: %v\nDescription: %v",
		city,
		data.ForecastDate,
		data.TMax,
		data.TMin,
		data.PrecipitaProb,
		MapWeather[data.IdWeatherType].DescIdWeatherTypeEN,
	)

}

//https://api.ipma.pt/open-data/forecast/meteorology/cities/daily/1060300.json
