package clients

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"net/http"
)

type ClientHelper struct {
	CodeSuccess int64
}

func (helper *ClientHelper) MethodGet(ctx context.Context, restyClient *resty.Client, url string, pathParams, query map[string]string, responseData interface{}) error {
	if restyClient == nil {
		return errors.Errorf("Client has not been initialized")
	}
	resp, err := constructRequest(ctx, restyClient, pathParams, query, nil).
		Get(url)
	raw, err := validateResponse(resp, url, helper.CodeSuccess, err)
	if err != nil {
		return err
	}
	err = dealWithResponse(raw, responseData)
	if err != nil {
		return err
	}
	return nil
}

func (helper *ClientHelper) MethodPut(ctx context.Context, restyClient *resty.Client, url string, pathParams, query map[string]string, body interface{}, responseData interface{}) ( error) {
	if restyClient == nil {
		return errors.Errorf("Client has not been initialized")
	}
	resp, err := constructRequest(ctx, restyClient, pathParams, query, body).
		Put(url)
	raw, err := validateResponse(resp, url, helper.CodeSuccess, err)
	if err != nil {
		return err
	}
	err = dealWithResponse(raw, responseData)
	if err != nil {
		return err
	}
	return nil
}

func (helper *ClientHelper) MethodPost(ctx context.Context, restyClient *resty.Client, url string, pathParams, query map[string]string, body interface{}, responseData interface{}) ( error) {
	if restyClient == nil {
		return errors.Errorf("Client has not been initialized")
	}
	resp, err := constructRequest(ctx, restyClient, pathParams, query, body).
		Post(url)
	raw, err := validateResponse(resp, url, helper.CodeSuccess, err)
	if err != nil {
		return err
	}
	err = dealWithResponse(raw, responseData)
	if err != nil {
		return  err
	}
	return nil
}

func (helper *ClientHelper) MethodPatch(ctx context.Context, restyClient *resty.Client, url string, pathParams, query map[string]string, body interface{}, responseData interface{}) ( error) {
	if restyClient == nil {
		return  errors.Errorf("Client has not been initialized")
	}
	resp, err := constructRequest(ctx, restyClient, pathParams, query, body).
		Patch(url)
	raw, err := validateResponse(resp, url, helper.CodeSuccess, err)
	if err != nil {
		return err
	}
	err = dealWithResponse(raw, responseData)
	if err != nil {
		return  err
	}
	return nil
}

func (helper *ClientHelper) MethodDelete(ctx context.Context, restyClient *resty.Client, url string, pathParams, query map[string]string, body interface{}, responseData interface{}) ( error) {
	if restyClient == nil {
		return errors.Errorf("Client has not been initialized")
	}
	resp, err := constructRequest(ctx, restyClient, pathParams, query, body).
		Delete(url)
	raw, err := validateResponse(resp, url, helper.CodeSuccess, err)
	if err != nil {
		return err
	}
	err = dealWithResponse(raw, responseData)
	if err != nil {
		return  err
	}
	return  nil
}

func constructRequest(ctx context.Context, restyClient *resty.Client, pathParams, query map[string]string, body interface{}) *resty.Request {
	req := restyClient.R().
		SetContext(ctx).
		SetResult(&BaseResp{})
	if body != nil {
		req = req.SetBody(body)
	}
	if query != nil {
		req = req.SetQueryParams(query)
	}
	if pathParams != nil {
		req = req.SetPathParams(pathParams)
	}
	return req
}

func dealWithResponse(raw *BaseResp, responseData interface{}) error {
	if responseData == nil {
		return nil
	}
	json, err := jsoniter.Marshal(raw.Data)
	if err != nil {
		return err
	}
	err = jsoniter.Unmarshal(json, responseData)
	if err != nil {
		return err
	}
	return nil
}

func validateResponse(resp *resty.Response, url string, restCodeSuccess int64, requestErr error) (*BaseResp, error) {
	var raw *BaseResp
	if requestErr != nil {
		raw := &BaseResp{
			Meta: &Meta{
				Code: CodeClientError,
			},
		}
		return raw, requestErr
	}
	if resp.StatusCode() != http.StatusOK {
		raw = &BaseResp{
			Meta: &Meta{},
		}
		raw.Meta.Code = int64(resp.StatusCode())
		raw.Meta.Message = string(resp.Body())
		return raw, errors.New(fmt.Sprintf("Fail to get from %+v.Remore code %+v.Message %+v", url, resp.StatusCode(), string(resp.Body())))
	}
	raw = resp.Result().(*BaseResp)
	if raw.Meta.Code != restCodeSuccess {
		return raw, errors.New(fmt.Sprintf("Fail to get from %+v.Remore code %+v.Message %+v", url, raw.Meta.Code, raw.Meta.Message))
	}
	return raw, nil
}
