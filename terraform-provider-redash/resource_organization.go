//
// Copyright (c) 2020 Snowplow Analytics Ltd. All rights reserved.
//
// This program is licensed to you under the Apache License Version 2.0,
// and you may not use this file except in compliance with the Apache License Version 2.0.
// You may obtain a copy of the Apache License Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the Apache License Version 2.0 is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the Apache License Version 2.0 for the specific language governing permissions and limitations there under.
//
package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/leandro-lorenzini/redash-client-go/redash"
)

func resourceRedashOrganization() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRedashOrganizationUpdate,
		ReadContext:   resourceRedashOrganizationRead,
		UpdateContext: resourceRedashOrganizationUpdate,
		DeleteContext: resourceRedashOrganizationUpdate,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"auth_password_login_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"auth_saml_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"auth_saml_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_saml_entity_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_saml_metadata_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_saml_nameid_format": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_saml_sso_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceRedashOrganizationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*redash.Client)

	var diags diag.Diagnostics

	payload := redash.Organization{
		AuthPasswordLoginEnabled: 	d.Get("auth_password_login_enabled").(bool),
		AuthSamlEnabled: 			d.Get("auth_saml_enabled").(bool),
		AuthSamlType: 				d.Get("auth_saml_type").(string),
		AuthSamlEntityId: 			d.Get("auth_saml_entity_id").(string),
		AuthSamlMetadataUrl: 		d.Get("auth_saml_metadata_url").(string),
		AuthSamlNameidFormat: 		d.Get("auth_saml_nameid_format").(string),
		AuthSamlSsoUrl: 			d.Get("auth_saml_sso_url").(string),
	}

	organization, err := c.UpdateOrganization(&payload)
	if organization != nil && err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprint(1))

	resourceRedashOrganizationRead(ctx, d, meta)

	return diags
}

func resourceRedashOrganizationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*redash.Client)

	var diags diag.Diagnostics

	id, err := strconv.Atoi(fmt.Sprint(1))
	if id == -1 || err != nil {
		return diag.FromErr(err)
	}

	organization, err := c.GetOrganization()
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("auth_password_login_enabled", &organization.AuthPasswordLoginEnabled)
	d.Set("auth_saml_enabled", &organization.AuthSamlEnabled)
	d.Set("auth_saml_type", &organization.AuthSamlType)
	d.Set("auth_saml_entity_id", &organization.AuthSamlEntityId)
	d.Set("auth_saml_metadata_url", &organization.AuthSamlMetadataUrl)
	d.Set("auth_saml_nameid_format", &organization.AuthSamlNameidFormat)
	d.Set("auth_saml_sso_url", &organization.AuthSamlSsoUrl)

	d.SetId(fmt.Sprint(1))

	return diags
}

func resourceRedashOrganizationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*redash.Client)

	id, err := strconv.Atoi(fmt.Sprint(1))
	if id == -1 || err != nil {
		return diag.FromErr(err)
	}

	payload := redash.Organization{
		AuthPasswordLoginEnabled: d.Get("auth_password_login_enabled").(bool),
		AuthSamlEnabled: d.Get("auth_saml_enabled").(bool),
		AuthSamlType: d.Get("auth_saml_type").(string),
		AuthSamlEntityId: d.Get("auth_saml_entity_id").(string),
		AuthSamlMetadataUrl: d.Get("auth_saml_metadata_url").(string),
		AuthSamlNameidFormat: d.Get("auth_saml_nameid_format").(string),
		AuthSamlSsoUrl: d.Get("auth_saml_sso_url").(string),
	}

	_, err = c.UpdateOrganization(&payload)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceRedashOrganizationRead(ctx, d, meta)
}
func resourceRedashOrganizationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*redash.Client)

	var diags diag.Diagnostics

	payload := redash.Organization{
		AuthPasswordLoginEnabled: 	true,
		AuthSamlEnabled: 			false,
		AuthSamlType: 				"",
		AuthSamlEntityId: 			"",
		AuthSamlMetadataUrl: 		"",
		AuthSamlNameidFormat: 		"",
		AuthSamlSsoUrl: 			"",
	}

	organization, err := c.UpdateOrganization(&payload)
	if organization == nil || err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
