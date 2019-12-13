package api

type API struct {
	// Name of the API
	Name string `json:"name"`
	// Path
	Path string `json:"path"`

	Delay       Delay       `json:"delay"`
	FailureRate FailureRate `json:"failureRate"`
	Response    Response    `json:"response"`
	Self        string      `json:"self"`
}

type Delay struct {
	MinTime int `json:"minTime"`
	MaxTime int `json:"maxTime"`
}

type FailureRate struct {
	Rate          int   `json:"rate"`
	ResponseCodes []int `json:"responseCodes"`
}

type Response struct {
	Type    string            `json:"type"`
	Static  string            `json:"static"`
	Headers map[string]string `json:"headers"`
}
