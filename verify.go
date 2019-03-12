package mobtexting_sms
import (
                "net/http"
                "io/ioutil"
       )

func VerifySend(to string, access_token string) string{
        var api_endpoint = "https://portal.mobtexting.com/api/v2/"
        var url = api_endpoint + "verify?to=" + to
        req, err := http.NewRequest("POST", url, nil)
        req.Header.Set("Authorization", "Bearer " + access_token)
        req.Header.Set("Accept", "application/json")
        client := &http.Client{}
        resp, err := client.Do(req)

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

func VerifyCheck(id string, token string, access_token string) string{
        return doGet("verify/check/" + id + "/" + token, access_token)
}

func VerifyCancel(id string, access_token string) string{
        return doGet("verify/cancel/" + id, access_token)
}

func doGet(url string, access_token string) string{
        var api_endpoint = "https://portal.mobtexting.com/api/v2/"
        url = api_endpoint + url

        req, err := http.NewRequest("GET", url, nil)
        req.Header.Set("Authorization", "Bearer " + access_token)
        req.Header.Set("Accept", "application/json")
        client := &http.Client{}
        resp, err := client.Do(req)

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
