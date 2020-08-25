package repository

import (
	"fmt"
	"photoalbum/models"
	"github.com/jinzhu/gorm"
)

type UserRepositoryError struct {
	error
}




type UserReposityory struct {
	DB *gorm.DB
}

func GetUserRespository(db *gorm.DB) UserReposityory{
	return UserReposityory {DB: db}
}

func (u *UserReposityory) GetAll() ([]models.User, error){
	var users []models.User
	u.DB.Find(&users)
	if len(users) == 0 || users ==nil{
		return users, &UserRepositoryError{fmt.Errorf("no record found")}
	}else{
		return users, nil

	}
	err := fmt.Errorf("Server error")
	return nil, &UserRepositoryError{err}
}

func (u UserReposityory) GetByID(id int) (models.User, error){
	var user models.User
	d := u.DB.Debug().First(&user, id)

	if d.RowsAffected == 0{
		return  user, &UserRepositoryError{fmt.Errorf("can not create user")}
	}else{
		return user, nil
	}
}

func (u UserReposityory) Save(user models.User) (models.User,error){
	d := u.DB.Save(&user)
	if d.Error !=nil{
		return  user, &UserRepositoryError{fmt.Errorf(d.Error.Error())}
	}else{
		return user, nil
	}
	return  user, &UserRepositoryError{fmt.Errorf("can not create user")}
}

func (u *UserReposityory) Delete(id int) (bool, error) {
	var user models.User
	d := u.DB.Debug().First(&user, id)
	fmt.Println(d.Value)
	if d.Error !=nil{
		return  false, &UserRepositoryError{fmt.Errorf(d.Error.Error())}
	}else{
		d = u.DB.Debug().Unscoped().Delete(&user)
	}
	if d.Error !=nil{
		return  false, &UserRepositoryError{fmt.Errorf(d.Error.Error())}
	}else{
		return true, nil
	}
}



