package dto

import (
	"gym/cmd/domain/admin/entity"
	"gym/pkg/auth/dto"
)

func CreateAdminResponse(admin *entity.Admin) AdminResponse {
	return AdminResponse{
		ID:        admin.ID,
		Name:      admin.Name,
		Phone:     admin.Phone,
		Email:     admin.Email,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}
}

func CreateAdminListResponse(admins *entity.AdminList) AdminListResponse {
	adminResp := AdminListResponse{}
	for _, p := range *admins {
		admin := CreateAdminResponse(p)
		adminResp = append(adminResp, &admin)
	}
	return adminResp
}

func CreateAdminAuthResponse(token dto.AccessToken, admin *entity.Admin) AdminAuthResponse {
	return AdminAuthResponse{
		Type:         token.Type,
		Token:        token.Token,
		RefreshToken: token.RefreshToken,
		Admin: AdminResponse{
			ID:        admin.ID,
			Name:      admin.Name,
			Phone:     admin.Phone,
			Email:     admin.Email,
			CreatedAt: admin.CreatedAt,
			UpdatedAt: admin.UpdatedAt,
		},
	}
}
