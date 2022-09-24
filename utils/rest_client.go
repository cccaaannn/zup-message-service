package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"zup-message-service/data/dtos"
)


func GetApiDataResult[T any](path string) dtos.DataResult[T] {
	result := dtos.DataResult[T]{Status: false, Message: "", Data: nil}

	res, err := http.Get(path)
	if err != nil {
		fmt.Printf("Error making http request: %s\n", err)
		return result
	}

	err = json.NewDecoder(res.Body).Decode(&result)

	if err != nil {
		result = dtos.DataResult[T]{Status: false, Message: "", Data: nil}
	}

	log.Println(result)
	return result
}

// func GetApiResult(path string) dtos.Result {
// 	result := dtos.Result{Status: false, Message: ""}

// 	req, err := http.Get(path)
// 	if err != nil {
// 		fmt.Printf("Error making http request: %s\n", err)
// 		return result
// 	}

// 	err = json.NewDecoder(req.Body).Decode(&result)

// 	if err != nil {
// 		result = dtos.Result{Status: false, Message: ""}
// 	}

// 	log.Println(result)
// 	return result
// }

func GetApiDataResultWithToken[T any](path string, accessToken string) dtos.DataResult[T] {
	result := dtos.DataResult[T]{Status: false, Message: "", Data: nil}

	var bearer = "Bearer " + accessToken

	req, err := http.NewRequest(http.MethodGet, path, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearer)

	if err != nil {
		fmt.Printf("Error making http request: %s\n", err)
		return result
	}


    client := &http.Client{}
  
  
    client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
        for key, val := range via[0].Header {
            req.Header[key] = val
        }
        return err
    }
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERRO] -", err)
    } else {
        defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&result)

		if err != nil {
			result = dtos.DataResult[T]{Status: false, Message: "", Data: nil}
		}

    }

	log.Println(result)
	return result
}

func GetApiResult(path string, accessToken string) dtos.Result {
	result := dtos.Result{Status: false, Message: ""}

	var bearer = "Bearer " + accessToken

	req, err := http.NewRequest(http.MethodGet, path, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearer)

	if err != nil {
		fmt.Printf("Error making http request: %s\n", err)
		return result
	}


    client := &http.Client{}
  
  
    client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
        for key, val := range via[0].Header {
            req.Header[key] = val
        }
        return err
    }
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERRO] -", err)
    } else {
        defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&result)

		if err != nil {
			result = dtos.Result{Status: false, Message: ""}
		}

    }


	log.Println(result)
	return result
}

func PutApiResult(path string, accessToken string) dtos.Result {
	result := dtos.Result{Status: false, Message: ""}

	var bearer = "Bearer " + accessToken

	req, err := http.NewRequest(http.MethodPut, path, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearer)

	if err != nil {
		fmt.Printf("Error making http request: %s\n", err)
		return result
	}


    client := &http.Client{}
  
  
    client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
        for key, val := range via[0].Header {
            req.Header[key] = val
        }
        return err
    }
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERRO] -", err)
    } else {
        defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&result)

		if err != nil {
			result = dtos.Result{Status: false, Message: ""}
		}

    }


	log.Println(result)
	return result
}
