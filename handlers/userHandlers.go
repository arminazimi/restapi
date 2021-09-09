package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/asdine/storm/v3"
	"gopkg.in/mgo.v2/bson"

	"github.com/arminazimi/restapi/user"
)

func usersGetAll(w http.ResponseWriter, r *http.Request) {
	users, err := user.All()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"users": users})

}

func userGetOne(w http.ResponseWriter, _ *http.Request, id bson.ObjectId) {
	u, err := user.One(id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"user": u})
}

func bodyToUser(r *http.Request, u *user.User) error {
	if r.Body == nil {
		return errors.New("no request body")
	}
	if u == nil {
		return errors.New("user is nil")
	}
	bd, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bd, u)
}

func userPostOne(w http.ResponseWriter, r *http.Request) {
	u := new(user.User)
	err := bodyToUser(r, u)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	u.Id = bson.NewObjectId()
	err = u.Save()
	if err != nil {
		if err == user.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Location", "/users/"+u.Id.Hex())
	w.WriteHeader(http.StatusCreated)

}

func userPutOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	u := new(user.User)
	err := bodyToUser(r, u)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	u.Id = id
	err = u.Save()
	if err != nil {
		if err == user.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"user": u})

}

func userPatchOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	u, err := user.One(id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	err = bodyToUser(r, u)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	u.Id = id
	err = u.Save()
	if err != nil {
		if err == user.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"user": u})

}

func userDeleteOne(w http.ResponseWriter, _ *http.Request, id bson.ObjectId) {
	err := user.Delete(id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}