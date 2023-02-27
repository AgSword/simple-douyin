package relationDB

import (
	// "log"

	// "github.com/karry-almond/model"
	"github.com/AgSword/simpleDouyin/cmd/favorite/packages/model"
	"gorm.io/gorm"
	"github.com/AgSword/simpleDouyin/kitex_gen/relation"
	// "golang.org/x/tools/go/analysis/passes/nilfunc"
)

// NewRelation 创建一条关注关系
func 	NewRelation(userID int64, toUserID int64) (status int32, err error) {

	relation := model.Relation{
		//TODO:ID这里不是逐主键
		UserId:   userID,
		ToUserId: toUserID}

	//新建关注、新增关注为同一事务
	err = Db.Transaction(func(tx *gorm.DB) error {
		// 创建一条关系数据
		if err := tx.Create(&relation).Error; err != nil {
			return err
		}

		//更改被关注user的粉丝数
		var toUser model.User
		if err := tx.Select("*").First(&model.User{ID: toUserID}).Scan(&toUser).Error; err != nil {
			return err
		}
		if err := tx.Debug().Model(&model.User{}).Where("id = ?", toUserID).Update("follower_count", toUser.FollowerCount+1).Error; err != nil {

			return err
			
		}	

		//更改当前用户user的关注数
		var user model.User
		if err := tx.Select("*").First(&model.User{ID: userID}).Scan(&user).Error; err != nil {
			return err
		}
		if err := tx.Debug().Model(&model.User{}).Where("id = ?", userID).Update("follow_count", user.FollowCount+1).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return 1, err
	}

	return 0, nil

}

// CancelRelation 取消一条关注信息
func CancelRelation(userID int64, toUserID int64) (status int32, err error) {

	var relation model.Relation

	err = Db.Transaction(func(tx *gorm.DB) error {
		//通过user_id和to_user_id找到要删除的relation记录
		if err := tx.Select("*").First(&model.Relation{UserId: userID, ToUserId: toUserID}).Scan(&relation).Error; err != nil {
			return err
		}
		//删除这条favorite记录
		if err := tx.Delete(&relation).Error; err != nil {
			return err
		}
		//更改被关注者to_user的follower_count
		var toUser model.User

		if err := tx.Select("*").First(&model.User{ID: toUserID}).Scan(&toUser).Error; err != nil {
			return err
		}

		if err := tx.Model(&model.User{ID: toUserID}).Update("follower_count", toUser.FollowerCount-1).Error; err != nil {
			return err
		}
		//更改当前user的follow_count
		var user model.User

		if err := tx.Select("*").First(&model.User{ID: userID}).Scan(&user).Error; err != nil {
			return err
		}

		if err := tx.Model(&model.User{ID: userID}).Update("follow_count", toUser.FollowCount-1).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return 1, err
	}

	return 0, nil
}

// GetFollowList 获取关注列表
func GetFollowList(user int64) (status int32, followList []*relation.User, err error) {

	var ids []int
	//err = Db.Where(&model.Relation{UserId: user}).Pluck("to_user_id", &ids).Error
	err = Db.Debug().Model(&model.Relation{UserId: user}).Where("user_id = ?", user).Pluck("to_user_id", &ids).Error

	if err != nil {
		return 1, nil, err
	}

	var list []model.User

	err = Db.Table("user").Where("id in ?", ids).Scan(&list).Error
	if err != nil {
		return 1, nil, err
	}

	var respList []*relation.User
	for _, m := range list {
		var tmp relation.User
		tmp.Id = m.ID
		tmp.Name = m.Name
		tmp.FollowCount = m.FollowCount
		tmp.FollowerCount = m.FollowerCount
		respList = append(respList, &tmp)
	}

	return 0, respList, nil
}

// GetFollowerList 获取粉丝列表
func GetFollowerList(user int64) (status int32, followList []*relation.User, err error) {

	var ids []int
	err = Db.Debug().Model(&model.Relation{}).Where("to_user_id = ?", user).Pluck("user_id", &ids).Error
	if err != nil {
		return 1, nil, err
	}

	var list []model.User
	err = Db.Table("user").Where("id in ?", ids).Find(&list).Error
	if err != nil {
		return 1, nil, err
	}

	var respList []*relation.User
	for _, m := range list {
		var tmp relation.User
		tmp.Id = m.ID
		tmp.Name = m.Name
		tmp.FollowCount = m.FollowCount
		tmp.FollowerCount = m.FollowerCount
		respList = append(respList, &tmp)
	}
	return 0, respList, nil
}
