package swipe

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"datink/internal/models/auth"
	"datink/internal/models/swipe/dto"
	"datink/internal/models/swipe/entity"
	"datink/pkg/util/httputil"
)

func (s service) Swiping(ctx context.Context, user *auth.UserJwtAuth, req dto.SwipingRequest) (dto.SwipingResponse, error) {
	//usr, err := s.repository.User.GetByID(ctx, s.resource.Database.GetMaster(), user.UserID)
	//if err != nil {
	//	return dto.SwipingResponse{}, err
	//}

	//if !usr.IsPremium {
	//	//check limit
	//	totalSwiped, err := s.repository.Swipe.GetSwipeCountNow(ctx, s.resource.Database.GetMaster(), user.UserID)
	//	if err != nil {
	//		return dto.SwipingResponse{}, err
	//	}
	//	if s.resource.Config.SwipeLimitFree < totalSwiped+1 {
	//		return dto.SwipingResponse{}, httputil.ErrorWrap(httputil.ErrBadRequest, errors.New("swipe limit exceeded, please upgrade packages"))
	//	}
	//}

	activePackage, err := s.repository.UserPackage.GetActivePackage(ctx, s.resource.Database.GetMaster(), user.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			//check limit
			totalSwiped, err := s.repository.Swipe.GetSwipeCountNow(ctx, s.resource.Database.GetMaster(), user.UserID)
			if err != nil {
				return dto.SwipingResponse{}, err
			}
			if s.resource.Config.SwipeLimitFree < totalSwiped+1 {
				return dto.SwipingResponse{}, httputil.ErrorWrap(httputil.ErrBadRequest, errors.New("swipe limit exceeded, please upgrade packages"))
			}
		}

		return dto.SwipingResponse{}, err
	}

	log.Println(activePackage)

	var swapID int64
	swipe, errGetSwap := s.repository.Swipe.GetByUserIdOrFriendID(ctx, s.resource.Database.GetMaster(), user.UserID, req.FriendID)

	tx, err := s.resource.Database.GetMaster().BeginTx(ctx, nil)
	if err != nil {
		return dto.SwipingResponse{}, err
	}

	if errGetSwap != nil {
		if errors.Is(errGetSwap, sql.ErrNoRows) {
			newSwap := entity.Swipe{
				UserID:         user.UserID,
				FriendID:       req.FriendID,
				SwipeDirection: req.SwipeDirection,
			}
			if err := s.repository.Swipe.Create(ctx, tx, &newSwap); err != nil {
				if err := tx.Rollback(); err != nil {
					s.resource.Logger.Error("error rollback", zap.Error(err))
				}
				return dto.SwipingResponse{}, err
			}

			swapID = newSwap.ID
		} else {
			return dto.SwipingResponse{}, err
		}
	}

	//update
	if swipe.SwipeDirection == entity.SwipeDirectionLike {
		if err := s.repository.Swipe.Update(ctx, s.resource.Database.GetMaster(), &entity.Swipe{
			ID:             swipe.ID,
			SwipeDirection: entity.SwipeDirectionMatch,
		}); err != nil {
			if err := tx.Rollback(); err != nil {
				s.resource.Logger.Error("error rollback", zap.Error(err))
			}
			return dto.SwipingResponse{}, err
		}
	} else if swipe.SwipeDirection == entity.SwipeDirectionDisLike {
		if err := s.repository.Swipe.Update(ctx, s.resource.Database.GetMaster(), &entity.Swipe{
			ID:             swipe.ID,
			SwipeDirection: entity.SwipeDirectionLike,
		}); err != nil {
			if err := tx.Rollback(); err != nil {
				s.resource.Logger.Error("error rollback", zap.Error(err))
			}
			return dto.SwipingResponse{}, err
		}
	}

	swapID = swipe.ID
	err = s.repository.Swipe.IncSwipeCount(ctx, tx, &entity.SwipeCount{
		UserID:    user.UserID,
		SwipeDate: time.Now(),
		Total:     1,
	})
	if err != nil {
		if err := tx.Rollback(); err != nil {
			s.resource.Logger.Error("error rollback", zap.Error(err))
		}
		return dto.SwipingResponse{}, err
	}

	if err := tx.Commit(); err != nil {
		return dto.SwipingResponse{}, err
	}
	//if err != nil {
	//	s.resource.Logger.Error("error commit",er)
	//}

	return dto.SwipingResponse{
		ID: swapID,
	}, nil

}
