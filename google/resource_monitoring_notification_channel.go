// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMonitoringNotificationChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceMonitoringNotificationChannelCreate,
		Read:   resourceMonitoringNotificationChannelRead,
		Update: resourceMonitoringNotificationChannelUpdate,
		Delete: resourceMonitoringNotificationChannelDelete,

		Importer: &schema.ResourceImporter{
			State: resourceMonitoringNotificationChannelImport,
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"user_labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"verification_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceMonitoringNotificationChannelCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	labelsProp, err := expandMonitoringNotificationChannelLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	typeProp, err := expandMonitoringNotificationChannelType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	userLabelsProp, err := expandMonitoringNotificationChannelUserLabels(d.Get("user_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_labels"); !isEmptyValue(reflect.ValueOf(userLabelsProp)) && (ok || !reflect.DeepEqual(v, userLabelsProp)) {
		obj["userLabels"] = userLabelsProp
	}
	descriptionProp, err := expandMonitoringNotificationChannelDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	displayNameProp, err := expandMonitoringNotificationChannelDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	enabledProp, err := expandMonitoringNotificationChannelEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !isEmptyValue(reflect.ValueOf(enabledProp)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}

	lockName, err := replaceVars(d, config, "stackdriver/notifications/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "https://monitoring.googleapis.com/v3/projects/{{project}}/notificationChannels")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new NotificationChannel: %#v", obj)
	res, err := sendRequest(config, "POST", url, obj)
	if err != nil {
		return fmt.Errorf("Error creating NotificationChannel: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating NotificationChannel %q: %#v", d.Id(), res)

	// `name` is autogenerated from the api so needs to be set post-create
	name, ok := res["name"]
	if !ok {
		return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
	}
	d.Set("name", name.(string))
	d.SetId(name.(string))

	return resourceMonitoringNotificationChannelRead(d, meta)
}

func resourceMonitoringNotificationChannelRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://monitoring.googleapis.com/v3/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("MonitoringNotificationChannel %q", d.Id()))
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading NotificationChannel: %s", err)
	}

	if err := d.Set("labels", flattenMonitoringNotificationChannelLabels(res["labels"])); err != nil {
		return fmt.Errorf("Error reading NotificationChannel: %s", err)
	}
	if err := d.Set("name", flattenMonitoringNotificationChannelName(res["name"])); err != nil {
		return fmt.Errorf("Error reading NotificationChannel: %s", err)
	}
	if err := d.Set("verification_status", flattenMonitoringNotificationChannelVerificationStatus(res["verificationStatus"])); err != nil {
		return fmt.Errorf("Error reading NotificationChannel: %s", err)
	}
	if err := d.Set("type", flattenMonitoringNotificationChannelType(res["type"])); err != nil {
		return fmt.Errorf("Error reading NotificationChannel: %s", err)
	}
	if err := d.Set("user_labels", flattenMonitoringNotificationChannelUserLabels(res["userLabels"])); err != nil {
		return fmt.Errorf("Error reading NotificationChannel: %s", err)
	}
	if err := d.Set("description", flattenMonitoringNotificationChannelDescription(res["description"])); err != nil {
		return fmt.Errorf("Error reading NotificationChannel: %s", err)
	}
	if err := d.Set("display_name", flattenMonitoringNotificationChannelDisplayName(res["displayName"])); err != nil {
		return fmt.Errorf("Error reading NotificationChannel: %s", err)
	}
	if err := d.Set("enabled", flattenMonitoringNotificationChannelEnabled(res["enabled"])); err != nil {
		return fmt.Errorf("Error reading NotificationChannel: %s", err)
	}

	return nil
}

func resourceMonitoringNotificationChannelUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	labelsProp, err := expandMonitoringNotificationChannelLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	typeProp, err := expandMonitoringNotificationChannelType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	userLabelsProp, err := expandMonitoringNotificationChannelUserLabels(d.Get("user_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, userLabelsProp)) {
		obj["userLabels"] = userLabelsProp
	}
	descriptionProp, err := expandMonitoringNotificationChannelDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	displayNameProp, err := expandMonitoringNotificationChannelDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	enabledProp, err := expandMonitoringNotificationChannelEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}

	lockName, err := replaceVars(d, config, "stackdriver/notifications/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "https://monitoring.googleapis.com/v3/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating NotificationChannel %q: %#v", d.Id(), obj)
	_, err = sendRequest(config, "PATCH", url, obj)

	if err != nil {
		return fmt.Errorf("Error updating NotificationChannel %q: %s", d.Id(), err)
	}

	return resourceMonitoringNotificationChannelRead(d, meta)
}

func resourceMonitoringNotificationChannelDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	lockName, err := replaceVars(d, config, "stackdriver/notifications/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "https://monitoring.googleapis.com/v3/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting NotificationChannel %q", d.Id())
	res, err := sendRequest(config, "DELETE", url, obj)
	if err != nil {
		return handleNotFoundError(err, d, "NotificationChannel")
	}

	log.Printf("[DEBUG] Finished deleting NotificationChannel %q: %#v", d.Id(), res)
	return nil
}

func resourceMonitoringNotificationChannelImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*Config)

	// current import_formats can't import id's with forward slashes in them.
	parseImportId([]string{"(?P<name>.+)"}, d, config)

	return []*schema.ResourceData{d}, nil
}

func flattenMonitoringNotificationChannelLabels(v interface{}) interface{} {
	return v
}

func flattenMonitoringNotificationChannelName(v interface{}) interface{} {
	return v
}

func flattenMonitoringNotificationChannelVerificationStatus(v interface{}) interface{} {
	return v
}

func flattenMonitoringNotificationChannelType(v interface{}) interface{} {
	return v
}

func flattenMonitoringNotificationChannelUserLabels(v interface{}) interface{} {
	return v
}

func flattenMonitoringNotificationChannelDescription(v interface{}) interface{} {
	return v
}

func flattenMonitoringNotificationChannelDisplayName(v interface{}) interface{} {
	return v
}

func flattenMonitoringNotificationChannelEnabled(v interface{}) interface{} {
	return v
}

func expandMonitoringNotificationChannelLabels(v interface{}, d *schema.ResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandMonitoringNotificationChannelType(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringNotificationChannelUserLabels(v interface{}, d *schema.ResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandMonitoringNotificationChannelDescription(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringNotificationChannelDisplayName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringNotificationChannelEnabled(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}