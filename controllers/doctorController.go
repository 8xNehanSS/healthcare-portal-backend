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

func UpdateAppointment(c *gin.Context) {
	var body struct {
		AppointmentID string `json:"appointment_id"`
		Status string `json:"status"`
	}
	c.Bind(&body)
	if(body.AppointmentID == "") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	var appointment models.Appointment
	err := initializers.DB.Where("id = ?", body.AppointmentID).First(&appointment)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return;
	}
	appointmentUpdated := models.Appointment{
		Status: body.Status,
	}
	initializers.DB.Model(&appointment).Updates(appointmentUpdated)
	c.JSON(http.StatusOK, gin.H{"message": "Appointment updated successfully"})
}

func GetDocDashboardData(c *gin.Context) {
	user, _ := c.Get("user")
	type AppointmentWithPatientName struct {
		models.Appointment
		PatientNamef string		
		PatientNamel string
	}

	var upComingappointmentsData []AppointmentWithPatientName
	initializers.DB.Limit(6).Table("appointments").
    	Select("appointments.*, patients.first_name AS patient_namef, patients.last_name AS patient_namel").
    	Joins("left join patients on patients.patient_id = appointments.patient_id").
    	Where("appointments.doctor_id = ? AND appointments.status = ?", user.(models.Login).UserID, "confirmed").
    	Scan(&upComingappointmentsData)
	var onGoingAppointments []AppointmentWithPatientName
	initializers.DB.Limit(6).Table("appointments").
    	Select("appointments.*, patients.first_name AS patient_namef, patients.last_name AS patient_namel").
    	Joins("left join patients on patients.patient_id = appointments.patient_id").
    	Where("appointments.doctor_id = ? AND appointments.status = ?", user.(models.Login).UserID, "ongoing").
    	Scan(&onGoingAppointments)
	var requestedAppointments []AppointmentWithPatientName
	initializers.DB.Limit(6).Table("appointments").
    	Select("appointments.*, patients.first_name AS patient_namef, patients.last_name AS patient_namel").
    	Joins("left join patients on patients.patient_id = appointments.patient_id").
    	Where("appointments.doctor_id = ? AND appointments.status = ?", user.(models.Login).UserID, "requested").
    	Scan(&requestedAppointments)
	var todayCount int64
	initializers.DB.Model(&models.Appointment{}).Where("doctor_id = ? AND appointments.status = ?", user.(models.Login).UserID, "confirmed").Count(&todayCount)
	var completedCount int64
	initializers.DB.Model(&models.Appointment{}).Where("doctor_id = ? AND appointments.status = ?", user.(models.Login).UserID, "completed").Count(&completedCount)
	var requestedCount int64
	initializers.DB.Model(&models.Appointment{}).Where("doctor_id = ? AND appointments.status = ?", user.(models.Login).UserID, "requested").Count(&requestedCount)
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