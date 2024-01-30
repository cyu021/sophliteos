package service

import (
	"encoding/base64"
	"sophliteos/global"
	"sophliteos/logger"
	"sophliteos/pkg/constant"
	"sophliteos/pkg/dto"
	"sophliteos/pkg/model"
	"sophliteos/pkg/utils/copier"
	"sophliteos/pkg/utils/ssh"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type HostService struct{}

func (u *HostService) TestByInfo(req ssh.ConnInfo) bool {
	if len(req.Password) == 0 {
		return false
	}

	var connInfo ssh.ConnInfo
	_ = copier.Copy(&connInfo, &req)
	connInfo.PrivateKey = []byte(req.PrivateKey)
	connInfo.AuthMode = "password"

	client, err := connInfo.NewClient()
	if err != nil {
		return false
	}
	defer client.Close()
	return true
}

func (u *HostService) List() ([]dto.CommandInfo, error) {
	commands, err := GetList()
	if err != nil {
		return nil, constant.ErrRecordNotFound
	}
	var dtoCommands []dto.CommandInfo
	for _, command := range commands {
		var item dto.CommandInfo
		if err := copier.Copy(&item, &command); err != nil {
			return nil, errors.WithMessage(constant.ErrStructTransform, err.Error())
		}
		dtoCommands = append(dtoCommands, item)
	}
	return dtoCommands, err
}
func GetList() ([]model.Command, error) {
	var commands []model.Command
	db := global.DB.Model(&model.Command{})

	err := db.Find(&commands).Error
	return commands, err
}

func (u *HostService) Create(commandDto dto.CommandInfo) error {
	var command model.Command
	if err := global.DB.Where("name = ?", commandDto.Name).First(&command); err.Error != gorm.ErrRecordNotFound {
		return constant.ErrRecordExist
	}

	command.Command = commandDto.Command
	command.Name = commandDto.Name
	if err := global.DB.Create(&command).Error; err != nil {
		return err
	}
	return nil
}

func (u *HostService) Delete(commandDto dto.CommandInfo) error {
	var command model.Command
	if err := global.DB.Where("name = ?", commandDto.Name).First(&command); err.Error == gorm.ErrRecordNotFound {
		return constant.ErrRecordNotFound
	}

	if err := global.DB.Delete(&command).Error; err != nil {
		logger.Error("删除数据时出错:%v", err)
		return err
	}
	return nil
}

func (u *HostService) Update(commandDto dto.CommandInfo) error {
	var command model.Command
	if err := global.DB.Where("name = ?", commandDto.Name).First(&command); err.Error == gorm.ErrRecordNotFound {
		return constant.ErrRecordNotFound
	}
	command.Command = commandDto.Command
	command.Name = commandDto.Name
	if err := global.DB.Save(&command).Error; err != nil {
		logger.Error("修改数据时出错:%v", err)
		return err
	}
	return nil
}

func (u *HostService) CreateHost(req dto.HostOperate) error {
	var host dto.HostOperate
	if err := global.DB.Where("addr = ?", req.Addr).First(&host); err.Error != gorm.ErrRecordNotFound {
		return constant.ErrRecordExist
	}

	var err error
	if len(req.Password) != 0 {
		req.Password, err = u.EncryptHost(req.Password)
		if err != nil {
			return err
		}
	}

	if err := global.DB.Create(&req).Error; err != nil {
		return err
	}
	return nil
}

func (u *HostService) DeleteHost(req dto.HostOperate) error {
	var host dto.HostOperate
	if err := global.DB.Where("addr = ?", req.Addr).First(&host); err.Error == gorm.ErrRecordNotFound {
		return constant.ErrRecordNotFound
	}

	if err := global.DB.Delete(&host).Error; err != nil {
		logger.Error("删除数据时出错:%v", err)
		return err
	}
	return nil
}

func (u *HostService) UpdateHost(req dto.HostOperate) error {
	var host dto.HostOperate
	if err := global.DB.Where("addr = ?", req.Addr).First(&host); err.Error == gorm.ErrRecordNotFound {
		return constant.ErrRecordNotFound
	}

	var err error
	if len(req.Password) != 0 {
		req.Password, err = u.EncryptHost(req.Password)
		if err != nil {
			return err
		}
	}

	host.Name = req.Name
	host.Description = req.Description
	host.Port = req.Port
	host.Password = req.Password
	host.User = req.User

	if err := global.DB.Save(&host).Error; err != nil {
		logger.Error("修改数据时出错:%v", err)
		return err
	}
	return nil
}

func (u *HostService) ListHost() ([]dto.HostOperate, error) {
	var hosts []dto.HostOperate
	if err := global.DB.Find(&hosts).Error; err != nil {
		return nil, err
	}
	return hosts, nil

}

func (u *HostService) GetHostInfo(ip string) (dto.HostOperate, error) {
	var host dto.HostOperate
	if err := global.DB.Where("addr = ?", ip).First(&host); err.Error == gorm.ErrRecordNotFound {
		return host, err.Error
	}

	return host, nil
}

func (u *HostService) EncryptHost(itemVal string) (string, error) {
	privateKey, err := base64.StdEncoding.DecodeString(itemVal)
	if err != nil {
		return "", err
	}

	return string(privateKey), nil
}
