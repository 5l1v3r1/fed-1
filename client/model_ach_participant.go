/*
 * FED API
 *
 * FED API is designed to create FEDACH and FEDWIRE dictionaries.  The FEDACH dictionary contains receiving depository financial institutions (RDFI’s) which are qualified to receive ACH entries.  The FEDWIRE dictionary contains receiving depository financial institutions (RDFI’s) which are qualified to receive WIRE entries.  This project implements a modern REST HTTP API for FEDACH Dictionary and FEDWIRE Dictionary.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ACHParticipant holds a FedACH dir routing record as defined by Fed ACH Format.  https://www.frbservices.org/EPaymentsDirectory/achFormat.html
type AchParticipant struct {
	// The institution's routing number
	RoutingNumber string `json:"routingNumber,omitempty"`
	// Main/Head Office or Branch. O=main B=branch
	OfficeCode string `json:"officeCode,omitempty"`
	// Servicing Fed's main office routing number
	ServicingFRBNumber string `json:"servicingFRBNumber,omitempty"`
	// The code indicating the ABA number to be used to route or send ACH items to the RDFI 0 = Institution is a Federal Reserve Bank 1 = Send items to customer routing number 2 = Send items to customer using new routing number field
	RecordTypeCode string `json:"recordTypeCode,omitempty"`
	// Date of last revision
	Revised string `json:"revised,omitempty"`
	// Financial Institution's new routing number resulting from a merger or renumber
	NewRoutingNumber string `json:"newRoutingNumber,omitempty"`
	// Financial Institution Name
	CustomerName string `json:"customerName,omitempty"`
	// FEDACH delivery address
	AchLocation []AchLocation `json:"achLocation,omitempty"`
	// The Financial Institution's phone number
	PhoneNumber string `json:"phoneNumber,omitempty"`
	// Code is based on the customers receiver code 1 = Receives Gov/Comm
	StatusCode string `json:"statusCode,omitempty"`
	// Code is current view 1 = Current view
	ViewCode string `json:"viewCode,omitempty"`
}