//go:generate goagen bootstrap -d local/pa_api/design

package main

import (
	"local/pa_api/app"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("PANOS API")

	service.Mux = NewMux()
	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "keygen" controller
	c := NewKeygenController(service)
	app.MountKeygenController(service, c)

	// Mount "userid" controller
	u := NewUserIDController(service)
	app.MountUserIDController(service, u)

	// Start service
	if err := service.ListenAndServeTLS(":8080", "server.crt", "server.key"); err != nil {
		service.LogError("startup", "err", err)
	}
}

var errorCodes = map[string]string{
	"400": "Bad request - Returned when a required parameter is missing, an illegal parameter value is used",
	"403": "Forbidden - Returned for authentication or authorization errors including invalid key, insufficient admin access rights",
	"1":   "Unknown command - The specific config or operational command is not recognized",
	"2":   "Internal error - Check with technical support when seeing these errors",
	"3":   "Internal error - Check with technical support when seeing these errors",
	"4":   "Internal error - Check with technical support when seeing these errors",
	"5":   "Internal error - Check with technical support when seeing these errors",
	"6":   "Bad Xpath - The xpath specified in one or more attributes of the command is invalid. Check the API browser for proper xpath values",
	"7":   "Object not present - Object specified by the xpath is not present. For example, entry[@name=’value’] where no object with name ‘value’ is present",
	"8":   "Object not unique - For commands that operate on a single object, the specified object is not unique",
	"9":   "Internal error - Check with technical support when seeing these errors",
	"10":  "Reference count not zero - Object cannot be deleted as there are other objects that refer to it. For example, address object still in use in policy",
	"11":  "Internal error - Check with technical support when seeing these errors",
	"12":  "Invalid object - Xpath or element values provided are not complete",
	"13":  "Operation failed - A descriptive error message is returned in the response",
	"14":  "Operation not possible - Operation is not possible. For example, moving a rule up one position when it is already at the top",
	"15":  "Operation denied - For example, Admin not allowed to delete own account, Running a command that is not allowed on a passive device",
	"16":  "Unauthorized - The API role does not have access rights to run this query",
	"17":  "Invalid command - Invalid command or parameters",
	"18":  "Malformed command - The XML is malformed",
	"19":  "Success - Command completed successfully",
	"20":  "Success - Command completed successfully",
	"21":  "Internal error - Check with technical support when seeing these errors",
	"22":  "Session timed out - The session for this query timed out",
}
