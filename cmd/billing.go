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

package cmd

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-validator/billing"
	"github.com/GoogleCloudPlatform/terraform-validator/tfgcv"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var billingCmd = &cobra.Command{
	Use:   "billing <tfplan>",
	Short: "Query list public prices for resources in a Terraform plan.",
	Long: `Billing (terraform-validator billing) query list public prices for resources in a Terraform plan.

Example:
  terraform-validator billing ./example/terraform.tfplan 
`,
	PreRunE: func(c *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("missing required argument <tfplan>")
		}
		return nil
	},
	RunE: func(c *cobra.Command, args []string) error {
		ctx := context.Background()
		assets, err := tfgcv.ReadPlannedAssets(ctx, args[0], "test", "organizations/0", false)
		if err != nil {
			return errors.Wrap(err, "converting tfplan to CAI assets")
		}

		prices, err := billing.GetAssetPrices(ctx, assets)
		if err != nil {
			return errors.Wrap(err, "getting prices")
		}
		fmt.Print(prices)

		return nil
	},
}
