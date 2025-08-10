package domain

import (
	"github.com/amorindev/headless-ecomm-cms/pkg/app/user/domain"
	"github.com/sirupsen/logrus"
)

func Logger(u domain.User, layerName string, action string) {
	greenColor := "\033[32m"
	resetColor := "\033[0m"
	logrus.Infof(
		"%s Layer: %s Action: %s %s",
		greenColor, layerName, action, resetColor,
	)
	logrus.Infof(
		"%s ID: %s %T %s",
		"\033[33m", u.ID, u.ID, resetColor,
	)
	logrus.Infof(
		"%s Email: %s %T %s",
		"\033[33m", u.Email, u.Email, resetColor,
	)
	logrus.Infof(
		"%s Email verified: %v %T %s",
		"\033[33m", u.EmailVerified, u.EmailVerified, resetColor,
	)
	if u.Name == nil {
		logrus.Infof(
			"%s Name: %v %T %s",
			"\033[33m", u.Name, u.Name, resetColor,
		)

	} else {
		logrus.Infof(
			"%s Name: %v %T %s",
			"\033[33m", *u.Name, u.Name, resetColor,
		)
	}
	if u.CreatedAt == nil {
		logrus.Infof(
			"%s CreatedAt: %v %T %s",
			"\033[33m", u.CreatedAt, u.CreatedAt, resetColor,
		)
	} else {
		logrus.Infof(
			"%s CreatedAt: %v %T %s",
			"\033[33m", *u.CreatedAt, u.CreatedAt, resetColor,
		)
	}

	if u.UpdatedAt == nil {
		logrus.Infof(
			"%s UpdatedAt: %v %T %s ",
			"\033[33m", u.UpdatedAt, u.UpdatedAt, resetColor,
		)
	} else {
		logrus.Infof(
			"%s UpdatedAt: %v %T %s",
			"\033[33m", *u.UpdatedAt, u.UpdatedAt, resetColor,
		)
	}

}
