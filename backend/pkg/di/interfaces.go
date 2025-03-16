package di

// import (
// 	"backend/internal/client"
// 	"backend/internal/file"
// 	"backend/internal/link"
// 	"backend/internal/product"
// )

// type IClientRepository interface {
// 	Create(name, telephone, mail, company string) (*client.Client, error)
// 	FindByData(name, telephone, mail, company string) (*client.Client, error)
// }

// type ILinkRepository interface {
// 	Create(valid int, count int, product_uid, file_uid, client_uid string) (*link.Link, error)
// }

// type IProductRepository interface {
// 	FindByUid(uid string) (*product.Product, error)
// 	CetProdsByUids(productUids []string) ([]*product.Product, error)
// }

// type IFileRepository interface {
// 	FindByUid(uid string) (*file.File, error)
// 	GetFilesByProdUid(productUids []string) ([]*file.File, error)
// }
