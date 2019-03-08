package mobtexting_sms

import (
		"net/http"
		"io/ioutil"
       )

func Send(to string, from string, body string, service string, access_token string) string{
	var api_endpoint = "https://portal.mobtexting.com/api/v2/"
        var url = api_endpoint + "sms/send?access_token=" + access_token + "&message=" + body  + "&sender="+ from + "&to=" + to + "&service=" + service

	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		return "request error!"
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
			bodyString := string(bodyBytes)
			if(err2 != nil){
				return "request error!!"
			}else{
				return bodyString
			}
	}
	return ""
}
