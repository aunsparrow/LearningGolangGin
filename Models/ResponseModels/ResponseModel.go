package ResponseModels

type ResponseModel struct {
	IsSuccess string      `json:"IsSuccess"`
	Status    int         `json:"Status"`
	Result    interface{} `josn:"Result"`
	Message   string      `json:"Message"`
}
