package controllers

// func CreateAppointment(c *gin.Context) {
// 	var body struct {
// 		PatientID uint      `json:"patient_id"`
// 		DoctorID  uint      `json:"doctor_id"`
// 		Date      time.Time `json:"date"`
// 		Reason    string    `json:"reason"`
// 	}

// 	c.Bind(&body)

// 	appointment := models.Appointment{
// 		PatientID: body.PatientID,
// 		DoctorID:  body.DoctorID,
// 		Date:     time.Now(),
// 		Reason:    body.Reason,
// 	}

// 	result := initializers.DB.Create(&appointment)
// 	if result.Error != nil {
// 		fmt.Println(result.Error)
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Appointment creation failed",
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"message": "Appointment created successfully",
// 	})

// }