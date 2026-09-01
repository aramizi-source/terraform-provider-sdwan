// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type ServiceAppQoE struct {
	Id                           types.String                                `tfsdk:"id"`
	Version                      types.Int64                                 `tfsdk:"version"`
	Name                         types.String                                `tfsdk:"name"`
	Description                  types.String                                `tfsdk:"description"`
	FeatureProfileId             types.String                                `tfsdk:"feature_profile_id"`
	AppqoeDeviceRole             types.String                                `tfsdk:"appqoe_device_role"`
	VirtualApplications          []ServiceAppQoEVirtualApplications          `tfsdk:"virtual_applications"`
	ForwarderControllerGroups    []ServiceAppQoEForwarderControllerGroups    `tfsdk:"forwarder_controller_groups"`
	ForwarderServiceNodeGroups   []ServiceAppQoEForwarderServiceNodeGroups   `tfsdk:"forwarder_service_node_groups"`
	ForwarderServiceContexts     []ServiceAppQoEForwarderServiceContexts     `tfsdk:"forwarder_service_contexts"`
	CombinedControllerGroups     []ServiceAppQoECombinedControllerGroups     `tfsdk:"combined_controller_groups"`
	CombinedServiceNodeGroups    []ServiceAppQoECombinedServiceNodeGroups    `tfsdk:"combined_service_node_groups"`
	CombinedServiceContexts      []ServiceAppQoECombinedServiceContexts      `tfsdk:"combined_service_contexts"`
	ServiceNodeServiceNodeGroups []ServiceAppQoEServiceNodeServiceNodeGroups `tfsdk:"service_node_service_node_groups"`
}

type ServiceAppQoEVirtualApplications struct {
	ResourceProfile         types.String `tfsdk:"resource_profile"`
	ResourceProfileVariable types.String `tfsdk:"resource_profile_variable"`
}

type ServiceAppQoEForwarderControllerGroups struct {
	AppnavControllers []ServiceAppQoEForwarderControllerGroupsAppnavControllers `tfsdk:"appnav_controllers"`
}

type ServiceAppQoEForwarderServiceNodeGroups struct {
	Name         types.String                                          `tfsdk:"name"`
	ServiceNodes []ServiceAppQoEForwarderServiceNodeGroupsServiceNodes `tfsdk:"service_nodes"`
}

type ServiceAppQoEForwarderServiceContexts struct {
	AppnavControllerGroup types.String `tfsdk:"appnav_controller_group"`
	ServiceNodeGroup      types.String `tfsdk:"service_node_group"`
	Enable                types.Bool   `tfsdk:"enable"`
	Vpn                   types.Int64  `tfsdk:"vpn"`
	VpnVariable           types.String `tfsdk:"vpn_variable"`
}

type ServiceAppQoECombinedControllerGroups struct {
	GroupName         types.String                                             `tfsdk:"group_name"`
	AppnavControllers []ServiceAppQoECombinedControllerGroupsAppnavControllers `tfsdk:"appnav_controllers"`
}

type ServiceAppQoECombinedServiceNodeGroups struct {
	Name         types.String                                         `tfsdk:"name"`
	ServiceNodes []ServiceAppQoECombinedServiceNodeGroupsServiceNodes `tfsdk:"service_nodes"`
}

type ServiceAppQoECombinedServiceContexts struct {
	AppnavControllerGroup types.String `tfsdk:"appnav_controller_group"`
	ServiceNodeGroup      types.String `tfsdk:"service_node_group"`
	Enable                types.Bool   `tfsdk:"enable"`
	Vpn                   types.Int64  `tfsdk:"vpn"`
	VpnVariable           types.String `tfsdk:"vpn_variable"`
}

type ServiceAppQoEServiceNodeServiceNodeGroups struct {
	Name         types.String                                            `tfsdk:"name"`
	ServiceNodes []ServiceAppQoEServiceNodeServiceNodeGroupsServiceNodes `tfsdk:"service_nodes"`
}

type ServiceAppQoEForwarderControllerGroupsAppnavControllers struct {
	Address         types.String `tfsdk:"address"`
	AddressVariable types.String `tfsdk:"address_variable"`
	Vpn             types.Int64  `tfsdk:"vpn"`
}

type ServiceAppQoEForwarderServiceNodeGroupsServiceNodes struct {
	Address types.String `tfsdk:"address"`
}

type ServiceAppQoECombinedControllerGroupsAppnavControllers struct {
	Address types.String `tfsdk:"address"`
}

type ServiceAppQoECombinedServiceNodeGroupsServiceNodes struct {
	Address types.String `tfsdk:"address"`
}

type ServiceAppQoEServiceNodeServiceNodeGroupsServiceNodes struct {
	Address types.String `tfsdk:"address"`
	VpgIp   types.String `tfsdk:"vpg_ip"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data ServiceAppQoE) getModel() string {
	return "service_appqoe"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ServiceAppQoE) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/service/%v/appqoe", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ServiceAppQoE) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if !data.AppqoeDeviceRole.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"appqoeDeviceRole.optionType", "global")
			body, _ = sjson.Set(body, path+"appqoeDeviceRole.value", data.AppqoeDeviceRole.ValueString())
		}
	}
	if true && strings.Contains(data.AppqoeDeviceRole.ValueString(), "Dre") {
		body, _ = sjson.Set(body, path+"dreopt.optionType", "global")
		body, _ = sjson.Set(body, path+"dreopt.value", true)
	}
	if true && !(strings.Contains(data.AppqoeDeviceRole.ValueString(), "Dre")) {
		body, _ = sjson.Set(body, path+"dreopt.optionType", "global")
		body, _ = sjson.Set(body, path+"dreopt.value", false)
	}
	if true && strings.Contains(data.AppqoeDeviceRole.ValueString(), "Dre") {
		body, _ = sjson.Set(body, path+"virtualApplication", []interface{}{})
		for _, item := range data.VirtualApplications {
			itemBody := ""
			if true {
				itemBody, _ = sjson.Set(itemBody, "instanceId.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "instanceId.value", 1)
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "applicationType.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "applicationType.value", "dreopt")
			}

			if !item.ResourceProfileVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "resourceProfile.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "resourceProfile.value", item.ResourceProfileVariable.ValueString())
				}
			} else if item.ResourceProfile.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "resourceProfile.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "resourceProfile.value", "default")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "resourceProfile.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "resourceProfile.value", item.ResourceProfile.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"virtualApplication.-1", itemBody)
		}
	}
	if true && data.AppqoeDeviceRole.ValueString() == "forwarder" {
		body, _ = sjson.Set(body, path+"forwarder.appnavControllerGroup", []interface{}{})
		for _, item := range data.ForwarderControllerGroups {
			itemBody := ""
			if true {
				itemBody, _ = sjson.Set(itemBody, "groupName.optionType", "default")
				itemBody, _ = sjson.Set(itemBody, "groupName.value", "ACG-APPQOE")
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "appnavControllers", []interface{}{})
				for _, childItem := range item.AppnavControllers {
					itemChildBody := ""

					if !childItem.AddressVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.AddressVariable.ValueString())
						}
					} else if !childItem.Address.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.Address.ValueString())
						}
					}
					if !childItem.Vpn.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "vpn.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "vpn.value", childItem.Vpn.ValueInt64())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "appnavControllers.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"forwarder.appnavControllerGroup.-1", itemBody)
		}
	}
	if true && data.AppqoeDeviceRole.ValueString() == "forwarder" {
		body, _ = sjson.Set(body, path+"forwarder.serviceNodeGroup", []interface{}{})
		for _, item := range data.ForwarderServiceNodeGroups {
			itemBody := ""
			if item.Name.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "name.value", "SNG-APPQOE")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.Name.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "internal.optionType", "default")
				itemBody, _ = sjson.Set(itemBody, "internal.value", false)
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "serviceNode", []interface{}{})
				for _, childItem := range item.ServiceNodes {
					itemChildBody := ""
					if !childItem.Address.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.Address.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "serviceNode.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"forwarder.serviceNodeGroup.-1", itemBody)
		}
	}
	if true && data.AppqoeDeviceRole.ValueString() == "forwarder" {
		body, _ = sjson.Set(body, path+"forwarder.serviceContext.appqoe", []interface{}{})
		for _, item := range data.ForwarderServiceContexts {
			itemBody := ""
			if true {
				itemBody, _ = sjson.Set(itemBody, "name.optionType", "default")
				itemBody, _ = sjson.Set(itemBody, "name.value", "/1")
			}
			if !item.AppnavControllerGroup.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "appnavControllerGroup.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "appnavControllerGroup.value", item.AppnavControllerGroup.ValueString())
				}
			}
			if !item.ServiceNodeGroup.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "serviceNodeGroup.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "serviceNodeGroup.value", item.ServiceNodeGroup.ValueString())
				}
			}
			if !item.Enable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "enable.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "enable.value", item.Enable.ValueBool())
				}
			}

			if !item.VpnVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", item.VpnVariable.ValueString())
				}
			} else if item.Vpn.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", item.Vpn.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"forwarder.serviceContext.appqoe.-1", itemBody)
		}
	}
	if true && strings.Contains(data.AppqoeDeviceRole.ValueString(), "forwarderAndServiceNode") {
		body, _ = sjson.Set(body, path+"forwarderAndServiceNode.appnavControllerGroup", []interface{}{})
		for _, item := range data.CombinedControllerGroups {
			itemBody := ""
			if item.GroupName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "groupName.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "groupName.value", "ACG-APPQOE")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "groupName.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "groupName.value", item.GroupName.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "appnavControllers", []interface{}{})
				for _, childItem := range item.AppnavControllers {
					itemChildBody := ""
					if childItem.Address.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", "192.168.2.1")
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.Address.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "appnavControllers.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"forwarderAndServiceNode.appnavControllerGroup.-1", itemBody)
		}
	}
	if true && strings.Contains(data.AppqoeDeviceRole.ValueString(), "forwarderAndServiceNode") {
		body, _ = sjson.Set(body, path+"forwarderAndServiceNode.serviceNodeGroup", []interface{}{})
		for _, item := range data.CombinedServiceNodeGroups {
			itemBody := ""
			if item.Name.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "name.value", "SNG-APPQOE")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.Name.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "internal.optionType", "default")
				itemBody, _ = sjson.Set(itemBody, "internal.value", true)
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "serviceNode", []interface{}{})
				for _, childItem := range item.ServiceNodes {
					itemChildBody := ""
					if childItem.Address.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", "192.168.2.2")
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.Address.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "serviceNode.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"forwarderAndServiceNode.serviceNodeGroup.-1", itemBody)
		}
	}
	if true && strings.Contains(data.AppqoeDeviceRole.ValueString(), "forwarderAndServiceNode") {
		body, _ = sjson.Set(body, path+"forwarderAndServiceNode.serviceContext.appqoe", []interface{}{})
		for _, item := range data.CombinedServiceContexts {
			itemBody := ""
			if true {
				itemBody, _ = sjson.Set(itemBody, "name.optionType", "default")
				itemBody, _ = sjson.Set(itemBody, "name.value", "/1")
			}
			if !item.AppnavControllerGroup.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "appnavControllerGroup.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "appnavControllerGroup.value", item.AppnavControllerGroup.ValueString())
				}
			}
			if !item.ServiceNodeGroup.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "serviceNodeGroup.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "serviceNodeGroup.value", item.ServiceNodeGroup.ValueString())
				}
			}
			if !item.Enable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "enable.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "enable.value", item.Enable.ValueBool())
				}
			}

			if !item.VpnVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", item.VpnVariable.ValueString())
				}
			} else if item.Vpn.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", item.Vpn.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"forwarderAndServiceNode.serviceContext.appqoe.-1", itemBody)
		}
	}
	if true && strings.Contains(data.AppqoeDeviceRole.ValueString(), "serviceNode") {
		body, _ = sjson.Set(body, path+"serviceNode.serviceNodeGroup", []interface{}{})
		for _, item := range data.ServiceNodeServiceNodeGroups {
			itemBody := ""
			if item.Name.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "name.value", "SNG-APPQOE")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.Name.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "externalNode.optionType", "default")
				itemBody, _ = sjson.Set(itemBody, "externalNode.value", true)
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "serviceNode", []interface{}{})
				for _, childItem := range item.ServiceNodes {
					itemChildBody := ""
					if childItem.Address.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", "192.168.2.2")
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.Address.ValueString())
						}
					}
					if childItem.VpgIp.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "vpgIp.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "vpgIp.value", "192.168.2.1/24")
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "vpgIp.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "vpgIp.value", childItem.VpgIp.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "serviceNode.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"serviceNode.serviceNodeGroup.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ServiceAppQoE) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.AppqoeDeviceRole = types.StringNull()

	if t := res.Get(path + "appqoeDeviceRole.optionType"); t.Exists() {
		va := res.Get(path + "appqoeDeviceRole.value")
		if t.String() == "global" {
			data.AppqoeDeviceRole = types.StringValue(va.String())
		}
	}
	if value := res.Get(path + "virtualApplication"); value.Exists() && len(value.Array()) > 0 {
		data.VirtualApplications = make([]ServiceAppQoEVirtualApplications, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceAppQoEVirtualApplications{}
			item.ResourceProfile = types.StringNull()
			item.ResourceProfileVariable = types.StringNull()
			if t := v.Get("resourceProfile.optionType"); t.Exists() {
				va := v.Get("resourceProfile.value")
				if t.String() == "variable" {
					item.ResourceProfileVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.ResourceProfile = types.StringValue(va.String())
				}
			}
			data.VirtualApplications = append(data.VirtualApplications, item)
			return true
		})
	}
	if value := res.Get(path + "forwarder.appnavControllerGroup"); value.Exists() && len(value.Array()) > 0 {
		data.ForwarderControllerGroups = make([]ServiceAppQoEForwarderControllerGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceAppQoEForwarderControllerGroups{}
			if cValue := v.Get("appnavControllers"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.AppnavControllers = make([]ServiceAppQoEForwarderControllerGroupsAppnavControllers, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceAppQoEForwarderControllerGroupsAppnavControllers{}
					cItem.Address = types.StringNull()
					cItem.AddressVariable = types.StringNull()
					if t := cv.Get("address.optionType"); t.Exists() {
						va := cv.Get("address.value")
						if t.String() == "variable" {
							cItem.AddressVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Address = types.StringValue(va.String())
						}
					}
					cItem.Vpn = types.Int64Null()

					if t := cv.Get("vpn.optionType"); t.Exists() {
						va := cv.Get("vpn.value")
						if t.String() == "global" {
							cItem.Vpn = types.Int64Value(va.Int())
						}
					}
					item.AppnavControllers = append(item.AppnavControllers, cItem)
					return true
				})
			}
			data.ForwarderControllerGroups = append(data.ForwarderControllerGroups, item)
			return true
		})
	}
	if value := res.Get(path + "forwarder.serviceNodeGroup"); value.Exists() && len(value.Array()) > 0 {
		data.ForwarderServiceNodeGroups = make([]ServiceAppQoEForwarderServiceNodeGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceAppQoEForwarderServiceNodeGroups{}
			item.Name = types.StringNull()

			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "global" {
					item.Name = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("serviceNode"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.ServiceNodes = make([]ServiceAppQoEForwarderServiceNodeGroupsServiceNodes, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceAppQoEForwarderServiceNodeGroupsServiceNodes{}
					cItem.Address = types.StringNull()

					if t := cv.Get("address.optionType"); t.Exists() {
						va := cv.Get("address.value")
						if t.String() == "global" {
							cItem.Address = types.StringValue(va.String())
						}
					}
					item.ServiceNodes = append(item.ServiceNodes, cItem)
					return true
				})
			}
			data.ForwarderServiceNodeGroups = append(data.ForwarderServiceNodeGroups, item)
			return true
		})
	}
	if value := res.Get(path + "forwarder.serviceContext.appqoe"); value.Exists() && len(value.Array()) > 0 {
		data.ForwarderServiceContexts = make([]ServiceAppQoEForwarderServiceContexts, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceAppQoEForwarderServiceContexts{}
			item.AppnavControllerGroup = types.StringNull()

			if t := v.Get("appnavControllerGroup.optionType"); t.Exists() {
				va := v.Get("appnavControllerGroup.value")
				if t.String() == "global" {
					item.AppnavControllerGroup = types.StringValue(va.String())
				}
			}
			item.ServiceNodeGroup = types.StringNull()

			if t := v.Get("serviceNodeGroup.optionType"); t.Exists() {
				va := v.Get("serviceNodeGroup.value")
				if t.String() == "global" {
					item.ServiceNodeGroup = types.StringValue(va.String())
				}
			}
			item.Enable = types.BoolNull()

			if t := v.Get("enable.optionType"); t.Exists() {
				va := v.Get("enable.value")
				if t.String() == "global" {
					item.Enable = types.BoolValue(va.Bool())
				}
			}
			item.Vpn = types.Int64Null()
			item.VpnVariable = types.StringNull()
			if t := v.Get("vpn.optionType"); t.Exists() {
				va := v.Get("vpn.value")
				if t.String() == "variable" {
					item.VpnVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Vpn = types.Int64Value(va.Int())
				}
			}
			data.ForwarderServiceContexts = append(data.ForwarderServiceContexts, item)
			return true
		})
	}
	if value := res.Get(path + "forwarderAndServiceNode.appnavControllerGroup"); value.Exists() && len(value.Array()) > 0 {
		data.CombinedControllerGroups = make([]ServiceAppQoECombinedControllerGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceAppQoECombinedControllerGroups{}
			item.GroupName = types.StringNull()

			if t := v.Get("groupName.optionType"); t.Exists() {
				va := v.Get("groupName.value")
				if t.String() == "global" || t.String() == "default" {
					item.GroupName = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("appnavControllers"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.AppnavControllers = make([]ServiceAppQoECombinedControllerGroupsAppnavControllers, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceAppQoECombinedControllerGroupsAppnavControllers{}
					cItem.Address = types.StringNull()

					if t := cv.Get("address.optionType"); t.Exists() {
						va := cv.Get("address.value")
						if t.String() == "global" || t.String() == "default" {
							cItem.Address = types.StringValue(va.String())
						}
					}
					item.AppnavControllers = append(item.AppnavControllers, cItem)
					return true
				})
			}
			data.CombinedControllerGroups = append(data.CombinedControllerGroups, item)
			return true
		})
	}
	if value := res.Get(path + "forwarderAndServiceNode.serviceNodeGroup"); value.Exists() && len(value.Array()) > 0 {
		data.CombinedServiceNodeGroups = make([]ServiceAppQoECombinedServiceNodeGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceAppQoECombinedServiceNodeGroups{}
			item.Name = types.StringNull()

			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "global" || t.String() == "default" {
					item.Name = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("serviceNode"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.ServiceNodes = make([]ServiceAppQoECombinedServiceNodeGroupsServiceNodes, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceAppQoECombinedServiceNodeGroupsServiceNodes{}
					cItem.Address = types.StringNull()

					if t := cv.Get("address.optionType"); t.Exists() {
						va := cv.Get("address.value")
						if t.String() == "global" || t.String() == "default" {
							cItem.Address = types.StringValue(va.String())
						}
					}
					item.ServiceNodes = append(item.ServiceNodes, cItem)
					return true
				})
			}
			data.CombinedServiceNodeGroups = append(data.CombinedServiceNodeGroups, item)
			return true
		})
	}
	if value := res.Get(path + "forwarderAndServiceNode.serviceContext.appqoe"); value.Exists() && len(value.Array()) > 0 {
		data.CombinedServiceContexts = make([]ServiceAppQoECombinedServiceContexts, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceAppQoECombinedServiceContexts{}
			item.AppnavControllerGroup = types.StringNull()

			if t := v.Get("appnavControllerGroup.optionType"); t.Exists() {
				va := v.Get("appnavControllerGroup.value")
				if t.String() == "global" {
					item.AppnavControllerGroup = types.StringValue(va.String())
				}
			}
			item.ServiceNodeGroup = types.StringNull()

			if t := v.Get("serviceNodeGroup.optionType"); t.Exists() {
				va := v.Get("serviceNodeGroup.value")
				if t.String() == "global" {
					item.ServiceNodeGroup = types.StringValue(va.String())
				}
			}
			item.Enable = types.BoolNull()

			if t := v.Get("enable.optionType"); t.Exists() {
				va := v.Get("enable.value")
				if t.String() == "global" {
					item.Enable = types.BoolValue(va.Bool())
				}
			}
			item.Vpn = types.Int64Null()
			item.VpnVariable = types.StringNull()
			if t := v.Get("vpn.optionType"); t.Exists() {
				va := v.Get("vpn.value")
				if t.String() == "variable" {
					item.VpnVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Vpn = types.Int64Value(va.Int())
				}
			}
			data.CombinedServiceContexts = append(data.CombinedServiceContexts, item)
			return true
		})
	}
	if value := res.Get(path + "serviceNode.serviceNodeGroup"); value.Exists() && len(value.Array()) > 0 {
		data.ServiceNodeServiceNodeGroups = make([]ServiceAppQoEServiceNodeServiceNodeGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceAppQoEServiceNodeServiceNodeGroups{}
			item.Name = types.StringNull()

			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "global" || t.String() == "default" {
					item.Name = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("serviceNode"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.ServiceNodes = make([]ServiceAppQoEServiceNodeServiceNodeGroupsServiceNodes, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceAppQoEServiceNodeServiceNodeGroupsServiceNodes{}
					cItem.Address = types.StringNull()

					if t := cv.Get("address.optionType"); t.Exists() {
						va := cv.Get("address.value")
						if t.String() == "global" || t.String() == "default" {
							cItem.Address = types.StringValue(va.String())
						}
					}
					cItem.VpgIp = types.StringNull()

					if t := cv.Get("vpgIp.optionType"); t.Exists() {
						va := cv.Get("vpgIp.value")
						if t.String() == "global" || t.String() == "default" {
							cItem.VpgIp = types.StringValue(va.String())
						}
					}
					item.ServiceNodes = append(item.ServiceNodes, cItem)
					return true
				})
			}
			data.ServiceNodeServiceNodeGroups = append(data.ServiceNodeServiceNodeGroups, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *ServiceAppQoE) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.AppqoeDeviceRole = types.StringNull()

	if t := res.Get(path + "appqoeDeviceRole.optionType"); t.Exists() {
		va := res.Get(path + "appqoeDeviceRole.value")
		if t.String() == "global" {
			data.AppqoeDeviceRole = types.StringValue(va.String())
		}
	}
	for i := range data.VirtualApplications {
		keys := [...]string{"resourceProfile"}
		keyValues := [...]string{data.VirtualApplications[i].ResourceProfile.ValueString()}
		keyValuesVariables := [...]string{data.VirtualApplications[i].ResourceProfileVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "virtualApplication").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
							continue
						}
						found = false
						break
					}
					continue
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if !r.Exists() {
			arr := res.Get(path + "virtualApplication").Array()
			if i < len(arr) {
				r = arr[i]
			}
		}
		data.VirtualApplications[i].ResourceProfile = types.StringNull()
		data.VirtualApplications[i].ResourceProfileVariable = types.StringNull()
		if t := r.Get("resourceProfile.optionType"); t.Exists() {
			va := r.Get("resourceProfile.value")
			if t.String() == "variable" {
				data.VirtualApplications[i].ResourceProfileVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.VirtualApplications[i].ResourceProfile = types.StringValue(va.String())
			}
		}
	}
	for i := range data.ForwarderControllerGroups {
		keys := [...]string{}
		keyValues := [...]string{}
		keyValuesVariables := [...]string{}

		var r gjson.Result
		res.Get(path + "forwarder.appnavControllerGroup").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
							continue
						}
						found = false
						break
					}
					continue
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if !r.Exists() {
			arr := res.Get(path + "forwarder.appnavControllerGroup").Array()
			if i < len(arr) {
				r = arr[i]
			}
		}
		for ci := range data.ForwarderControllerGroups[i].AppnavControllers {
			keys := [...]string{"address"}
			keyValues := [...]string{data.ForwarderControllerGroups[i].AppnavControllers[ci].Address.ValueString()}
			keyValuesVariables := [...]string{data.ForwarderControllerGroups[i].AppnavControllers[ci].AddressVariable.ValueString()}

			var cr gjson.Result
			r.Get("appnavControllers").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType")
						vv := v.Get(keys[ik] + ".value")
						if tt.Exists() && vv.Exists() {
							if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
								found = true
								continue
							} else if tt.String() == "default" {
								continue
							}
							found = false
							break
						}
						continue
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			if !cr.Exists() {
				arr := r.Get("appnavControllers").Array()
				if ci < len(arr) {
					cr = arr[ci]
				}
			}
			data.ForwarderControllerGroups[i].AppnavControllers[ci].Address = types.StringNull()
			data.ForwarderControllerGroups[i].AppnavControllers[ci].AddressVariable = types.StringNull()
			if t := cr.Get("address.optionType"); t.Exists() {
				va := cr.Get("address.value")
				if t.String() == "variable" {
					data.ForwarderControllerGroups[i].AppnavControllers[ci].AddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.ForwarderControllerGroups[i].AppnavControllers[ci].Address = types.StringValue(va.String())
				}
			}
			data.ForwarderControllerGroups[i].AppnavControllers[ci].Vpn = types.Int64Null()

			if t := cr.Get("vpn.optionType"); t.Exists() {
				va := cr.Get("vpn.value")
				if t.String() == "global" {
					data.ForwarderControllerGroups[i].AppnavControllers[ci].Vpn = types.Int64Value(va.Int())
				}
			}
		}
	}
	for i := range data.ForwarderServiceNodeGroups {
		keys := [...]string{"name"}
		keyValues := [...]string{data.ForwarderServiceNodeGroups[i].Name.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "forwarder.serviceNodeGroup").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
							continue
						}
						found = false
						break
					}
					continue
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if !r.Exists() {
			arr := res.Get(path + "forwarder.serviceNodeGroup").Array()
			if i < len(arr) {
				r = arr[i]
			}
		}
		data.ForwarderServiceNodeGroups[i].Name = types.StringNull()

		if t := r.Get("name.optionType"); t.Exists() {
			va := r.Get("name.value")
			if t.String() == "global" {
				data.ForwarderServiceNodeGroups[i].Name = types.StringValue(va.String())
			}
		}
		for ci := range data.ForwarderServiceNodeGroups[i].ServiceNodes {
			keys := [...]string{"address"}
			keyValues := [...]string{data.ForwarderServiceNodeGroups[i].ServiceNodes[ci].Address.ValueString()}
			keyValuesVariables := [...]string{""}

			var cr gjson.Result
			r.Get("serviceNode").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType")
						vv := v.Get(keys[ik] + ".value")
						if tt.Exists() && vv.Exists() {
							if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
								found = true
								continue
							} else if tt.String() == "default" {
								continue
							}
							found = false
							break
						}
						continue
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			if !cr.Exists() {
				arr := r.Get("serviceNode").Array()
				if ci < len(arr) {
					cr = arr[ci]
				}
			}
			data.ForwarderServiceNodeGroups[i].ServiceNodes[ci].Address = types.StringNull()

			if t := cr.Get("address.optionType"); t.Exists() {
				va := cr.Get("address.value")
				if t.String() == "global" {
					data.ForwarderServiceNodeGroups[i].ServiceNodes[ci].Address = types.StringValue(va.String())
				}
			}
		}
	}
	for i := range data.ForwarderServiceContexts {
		keys := [...]string{"appnavControllerGroup"}
		keyValues := [...]string{data.ForwarderServiceContexts[i].AppnavControllerGroup.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "forwarder.serviceContext.appqoe").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
							continue
						}
						found = false
						break
					}
					continue
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if !r.Exists() {
			arr := res.Get(path + "forwarder.serviceContext.appqoe").Array()
			if i < len(arr) {
				r = arr[i]
			}
		}
		data.ForwarderServiceContexts[i].AppnavControllerGroup = types.StringNull()

		if t := r.Get("appnavControllerGroup.optionType"); t.Exists() {
			va := r.Get("appnavControllerGroup.value")
			if t.String() == "global" {
				data.ForwarderServiceContexts[i].AppnavControllerGroup = types.StringValue(va.String())
			}
		}
		data.ForwarderServiceContexts[i].ServiceNodeGroup = types.StringNull()

		if t := r.Get("serviceNodeGroup.optionType"); t.Exists() {
			va := r.Get("serviceNodeGroup.value")
			if t.String() == "global" {
				data.ForwarderServiceContexts[i].ServiceNodeGroup = types.StringValue(va.String())
			}
		}
		data.ForwarderServiceContexts[i].Enable = types.BoolNull()

		if t := r.Get("enable.optionType"); t.Exists() {
			va := r.Get("enable.value")
			if t.String() == "global" {
				data.ForwarderServiceContexts[i].Enable = types.BoolValue(va.Bool())
			}
		}
		data.ForwarderServiceContexts[i].Vpn = types.Int64Null()
		data.ForwarderServiceContexts[i].VpnVariable = types.StringNull()
		if t := r.Get("vpn.optionType"); t.Exists() {
			va := r.Get("vpn.value")
			if t.String() == "variable" {
				data.ForwarderServiceContexts[i].VpnVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.ForwarderServiceContexts[i].Vpn = types.Int64Value(va.Int())
			}
		}
	}
	for i := range data.CombinedControllerGroups {
		keys := [...]string{"groupName"}
		keyValues := [...]string{data.CombinedControllerGroups[i].GroupName.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "forwarderAndServiceNode.appnavControllerGroup").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
							continue
						}
						found = false
						break
					}
					continue
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if !r.Exists() {
			arr := res.Get(path + "forwarderAndServiceNode.appnavControllerGroup").Array()
			if i < len(arr) {
				r = arr[i]
			}
		}
		tempGroupName := data.CombinedControllerGroups[i].GroupName
		data.CombinedControllerGroups[i].GroupName = types.StringNull()

		if t := r.Get("groupName.optionType"); t.Exists() {
			va := r.Get("groupName.value")
			if t.String() == "global" || (t.String() == "default" && !tempGroupName.IsNull()) {
				data.CombinedControllerGroups[i].GroupName = types.StringValue(va.String())
			}
		}
		for ci := range data.CombinedControllerGroups[i].AppnavControllers {
			keys := [...]string{"address"}
			keyValues := [...]string{data.CombinedControllerGroups[i].AppnavControllers[ci].Address.ValueString()}
			keyValuesVariables := [...]string{""}

			var cr gjson.Result
			r.Get("appnavControllers").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType")
						vv := v.Get(keys[ik] + ".value")
						if tt.Exists() && vv.Exists() {
							if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
								found = true
								continue
							} else if tt.String() == "default" {
								continue
							}
							found = false
							break
						}
						continue
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			if !cr.Exists() {
				arr := r.Get("appnavControllers").Array()
				if ci < len(arr) {
					cr = arr[ci]
				}
			}
			tempAddress := data.CombinedControllerGroups[i].AppnavControllers[ci].Address
			data.CombinedControllerGroups[i].AppnavControllers[ci].Address = types.StringNull()

			if t := cr.Get("address.optionType"); t.Exists() {
				va := cr.Get("address.value")
				if t.String() == "global" || (t.String() == "default" && !tempAddress.IsNull()) {
					data.CombinedControllerGroups[i].AppnavControllers[ci].Address = types.StringValue(va.String())
				}
			}
		}
	}
	for i := range data.CombinedServiceNodeGroups {
		keys := [...]string{"name"}
		keyValues := [...]string{data.CombinedServiceNodeGroups[i].Name.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "forwarderAndServiceNode.serviceNodeGroup").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
							continue
						}
						found = false
						break
					}
					continue
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if !r.Exists() {
			arr := res.Get(path + "forwarderAndServiceNode.serviceNodeGroup").Array()
			if i < len(arr) {
				r = arr[i]
			}
		}
		tempName := data.CombinedServiceNodeGroups[i].Name
		data.CombinedServiceNodeGroups[i].Name = types.StringNull()

		if t := r.Get("name.optionType"); t.Exists() {
			va := r.Get("name.value")
			if t.String() == "global" || (t.String() == "default" && !tempName.IsNull()) {
				data.CombinedServiceNodeGroups[i].Name = types.StringValue(va.String())
			}
		}
		for ci := range data.CombinedServiceNodeGroups[i].ServiceNodes {
			keys := [...]string{"address"}
			keyValues := [...]string{data.CombinedServiceNodeGroups[i].ServiceNodes[ci].Address.ValueString()}
			keyValuesVariables := [...]string{""}

			var cr gjson.Result
			r.Get("serviceNode").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType")
						vv := v.Get(keys[ik] + ".value")
						if tt.Exists() && vv.Exists() {
							if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
								found = true
								continue
							} else if tt.String() == "default" {
								continue
							}
							found = false
							break
						}
						continue
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			if !cr.Exists() {
				arr := r.Get("serviceNode").Array()
				if ci < len(arr) {
					cr = arr[ci]
				}
			}
			tempAddress := data.CombinedServiceNodeGroups[i].ServiceNodes[ci].Address
			data.CombinedServiceNodeGroups[i].ServiceNodes[ci].Address = types.StringNull()

			if t := cr.Get("address.optionType"); t.Exists() {
				va := cr.Get("address.value")
				if t.String() == "global" || (t.String() == "default" && !tempAddress.IsNull()) {
					data.CombinedServiceNodeGroups[i].ServiceNodes[ci].Address = types.StringValue(va.String())
				}
			}
		}
	}
	for i := range data.CombinedServiceContexts {
		keys := [...]string{"appnavControllerGroup"}
		keyValues := [...]string{data.CombinedServiceContexts[i].AppnavControllerGroup.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "forwarderAndServiceNode.serviceContext.appqoe").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
							continue
						}
						found = false
						break
					}
					continue
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if !r.Exists() {
			arr := res.Get(path + "forwarderAndServiceNode.serviceContext.appqoe").Array()
			if i < len(arr) {
				r = arr[i]
			}
		}
		data.CombinedServiceContexts[i].AppnavControllerGroup = types.StringNull()

		if t := r.Get("appnavControllerGroup.optionType"); t.Exists() {
			va := r.Get("appnavControllerGroup.value")
			if t.String() == "global" {
				data.CombinedServiceContexts[i].AppnavControllerGroup = types.StringValue(va.String())
			}
		}
		data.CombinedServiceContexts[i].ServiceNodeGroup = types.StringNull()

		if t := r.Get("serviceNodeGroup.optionType"); t.Exists() {
			va := r.Get("serviceNodeGroup.value")
			if t.String() == "global" {
				data.CombinedServiceContexts[i].ServiceNodeGroup = types.StringValue(va.String())
			}
		}
		data.CombinedServiceContexts[i].Enable = types.BoolNull()

		if t := r.Get("enable.optionType"); t.Exists() {
			va := r.Get("enable.value")
			if t.String() == "global" {
				data.CombinedServiceContexts[i].Enable = types.BoolValue(va.Bool())
			}
		}
		data.CombinedServiceContexts[i].Vpn = types.Int64Null()
		data.CombinedServiceContexts[i].VpnVariable = types.StringNull()
		if t := r.Get("vpn.optionType"); t.Exists() {
			va := r.Get("vpn.value")
			if t.String() == "variable" {
				data.CombinedServiceContexts[i].VpnVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.CombinedServiceContexts[i].Vpn = types.Int64Value(va.Int())
			}
		}
	}
	for i := range data.ServiceNodeServiceNodeGroups {
		keys := [...]string{"name"}
		keyValues := [...]string{data.ServiceNodeServiceNodeGroups[i].Name.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "serviceNode.serviceNodeGroup").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
							continue
						}
						found = false
						break
					}
					continue
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if !r.Exists() {
			arr := res.Get(path + "serviceNode.serviceNodeGroup").Array()
			if i < len(arr) {
				r = arr[i]
			}
		}
		tempName := data.ServiceNodeServiceNodeGroups[i].Name
		data.ServiceNodeServiceNodeGroups[i].Name = types.StringNull()

		if t := r.Get("name.optionType"); t.Exists() {
			va := r.Get("name.value")
			if t.String() == "global" || (t.String() == "default" && !tempName.IsNull()) {
				data.ServiceNodeServiceNodeGroups[i].Name = types.StringValue(va.String())
			}
		}
		for ci := range data.ServiceNodeServiceNodeGroups[i].ServiceNodes {
			keys := [...]string{"address"}
			keyValues := [...]string{data.ServiceNodeServiceNodeGroups[i].ServiceNodes[ci].Address.ValueString()}
			keyValuesVariables := [...]string{""}

			var cr gjson.Result
			r.Get("serviceNode").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType")
						vv := v.Get(keys[ik] + ".value")
						if tt.Exists() && vv.Exists() {
							if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
								found = true
								continue
							} else if tt.String() == "default" {
								continue
							}
							found = false
							break
						}
						continue
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			if !cr.Exists() {
				arr := r.Get("serviceNode").Array()
				if ci < len(arr) {
					cr = arr[ci]
				}
			}
			tempAddress := data.ServiceNodeServiceNodeGroups[i].ServiceNodes[ci].Address
			data.ServiceNodeServiceNodeGroups[i].ServiceNodes[ci].Address = types.StringNull()

			if t := cr.Get("address.optionType"); t.Exists() {
				va := cr.Get("address.value")
				if t.String() == "global" || (t.String() == "default" && !tempAddress.IsNull()) {
					data.ServiceNodeServiceNodeGroups[i].ServiceNodes[ci].Address = types.StringValue(va.String())
				}
			}
			tempVpgIp := data.ServiceNodeServiceNodeGroups[i].ServiceNodes[ci].VpgIp
			data.ServiceNodeServiceNodeGroups[i].ServiceNodes[ci].VpgIp = types.StringNull()

			if t := cr.Get("vpgIp.optionType"); t.Exists() {
				va := cr.Get("vpgIp.value")
				if t.String() == "global" || (t.String() == "default" && !tempVpgIp.IsNull()) {
					data.ServiceNodeServiceNodeGroups[i].ServiceNodes[ci].VpgIp = types.StringValue(va.String())
				}
			}
		}
	}
}

// End of section. //template:end updateFromBody
