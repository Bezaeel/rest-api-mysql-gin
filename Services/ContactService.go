package Services

import (
	"github.com/bezaeel/rest-api-mysql-gin/Models"

	"github.com/bezaeel/rest-api-mysql-gin/Config"
)

//GetAllContacts
func GetAllContacts(contact *[]Models.Contact) (err error) {
	if err = Config.DB.Find(contact).Error; err != nil {
		return err
	}
	return nil
}

//CreateContact
func CreateContact(contact *Models.Contact) (err error) {
	if err = Config.DB.Create(contact).Error; err != nil {
		return err
	}
	return nil
}

///GetContactById
func GetContactById(id string, contact *Models.Contact) (err error) {
	if err = Config.DB.Where("id = ?", id).First(&contact).Error; err != nil {
		return err
	}
	return nil
}

//UpdateContact
func UpdateContact(id string, contact *Models.Contact) (err error) {
	//Config.DB
	Config.DB.Where("id = ?", id).Save(contact)
	return nil
}

//DeleteContact
func DeleteContact(id string, contact *Models.Contact) (err error) {
	Config.DB.Where("id = ?", id).Delete(contact)
	return nil
}
