/*
 * NetSchool
 *
 * The API for the NetSchool irTech project
 *
 * API version: 4.30.43656
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Diary struct {
	WeekStart string          `json:"weekStart,omitempty"`
	WeekEnd   string          `json:"weekEnd,omitempty"`
	WeekDays  []DiaryWeekDays `json:"weekDays,omitempty"`
	TermName  string          `json:"termName,omitempty"`
	ClassName string          `json:"className,omitempty"`
}
