package dtos

type Info struct {
	PersonalInfo struct {
		Name  string `json:"name"`
		Title string `json:"title"`
		Phone string `json:"phone"`
		Email string `json:"email"`
	} `json:"personal_info"`
	Links          map[string]string `json:"links"`
	Profile        map[string]string `json:"profile"`
	WorkExperience []struct {
		Position               string   `json:"position"`
		Company                string   `json:"company"`
		Location               string   `json:"location"`
		StartDate              string   `json:"start_date"`
		EndDate                string   `json:"end_date"`
		Technologies           []string `json:"technologies"`
		Responsibilities       []string `json:"responsibilities"`
		AchievementsAndDetails []string `json:"achievements_and_details"`
	} `json:"work_experience"`
	Education []struct {
		Institution  string   `json:"institution"`
		Location     string   `json:"location"`
		Degree       string   `json:"degree"`
		StartDate    string   `json:"start_date"`
		EndDate      string   `json:"end_date"`
		Achievements []string `json:"achievements"`
	} `json:"education"`
	Skills map[string]interface{} `json:"skills"`
}
