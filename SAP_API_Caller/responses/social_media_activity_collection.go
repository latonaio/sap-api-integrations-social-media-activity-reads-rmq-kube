package responses

type SocialMediaActivityCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                             string `json:"ObjectID"`
			SocialMediaMessageAuthor             string `json:"SocialMediaMessageAuthor"`
			CategoryCode                         string `json:"CategoryCode"`
			PrivateSocialMediaMessageIndicator   bool   `json:"PrivateSocialMediaMessageIndicator"`
			SocialMediaMessageDomain             string `json:"SocialMediaMessageDomain"`
			ID                                   string `json:"ID"`
			RootSocialMediaActivityUUID          string `json:"RootSocialMediaActivityUUID"`
			SocialMediaMessageURI                string `json:"SocialMediaMessageURI"`
			ParentSocialMediaActivityUUID        string `json:"ParentSocialMediaActivityUUID"`
			SocialMediaActivityProviderUUID      string `json:"SocialMediaActivityProviderUUID"`
			SocialMediaUserProfileUUID           string `json:"SocialMediaUserProfileUUID"`
			SocialFeedbackLikeUnlike             bool   `json:"SocialFeedbackLikeUnlike"`
			SocialMediaChannelCode               string `json:"SocialMediaChannelCode"`
			SocialMediaMessageID                 string `json:"SocialMediaMessageID"`
			SocialMediaMessageCreationDateTime   string `json:"SocialMediaMessageCreationDateTime"`
			UUID                                 string `json:"UUID"`
			EntityLastChangedOn                  string `json:"EntityLastChangedOn"`
			ETag                                 string `json:"ETag"`
			Text                                 string `json:"Text"`
			LanguageCode                         string `json:"languageCode"`
			PriorityCode                         string `json:"PriorityCode"`
			InitiatorCode                        string `json:"InitiatorCode"`
			SocialMediaActivityProviderName      string `json:"SocialMediaActivityProviderName"`
			SocialMediaActivityProviderID        string `json:"SocialMediaActivityProviderID"`
			SocialMediaActivityAccessControlList struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SocialMediaActivityAccessControlList"`
			SocialMediaActivityAttachmentFolder struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SocialMediaActivityAttachmentFolder"`
			SocialMediaActivityToChildSocialMediaActivity struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SocialMediaActivityToChildSocialMediaActivity"`
		} `json:"results"`
	} `json:"d"`
}
