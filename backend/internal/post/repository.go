package post

import (
	"backend/pkg/db"

	"gorm.io/gorm/clause"
)

type PostRepository struct {
	Db *db.Db
}

func NewPostRepository(database *db.Db) *PostRepository {
	return &PostRepository{
		Db: database,
	}
}

func (repo *PostRepository) Create(post *Post) (*Post, error) {
	result := repo.Db.
		Table("posts").
		Create(post)
	if result.Error != nil {
		return nil, result.Error
	}
	return post, nil
}

func (repo *PostRepository) FindByName(name string) (*Post, error) {
	var post Post
	result := repo.Db.
		Table("posts").
		First(&post, "name = ?", name)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}

func (repo *PostRepository) FindByUid(uid string) (*Post, error) {
	var post Post
	result := repo.Db.
		Table("posts").
		First(&post, "uid = ?", uid)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}

func (repo *PostRepository) Update(uid string, post *Post) (*Post, error) {
	result := repo.Db.
		Table("posts").
		Where("uid = ?", uid).
		Clauses(clause.Returning{}).
		Updates(post)
	if result.Error != nil {
		return nil, result.Error
	}
	return post, nil
}

func (repo *PostRepository) Delete(uid string) error {
	var product Post
	// var f file.File
	result := repo.Db.
		Table("posts").
		Delete(&product, "uid = ?", uid)
	if result.Error != nil {
		return result.Error
	}
	// resultFiles := repo.Db.
	// 	Table("files").
	// 	Delete(&f, "product_uid = ?", uid)
	// if resultFiles.Error != nil {
	// 	return result.Error
	// }

	return nil
}

func (repo *PostRepository) Count() (int64, error) {
	var count int64
	result := repo.Db.
		Table("posts").
		Where("deleted_at is null").
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil

}

func (repo *PostRepository) GetProds(limit, offset int, columns []string) ([]PostResponse, error) {
	var posts []PostResponse

	if len(columns) == 0 {
		return posts, nil
	}

	result := repo.Db.
		Table("posts").
		Where("deleted_at is null").
		Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Scan(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (repo *PostRepository) CetPostsByUids(postUids []string) ([]*Post, error) {
	var posts []*Post

	result := repo.Db.
		Table("posts").
		Where("deleted_at is null").
		Where("uid IN ?", postUids).
		Scan(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}
