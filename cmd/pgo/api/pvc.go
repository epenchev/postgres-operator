package api

/*
 Copyright 2018 - 2022 Crunchy Data Solutions, Inc.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	msgs "github.com/crunchydata/postgres-operator/pkg/apiservermsgs"
	log "github.com/sirupsen/logrus"
)

func ShowPVC(httpclient *http.Client, request *msgs.ShowPVCRequest, SessionCredentials *msgs.BasicAuthCredentials) (msgs.ShowPVCResponse, error) {
	var response msgs.ShowPVCResponse

	ctx := context.TODO()
	url := SessionCredentials.APIServerURL + "/showpvc"
	log.Debugf("ShowPVC called...[%s]", url)

	jsonValue, _ := json.Marshal(request)
	log.Debugf("ShowPVC called...[%s]", url)

	action := "POST"
	req, err := http.NewRequestWithContext(ctx, action, url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return response, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(SessionCredentials.Username, SessionCredentials.Password)

	resp, err := httpclient.Do(req)
	if err != nil {
		fmt.Println("Error: Do: ", err)
		return response, err
	}
	defer resp.Body.Close()
	log.Debugf("%v", resp)
	err = StatusCheck(resp)
	if err != nil {
		return response, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("%v\n", resp.Body)
		log.Println(err)
		return response, err
	}

	return response, err
}
