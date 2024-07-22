package controllers

import (
	"healthcare-portal/initializers"
	"healthcare-portal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func GetAllAppointments() {

}

func GetAppointmentByID() {

}

func UpdateAppointment() {

}

func GetDocDashboardData(c *gin.Context) {
	user, _ := c.Get("user")
	type AppointmentWithPatientName struct {
		models.Appointment
		PatientNamef string		
		PatientNamel string
	}

	var appointmentsData []AppointmentWithPatientName
	initializers.DB.Table("appointments").
    	Select("appointments.*, users.first_name AS patient_namef, users.last_name AS patient_namel").
    	Joins("left join users on users.id = appointments.patient_id").
    	Where("appointments.doctor_id = ? AND appointments.is_completed = ? AND appointments.is_accepted = ?", user.(models.User).ID, false, false).
    	Scan(&appointmentsData)
	var count int64
	initializers.DB.Model(&models.Appointment{}).Where("doctor_id = ? AND is_completed = ? AND is_accepted = ?", user.(models.User).ID, false, false).Count(&count)
	var completedCount int64
	initializers.DB.Model(&models.Appointment{}).Where("doctor_id = ? AND is_completed = ? AND is_accepted = ?", user.(models.User).ID, true, true).Count(&completedCount)
	var requestedCount int64
	initializers.DB.Model(&models.Appointment{}).Where("doctor_id = ? AND is_completed = ? AND is_accepted = ?", user.(models.User).ID, false, false).Count(&requestedCount)
	if(count == 0) {
		c.JSON(http.StatusOK, gin.H{
			"upcomingA": nil,
			"todayCount": count,
			"completedCount": completedCount,
			"requestedCount": requestedCount,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
		"upcomingA": appointmentsData,
		"todayCount": count,
		"completedCount": completedCount,
		"requestedCount": requestedCount,
	})
	}
}

func GetDoctorDashUpdates() {

}

func GetDoctorUpdateByID() {

}