package main

type State string

const (
	ACTIVE        State = "ACTIVE"
	BANNED        State = "BANNED"
	VOTE_BANNED   State = "VOTE_BANNED"
	FOLLOW_BANNED State = "FOLLOW_BANNED"
	PRIVATE       State = "PRIVATE"
)

type PartnerLinks struct {
	ID          string   `json:"id"`
	PartnerName string   `json:"partnerName"`
	Partner     Partners `json:"partner"`
	Name        string   `json:"name"`
	Emoji       string   `json:"emoji"`
	Link        string   `json:"link"`
}

type Partners struct {
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	Logo            string        `json:"logo"`
	Category        string        `json:"category"`
	Owner           string        `json:"owner"`
	OwnerImage      string        `json:"ownerImage"`
	OwnerLink       *string       `json:"ownerLink,omitempty"`
	Description     string        `json:"description"`
	LongDescription string        `json:"long_description"`
	Links           []PartnerLinks `json:"links"`
}

type Applications struct {
	CreatorID   string   `json:"creatorid"`
	Owner       Users    `json:"owner"`
	Name        string   `json:"name"`
	Logo        string   `json:"logo"`
	Token       string   `json:"token"`
	Active      bool     `json:"active"`
	Permissions []string `json:"permissions"`
}

type Plugins struct {
	ID       int         `json:"id"`
	PostID   string      `json:"postid"`
	Post     Posts       `json:"post"`
	Type     string      `json:"type"`
	Href     *string     `json:"href,omitempty"`
	JsonData interface{} `json:"jsonData,omitempty"`
}

type FcmKeys struct {
	ID     string `json:"id"`
	UserID string `json:"userid"`
	User   Users  `json:"user"`
	Name   string `json:"name"`
	Key    string `json:"key"`
}

type Comments struct {
	CreatorID string `json:"creatorid"`
	User      Users  `json:"user"`
	Caption   string `json:"caption"`
	Image     *string `json:"image,omitempty"`
	Post      Posts   `json:"post"`
	PostID    string  `json:"postid"`
	CommentID string  `json:"commentid"`
}

type Upvotes struct {
	ID     string `json:"id"`
	UserID string `json:"userid"`
	PostID string `json:"postid"`
	Post   Posts  `json:"post"`
}

type Downvotes struct {
	ID     string `json:"id"`
	UserID string `json:"userid"`
	PostID string `json:"postid"`
	Post   Posts  `json:"post"`
}

type Following struct {
	ID       string `json:"id"`
	UserID   string `json:"userid"`
	User     Users  `json:"user"`
	TargetID string `json:"targetid"`
	Target   Users  `json:"target"`
}

type Posts struct {
	UserID    string     `json:"userid"`
	User      Users      `json:"user"`
	Caption   string     `json:"caption"`
	Image     *string    `json:"image,omitempty"`
	Plugins   []Plugins  `json:"plugins"`
	Type      int        `json:"type"`
	PostID    string     `json:"postid"`
	Upvotes   []Upvotes  `json:"upvotes"`
	Downvotes []Downvotes `json:"downvotes"`
	Comments  []Comments  `json:"comments"`
}

type Users struct {
	Name         *string     `json:"name,omitempty"`
	UserID       string      `json:"userid"`
	UserTag      string      `json:"usertag"`
	Bio          string      `json:"bio"`
	Avatar       string      `json:"avatar"`
	Followers    []Following `json:"followers"`
	Following    []Following `json:"following"`
	Badges       []string    `json:"badges"`
	State        State       `json:"state"`
	StaffPerms   []string    `json:"staff_perms"`
	Applications []Applications `json:"applications"`
	Posts        []Posts       `json:"posts"`
	Comments     []Comments    `json:"comments"`
	FcmKeys      []FcmKeys     `json:"fcm_keys"`
}