package nestsensors

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"

// 	"jaytaylor.com/nestsensors/domain"
// )

// func Raw(auth *domain.Auth, client *http.Client) error {
// 	u := fmt.Sprintf("%v/api/0.1/user/%v/app_launch", BaseURL, auth.UserID)
// 	body, err := checkedGet(client, u)
// 	if err != nil {
// 		return err
// 	}
// 	defer body.Close()

// 	data, err := ioutil.ReadAll(body)
// 	if err != nil {
// 		return err
// 	}

// 	m := map[string]interface{}{}
// 	if err := json.Unmarshal(data, &m); err != nil {
// 		return err
// 	}
// 	bs, err := json.MarshalIndent(&m, "", "    ")
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("%v\n", string(bs))

// 	return nil
// }
