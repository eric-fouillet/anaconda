package anaconda

import (
	"net/url"
	"strconv"
	"strings"
)

// CreateList implements /lists/create.json
func (a TwitterApi) CreateList(name, description string, v url.Values) (list List, err error) {
	if v == nil {
		v = url.Values{}
	}
	v.Set("name", name)
	v.Set("description", description)

	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/lists/create.json", v, &list, _POST, response_ch}
	return list, (<-response_ch).err
}

// AddUserToList implements /lists/members/create.json
func (a TwitterApi) AddUserToList(screenName string, listID int64, v url.Values) (users []User, err error) {
	if v == nil {
		v = url.Values{}
	}
	v.Set("list_id", strconv.FormatInt(listID, 10))
	v.Set("screen_name", screenName)

	var addUserToListResponse AddUserToListResponse

	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/lists/members/create.json", v, &addUserToListResponse, _POST, response_ch}
	return addUserToListResponse.Users, (<-response_ch).err
}

// AddUsersToList implements /lists/members/create_all.json
func (a TwitterApi) AddUsersToList(listID int64, userIDs []int64, v url.Values) (users []User, err error) {
	if v == nil {
		v = url.Values{}
	}
	v.Set("list_id", strconv.FormatInt(listID, 10))
	userIDsParam := make([]string, len(userIDs))
	for i, id := range userIDs {
		userIDsParam[i] = strconv.FormatInt(id, 10)
	}
	v.Set("user_id", strings.Join(userIDsParam, ","))

	var addUserToListResponse AddUserToListResponse

	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/lists/members/create_all.json", v, &addUserToListResponse, _POST, response_ch}
	return addUserToListResponse.Users, (<-response_ch).err
}

// RemoveUsersFromList implements /lists/members/destroy_all.json
func (a TwitterApi) RemoveUsersFromList(listID int64, userIDs []int64, v url.Values) (users []User, err error) {
	if v == nil {
		v = url.Values{}
	}
	v.Set("list_id", strconv.FormatInt(listID, 10))
	userIDsParam := make([]string, len(userIDs))
	for i, id := range userIDs {
		userIDsParam[i] = strconv.FormatInt(id, 10)
	}
	v.Set("user_id", strings.Join(userIDsParam, ","))

	var addUserToListResponse AddUserToListResponse

	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/lists/members/destroy_all.json", v, &addUserToListResponse, _POST, response_ch}
	return addUserToListResponse.Users, (<-response_ch).err
}

// GetListMembers implements /lists/members
// slug, owner_screen_name, owner_id, count, cursor are all optional values
func (a TwitterApi) GetListMembers(listId int64, v url.Values) (users []User, err error) {
	if v == nil {
		v = url.Values{}
	}
	v.Set("list_id", strconv.FormatInt(listId, 10))

	var listMembersResponse GetListMembersResponse

	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/lists/members.json", v, &listMembersResponse, _GET, response_ch}
	return listMembersResponse.Users, (<-response_ch).err
}

// GetListsOwnedBy implements /lists/ownerships.json
// screen_name, count, and cursor are all optional values
func (a TwitterApi) GetListsOwnedBy(userID int64, v url.Values) (lists []List, err error) {
	if v == nil {
		v = url.Values{}
	}
	v.Set("user_id", strconv.FormatInt(userID, 10))

	var listResponse ListResponse

	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/lists/ownerships.json", v, &listResponse, _GET, response_ch}
	return listResponse.Lists, (<-response_ch).err
}

func (a TwitterApi) GetListTweets(listID int64, includeRTs bool, v url.Values) (tweets []Tweet, err error) {
	if v == nil {
		v = url.Values{}
	}
	v.Set("list_id", strconv.FormatInt(listID, 10))
	v.Set("include_rts", strconv.FormatBool(includeRTs))

	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/lists/statuses.json", v, &tweets, _GET, response_ch}
	return tweets, (<-response_ch).err
}
