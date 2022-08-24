package sap_api_output_formatter

type SocialMediaActivity struct {
	ConnectionKey           string `json:"connection_key"`
	Result                  bool   `json:"result"`
	RedisKey                string `json:"redis_key"`
	Filepath                string `json:"filepath"`
	APISchema               string `json:"api_schema"`
	SocialMediaActivityCode string `json:"social_media_activity_code"`
	Deleted                 bool   `json:"deleted"`
}

type SocialMediaActivityCollection struct {
	ObjectID                           string `json:"ObjectID"`
	SocialMediaMessageAuthor           string `json:"SocialMediaMessageAuthor"`
	CategoryCode                       string `json:"CategoryCode"`
	PrivateSocialMediaMessageIndicator bool   `json:"PrivateSocialMediaMessageIndicator"`
	SocialMediaMessageDomain           string `json:"SocialMediaMessageDomain"`
	ID                                 string `json:"ID"`
	RootSocialMediaActivityUUID        string `json:"RootSocialMediaActivityUUID"`
	SocialMediaMessageURI              string `json:"SocialMediaMessageURI"`
	ParentSocialMediaActivityUUID      string `json:"ParentSocialMediaActivityUUID"`
	SocialMediaActivityProviderUUID    string `json:"SocialMediaActivityProviderUUID"`
	SocialMediaUserProfileUUID         string `json:"SocialMediaUserProfileUUID"`
	SocialFeedbackLikeUnlike           bool   `json:"SocialFeedbackLikeUnlike"`
	SocialMediaChannelCode             string `json:"SocialMediaChannelCode"`
	SocialMediaMessageID               string `json:"SocialMediaMessageID"`
	SocialMediaMessageCreationDateTime string `json:"SocialMediaMessageCreationDateTime"`
	UUID                               string `json:"UUID"`
	EntityLastChangedOn                string `json:"EntityLastChangedOn"`
	ETag                               string `json:"ETag"`
	Text                               string `json:"Text"`
	LanguageCode                       string `json:"languageCode"`
	PriorityCode                       string `json:"PriorityCode"`
	InitiatorCode                      string `json:"InitiatorCode"`
	SocialMediaActivityProviderName    string `json:"SocialMediaActivityProviderName"`
	SocialMediaActivityProviderID      string `json:"SocialMediaActivityProviderID"`
}
