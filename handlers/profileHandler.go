package handlers

import "net/http"

func GetProfileData(r *http.Request) map[string]interface{} {
	profileData := make(map[string]interface{})

	session, _ := Store.Get(r, "session-name")

	profileData["username"] = session.Values["TOKEN"]

	return profileData
}
