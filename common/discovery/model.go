package discovery

import "encoding/json"

type EndpointInfoModel struct {
	IP       string                 `json:"ip"`
	Port     string                 `json:"port"`
	MetaData map[string]interface{} `json:"meta"`
}

func UnMarshal(data []byte) (*EndpointInfoModel, error) {
	ed := &EndpointInfoModel{}
	err := json.Unmarshal(data, ed)
	if err != nil {
		return nil, err
	}
	return ed, nil
}

func (edi *EndpointInfoModel) Marshal() string {
	data, err := json.Marshal(edi)
	if err != nil {
		panic(err)
	}
	return string(data)
}
