package Models

import (
	"github.com/bezaeel/rest-api-mysql-gin/Config"
)

//GetAllContacts
func GetAllContacts(contact *[]Contact) (err error) {
	if err = Config.DB.Find(contact).Error; err != nil {
		return err
	}
	return nil
}

//CreateContact
func CreateContact(contact *Contact) (err error) {
	if err = Config.DB.Create(contact).Error; err != nil {
		return err
	}
	return nil
}

///GetContactById
func GetContactById(id string) (err error) {
	// if err = Config.DB.Where("id = ?", id).First().Error; err != nil {
	// 	return err
	// }
	return nil
}

//UpdateContact
func UpdateContact(id string) (err error) {
	//Config.DB
	return err
}
