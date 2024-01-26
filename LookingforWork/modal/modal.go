package modal

type Candidate struct {
	FullName    string       `json:"fullname"`
	ID          int          `json:"id"`
	About       string       `json:"about"`
	JobTitle    string       `json:"jobtitle"`
	Experience  string       `json:"experience"`
	Country     string       `json:"country"`
	Job         *Job         `json:"job"`
	Profile     *Profile     `json:"profile"`
	Skills      *Skills      `json:"skills"`
	Location    *Location    `json:"location"`
	SocialMedia *SocialMedia `json:"socialmedia"`
}

type Profile struct {
	ID          int    `json:"id"`
	CandidateID int    `json:"candidate_id"` // Foreign key referencing the Candidate's ID
	ProfileLink string `json:"profilelink"`
}
type Location struct {
	ID           int    `json:"id"`
	CandidateID  int    `json:"candidate_id"` // Foreign key referencing the Candidate's ID
	LocationName string `json:"locationname"`
	LocationType string `json:"locationtype"`
}
type Job struct {
	ID          int    `json:"jobId"`
	CandidateID int    `json:"candidate_id"` // Foreign key referencing the Candidate's ID
	JobName     string `json:"jobName"`
}
type Skills struct {
	ID          int    `json:"id"`
	CandidateID int    `json:"candidate_id"` // Foreign key referencing the Candidate's ID
	SkillsName  string `json:"skillsname"`
}
type SocialMedia struct {
	ID          int    `json:"id"`
	CandidateID int    `json:"candidate_id"` // Foreign key referencing the Candidate's ID
	Platform    string `json:"platform"`
	URL         string `json:"url"`
}
