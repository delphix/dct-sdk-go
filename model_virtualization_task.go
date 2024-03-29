/*
Delphix DCT API

Delphix DCT API

API version: 3.9.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
	"time"
)

// checks if the VirtualizationTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualizationTask{}

// VirtualizationTask struct for VirtualizationTask
type VirtualizationTask struct {
	Id *string `json:"id,omitempty"`
	ParentJobId *string `json:"parent_job_id,omitempty"`
	StartTime *time.Time `json:"start_time,omitempty"`
	EndTime *time.Time `json:"end_time,omitempty"`
	Title *string `json:"title,omitempty"`
	PercentComplete *int32 `json:"percent_complete,omitempty"`
	Events []VirtualizationTaskEvent `json:"events,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewVirtualizationTask instantiates a new VirtualizationTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationTask() *VirtualizationTask {
	this := VirtualizationTask{}
	return &this
}

// NewVirtualizationTaskWithDefaults instantiates a new VirtualizationTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationTaskWithDefaults() *VirtualizationTask {
	this := VirtualizationTask{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VirtualizationTask) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationTask) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VirtualizationTask) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VirtualizationTask) SetId(v string) {
	o.Id = &v
}

// GetParentJobId returns the ParentJobId field value if set, zero value otherwise.
func (o *VirtualizationTask) GetParentJobId() string {
	if o == nil || IsNil(o.ParentJobId) {
		var ret string
		return ret
	}
	return *o.ParentJobId
}

// GetParentJobIdOk returns a tuple with the ParentJobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationTask) GetParentJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentJobId) {
		return nil, false
	}
	return o.ParentJobId, true
}

// HasParentJobId returns a boolean if a field has been set.
func (o *VirtualizationTask) HasParentJobId() bool {
	if o != nil && !IsNil(o.ParentJobId) {
		return true
	}

	return false
}

// SetParentJobId gets a reference to the given string and assigns it to the ParentJobId field.
func (o *VirtualizationTask) SetParentJobId(v string) {
	o.ParentJobId = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *VirtualizationTask) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationTask) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *VirtualizationTask) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *VirtualizationTask) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *VirtualizationTask) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationTask) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *VirtualizationTask) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *VirtualizationTask) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *VirtualizationTask) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationTask) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *VirtualizationTask) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *VirtualizationTask) SetTitle(v string) {
	o.Title = &v
}

// GetPercentComplete returns the PercentComplete field value if set, zero value otherwise.
func (o *VirtualizationTask) GetPercentComplete() int32 {
	if o == nil || IsNil(o.PercentComplete) {
		var ret int32
		return ret
	}
	return *o.PercentComplete
}

// GetPercentCompleteOk returns a tuple with the PercentComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationTask) GetPercentCompleteOk() (*int32, bool) {
	if o == nil || IsNil(o.PercentComplete) {
		return nil, false
	}
	return o.PercentComplete, true
}

// HasPercentComplete returns a boolean if a field has been set.
func (o *VirtualizationTask) HasPercentComplete() bool {
	if o != nil && !IsNil(o.PercentComplete) {
		return true
	}

	return false
}

// SetPercentComplete gets a reference to the given int32 and assigns it to the PercentComplete field.
func (o *VirtualizationTask) SetPercentComplete(v int32) {
	o.PercentComplete = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *VirtualizationTask) GetEvents() []VirtualizationTaskEvent {
	if o == nil || IsNil(o.Events) {
		var ret []VirtualizationTaskEvent
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationTask) GetEventsOk() ([]VirtualizationTaskEvent, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *VirtualizationTask) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []VirtualizationTaskEvent and assigns it to the Events field.
func (o *VirtualizationTask) SetEvents(v []VirtualizationTaskEvent) {
	o.Events = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VirtualizationTask) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationTask) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VirtualizationTask) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VirtualizationTask) SetStatus(v string) {
	o.Status = &v
}

func (o VirtualizationTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualizationTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ParentJobId) {
		toSerialize["parent_job_id"] = o.ParentJobId
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.PercentComplete) {
		toSerialize["percent_complete"] = o.PercentComplete
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableVirtualizationTask struct {
	value *VirtualizationTask
	isSet bool
}

func (v NullableVirtualizationTask) Get() *VirtualizationTask {
	return v.value
}

func (v *NullableVirtualizationTask) Set(val *VirtualizationTask) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationTask) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationTask(val *VirtualizationTask) *NullableVirtualizationTask {
	return &NullableVirtualizationTask{value: val, isSet: true}
}

func (v NullableVirtualizationTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


