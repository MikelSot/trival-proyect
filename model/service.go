package model

type ApiResponse struct {
	Results Results `json:"results"`
}

type Service struct {
	Gender string `json:"gender"`
	Name   Name   `json:"name"`
	Email  string `json:"email"`
	Login  Login  `json:"login"`
}

type Results []Service

func (r Results) UniqueMap() map[string]Service {
	resultsMap := make(map[string]Service)

	for _, result := range r {
		resultsMap[result.Login.Uuid] = result
	}

	return resultsMap
}

type Name struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

type Login struct {
	Uuid string `json:"uuid"`
}
