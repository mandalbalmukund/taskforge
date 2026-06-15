package dto

type ApprovalRequest struct {
	Approved bool   `json:"approved"`
	Approver string `json:"approver"`
}