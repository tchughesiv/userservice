package user_handles

import (
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	tokenhandlers "userservice-go/handlers/token-handlers"
	"userservice-go/types"
)

func FindUsers(findUsersCriteria types.FindUsersCriteria) (error, []types.User) {
	var usersList []types.User
	if len(findUsersCriteria.OrgId) == 0 || (len(findUsersCriteria.Emails) == 0 && len(findUsersCriteria.UserIds) == 0 && len(findUsersCriteria.Usernames) == 0) {
		return errors.New("unusable find users criteria"), usersList
	}

	if len(findUsersCriteria.Emails) > 0 {
		err, usersList := findUsersByEmails(findUsersCriteria)
		return err, usersList
	} else if len(findUsersCriteria.Usernames) > 0 {
		err, usersList := findUsersByUserNames(findUsersCriteria)
		return err, usersList
	} else if len(findUsersCriteria.UserIds) > 0 {
		err, usersList := findUsersByUserIds(findUsersCriteria)
		return err, usersList
	}

	return nil, usersList
}

func findUsersByEmails(findUsersCriteria types.FindUsersCriteria) (error, []types.User) {
	var usersList []types.User

	qPart := "q=org_id:" + findUsersCriteria.OrgId
	hostPath := types.KEYCLOAK_BACKEND_URL + types.KEYCLOAK_GET_BY_USERS + "?" + qPart

	for _, email := range findUsersCriteria.Emails {
		if len(email) != 0 {
			url := hostPath + "&" + "email=" + email
			log.Info().Msg(url)
			err, req, client := tokenhandlers.GetHttpClientAndRequestWithToken(http.MethodGet, url, nil)
			if err != nil {
				log.Error().Msg(err.Error())
				return err, usersList
			}

			if client != nil && req != nil {
				response, err := client.Do(req)
				if err != nil {
					log.Error().Msg(err.Error())
					return err, usersList
				}

				if response.StatusCode == http.StatusOK {
					responseData, err := ioutil.ReadAll(response.Body)
					if err != nil {
						log.Error().Msg(err.Error())
						return err, usersList
					}
					var users []types.User
					err = json.Unmarshal(responseData, &users)
					if err != nil {
						log.Error().Msg(err.Error())
						return err, usersList
					}
					usersList = append(usersList, users...)
				}
			}
		}
	}
	return nil, usersList
}

func findUsersByUserNames(findUsersCriteria types.FindUsersCriteria) (error, []types.User) {
	var usersList []types.User

	qPart := "q=org_id:" + findUsersCriteria.OrgId
	hostPath := types.KEYCLOAK_BACKEND_URL + types.KEYCLOAK_GET_BY_USERS + "?" + qPart

	for _, userName := range findUsersCriteria.Usernames {
		if len(userName) != 0 {
			url := hostPath + "&" + "username=" + userName
			log.Info().Msg(url)
			err, req, client := tokenhandlers.GetHttpClientAndRequestWithToken(http.MethodGet, url, nil)
			if err != nil {
				log.Error().Msg(err.Error())
				return err, usersList
			}

			if client != nil && req != nil {
				response, err := client.Do(req)
				if err != nil {
					log.Error().Msg(err.Error())
					return err, usersList
				}

				if response.StatusCode == http.StatusOK {
					responseData, err := ioutil.ReadAll(response.Body)

					if err != nil {
						log.Error().Msg(err.Error())
						return err, usersList
					}
					var users []types.User
					err = json.Unmarshal(responseData, &users)
					if err != nil {
						log.Error().Msg(err.Error())
						return err, usersList
					}
					usersList = append(usersList, users...)
				}
			}
		}
	}
	return nil, usersList
}

func findUsersByUserIds(findUsersCriteria types.FindUsersCriteria) (error, []types.User) {
	var usersList []types.User

	qPart := "q=org_id:" + findUsersCriteria.OrgId
	hostPath := types.KEYCLOAK_BACKEND_URL + types.KEYCLOAK_GET_BY_USERS + "?" + qPart

	for _, userId := range findUsersCriteria.UserIds {
		if len(userId) != 0 {
			// url := hostPath + "&" + "email=" + email
			url := hostPath + "&" + "id=" + userId
			log.Info().Msg(url)
			err, req, client := tokenhandlers.GetHttpClientAndRequestWithToken(http.MethodGet, url, nil)
			if err != nil {
				log.Error().Msg(err.Error())
				return err, usersList
			}

			if client != nil && req != nil {
				response, err := client.Do(req)
				if err != nil {
					log.Error().Msg(err.Error())
					return err, usersList
				}

				if response.StatusCode == http.StatusOK {
					responseData, err := ioutil.ReadAll(response.Body)

					if err != nil {
						log.Error().Msg(err.Error())
						return err, usersList
					}
					var users []types.User
					err = json.Unmarshal(responseData, &users)
					if err != nil {
						log.Error().Msg(err.Error())
						return err, usersList
					}
					usersList = append(usersList, users...)
				}
			}
		}
	}
	return nil, usersList
}