// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MultiTaxFilingReturnModel Filing Returns Model
// Example: {"accrualType":0,"adjustments":[{"accountType":"AccountsReceivableAccountsPayable","amount":100,"id":15634,"isCalculated":false,"period":"CurrentPeriod","type":"Discount"}],"description":"Sales Tax Return","filingCalendarId":513256,"filingFrequency":"Monthly","filingType":"ElectronicReturn","formCode":"AL2100","formName":"AL 2100","id":5233103,"returnTaxDetails":[{"nonTaxableAmount":100,"numberOfNights":1,"salesAmount":200,"taxAmount":10,"taxType":"S"}],"returnTaxSummary":{"collectAmount":10,"nonTaxableAccrualAmount":100,"nonTaxableAmount":100,"remittanceAmount":10,"reportableNonTaxableAmount":50,"reportableSalesAmount":100,"reportableTaxAmount":5,"reportableTaxableAmount":5,"salesAccrualAmount":100,"salesAmount":200,"taxAccrualAmount":10,"taxAmount":10,"taxableAccrualAmount":0,"taxableAmount":0},"status":"PendingApproval","totalAdjustments":0,"totalAugmentations":0,"totalPayments":0}
//
// swagger:model MultiTaxFilingReturnModel
type MultiTaxFilingReturnModel struct {

	// Accrual type of the return
	// Example: Filing
	// Enum: [Filing Accrual]
	AccrualType string `json:"accrualType,omitempty"`

	// The Adjustments for this return.
	// Example: [{"accountType":"AccountsReceivableAccountsPayable","amount":100,"id":15634,"isCalculated":false,"period":"CurrentPeriod","type":"Discount"}]
	Adjustments []*FilingAdjustmentModel `json:"adjustments"`

	// The applied carry over credit documents
	// Example: {"totalExempt":0,"totalSales":0,"totalTax":0,"totalTaxable":0}
	AppliedCarryOverCredits *FilingReturnCreditModel `json:"appliedCarryOverCredits,omitempty"`

	// The attachments for this return.
	Attachments []*FilingAttachmentModel `json:"attachments"`

	// The Augmentations for this return.
	Augmentations []*FilingAugmentationModel `json:"augmentations"`

	// A description for the return.
	// Example: Sales Tax Return
	Description string `json:"description,omitempty"`

	// The end date of this return
	// Format: date-time
	EndPeriod strfmt.DateTime `json:"endPeriod,omitempty"`

	// The excluded carry over credit documents
	// Example: {"totalExempt":0,"totalSales":0,"totalTax":0,"totalTaxable":0}
	ExcludedCarryOverCredits *FilingReturnCreditModel `json:"excludedCarryOverCredits,omitempty"`

	// The date the return was filed by Avalara.
	// Format: date-time
	FiledDate strfmt.DateTime `json:"filedDate,omitempty"`

	// The unique ID number of the filing calendar associated with this return.
	// Example: 513256
	FilingCalendarID int64 `json:"filingCalendarId,omitempty"`

	// The filing frequency of the return.
	// Example: Monthly
	// Enum: [Monthly Quarterly SemiAnnually Annually Bimonthly Occasional InverseQuarterly Weekly]
	FilingFrequency string `json:"filingFrequency,omitempty"`

	// The filing type of the return.
	// Example: PaperReturn
	// Enum: [PaperReturn ElectronicReturn SER EFTPaper PhonePaper SignatureReady EfileCheck]
	FilingType string `json:"filingType,omitempty"`

	// The unique code of the form.
	// Example: AL2100
	FormCode string `json:"formCode,omitempty"`

	// The name of the form.
	// Example: AL 2100
	FormName string `json:"formName,omitempty"`

	// The unique ID number of this filing return.
	// Example: 5233103
	ID int64 `json:"id,omitempty"`

	// The payments for this return.
	Payments []*FilingPaymentModel `json:"payments"`

	// A detailed breakdown of the taxes in this filing
	// Example: [{"nonTaxableAmount":100,"numberOfNights":1,"salesAmount":200,"taxAmount":10,"taxType":"S"}]
	ReturnTaxDetails []*FilingsTaxDetailsModel `json:"returnTaxDetails"`

	// A summary of all taxes compbined for this period
	// Example: {"collectAmount":10,"nonTaxableAccrualAmount":100,"nonTaxableAmount":100,"remittanceAmount":10,"reportableNonTaxableAmount":50,"reportableSalesAmount":100,"reportableTaxAmount":5,"reportableTaxableAmount":5,"salesAccrualAmount":100,"salesAmount":200,"taxAccrualAmount":10,"taxAmount":10,"taxableAccrualAmount":0,"taxableAmount":0}
	ReturnTaxSummary *FilingsTaxSummaryModel `json:"returnTaxSummary,omitempty"`

	// The start date of this return
	// Format: date-time
	StartPeriod strfmt.DateTime `json:"startPeriod,omitempty"`

	// The current status of the filing return.
	// Example: PendingApproval
	// Enum: [PendingApproval Dirty ApprovedToFile PendingFiling PendingFilingOnBehalf Filed FiledOnBehalf ReturnAccepted ReturnAcceptedOnBehalf PaymentRemitted Voided PendingReturn PendingReturnOnBehalf DoNotFile ReturnRejected ReturnRejectedOnBehalf ApprovedToFileOnBehalf]
	Status string `json:"status,omitempty"`

	// Tax Authority ID of this return
	TaxAuthorityID int32 `json:"taxAuthorityId,omitempty"`

	// Total amount of adjustments on this return
	// Example: 0
	TotalAdjustments float64 `json:"totalAdjustments,omitempty"`

	// Total amount of augmentations on this return
	// Example: 0
	TotalAugmentations float64 `json:"totalAugmentations,omitempty"`

	// Total amount of payments on this return
	// Example: 0
	TotalPayments float64 `json:"totalPayments,omitempty"`

	// The FilingTaskType for this return.
	Type string `json:"type,omitempty"`
}

// Validate validates this multi tax filing return model
func (m *MultiTaxFilingReturnModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccrualType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdjustments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppliedCarryOverCredits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttachments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAugmentations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExcludedCarryOverCredits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFiledDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilingFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilingType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnTaxDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnTaxSummary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var multiTaxFilingReturnModelTypeAccrualTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Filing","Accrual"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		multiTaxFilingReturnModelTypeAccrualTypePropEnum = append(multiTaxFilingReturnModelTypeAccrualTypePropEnum, v)
	}
}

const (

	// MultiTaxFilingReturnModelAccrualTypeFiling captures enum value "Filing"
	MultiTaxFilingReturnModelAccrualTypeFiling string = "Filing"

	// MultiTaxFilingReturnModelAccrualTypeAccrual captures enum value "Accrual"
	MultiTaxFilingReturnModelAccrualTypeAccrual string = "Accrual"
)

// prop value enum
func (m *MultiTaxFilingReturnModel) validateAccrualTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, multiTaxFilingReturnModelTypeAccrualTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MultiTaxFilingReturnModel) validateAccrualType(formats strfmt.Registry) error {
	if swag.IsZero(m.AccrualType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccrualTypeEnum("accrualType", "body", m.AccrualType); err != nil {
		return err
	}

	return nil
}

func (m *MultiTaxFilingReturnModel) validateAdjustments(formats strfmt.Registry) error {
	if swag.IsZero(m.Adjustments) { // not required
		return nil
	}

	for i := 0; i < len(m.Adjustments); i++ {
		if swag.IsZero(m.Adjustments[i]) { // not required
			continue
		}

		if m.Adjustments[i] != nil {
			if err := m.Adjustments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("adjustments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("adjustments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MultiTaxFilingReturnModel) validateAppliedCarryOverCredits(formats strfmt.Registry) error {
	if swag.IsZero(m.AppliedCarryOverCredits) { // not required
		return nil
	}

	if m.AppliedCarryOverCredits != nil {
		if err := m.AppliedCarryOverCredits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appliedCarryOverCredits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appliedCarryOverCredits")
			}
			return err
		}
	}

	return nil
}

func (m *MultiTaxFilingReturnModel) validateAttachments(formats strfmt.Registry) error {
	if swag.IsZero(m.Attachments) { // not required
		return nil
	}

	for i := 0; i < len(m.Attachments); i++ {
		if swag.IsZero(m.Attachments[i]) { // not required
			continue
		}

		if m.Attachments[i] != nil {
			if err := m.Attachments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attachments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MultiTaxFilingReturnModel) validateAugmentations(formats strfmt.Registry) error {
	if swag.IsZero(m.Augmentations) { // not required
		return nil
	}

	for i := 0; i < len(m.Augmentations); i++ {
		if swag.IsZero(m.Augmentations[i]) { // not required
			continue
		}

		if m.Augmentations[i] != nil {
			if err := m.Augmentations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("augmentations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("augmentations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MultiTaxFilingReturnModel) validateEndPeriod(formats strfmt.Registry) error {
	if swag.IsZero(m.EndPeriod) { // not required
		return nil
	}

	if err := validate.FormatOf("endPeriod", "body", "date-time", m.EndPeriod.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MultiTaxFilingReturnModel) validateExcludedCarryOverCredits(formats strfmt.Registry) error {
	if swag.IsZero(m.ExcludedCarryOverCredits) { // not required
		return nil
	}

	if m.ExcludedCarryOverCredits != nil {
		if err := m.ExcludedCarryOverCredits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("excludedCarryOverCredits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("excludedCarryOverCredits")
			}
			return err
		}
	}

	return nil
}

func (m *MultiTaxFilingReturnModel) validateFiledDate(formats strfmt.Registry) error {
	if swag.IsZero(m.FiledDate) { // not required
		return nil
	}

	if err := validate.FormatOf("filedDate", "body", "date-time", m.FiledDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var multiTaxFilingReturnModelTypeFilingFrequencyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Monthly","Quarterly","SemiAnnually","Annually","Bimonthly","Occasional","InverseQuarterly","Weekly"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		multiTaxFilingReturnModelTypeFilingFrequencyPropEnum = append(multiTaxFilingReturnModelTypeFilingFrequencyPropEnum, v)
	}
}

const (

	// MultiTaxFilingReturnModelFilingFrequencyMonthly captures enum value "Monthly"
	MultiTaxFilingReturnModelFilingFrequencyMonthly string = "Monthly"

	// MultiTaxFilingReturnModelFilingFrequencyQuarterly captures enum value "Quarterly"
	MultiTaxFilingReturnModelFilingFrequencyQuarterly string = "Quarterly"

	// MultiTaxFilingReturnModelFilingFrequencySemiAnnually captures enum value "SemiAnnually"
	MultiTaxFilingReturnModelFilingFrequencySemiAnnually string = "SemiAnnually"

	// MultiTaxFilingReturnModelFilingFrequencyAnnually captures enum value "Annually"
	MultiTaxFilingReturnModelFilingFrequencyAnnually string = "Annually"

	// MultiTaxFilingReturnModelFilingFrequencyBimonthly captures enum value "Bimonthly"
	MultiTaxFilingReturnModelFilingFrequencyBimonthly string = "Bimonthly"

	// MultiTaxFilingReturnModelFilingFrequencyOccasional captures enum value "Occasional"
	MultiTaxFilingReturnModelFilingFrequencyOccasional string = "Occasional"

	// MultiTaxFilingReturnModelFilingFrequencyInverseQuarterly captures enum value "InverseQuarterly"
	MultiTaxFilingReturnModelFilingFrequencyInverseQuarterly string = "InverseQuarterly"

	// MultiTaxFilingReturnModelFilingFrequencyWeekly captures enum value "Weekly"
	MultiTaxFilingReturnModelFilingFrequencyWeekly string = "Weekly"
)

// prop value enum
func (m *MultiTaxFilingReturnModel) validateFilingFrequencyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, multiTaxFilingReturnModelTypeFilingFrequencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MultiTaxFilingReturnModel) validateFilingFrequency(formats strfmt.Registry) error {
	if swag.IsZero(m.FilingFrequency) { // not required
		return nil
	}

	// value enum
	if err := m.validateFilingFrequencyEnum("filingFrequency", "body", m.FilingFrequency); err != nil {
		return err
	}

	return nil
}

var multiTaxFilingReturnModelTypeFilingTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PaperReturn","ElectronicReturn","SER","EFTPaper","PhonePaper","SignatureReady","EfileCheck"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		multiTaxFilingReturnModelTypeFilingTypePropEnum = append(multiTaxFilingReturnModelTypeFilingTypePropEnum, v)
	}
}

const (

	// MultiTaxFilingReturnModelFilingTypePaperReturn captures enum value "PaperReturn"
	MultiTaxFilingReturnModelFilingTypePaperReturn string = "PaperReturn"

	// MultiTaxFilingReturnModelFilingTypeElectronicReturn captures enum value "ElectronicReturn"
	MultiTaxFilingReturnModelFilingTypeElectronicReturn string = "ElectronicReturn"

	// MultiTaxFilingReturnModelFilingTypeSER captures enum value "SER"
	MultiTaxFilingReturnModelFilingTypeSER string = "SER"

	// MultiTaxFilingReturnModelFilingTypeEFTPaper captures enum value "EFTPaper"
	MultiTaxFilingReturnModelFilingTypeEFTPaper string = "EFTPaper"

	// MultiTaxFilingReturnModelFilingTypePhonePaper captures enum value "PhonePaper"
	MultiTaxFilingReturnModelFilingTypePhonePaper string = "PhonePaper"

	// MultiTaxFilingReturnModelFilingTypeSignatureReady captures enum value "SignatureReady"
	MultiTaxFilingReturnModelFilingTypeSignatureReady string = "SignatureReady"

	// MultiTaxFilingReturnModelFilingTypeEfileCheck captures enum value "EfileCheck"
	MultiTaxFilingReturnModelFilingTypeEfileCheck string = "EfileCheck"
)

// prop value enum
func (m *MultiTaxFilingReturnModel) validateFilingTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, multiTaxFilingReturnModelTypeFilingTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MultiTaxFilingReturnModel) validateFilingType(formats strfmt.Registry) error {
	if swag.IsZero(m.FilingType) { // not required
		return nil
	}

	// value enum
	if err := m.validateFilingTypeEnum("filingType", "body", m.FilingType); err != nil {
		return err
	}

	return nil
}

func (m *MultiTaxFilingReturnModel) validatePayments(formats strfmt.Registry) error {
	if swag.IsZero(m.Payments) { // not required
		return nil
	}

	for i := 0; i < len(m.Payments); i++ {
		if swag.IsZero(m.Payments[i]) { // not required
			continue
		}

		if m.Payments[i] != nil {
			if err := m.Payments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("payments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("payments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MultiTaxFilingReturnModel) validateReturnTaxDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.ReturnTaxDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.ReturnTaxDetails); i++ {
		if swag.IsZero(m.ReturnTaxDetails[i]) { // not required
			continue
		}

		if m.ReturnTaxDetails[i] != nil {
			if err := m.ReturnTaxDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("returnTaxDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("returnTaxDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MultiTaxFilingReturnModel) validateReturnTaxSummary(formats strfmt.Registry) error {
	if swag.IsZero(m.ReturnTaxSummary) { // not required
		return nil
	}

	if m.ReturnTaxSummary != nil {
		if err := m.ReturnTaxSummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("returnTaxSummary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("returnTaxSummary")
			}
			return err
		}
	}

	return nil
}

func (m *MultiTaxFilingReturnModel) validateStartPeriod(formats strfmt.Registry) error {
	if swag.IsZero(m.StartPeriod) { // not required
		return nil
	}

	if err := validate.FormatOf("startPeriod", "body", "date-time", m.StartPeriod.String(), formats); err != nil {
		return err
	}

	return nil
}

var multiTaxFilingReturnModelTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PendingApproval","Dirty","ApprovedToFile","PendingFiling","PendingFilingOnBehalf","Filed","FiledOnBehalf","ReturnAccepted","ReturnAcceptedOnBehalf","PaymentRemitted","Voided","PendingReturn","PendingReturnOnBehalf","DoNotFile","ReturnRejected","ReturnRejectedOnBehalf","ApprovedToFileOnBehalf"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		multiTaxFilingReturnModelTypeStatusPropEnum = append(multiTaxFilingReturnModelTypeStatusPropEnum, v)
	}
}

const (

	// MultiTaxFilingReturnModelStatusPendingApproval captures enum value "PendingApproval"
	MultiTaxFilingReturnModelStatusPendingApproval string = "PendingApproval"

	// MultiTaxFilingReturnModelStatusDirty captures enum value "Dirty"
	MultiTaxFilingReturnModelStatusDirty string = "Dirty"

	// MultiTaxFilingReturnModelStatusApprovedToFile captures enum value "ApprovedToFile"
	MultiTaxFilingReturnModelStatusApprovedToFile string = "ApprovedToFile"

	// MultiTaxFilingReturnModelStatusPendingFiling captures enum value "PendingFiling"
	MultiTaxFilingReturnModelStatusPendingFiling string = "PendingFiling"

	// MultiTaxFilingReturnModelStatusPendingFilingOnBehalf captures enum value "PendingFilingOnBehalf"
	MultiTaxFilingReturnModelStatusPendingFilingOnBehalf string = "PendingFilingOnBehalf"

	// MultiTaxFilingReturnModelStatusFiled captures enum value "Filed"
	MultiTaxFilingReturnModelStatusFiled string = "Filed"

	// MultiTaxFilingReturnModelStatusFiledOnBehalf captures enum value "FiledOnBehalf"
	MultiTaxFilingReturnModelStatusFiledOnBehalf string = "FiledOnBehalf"

	// MultiTaxFilingReturnModelStatusReturnAccepted captures enum value "ReturnAccepted"
	MultiTaxFilingReturnModelStatusReturnAccepted string = "ReturnAccepted"

	// MultiTaxFilingReturnModelStatusReturnAcceptedOnBehalf captures enum value "ReturnAcceptedOnBehalf"
	MultiTaxFilingReturnModelStatusReturnAcceptedOnBehalf string = "ReturnAcceptedOnBehalf"

	// MultiTaxFilingReturnModelStatusPaymentRemitted captures enum value "PaymentRemitted"
	MultiTaxFilingReturnModelStatusPaymentRemitted string = "PaymentRemitted"

	// MultiTaxFilingReturnModelStatusVoided captures enum value "Voided"
	MultiTaxFilingReturnModelStatusVoided string = "Voided"

	// MultiTaxFilingReturnModelStatusPendingReturn captures enum value "PendingReturn"
	MultiTaxFilingReturnModelStatusPendingReturn string = "PendingReturn"

	// MultiTaxFilingReturnModelStatusPendingReturnOnBehalf captures enum value "PendingReturnOnBehalf"
	MultiTaxFilingReturnModelStatusPendingReturnOnBehalf string = "PendingReturnOnBehalf"

	// MultiTaxFilingReturnModelStatusDoNotFile captures enum value "DoNotFile"
	MultiTaxFilingReturnModelStatusDoNotFile string = "DoNotFile"

	// MultiTaxFilingReturnModelStatusReturnRejected captures enum value "ReturnRejected"
	MultiTaxFilingReturnModelStatusReturnRejected string = "ReturnRejected"

	// MultiTaxFilingReturnModelStatusReturnRejectedOnBehalf captures enum value "ReturnRejectedOnBehalf"
	MultiTaxFilingReturnModelStatusReturnRejectedOnBehalf string = "ReturnRejectedOnBehalf"

	// MultiTaxFilingReturnModelStatusApprovedToFileOnBehalf captures enum value "ApprovedToFileOnBehalf"
	MultiTaxFilingReturnModelStatusApprovedToFileOnBehalf string = "ApprovedToFileOnBehalf"
)

// prop value enum
func (m *MultiTaxFilingReturnModel) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, multiTaxFilingReturnModelTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MultiTaxFilingReturnModel) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this multi tax filing return model based on the context it is used
func (m *MultiTaxFilingReturnModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdjustments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAppliedCarryOverCredits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAttachments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAugmentations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExcludedCarryOverCredits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePayments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReturnTaxDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReturnTaxSummary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MultiTaxFilingReturnModel) contextValidateAdjustments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Adjustments); i++ {

		if m.Adjustments[i] != nil {
			if err := m.Adjustments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("adjustments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("adjustments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MultiTaxFilingReturnModel) contextValidateAppliedCarryOverCredits(ctx context.Context, formats strfmt.Registry) error {

	if m.AppliedCarryOverCredits != nil {
		if err := m.AppliedCarryOverCredits.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appliedCarryOverCredits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appliedCarryOverCredits")
			}
			return err
		}
	}

	return nil
}

func (m *MultiTaxFilingReturnModel) contextValidateAttachments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attachments); i++ {

		if m.Attachments[i] != nil {
			if err := m.Attachments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attachments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MultiTaxFilingReturnModel) contextValidateAugmentations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Augmentations); i++ {

		if m.Augmentations[i] != nil {
			if err := m.Augmentations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("augmentations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("augmentations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MultiTaxFilingReturnModel) contextValidateExcludedCarryOverCredits(ctx context.Context, formats strfmt.Registry) error {

	if m.ExcludedCarryOverCredits != nil {
		if err := m.ExcludedCarryOverCredits.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("excludedCarryOverCredits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("excludedCarryOverCredits")
			}
			return err
		}
	}

	return nil
}

func (m *MultiTaxFilingReturnModel) contextValidatePayments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Payments); i++ {

		if m.Payments[i] != nil {
			if err := m.Payments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("payments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("payments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MultiTaxFilingReturnModel) contextValidateReturnTaxDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ReturnTaxDetails); i++ {

		if m.ReturnTaxDetails[i] != nil {
			if err := m.ReturnTaxDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("returnTaxDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("returnTaxDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MultiTaxFilingReturnModel) contextValidateReturnTaxSummary(ctx context.Context, formats strfmt.Registry) error {

	if m.ReturnTaxSummary != nil {
		if err := m.ReturnTaxSummary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("returnTaxSummary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("returnTaxSummary")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MultiTaxFilingReturnModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MultiTaxFilingReturnModel) UnmarshalBinary(b []byte) error {
	var res MultiTaxFilingReturnModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}