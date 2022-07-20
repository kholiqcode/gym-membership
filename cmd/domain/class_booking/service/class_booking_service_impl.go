package service

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
	"github.com/rs/zerolog/log"
	_rClass "gym/cmd/domain/class/repository"
	"gym/cmd/domain/class_booking/dto"
	"gym/cmd/domain/class_booking/entity"
	"gym/cmd/domain/class_booking/repository"
	_rMember "gym/cmd/domain/member/repository"
	"gym/pkg/nullstring"
	"strings"
	"time"
)

type ClassBookingServiceImpl struct {
	RepoClassBooking repository.ClassBookingRepository
	RepoClass        _rClass.ClassRepository
	RepoMember       _rMember.MemberRepository
}

func (s *ClassBookingServiceImpl) GetAll(ctx echo.Context) (*dto.ClassBookingListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ClassBookingServiceImpl) GetByInvoice(ctx echo.Context, invoice string) (*dto.ClassBookingResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ClassBookingServiceImpl) Create(ctx echo.Context, request *dto.ClassBookingCreateRequest) (*dto.ClassBookingResponse, error) {
	memberId := ctx.Get("user_id").(float64)

	member, err := s.RepoMember.Find(uint(memberId))
	if err != nil {
		log.Err(err).Msg("Error fetch member detail from DB")
		return nil, err
	}

	classes, err := s.RepoClass.FindByIds(ctx, request.ClassIds)

	if err != nil {
		log.Err(err).Msg("class is not found")
		return nil, err
	}

	var classBookingDetails []*entity.ClassBookingDetail
	var total float64

	for _, c := range *classes {
		total += c.Price
		t, _ := time.Parse("2006-01-02T15:04:05-0700", c.Date)
		classBookingDetails = append(classBookingDetails, &entity.ClassBookingDetail{
			ClassID:          c.ID,
			ClassName:        c.Name,
			ClassDescription: c.Description,
			ClassImage:       c.Image,
			ClassPrice:       c.Price,
			ClassCategory:    c.ClassCategory.ID,
			ClassDate:        t.Format("2006-01-02"),
			ClassStartTime:   c.StartTime,
			ClassEndTime:     c.EndTime,
			ClassTrainerName: c.TrainerName,
		})
	}

	invoiceNo := fmt.Sprintf("INVC-%s", strings.ToUpper(random.String(16)))
	classBooking, err := s.RepoClassBooking.Insert(ctx, &entity.ClassBooking{
		MemberID:           uint(memberId),
		InvoiceNo:          invoiceNo,
		MemberName:         member.Name,
		MemberPhone:        member.Phone,
		MemberEmail:        member.Email,
		PaymentMethod:      "BCA",
		Status:             1,
		Note:               nullstring.NewNullString(request.Note),
		Total:              total,
		ClassBookingDetail: classBookingDetails,
	})

	if err != nil {
		log.Err(err).Msg("Error insert member type to DB")
		return nil, err
	}

	classBookingResp := dto.CreateClassBookingResponse(classBooking)
	log.Info().Msg("Successfully insert to to DB")
	return &classBookingResp, nil
}
