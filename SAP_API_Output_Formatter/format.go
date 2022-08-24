package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-social-media-activity-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToSocialMediaActivityCollection(raw []byte, l *logger.Logger) ([]SocialMediaActivityCollection, error) {
	pm := &responses.SocialMediaActivityCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SocialMediaActivityCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	socialMediaActivityCollection := make([]SocialMediaActivityCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		socialMediaActivityCollection = append(socialMediaActivityCollection, SocialMediaActivityCollection{
			ObjectID:                           data.ObjectID,
			SocialMediaMessageAuthor:           data.SocialMediaMessageAuthor,
			CategoryCode:                       data.CategoryCode,
			PrivateSocialMediaMessageIndicator: data.PrivateSocialMediaMessageIndicator,
			SocialMediaMessageDomain:           data.SocialMediaMessageDomain,
			ID:                                 data.ID,
			RootSocialMediaActivityUUID:        data.RootSocialMediaActivityUUID,
			SocialMediaMessageURI:              data.SocialMediaMessageURI,
			ParentSocialMediaActivityUUID:      data.ParentSocialMediaActivityUUID,
			SocialMediaActivityProviderUUID:    data.SocialMediaActivityProviderUUID,
			SocialMediaUserProfileUUID:         data.SocialMediaUserProfileUUID,
			SocialFeedbackLikeUnlike:           data.SocialFeedbackLikeUnlike,
			SocialMediaChannelCode:             data.SocialMediaChannelCode,
			SocialMediaMessageID:               data.SocialMediaMessageID,
			SocialMediaMessageCreationDateTime: data.SocialMediaMessageCreationDateTime,
			UUID:                               data.UUID,
			EntityLastChangedOn:                data.EntityLastChangedOn,
			ETag:                               data.ETag,
			Text:                               data.Text,
			LanguageCode:                       data.LanguageCode,
			PriorityCode:                       data.PriorityCode,
			InitiatorCode:                      data.InitiatorCode,
			SocialMediaActivityProviderName:    data.SocialMediaActivityProviderName,
			SocialMediaActivityProviderID:      data.SocialMediaActivityProviderID,
		})
	}

	return socialMediaActivityCollection, nil
}
