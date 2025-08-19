package entity

// UserBase holds common fields for all user types.
type UserBase struct {
	ID              string // 一意ID
	Name            string // 表示名
	ProfileImageURL string // プロフィール画像URL
	Bio             string // 自己紹介や方針
}

// User represents a general user in the AI Coaching Service.
type User struct {
	UserBase
	// User固有の属性を追加可能
}

// Coach represents a coach in the AI Coaching Service.
type Coach struct {
	UserBase
	Specialty    string // コーチの専門分野
	Achievements string // 実績や資格
	// Coach固有の属性を追加可能
}
