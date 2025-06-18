package group

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/onion0904/CarShareSystem/app/config" // config をインポート
	groupDomain "github.com/onion0904/CarShareSystem/app/domain/group"
	"github.com/onion0904/CarShareSystem/pkg/jwt" // jwt をインポート
	"github.com/skip2/go-qrcode"                  // QRコード生成ライブラリ
)

type GroupInviteService struct {
	groupRepo groupDomain.GroupRepository // グループのリポジトリ
	baseURL   string                      // 招待リンクのベースURL
}

func NewGroupInviteService(
	groupRepo groupDomain.GroupRepository, baseURL string,
) *GroupInviteService {
	return &GroupInviteService{
		groupRepo: groupRepo,
		baseURL:   baseURL,
	}
}

// 招待リンクを生成
func (s *GroupInviteService) GenerateInviteLink(ctx context.Context, groupID string) (string, error) {
	// JWTの秘密鍵を取得
	cfg := config.GetConfig()
	jwtSecret := []byte(cfg.JWT.Secret)

	// 招待用トークンを生成
	claims := jwt.NewInviteClaims(groupID)
	token := jwt.CreateToken(claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("招待トークンの生成に失敗しました: %w", err)
	}

	// トークンを含む招待リンクを生成
	inviteLink := fmt.Sprintf("%s/invite?token=%s", s.baseURL, tokenString) // フロントエンドの招待承諾ページのURL
	return inviteLink, nil
}

// QRコードを生成
func (s *GroupInviteService) GenerateQRCode(ctx context.Context, groupID string) (string, error) {
	// 招待リンク生成
	inviteLink, err := s.GenerateInviteLink(ctx, groupID)
	if err != nil {
		return "", err
	}

	// QRコードを生成
	pngData, err := qrcode.Encode(inviteLink, qrcode.Medium, 256)
	if err != nil {
		return "", fmt.Errorf("QRコード生成エラー: %w", err)
	}

	// Base64エンコードしたデータを返す
	encoded := base64.StdEncoding.EncodeToString(pngData)
	return fmt.Sprintf("data:image/png;base64,%s", encoded), nil
}
