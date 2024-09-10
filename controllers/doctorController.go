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

	var upComingappointmentsData []AppointmentWithPatientName
	initializers.DB.Table("appointments").
    	Select("appointments.*, users.first_name AS patient_namef, users.last_name AS patient_namel").
    	Joins("left join users on users.id = appointments.patient_id").
    	Where("appointments.doctor_id = ? AND appointments.is_completed = ? AND appointments.is_accepted = ?", user.(models.User).ID, false, true).
    	Scan(&upComingappointmentsData)
	var onGoingAppointments []AppointmentWithPatientName
	initializers.DB.Table("appointments").
		Select("appointments.*, users.first_name AS patient_namef, users.last_name AS patient_namel").
    	Joins("left join users on users.id = appointments.patient_id").
    	Where("appointments.doctor_id = ? AND appointments.is_ongoing = ?", user.(models.User).ID, true).
    	Scan(&onGoingAppointments)
	var requestedAppointments []AppointmentWithPatientName
	initializers.DB.Table("appointments").
		Select("appointments.*, users.first_name AS patient_namef, users.last_name AS patient_namel").
    	Joins("left join users on users.id = appointments.patient_id").
    	Where("appointments.doctor_id = ? AND appointments.is_accepted = ?", user.(models.User).ID, false).
    	Scan(&requestedAppointments)
	var todayCount int64
	initializers.DB.Model(&models.Appointment{}).Where("doctor_id = ? AND is_accepted = ?", user.(models.User).ID, true).Count(&todayCount)
	var completedCount int64
	initializers.DB.Model(&models.Appointment{}).Where("doctor_id = ? AND is_completed = ?", user.(models.User).ID, true).Count(&completedCount)
	var requestedCount int64
	initializers.DB.Model(&models.Appointment{}).Where("doctor_id = ? AND is_accepted = ?", user.(models.User).ID, false).Count(&requestedCount)
		c.JSON(http.StatusOK, gin.H{
		"upcomingA": upComingappointmentsData,
		"todayCount": todayCount,
		"completedCount": completedCount,
		"requestedCount": requestedCount,
		"ongoingA": onGoingAppointments,
		"requestedA": requestedAppointments,
	})

}

func GetDoctorDashUpdates() {

}

func GetDoctorUpdateByID() {

}