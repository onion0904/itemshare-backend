package group

import (
    "context"
    "encoding/base64"
    "fmt"

    groupDomain "github.com/onion0904/CarShareSystem/app/domain/group"
    "github.com/skip2/go-qrcode" // QRコード生成ライブラリ
)

type GroupInviteService struct {
    groupRepo groupDomain.GroupRepository // グループのリポジトリ
    baseURL   string          // 招待リンクのベースURL
}

func NewGroupInviteService(
	groupRepo groupDomain.GroupRepository, baseURL string,
) *GroupInviteService {
    return &GroupInviteService{
        groupRepo: groupRepo,
        baseURL:   baseURL,
    }
}

// Lineで招待
func (s *GroupInviteService) InviteByLine(ctx context.Context, groupID, userID string) error {
    group, err := s.groupRepo.FindGroup(ctx, groupID)
    if err != nil {
        return fmt.Errorf("グループが見つかりません: %w", err)
    }

    // Lineでの招待リンクを生成（仮）
    inviteLink := fmt.Sprintf("%s/groups/%s/invite", s.baseURL, group.ID())
    // LineのAPIを呼び出すロジックをここに記述
    fmt.Printf("Lineで招待リンクを送信: %s\n", inviteLink)

    return nil
}

// QRコードを生成
func (s *GroupInviteService) GenerateQRCode(ctx context.Context, groupID string) (string, error) {
    group, err := s.groupRepo.FindGroup(ctx, groupID)
    if err != nil {
        return "", fmt.Errorf("グループが見つかりません: %w", err)
    }

    // 招待リンク生成
    inviteLink := fmt.Sprintf("%s/groups/%s/invite", s.baseURL, group.ID())
    
    // QRコードを生成
    pngData, err := qrcode.Encode(inviteLink, qrcode.Medium, 256)
    if err != nil {
        return "", fmt.Errorf("QRコード生成エラー: %w", err)
    }

    // Base64エンコードしたデータを返す
    encoded := base64.StdEncoding.EncodeToString(pngData)
    return fmt.Sprintf("data:image/png;base64,%s", encoded), nil
}

// 招待リンクを生成
func (s *GroupInviteService) GenerateInviteLink(ctx context.Context, groupID string) (string, error) {
    group, err := s.groupRepo.FindGroup(ctx, groupID)
    if err != nil {
        return "", fmt.Errorf("グループが見つかりません: %w", err)
    }

    // 招待リンク生成
    inviteLink := fmt.Sprintf("%s/groups/%s/invite", s.baseURL, group.ID())
    return inviteLink, nil
}
