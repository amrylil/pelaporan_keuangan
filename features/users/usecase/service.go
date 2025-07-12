package usecase

import (
	"errors"
	user "pelaporan_keuangan/features/users"
	"pelaporan_keuangan/features/users/dtos"
	"pelaporan_keuangan/helpers"
	"strconv"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model     user.Repository
	hash      helpers.HashInterface
	validator helpers.ValidationInterface
	jwt       helpers.JWTInterface
}

func New(model user.Repository, hash helpers.HashInterface, validator helpers.ValidationInterface, jwt helpers.JWTInterface) user.Usecase {
	return &service{
		model:     model,
		hash:      hash,
		validator: validator,
		jwt:       jwt,
	}
}

func (svc *service) FindAll(page, size int) ([]dtos.ResUser, int64, error) {
	var users []dtos.ResUser

	usersEnt, total, err := svc.model.GetAll(page, size)
	if err != nil {
		log.Error(err.Error())
		return nil, 0, err
	}

	for _, user := range usersEnt {
		var data dtos.ResUser

		if err := smapping.FillStruct(&data, smapping.MapFields(user)); err != nil {
			log.Error(err.Error())
			return nil, 0, err
		}

		users = append(users, data)
	}

	return users, total, nil
}

func (svc *service) FindByID(userID uint64) (*dtos.ResUser, error) {
	res := dtos.ResUser{}
	user, err := svc.model.SelectByID(userID)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	err = smapping.FillStruct(&res, smapping.MapFields(user))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &res, nil
}

func (svc *service) Create(newUser dtos.InputUser) error {
	user := user.User{}

	err := smapping.FillStruct(&user, smapping.MapFields(newUser))
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	user.ID = helpers.GenerateID()
	user.Password = svc.hash.HashPassword(newUser.Password)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.IsActive = true
	user.UserType = "user"
	err = svc.model.Insert(user)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Login(user dtos.LoginRequest) (*dtos.ResUser, error) {
	errMap, err := svc.validator.ValidateRequest(user)
	if err != nil {
		log.Error("validation error: ", err)
		if errMap != nil {
			log.Error("validation map: ", errMap)
		}
		// Langsung hentikan jika validasi gagal
		return nil, errors.New("validation failed")
	}

	// 2. Dapatkan user berdasarkan email
	userData, err := svc.model.GetUserByEmail(user.Email)
	if err != nil {
		// Jika user tidak ditemukan atau ada error DB, LOG dan KEMBALIKAN error.
		// Ini akan mencegah nil pointer di langkah berikutnya.
		log.Error("error getting user by email: ", err.Error())
		return nil, errors.New("login failed: invalid email or password")
	}

	// 3. Bandingkan password
	if !svc.hash.CompareHash(user.Password, userData.Password) {
		// Jika password tidak cocok, kembalikan error yang sama agar lebih aman.
		return nil, errors.New("login failed: invalid email or password")
	}

	// 4. Buat response dan generate token
	resUser := dtos.ResUser{}
	userType := userData.UserType
	tokenData := svc.jwt.GenerateJWT(strconv.FormatUint(uint64(userData.ID), 10), userType)

	// Gunakan "comma ok" idiom untuk type assertion yang lebih aman untuk mencegah panic
	accessToken, ok := tokenData["access_token"].(string)
	if !ok {
		log.Error("failed to assert access_token to string")
		return nil, errors.New("internal server error: failed to generate token")
	}

	refreshToken, ok := tokenData["refresh_token"].(string)
	if !ok {
		log.Error("failed to assert refresh_token to string")
		return nil, errors.New("internal server error: failed to generate token")
	}

	resUser.AccessToken = accessToken
	resUser.RefreshToken = refreshToken
	resUser.UserType = userType
	resUser.Name = userData.Name

	return &resUser, nil
}

func (svc *service) Modify(userData dtos.InputUser, userID uint64) error {
	newUser := user.User{}

	err := smapping.FillStruct(&newUser, smapping.MapFields(userData))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	newUser.ID = userID
	err = svc.model.Update(newUser)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Remove(userID uint64) error {
	err := svc.model.DeleteByID(userID)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
