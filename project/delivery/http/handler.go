package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Gidraff/task-manager-service/pkg/utils/helpers"
	"github.com/Gidraff/task-manager-service/project"
	"github.com/gorilla/mux"
)

const (
	// ServerErrorMessage http error
	ServerErrorMessage = "Server failed to process request. Try again"

	// BadRequestMessage http error
	BadRequestMessage = "Failed to process request. Please check JSON request format."

	// ResourceCreatedMessage success message
	ResourceCreatedMessage = "Resource successfully created."

	// ResourceUpdatedMessage success message
	ResourceUpdatedMessage = "Resource successfully updated."

	// ResourceDeletedMessage success message
	ResourceDeletedMessage = "Resource successfully deleted."

	// ResourceNotAvailableMessage http message
	ResourceNotAvailableMessage = "Resource not found"
)

// ProjectInput encapsulate request body
type ProjectInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ProjectHandler encapsulate usecase
type ProjectHandler struct {
	projectUC project.IProjectUsecase
}

// NewProjectHandler will initialize user resources endpoint
func NewProjectHandler(projectUC project.IProjectUsecase) *ProjectHandler {
	return &ProjectHandler{projectUC: projectUC}
}

// CreateProject creates a bew project
func (ph *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	userID, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerErrorMessage, w)
		return
	}
	var project ProjectInput
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&project); err != nil {
		helpers.Response(http.StatusBadRequest, BadRequestMessage, w)
		return
	}

	isProjectNameValid, err := helpers.ValidateProjectName(project.Name)
	if !isProjectNameValid {
		helpers.Response(http.StatusBadRequest, err.Error(), w)
		return
	}
	duplicateProject, _ := ph.projectUC.FetchProjectByName(project.Name, userID)
	if duplicateProject != nil {
		helpers.Response(http.StatusConflict, "A resource with the name already exist", w)
		return
	}

	err = ph.projectUC.CreateNewProject(userID, project.Name, project.Description)
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerErrorMessage, w)
		return
	}
	helpers.Response(http.StatusCreated, ResourceCreatedMessage, w)
	return
}

// GetAllProjects handles get request
func (ph *ProjectHandler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	userID, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerErrorMessage, w)
		return
	}
	projects, err := ph.projectUC.FetchAllProjects(userID)
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerErrorMessage, w)
		return
	}
	helpers.PayloadResponse(projects, w)
	return
}

// GetProjectByID GET project by ID
func (ph *ProjectHandler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	userID, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerErrorMessage, w)
		return
	}
	params := mux.Vars(r)
	id := params["id"]
	projectID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("delivery/http: Failed to convert id", err)
	}

	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerErrorMessage, w)
		return
	}
	project, err := ph.projectUC.FetchProjectByID(uint64(projectID), userID)
	if err != nil {
		helpers.Response(http.StatusBadRequest, ResourceNotAvailableMessage, w)
		return
	}
	helpers.PayloadResponse(project, w)
	return
}

// GetProjectByName GET project by ID
func (ph *ProjectHandler) GetProjectByName(w http.ResponseWriter, r *http.Request) {
	userID, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerErrorMessage, w)
		return
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		helpers.Response(http.StatusBadRequest, "invalid query params", w)
		return
	}

	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerErrorMessage, w)
		return
	}
	project, err := ph.projectUC.FetchProjectByName(name, userID)
	if err != nil {
		helpers.Response(http.StatusBadRequest, ResourceNotAvailableMessage, w)
		return
	}
	helpers.PayloadResponse(project, w)
	return
}

// UpdateProject GET project by ID
func (ph *ProjectHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	userID, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		helpers.Response(http.StatusUnauthorized, "Unauthorized", w)
		return
	}
	params := mux.Vars(r)
	id := params["id"]
	projectID, err := strconv.Atoi(id)
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerErrorMessage, w)
		return
	}

	var project ProjectInput
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&project); err != nil {
		helpers.Response(http.StatusBadRequest, BadRequestMessage, w)
		return
	}
	err = ph.projectUC.UpdateProject(uint64(projectID), userID, project.Name, project.Description)
	if err != nil {
		helpers.Response(http.StatusBadRequest, ResourceNotAvailableMessage, w)
		return
	}
	helpers.Response(http.StatusOK, ResourceUpdatedMessage, w)
	return
}

// DeleteProject handles delete request
func (ph *ProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	userID, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerErrorMessage, w)
		return
	}
	params := mux.Vars(r)
	id := params["id"]
	projectID, _ := strconv.Atoi(id)

	err = ph.projectUC.DeleteProject(uint64(projectID), userID)
	if err != nil {
		helpers.Response(http.StatusBadRequest, ResourceNotAvailableMessage, w)
		return
	}
	helpers.Response(http.StatusOK, ResourceDeletedMessage, w)
	return
}
