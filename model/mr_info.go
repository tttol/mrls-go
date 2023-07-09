package model

type MrInfo struct {
	Assignee        User
	MrDetails       []MrDetail
	AssignedMrCount int
}
