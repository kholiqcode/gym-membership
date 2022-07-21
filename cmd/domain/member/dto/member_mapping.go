package dto

import (
	"gym/cmd/domain/member/entity"
	"gym/pkg/auth/dto"
)

func CreateMemberResponse(member *entity.Member) MemberResponse {
	return MemberResponse{
		ID:         member.ID,
		Name:       member.Name,
		Phone:      member.Phone,
		Email:      member.Email,
		OrderCount: uint(len(member.MemberOrder)),
		CreatedAt:  member.CreatedAt,
		UpdatedAt:  member.UpdatedAt,
	}
}

func CreateMemberListResponse(members *entity.MemberList) MemberListResponse {
	memberResp := MemberListResponse{}
	for _, p := range *members {
		member := CreateMemberResponse(p)
		memberResp = append(memberResp, &member)
	}
	return memberResp
}

func CreateMemberTypeResponse(memberType *entity.MemberType) MemberTypeResponse {
	return MemberTypeResponse{
		ID:          memberType.ID,
		Name:        memberType.Name,
		Description: memberType.Description,
		Image:       memberType.Image,
		Duration:    memberType.Duration,
		Price:       memberType.Price,
		CreatedAt:   memberType.CreatedAt,
		UpdatedAt:   memberType.UpdatedAt,
	}
}

func CreateMemberJoinResponse(memberJoin *entity.MemberJoin) MemberJoinResponse {
	return MemberJoinResponse{
		ID:                 memberJoin.ID,
		InvoiceNo:          memberJoin.InvoiceNo,
		StartAt:            memberJoin.StartAt,
		MemberName:         memberJoin.MemberName,
		MemberNik:          memberJoin.MemberNik,
		MemberPhone:        memberJoin.MemberPhone,
		MemberGender:       memberJoin.MemberGender,
		MemberAddress:      memberJoin.MemberAddress,
		MemberCity:         memberJoin.MemberCity,
		MemberPostalCode:   memberJoin.MemberPostalCode,
		MemberTypeName:     memberJoin.MemberTypeName,
		MemberTypeDuration: memberJoin.MemberTypeDuration,
		MemberTypePrice:    memberJoin.MemberTypePrice,
		MemberTypeImage:    memberJoin.MemberTypeImage,
		Status:             memberJoin.Status,
		Total:              memberJoin.Total,
	}
}

func CreateMemberJoinListResponse(memberOrders *entity.MemberJoinList) MemberOrderListResponse {
	memberOrderResp := MemberOrderListResponse{}
	for _, p := range *memberOrders {
		memberOrder := CreateMemberJoinResponse(p)
		memberOrderResp = append(memberOrderResp, &memberOrder)
	}
	return memberOrderResp
}

func CreateMemberAuthResponse(token dto.AccessToken, member *entity.Member) MemberAuthResponse {
	return MemberAuthResponse{
		Type:         token.Type,
		Token:        token.Token,
		RefreshToken: token.RefreshToken,
		Member: MemberResponse{
			ID:    member.ID,
			Name:  member.Name,
			Phone: member.Phone,
			Email: member.Email,
		},
	}
}
