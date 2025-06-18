package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	errDomain "github.com/onion0904/CarShareSystem/app/domain/error"
	"github.com/onion0904/CarShareSystem/app/domain/user"
	"github.com/onion0904/CarShareSystem/app/infrastructure/db"
	dbgen "github.com/onion0904/CarShareSystem/app/infrastructure/db/sqlc/dbgen"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) user.UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) Save(ctx context.Context, user *user.User) error {
	query := db.GetQuery(ctx)
	if err := query.UpsertUser(ctx, dbgen.UpsertUserParams{
		ID:        user.ID(),
		LastName:  user.LastName(),
		FirstName: user.FirstName(),
		Email:     user.Email(),
		Password:  user.Password(),
		Icon:      user.Icon(),
	}); err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) Update(ctx context.Context, user *user.User) error {
	query := db.GetQuery(ctx)
	if err := query.UpsertUser(ctx, dbgen.UpsertUserParams{
		ID:        user.ID(),
		LastName:  user.LastName(),
		FirstName: user.FirstName(),
		Email:     user.Email(),
		Password:  user.Password(),
		Icon:      user.Icon(),
	}); err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) FindUser(ctx context.Context, UserID string) (*user.User, error) {
	DB := db.GetDB()                // DB インスタンス取得
	tx, err := DB.BeginTx(ctx, nil) // トランザクション開始
	if err != nil {
		return nil, err
	}
	defer func(){// エラー時のロールバック保証
		if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
			log.Printf("rollback failed: %v", err)
		}
	}()

	query := db.GetQuery(ctx).WithTx(tx) // query変数にトランザクション適用

	// ユーザー情報の取得
	u, err := query.FindUserByID(ctx, UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errDomain.NewError("User not found")
		}
		return nil, err
	}

	// ユーザーの所属グループID取得
	groupIDs, err := query.GetGroupIDsByUserID(ctx, UserID)
	if err != nil {
		return nil, err
	}

	// ユーザーの関連イベントID取得
	eventIDs, err := query.GetEventIDsByUserID(ctx, UserID)
	if err != nil {
		return nil, err
	}

	// ユーザーをドメインモデルとして再構築
	nu, err := user.Reconstruct(
		u.ID,
		u.LastName,
		u.FirstName,
		u.Email,
		u.Password,
		u.Icon,
		groupIDs,
		eventIDs,
	)
	if err != nil {
		return nil, err
	}
	nu.SetCreatedAt(u.CreatedAt)
	nu.SetUpdatedAt(u.UpdatedAt)

	// コミット
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return nu, nil
}

func (ur *userRepository) FindUserByEmail(ctx context.Context, email string) (*user.User, error) {
	DB := db.GetDB()                // DB インスタンス取得
	tx, err := DB.BeginTx(ctx, nil) // トランザクション開始
	if err != nil {
		return nil, err
	}
	defer func(){// エラー時のロールバック保証
		if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
			log.Printf("rollback failed: %v", err)
		}
	}()

	query := db.GetQuery(ctx).WithTx(tx)

	u, err := query.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// ユーザーの所属グループID取得
	groupIDs, err := query.GetGroupIDsByUserID(ctx, u.ID)
	if err != nil {
		return nil, err
	}

	// ユーザーの関連イベントID取得
	eventIDs, err := query.GetEventIDsByUserID(ctx, u.ID)
	if err != nil {
		return nil, err
	}

	// ユーザーをドメインモデルとして再構築
	nu, err := user.Reconstruct(
		u.ID,
		u.LastName,
		u.FirstName,
		u.Email,
		u.Password,
		u.Icon,
		groupIDs,
		eventIDs,
	)
	if err != nil {
		return nil, err
	}
	nu.SetCreatedAt(u.CreatedAt)
	nu.SetUpdatedAt(u.UpdatedAt)

	// コミット
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return nu, nil
}

func (ur *userRepository) Delete(ctx context.Context, UserID string) error {
	query := db.GetQuery(ctx)

	if err := query.DeleteUser(ctx, UserID); err != nil {
		return err
	}
	return nil
}
