// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package billing

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-validator/converters/google"
	"github.com/pkg/errors"
	"google.golang.org/api/cloudbilling/v1"
)

// To be set by Go build tools.
var buildVersion string

// BuildVersion returns the build version of Terraform Validator.
func BuildVersion() string {
	return buildVersion
}

type assetUnitPrice struct {
	assetName string
	assetType string
}

// getServices get listed public prices of assets
func getServices(ctx context.Context, assets []google.Asset) (map[string]string, error) {
	allServiceMap := make(map[string]string)
	assetServiceMap := make(map[string]string)
	cloudbillingService, err := cloudbilling.NewService(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "initializing cloud billing service")
	}
	servicesService := cloudbilling.NewServicesService(cloudbillingService)
	servicesServiceCall := servicesService.List()
	servicesServiceCall = servicesServiceCall.Context(ctx)
	servicesServiceResponse, err := servicesServiceCall.Do()
	if err != nil {
		return nil, errors.Wrap(err, "getting response from list services call")
	}
	for _, s := range servicesServiceResponse.Services {
		allServiceMap[s.DisplayName] = s.Name
	}
	for servicesServiceResponse.NextPageToken != "" {
		servicesServiceCall := servicesServiceCall.PageToken(servicesServiceResponse.NextPageToken)
		servicesServiceResponse, err = servicesServiceCall.Do()
		if err != nil {
			return nil, errors.Wrap(err, "getting response from list services call")
		}
		for _, s := range servicesServiceResponse.Services {
			allServiceMap[s.DisplayName] = s.Name
		}
	}
	for _, a := range assets {
		assetServiceDisplayName := mappers()[a.Type].serviceDisplayName
		if assetServiceDisplayName != "" {
			assetServiceMap[assetServiceDisplayName] = allServiceMap[assetServiceDisplayName]
		}
	}
	return assetServiceMap, nil
}

// GetAssetPrices get listed public prices of assets
func GetAssetPrices(ctx context.Context, assets []google.Asset) ([]assetUnitPrice, error) {
	var prices []assetUnitPrice
	assetServiceMap, err := getServices(ctx, assets)
	if err != nil {
		return nil, errors.Wrap(err, "getting services")
	}
	cloudbillingService, err := cloudbilling.NewService(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "initializing cloud billing service")
	}
	var skus []*cloudbilling.Sku
	servicesSkusService := cloudbilling.NewServicesSkusService(cloudbillingService)
	for _, service := range assetServiceMap {
		fmt.Println("getting sku for: ", service)
		servicesSkusCall := servicesSkusService.List(service)
		servicesSkusCall = servicesSkusCall.Context(ctx)
		servicesSkusCall = servicesSkusCall.CurrencyCode("SGD")
		listSkusResponse, err := servicesSkusCall.Do()
		if err != nil {
			return nil, errors.Wrap(err, "getting response from list skus call")
		}
		for _, s := range listSkusResponse.Skus {
			skus = append(skus, s)
		}
		for listSkusResponse.NextPageToken != "" {
			servicesSkusCall := servicesSkusCall.PageToken(listSkusResponse.NextPageToken)
			listSkusResponse, err = servicesSkusCall.Do()
			if err != nil {
				return nil, errors.Wrap(err, "getting response from list skus call")
			}
			for _, s := range listSkusResponse.Skus {
				skus = append(skus, s)
			}
		}
	}

	for _, s := range skus {
		fmt.Println(s.Description, ",", s.ServiceRegions, ",", s.PricingInfo)
	}
	//TODO: get those relevant for assets only or even hardcode
	return nil, nil

	/*
		for i := range assets {
		}

		if err := valid.AddData(&validator.AddDataRequest{
			Assets: pbAssets,
		}); err != nil {
			return nil, errors.Wrap(err, "adding data to validator")
		}

		auditResult, err := valid.Audit(context.Background())
		if err != nil {
			return nil, errors.Wrap(err, "auditing")
		}
	*/
	return prices, nil
}
