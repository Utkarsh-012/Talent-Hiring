package modal

type Candidate struct {
	FullName    string       `json:"fullname"`
	ID          string       `json:"id"`
	About       string       `json:"about"`
	JobTitle    string       `json:"jobtitle"`
	Experience  string       `json:"experience"`
	Country     string       `json:"country"`
	Location    *Location    `json:"location"`
	Skills      *Skills      `json:"skills"`
	Profile     *Profile     `json:"profile"`
	Job         *Job         `json:"job"`
	SocialMedia *SocialMedia `json:"socialmedia"`
}

type Company struct {
	ID             string    `json:"id"`
	CompanyName    string    `json:"companyname"`
	EMail          string    `json:"e-mail"`
	VerifiedEmail  string    `json:"verifiedemail"`
	CompanyWebsite string    `json:"companywebsite"`
	Location       *Location `json:"location"`
	Profile        *Profile  `json:"profile"`
}

type Profile struct {
	ID          string `json:"id"`
	ProfileLink string `json:"profilelink"`
}
type Location struct {
	ID           string `json:"id"`
	LocationName string `json:"locationname"`
	LocationType string `json:"locationtype"`
}
type Job struct {
	ID      string `json:"id"`
	Jobname string `json:"jobname"`
}
type Skills struct {
	ID         string `json:"id"`
	SkillsName string `json:"skillsname"`
}
type SocialMedia struct {
	ID       string `json:"id"`
	Platform string `json:"platform"`
	URL      string `json:"url"`
}
