package model

type MrInfo struct {
	Assignee        User
	MrDetailForms   []MrDetail
	AssignedMrCount int
}
