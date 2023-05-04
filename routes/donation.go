package routes

import (
	"holyways/handlers"
	"holyways/pkg/middleware"
	"holyways/pkg/mysql"
	"holyways/repositories"

	"github.com/labstack/echo/v4"
)

func DonationRoute(e *echo.Group) {
	donationRepository := repositories.RepositoryDonation(mysql.DB)
	h := handlers.HandlerDonation(donationRepository)

	e.GET("/donations", h.FindDonation)
	e.GET("/donation/:id", h.GetDonation)
	e.POST("/donation", middleware.Auth(middleware.UploadFile(h.CreateDonation)))
	e.PATCH("/donation/:id", middleware.Auth(middleware.UploadFile(h.Updatedonation)))
	e.DELETE("/donation/:id", middleware.Auth(h.DeleteDonation))
}
