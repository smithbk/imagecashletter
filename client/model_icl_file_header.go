/*
 * ImageCashLetter API
 *
 * Moov Image Cash Letter (ICL) implements an HTTP API for creating, parsing, and validating ImageCashLetter files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// IclFileHeader struct for IclFileHeader
type IclFileHeader struct {
	// FileHeader ID
	ID string `json:"ID,omitempty"`
	// StandardLevel identifies the standard level of the file.  * `03` - DSTU X9.37-2003 * `30` - X9.100-187-2008 * `35` - X9.100-187-2013 and X9.100-187-2016
	StandardLevel string `json:"standardLevel"`
	// TestIndicator identifies whether the file is a test or production file.  * `T` - Test File * `P` - Production File
	TestIndicator string `json:"testIndicator"`
	// ImmediateDestination is the routing and transit number of the Federal Reserve Bank (FRB) or receiver to which the file is being sent.
	ImmediateDestination string `json:"immediateDestination"`
	// ImmediateOrigin is the routing and transit number of the Federal Reserve Bank (FRB) or originator from which the file is being sent.
	ImmediateOrigin string `json:"immediateOrigin"`
	// FileCreationDate is the date the immediate origin institution creates the file.
	FileCreationDate time.Time `json:"fileCreationDate"`
	// FileCreationTime is the time the immediate origin institution creates the file.
	FileCreationTime time.Time `json:"fileCreationTime"`
	// ResendIndicator indicates whether the file has been previously transmitted. (Y - Yes, N - No)
	ResendIndicator string `json:"resendIndicator"`
	// ImmediateDestinationName identifies the short name of the institution that receives the file.
	ImmediateDestinationName string `json:"immediateDestinationName,omitempty"`
	// immediateOriginName identifies the short name of the institution that sends the file.
	ImmediateOriginName string `json:"immediateOriginName,omitempty"`
	// FileIDModifier is a code that permits multiple files, created on the same date, at the same time, and sent between the same institutions, to be distinguished from one another. If FileHeader ImmediateDestination, ImmediateOrigin, FileCreationDate, and FileCreationTime in a previous file are equal to the same fields in this file, FileIDModifier must be defined.
	FileIDModifier string `json:"fileIDModifier,omitempty"`
	// CountryCode is a 2-character code as approved by the International Organization for Standardization (ISO) used to identify the country in which the payer bank is located.
	CountryCode string `json:"countryCode,omitempty"`
	// UserField identifies a field used at the discretion of users of the standard.
	UserField string `json:"userField,omitempty"`
	// CompanionDocumentIndicator indicates the Companion Document being used. It shall be present only under clearing arrangements, where Companion Document usage and values are defined. Values: * 0–9 - Reserved for United States use * A–J - Reserved for Canadian use * Other - Defined by clearing arrangements
	CompanionDocumentIndicator string `json:"companionDocumentIndicator,omitempty"`
}
