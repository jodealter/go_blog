package dao

import (
	"github.com/jodealter/go_blog/models"
	"log"
)

func UpdatePost(post *models.Post) {
	_, err := DB.Exec("update blog_post set title = ?,content = ?,markdown = ?,category_id = ?,type = ?,slug = ?,update_at = ? where pid = ?",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.Type,
		post.Slug,
		post.UpdateAt,
		post.Pid,
	)
	if err != nil {
		log.Println(err)
	}
}
func SavePost(post *models.Post) {
	res, err := DB.Exec("insert into blog_post "+
		"(title,content,markdown,category_id,user_id,view_count,type,slug,create_at,update_at)"+
		"values(?,?,?,?,?,?,?,?,?,?)",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.UserId,
		post.ViewCount,
		post.Type,
		post.Slug,
		post.CreateAt,
		post.UpdateAt,
	)
	if err != nil {
		log.Println(err)
	}
	Pid, _ := res.LastInsertId()
	post.Pid = int(Pid)
}
func CountGetAllPostByCategoryId(cid int) (count int) {
	rows := DB.QueryRow("select count(1) from blog_post where category_id = ?", cid)
	_ = rows.Scan(&count)
	return
}

func CountGetAllPost() (count int) {
	rows := DB.QueryRow("select count(1) from blog_post")
	_ = rows.Scan(&count)
	return
}

func CountGetAllPostBySlug(slug string) (count int) {
	rows := DB.QueryRow("select count(1) from blog_post where slug = ?", slug)
	_ = rows.Scan(&count)
	return
}
func GetPostPage(page, pagesize int) ([]models.Post, error) {
	page = (page - 1) * pagesize
	row, err := DB.Query("select * from blog_post limit ?,?", page, pagesize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for row.Next() {
		var post models.Post
		err := row.Scan(&post.Pid, &post.Title, &post.Content, &post.Markdown, &post.CategoryId, &post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostAll() ([]models.Post, error) {
	row, err := DB.Query("select * from blog_post")
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for row.Next() {
		var post models.Post
		err := row.Scan(&post.Pid, &post.Title, &post.Content, &post.Markdown, &post.CategoryId, &post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostPageBCategoryId(cid int, page, pagesize int) ([]models.Post, error) {
	page = (page - 1) * pagesize
	row, err := DB.Query("select * from blog_post where category_id = ? limit ?,?", cid, page, pagesize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for row.Next() {
		var post models.Post
		err := row.Scan(&post.Pid, &post.Title, &post.Content, &post.Markdown, &post.CategoryId, &post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func GetPostByid(pid int) (models.Post, error) {
	row := DB.QueryRow("select * from blog_post where pid = ?", pid)
	var post models.Post
	if row.Err() != nil {
		return post, row.Err()
	}

	err := row.Scan(&post.Pid, &post.Title, &post.Content, &post.Markdown, &post.CategoryId, &post.UserId,
		&post.ViewCount,
		&post.Type,
		&post.Slug,
		&post.CreateAt,
		&post.UpdateAt,
	)
	if err != nil {
		return post, row.Err()
	}
	return post, nil
}

func GetPostPageBySlug(slug string, page, pagesize int) ([]models.Post, error) {
	page = (page - 1) * pagesize
	row, err := DB.Query("select * from blog_post where slug = ? limit ?,?", slug, page, pagesize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for row.Next() {
		var post models.Post
		err := row.Scan(&post.Pid, &post.Title, &post.Content, &post.Markdown, &post.CategoryId, &post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func GetPostSearch(condition string) ([]models.Post, error) {

	row, err := DB.Query("select * from blog_post where title like ?", "%"+condition+"%")
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for row.Next() {
		var post models.Post
		err := row.Scan(&post.Pid, &post.Title, &post.Content, &post.Markdown, &post.CategoryId, &post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
