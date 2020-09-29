package common

import "github.com/sirupsen/logrus"

func BuildContext(context string, fields map[string]interface{}) logrus.Fields {
	logFields := logrus.Fields{
		"context": context,
	}

	for key, value := range fields {
		logFields[key] = value
	}

	return logFields
}
