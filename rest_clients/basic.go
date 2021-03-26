package clients

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"net/http"
	`reflect`
)

type ClientHelper struct {
	CodeSuccess int
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

func constructRequest(ctx context.Context, restyClient *resty.Client, pathParams, query map[string]string, body interface{}, basicResponseType reflect.Type) *resty.Request {
	req := restyClient.R().
		SetContext(ctx).
		SetResult(&tspModel.BaseResp{})
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

func dealWithResponse(raw *tspModel.BaseResp, responseData interface{}) error {
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

func validateResponse(resp *resty.Response, url string, restCodeSuccess int, requestErr error) (*tspModel.BaseResp, error) {
	var raw *tspModel.BaseResp
	if requestErr != nil {
		raw := &tspModel.BaseResp{
			Meta: &tspModel.Meta{
				Code: tspModel.CodeClientError,
			},
		}
		return raw, requestErr
	}
	if resp.StatusCode() != http.StatusOK {
		raw = &tspModel.BaseResp{
			Meta: &tspModel.Meta{},
		}
		raw.Meta.Code = resp.StatusCode()
		raw.Meta.Message = string(resp.Body())
		return raw, errors.New(fmt.Sprintf("Fail to get from %+v.Remore code %+v.Message %+v", url, resp.StatusCode(), string(resp.Body())))
	}
	raw = resp.Result().(*tspModel.BaseResp)
	if raw.Meta.Code != restCodeSuccess {
		return raw, errors.New(fmt.Sprintf("Fail to get from %+v.Remore code %+v.Message %+v", url, raw.Meta.Code, raw.Meta.Message))
	}
	return raw, nil
}
