package post

import (
	"backend/pkg/logger"
	"backend/pkg/request"
	"backend/pkg/response"
	"net/http"

	"gorm.io/gorm"
)

type PostHandlerDeps struct {
	PostRepository *PostRepository
}

type PostHandler struct {
	PostRepository *PostRepository
}

func NewPostHandler(router *http.ServeMux, deps *PostHandlerDeps) {
	handler := &PostHandler{
		PostRepository: deps.PostRepository,
	}
	router.HandleFunc("POST /api/post", handler.Create())
	router.HandleFunc("GET /api/post/{uid}", handler.Read())
	router.HandleFunc("PATCH /api/post/{uid}", handler.Update())
	router.HandleFunc("DELETE /api/post/{uid}", handler.Delete())

	router.HandleFunc("POST /api/posts", handler.GetProds())

}

func (handler *PostHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[PostCreateRequest](&w, r)
		if err != nil {
			logger.ERROR(err)
			return
		}

		post := NewPost(body.Name, body.Description)

		existedPost, _ := handler.PostRepository.FindByName(post.Name)
		if existedPost != nil {
			http.Error(w, existedPost.Name+" is already exists", http.StatusBadRequest)
			logger.ERROR(existedPost.Name+" is already exists", http.StatusBadRequest)
			return
		}

		for {
			existedPost, _ = handler.PostRepository.FindByUid(post.Uid)
			if existedPost == nil {
				break
			}
			post.GenerateHash()
		}

		createdProd, err := handler.PostRepository.Create(post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.ERROR(err.Error(), http.StatusBadRequest)
			return
		}

		response.Json(w, createdProd.Name+" has been added successfully", http.StatusOK)

	}
}

func (handler *PostHandler) Read() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := r.PathValue("uid")

		existedProd, err := handler.PostRepository.FindByUid(uid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.ERROR(err.Error(), http.StatusBadRequest)
			return
		}

		response.Json(w, existedProd, http.StatusOK)
	}
}

func (handler *PostHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[PostUpdateRequest](&w, r)
		if err != nil {
			logger.ERROR(err)
			return
		}

		uid := r.PathValue("uid")

		_, err = handler.PostRepository.FindByUid(uid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.ERROR(err.Error(), http.StatusBadRequest)
			return
		}

		post, err := handler.PostRepository.Update(uid, &Post{
			Model: gorm.Model{},
			Name:  body.Name,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.ERROR(err.Error(), http.StatusBadRequest)
			return
		}

		response.Json(w, post.Name+" has been updated successfully", http.StatusOK)
	}
}

func (handler *PostHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := r.PathValue("uid")

		_, err := handler.PostRepository.FindByUid(uid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.ERROR(err.Error(), http.StatusBadRequest)
			return
		}

		err = handler.PostRepository.Delete(uid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.ERROR(err.Error(), http.StatusBadRequest)
			return
		}

		response.Json(w, "The post has been successfully removed", http.StatusOK)
	}
}

func (handler *PostHandler) GetProds() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[GetPostsRequest](&w, r)
		if err != nil {
			logger.ERROR(err)
			return
		}

		count, err := handler.PostRepository.Count()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.ERROR(err.Error(), http.StatusBadRequest)
			return
		}

		posts, err := handler.PostRepository.GetProds(body.Limit, body.Offset, body.Columns)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.ERROR(err.Error(), http.StatusBadRequest)
			return
		}

		response.Json(w, &GetPostsResponse{
			Columns: []string{"uid", "name", "text", "created_at", "updated_at"},
			Data:    posts,
			Count:   count,
		}, http.StatusOK)
	}
}
