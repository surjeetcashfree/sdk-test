/*
Cashfree Payout APIs

Cashfree's Payout APIs provide developers with a streamlined pathway to integrate advanced payout capabilities into their applications, platforms and websites.

API version: 2024-01-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_payout

import (
	"encoding/json"
	"strings"
)

// checks if the CreateBeneficiaryRequestBeneficiaryContactDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBeneficiaryRequestBeneficiaryContactDetails{}

// CreateBeneficiaryRequestBeneficiaryContactDetails It should contain the contact details of the beneficiary.
type CreateBeneficiaryRequestBeneficiaryContactDetails struct {
	// It is the email address of the beneficiary. The maximum character limit is 200. It should contain dot (.) and at the rate of (@).
	BeneficiaryEmail *string `json:"beneficiary_email,omitempty"`
	// It is the phone number of the beneficiary. Only reigstered Indian phone numbers are allowed. The value should be between 8 and 12 characters after stripping +91.
	BeneficiaryPhone *string `json:"beneficiary_phone,omitempty"`
	// It is the country code (+91) of the beneficiary's phone number
	BeneficiaryCountryCode *string `json:"beneficiary_country_code,omitempty"`
	// It is the address information of the beneficiary.
	BeneficiaryAddress *string `json:"beneficiary_address,omitempty"`
	// It is the name of the city as present in the beneficiary's address.
	BeneficiaryCity *string `json:"beneficiary_city,omitempty"`
	// It is the name of the state as present in the beneficiary's address.
	BeneficiaryState *string `json:"beneficiary_state,omitempty"`
	// It is the PIN code information as present in the beneficiary's address. The maximum character limit is 6. Only numbers are allowed.
	BeneficiaryPostalCode *string `json:"beneficiary_postal_code,omitempty"`
}


func (o CreateBeneficiaryRequestBeneficiaryContactDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBeneficiaryRequestBeneficiaryContactDetails) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BeneficiaryEmail) {
		toSerialize["beneficiary_email"] = o.BeneficiaryEmail
	}
	if !IsNil(o.BeneficiaryPhone) {
		toSerialize["beneficiary_phone"] = o.BeneficiaryPhone
	}
	if !IsNil(o.BeneficiaryCountryCode) {
		toSerialize["beneficiary_country_code"] = o.BeneficiaryCountryCode
	}
	if !IsNil(o.BeneficiaryAddress) {
		toSerialize["beneficiary_address"] = o.BeneficiaryAddress
	}
	if !IsNil(o.BeneficiaryCity) {
		toSerialize["beneficiary_city"] = o.BeneficiaryCity
	}
	if !IsNil(o.BeneficiaryState) {
		toSerialize["beneficiary_state"] = o.BeneficiaryState
	}
	if !IsNil(o.BeneficiaryPostalCode) {
		toSerialize["beneficiary_postal_code"] = o.BeneficiaryPostalCode
	}
	return toSerialize, nil
}



