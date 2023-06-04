package estuary

import (
	"github.com/kenlabs/permastar/pkg/util/log"
	"net/http"
	"time"
)

var logger = log.NewSubsystemLogger()

type API struct {
	Client     *http.Client
	AuthHeader map[string][]string
	BaseURL    string
	ShuttleURL string
}

func NewEstuary(apiKey string, baseURL string, shuttleURL string) *API {
	return &API{
		Client: &http.Client{Timeout: 60 * time.Second},
		AuthHeader: map[string][]string{
			"Authorization": {"Bearer " + apiKey},
		},
		BaseURL:    baseURL,
		ShuttleURL: shuttleURL,
	}
}

//func (a *API) NewAPIKeyAndRootCollectionForUser(accountAddr string) (*model.UserInfo, error) {
//	userInfo := &model.UserInfo{}
//
//	userInfo.AccountAddress = accountAddr
//
//	createUserAPIKeyParams := fmt.Sprintf("expiry=false&label=%s", accountAddr)
//	createUserAPIKeyURL, err := url.JoinPath(a.BaseURL, "/user/api-keys")
//	if err != nil {
//		return nil, err
//	}
//	createUserAPIKeyURL = fmt.Sprintf("%s?%s", createUserAPIKeyURL, createUserAPIKeyParams)
//	createUserAPIKeyRequest, err := http.NewRequest(
//		"POST",
//		createUserAPIKeyURL,
//		strings.NewReader("{\"expiry\":false}"),
//	)
//	createUserAPIKeyRequest.Header = a.AuthHeader
//	createUserAPIKeyResp, err := a.Client.Do(createUserAPIKeyRequest)
//	if err != nil {
//		return nil, err
//	}
//	respBytes, err := io.ReadAll(createUserAPIKeyResp.Body)
//	if err != nil {
//		return nil, err
//	}
//	if respErr := checkRespErr(createUserAPIKeyResp.StatusCode, respBytes); respErr != nil {
//		return nil, respErr
//	}
//	createUserAPIKeyResult := &CreateUserAPIKeyResp{}
//	if err = json.Unmarshal(respBytes, createUserAPIKeyResult); err != nil {
//		return nil, err
//	}
//	userInfo.EstuaryAPIKey = createUserAPIKeyResult.Token
//
//	createCollectionURL, err := url.JoinPath(a.BaseURL, "/collections")
//	if err != nil {
//		return nil, err
//	}
//	createCollectionReq, err := http.NewRequest(
//		"POST",
//		createCollectionURL,
//		strings.NewReader(fmt.Sprintf("{\"description\":\"%s\", \"name\":\"%s\"}", accountAddr, accountAddr)),
//	)
//	// use user's api token
//	createCollectionReq.Header = map[string][]string{
//		"Authorization": {"Bearer " + userInfo.EstuaryAPIKey},
//	}
//
//	createRootCollectionResp, err := a.Client.Do(createCollectionReq)
//	if err != nil {
//		return nil, err
//	}
//	rootCollectionRespBody, err := io.ReadAll(createRootCollectionResp.Body)
//	if err != nil {
//		return nil, err
//	}
//	if respErr := checkRespErr(createRootCollectionResp.StatusCode, rootCollectionRespBody); err != nil {
//		return nil, respErr
//	}
//	rootCollection := &Collection{}
//	if err = json.Unmarshal(rootCollectionRespBody, &rootCollection); err != nil {
//		return nil, err
//	}
//	userInfo.RootCollection = rootCollection
//
//	return userInfo, nil
//}
//
//func (a *API) UploadFileToCollection(localPath string, dstPath string, colID string, userToken string) error {
//	payload := &bytes.Buffer{}
//	writer := multipart.NewWriter(payload)
//	file, err := os.Open(localPath)
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//	fileParts, err := writer.CreateFormFile("data", file.Name())
//	if err != nil {
//		return err
//	}
//	_, err = io.Copy(fileParts, file)
//	if err != nil {
//		return err
//	}
//	_ = writer.WriteField("filename", file.Name())
//	err = writer.Close()
//	if err != nil {
//		return err
//	}
//	uploadFileToEstuaryURL, err := url.JoinPath(a.ShuttleURL, "/content/add")
//	if err != nil {
//		return err
//	}
//	uploadFileToEstuaryReq, err := http.NewRequest("POST", uploadFileToEstuaryURL, payload)
//	if err != nil {
//		return err
//	}
//	uploadFileToEstuaryReq.Header.Set("Accept", "application/json")
//	uploadFileToEstuaryReq.Header.Set("Content-Type", writer.FormDataContentType())
//	uploadFileToEstuaryReq.Header.Set("Authorization", "Bearer "+userToken)
//	uploadFileToEstuaryResp, err := a.Client.Do(uploadFileToEstuaryReq)
//	if err != nil {
//		return err
//	}
//	respBody, err := io.ReadAll(uploadFileToEstuaryResp.Body)
//	logger.Debugf("%v", respBody)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func checkRespErr(respCode int, respBody []byte) (err error) {
//	if respCode != 200 {
//		apiErr := &APIErrorResp{}
//		if err = json.Unmarshal(respBody, &apiErr); err != nil {
//			return
//		}
//		return errors.Newf("create user failed: code=%s details=%s reason=%s",
//			apiErr.Code, apiErr.Details, apiErr.Reason)
//	}
//
//	return
//}
