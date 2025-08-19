package repository

import (
	"api/domain/entity"
	"api/domain/repository"
)

type ChatRepositoryImpl struct {
	// DB接続や依存リソースをここに持つ（例: db *gorm.DB）
	
}

func (r *ChatRepositoryImpl) CreateChat(chat *entity.Chat) (*entity.Chat, error) {
	// 実装例: DBにChatを保存
	return chat, nil
}

func (r *ChatRepositoryImpl) FindChatByID(chatId string) (*entity.Chat, error) {
	// 実装例: DBからChatを取得
	return &entity.Chat{}, nil
}

func (r *ChatRepositoryImpl) UpdateChat(chat *entity.Chat) (*entity.Chat, error) {
	// 実装例: DBのChatを更新
	return chat, nil
}

// domain/repository.ChatRepositoryインターフェースを満たす
var _ repository.ChatRepositoryInterface = (*ChatRepositoryImpl)(nil)
