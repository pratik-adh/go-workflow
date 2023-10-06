package errors


func GetErrorResponse(input int, message string) error {
    switch input {
    case 1062:
		// c.JSON(http.StatusNotFound, gin.H{"status": 404, "error": "Data already exists, please enter a new one."})
		return nil;
    default:
        return nil // Default value if none of the cases match
    }
}