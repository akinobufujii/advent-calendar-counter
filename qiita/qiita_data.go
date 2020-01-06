package qiita

// User Qiitaのユーザー情報
type User struct {
	Description       *string `json:"description"`
	FacebookID        *string `json:"facebook_id"`
	FolloweesCount    int     `json:"followees_count"`
	FollowersCount    int     `json:"followers_count"`
	GithubLoginName   *string `json:"github_login_name"`
	ID                string  `json:"id"`
	ItemsCount        int     `json:"items_count"`
	LinkedinID        *string `json:"linkedin_id"`
	Location          *string `json:"location"`
	Name              *string `json:"name"`
	Organization      *string `json:"organization"`
	PermanentID       int     `json:"permanent_id"`
	ProfileImageURL   string  `json:"profile_image_url"`
	TeamOnly          bool    `json:"team_only"`
	TwitterScreenName *string `json:"twitter_screen_name"`
	WebsiteURL        *string `json:"website_url"`
}

// Group Qiitaのグループを
type Group struct {
	CreatedAt string `json:"created_at"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Private   bool   `json:"private"`
	UpdatedAt string `json:"updated_at"`
	URLName   string `json:"url_name"`
}

// Tag Qiitaの記事についているタグ情報
type Tag struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
}

// Item Qiitaの記事情報
type Item struct {
	RenderedBody   string `json:"rendered_body"`
	Body           string `json:"body"`
	Coediting      bool   `json:"coediting"`
	CommentsCount  int    `json:"comments_count"`
	CreatedAt      string `json:"created_at"`
	Group          Group  `json:"group"`
	ID             string `json:"id"`
	LikesCount     int    `json:"likes_count"`
	Private        bool   `json:"private"`
	ReactionsCount int    `json:"reactions_count"`
	Tags           []Tag  `json:"tags"`
	Title          string `json:"title"`
	UpdatedAt      string `json:"updated_at"`
	URL            string `json:"url"`
	User           User   `json:"user"`
	PageViewsCount int    `json:"page_views_count"`
}
