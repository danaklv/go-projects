package repositories

import "td/back/models"

func TodoRep(todo *models.ToDo) error {
	err := Db.AutoMigrate(&models.ToDo{})

	if err != nil {
		return err
	}

	res := Db.Create(&models.ToDo{
		Id: todo.Id,
		Type: todo.Type,
		Descrription: todo.Descrription,
	})

	if res.Error != nil {
		return res.Error
	}
	return nil

}
