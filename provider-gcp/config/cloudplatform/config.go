package cloudplatform

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-gcp/config/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_project", func(r *config.Resource) {
		r.TerraformResource.Schema["org_id"].Description =
			"The numeric ID of the organization this project belongs to."
		r.UseAsync = true
	})
	p.AddResourceConfigurator("google_service_account_key", func(r *config.Resource) {
		// Note(turkenh): We have to modify schema of "keepers", since it is a
		// map where elements configured as nil, but needs to be String:
		r.TerraformResource.
			Schema["keepers"].Elem = schema.TypeString

		r.References["service_account_id"] = config.Reference{
			Type:      "ServiceAccount",
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("google_service_account", func(r *config.Resource) {
		r.Kind = "ServiceAccount"
	})
}