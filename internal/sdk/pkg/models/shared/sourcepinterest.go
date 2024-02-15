// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/types"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type SourcePinterestAuthMethod string

const (
	SourcePinterestAuthMethodOauth20 SourcePinterestAuthMethod = "oauth2.0"
)

func (e SourcePinterestAuthMethod) ToPointer() *SourcePinterestAuthMethod {
	return &e
}

func (e *SourcePinterestAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourcePinterestAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePinterestAuthMethod: %v", v)
	}
}

type SourcePinterestOAuth20 struct {
	authMethod SourcePinterestAuthMethod `const:"oauth2.0" json:"auth_method"`
	// The Client ID of your OAuth application
	ClientID string `json:"client_id"`
	// The Client Secret of your OAuth application.
	ClientSecret string `json:"client_secret"`
	// Refresh Token to obtain new Access Token, when it's expired.
	RefreshToken string `json:"refresh_token"`
}

func (s SourcePinterestOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourcePinterestOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourcePinterestOAuth20) GetAuthMethod() SourcePinterestAuthMethod {
	return SourcePinterestAuthMethodOauth20
}

func (o *SourcePinterestOAuth20) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourcePinterestOAuth20) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourcePinterestOAuth20) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

// SourcePinterestValidEnums - An enumeration.
type SourcePinterestValidEnums string

const (
	SourcePinterestValidEnumsIndividual SourcePinterestValidEnums = "INDIVIDUAL"
	SourcePinterestValidEnumsHousehold  SourcePinterestValidEnums = "HOUSEHOLD"
)

func (e SourcePinterestValidEnums) ToPointer() *SourcePinterestValidEnums {
	return &e
}

func (e *SourcePinterestValidEnums) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "INDIVIDUAL":
		fallthrough
	case "HOUSEHOLD":
		*e = SourcePinterestValidEnums(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePinterestValidEnums: %v", v)
	}
}

// SourcePinterestClickWindowDays - Number of days to use as the conversion attribution window for a pin click action.
type SourcePinterestClickWindowDays int64

const (
	SourcePinterestClickWindowDaysZero     SourcePinterestClickWindowDays = 0
	SourcePinterestClickWindowDaysOne      SourcePinterestClickWindowDays = 1
	SourcePinterestClickWindowDaysSeven    SourcePinterestClickWindowDays = 7
	SourcePinterestClickWindowDaysFourteen SourcePinterestClickWindowDays = 14
	SourcePinterestClickWindowDaysThirty   SourcePinterestClickWindowDays = 30
	SourcePinterestClickWindowDaysSixty    SourcePinterestClickWindowDays = 60
)

func (e SourcePinterestClickWindowDays) ToPointer() *SourcePinterestClickWindowDays {
	return &e
}

func (e *SourcePinterestClickWindowDays) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 7:
		fallthrough
	case 14:
		fallthrough
	case 30:
		fallthrough
	case 60:
		*e = SourcePinterestClickWindowDays(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePinterestClickWindowDays: %v", v)
	}
}

// SourcePinterestSchemasValidEnums - An enumeration.
type SourcePinterestSchemasValidEnums string

const (
	SourcePinterestSchemasValidEnumsAdvertiserID                                 SourcePinterestSchemasValidEnums = "ADVERTISER_ID"
	SourcePinterestSchemasValidEnumsAdAccountID                                  SourcePinterestSchemasValidEnums = "AD_ACCOUNT_ID"
	SourcePinterestSchemasValidEnumsAdGroupEntityStatus                          SourcePinterestSchemasValidEnums = "AD_GROUP_ENTITY_STATUS"
	SourcePinterestSchemasValidEnumsAdGroupID                                    SourcePinterestSchemasValidEnums = "AD_GROUP_ID"
	SourcePinterestSchemasValidEnumsAdID                                         SourcePinterestSchemasValidEnums = "AD_ID"
	SourcePinterestSchemasValidEnumsCampaignDailySpendCap                        SourcePinterestSchemasValidEnums = "CAMPAIGN_DAILY_SPEND_CAP"
	SourcePinterestSchemasValidEnumsCampaignEntityStatus                         SourcePinterestSchemasValidEnums = "CAMPAIGN_ENTITY_STATUS"
	SourcePinterestSchemasValidEnumsCampaignID                                   SourcePinterestSchemasValidEnums = "CAMPAIGN_ID"
	SourcePinterestSchemasValidEnumsCampaignLifetimeSpendCap                     SourcePinterestSchemasValidEnums = "CAMPAIGN_LIFETIME_SPEND_CAP"
	SourcePinterestSchemasValidEnumsCampaignName                                 SourcePinterestSchemasValidEnums = "CAMPAIGN_NAME"
	SourcePinterestSchemasValidEnumsCheckoutRoas                                 SourcePinterestSchemasValidEnums = "CHECKOUT_ROAS"
	SourcePinterestSchemasValidEnumsClickthrough1                                SourcePinterestSchemasValidEnums = "CLICKTHROUGH_1"
	SourcePinterestSchemasValidEnumsClickthrough1Gross                           SourcePinterestSchemasValidEnums = "CLICKTHROUGH_1_GROSS"
	SourcePinterestSchemasValidEnumsClickthrough2                                SourcePinterestSchemasValidEnums = "CLICKTHROUGH_2"
	SourcePinterestSchemasValidEnumsCpcInMicroDollar                             SourcePinterestSchemasValidEnums = "CPC_IN_MICRO_DOLLAR"
	SourcePinterestSchemasValidEnumsCpmInDollar                                  SourcePinterestSchemasValidEnums = "CPM_IN_DOLLAR"
	SourcePinterestSchemasValidEnumsCpmInMicroDollar                             SourcePinterestSchemasValidEnums = "CPM_IN_MICRO_DOLLAR"
	SourcePinterestSchemasValidEnumsCtr                                          SourcePinterestSchemasValidEnums = "CTR"
	SourcePinterestSchemasValidEnumsCtr2                                         SourcePinterestSchemasValidEnums = "CTR_2"
	SourcePinterestSchemasValidEnumsEcpcvInDollar                                SourcePinterestSchemasValidEnums = "ECPCV_IN_DOLLAR"
	SourcePinterestSchemasValidEnumsEcpcvP95InDollar                             SourcePinterestSchemasValidEnums = "ECPCV_P95_IN_DOLLAR"
	SourcePinterestSchemasValidEnumsEcpcInDollar                                 SourcePinterestSchemasValidEnums = "ECPC_IN_DOLLAR"
	SourcePinterestSchemasValidEnumsEcpcInMicroDollar                            SourcePinterestSchemasValidEnums = "ECPC_IN_MICRO_DOLLAR"
	SourcePinterestSchemasValidEnumsEcpeInDollar                                 SourcePinterestSchemasValidEnums = "ECPE_IN_DOLLAR"
	SourcePinterestSchemasValidEnumsEcpmInMicroDollar                            SourcePinterestSchemasValidEnums = "ECPM_IN_MICRO_DOLLAR"
	SourcePinterestSchemasValidEnumsEcpvInDollar                                 SourcePinterestSchemasValidEnums = "ECPV_IN_DOLLAR"
	SourcePinterestSchemasValidEnumsEctr                                         SourcePinterestSchemasValidEnums = "ECTR"
	SourcePinterestSchemasValidEnumsEengagementRate                              SourcePinterestSchemasValidEnums = "EENGAGEMENT_RATE"
	SourcePinterestSchemasValidEnumsEngagement1                                  SourcePinterestSchemasValidEnums = "ENGAGEMENT_1"
	SourcePinterestSchemasValidEnumsEngagement2                                  SourcePinterestSchemasValidEnums = "ENGAGEMENT_2"
	SourcePinterestSchemasValidEnumsEngagementRate                               SourcePinterestSchemasValidEnums = "ENGAGEMENT_RATE"
	SourcePinterestSchemasValidEnumsIdeaPinProductTagVisit1                      SourcePinterestSchemasValidEnums = "IDEA_PIN_PRODUCT_TAG_VISIT_1"
	SourcePinterestSchemasValidEnumsIdeaPinProductTagVisit2                      SourcePinterestSchemasValidEnums = "IDEA_PIN_PRODUCT_TAG_VISIT_2"
	SourcePinterestSchemasValidEnumsImpression1                                  SourcePinterestSchemasValidEnums = "IMPRESSION_1"
	SourcePinterestSchemasValidEnumsImpression1Gross                             SourcePinterestSchemasValidEnums = "IMPRESSION_1_GROSS"
	SourcePinterestSchemasValidEnumsImpression2                                  SourcePinterestSchemasValidEnums = "IMPRESSION_2"
	SourcePinterestSchemasValidEnumsInappCheckoutCostPerAction                   SourcePinterestSchemasValidEnums = "INAPP_CHECKOUT_COST_PER_ACTION"
	SourcePinterestSchemasValidEnumsOutboundClick1                               SourcePinterestSchemasValidEnums = "OUTBOUND_CLICK_1"
	SourcePinterestSchemasValidEnumsOutboundClick2                               SourcePinterestSchemasValidEnums = "OUTBOUND_CLICK_2"
	SourcePinterestSchemasValidEnumsPageVisitCostPerAction                       SourcePinterestSchemasValidEnums = "PAGE_VISIT_COST_PER_ACTION"
	SourcePinterestSchemasValidEnumsPageVisitRoas                                SourcePinterestSchemasValidEnums = "PAGE_VISIT_ROAS"
	SourcePinterestSchemasValidEnumsPaidImpression                               SourcePinterestSchemasValidEnums = "PAID_IMPRESSION"
	SourcePinterestSchemasValidEnumsPinID                                        SourcePinterestSchemasValidEnums = "PIN_ID"
	SourcePinterestSchemasValidEnumsPinPromotionID                               SourcePinterestSchemasValidEnums = "PIN_PROMOTION_ID"
	SourcePinterestSchemasValidEnumsRepin1                                       SourcePinterestSchemasValidEnums = "REPIN_1"
	SourcePinterestSchemasValidEnumsRepin2                                       SourcePinterestSchemasValidEnums = "REPIN_2"
	SourcePinterestSchemasValidEnumsRepinRate                                    SourcePinterestSchemasValidEnums = "REPIN_RATE"
	SourcePinterestSchemasValidEnumsSpendInDollar                                SourcePinterestSchemasValidEnums = "SPEND_IN_DOLLAR"
	SourcePinterestSchemasValidEnumsSpendInMicroDollar                           SourcePinterestSchemasValidEnums = "SPEND_IN_MICRO_DOLLAR"
	SourcePinterestSchemasValidEnumsTotalCheckout                                SourcePinterestSchemasValidEnums = "TOTAL_CHECKOUT"
	SourcePinterestSchemasValidEnumsTotalCheckoutValueInMicroDollar              SourcePinterestSchemasValidEnums = "TOTAL_CHECKOUT_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestSchemasValidEnumsTotalClickthrough                            SourcePinterestSchemasValidEnums = "TOTAL_CLICKTHROUGH"
	SourcePinterestSchemasValidEnumsTotalClickAddToCart                          SourcePinterestSchemasValidEnums = "TOTAL_CLICK_ADD_TO_CART"
	SourcePinterestSchemasValidEnumsTotalClickCheckout                           SourcePinterestSchemasValidEnums = "TOTAL_CLICK_CHECKOUT"
	SourcePinterestSchemasValidEnumsTotalClickCheckoutValueInMicroDollar         SourcePinterestSchemasValidEnums = "TOTAL_CLICK_CHECKOUT_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestSchemasValidEnumsTotalClickLead                               SourcePinterestSchemasValidEnums = "TOTAL_CLICK_LEAD"
	SourcePinterestSchemasValidEnumsTotalClickSignup                             SourcePinterestSchemasValidEnums = "TOTAL_CLICK_SIGNUP"
	SourcePinterestSchemasValidEnumsTotalClickSignupValueInMicroDollar           SourcePinterestSchemasValidEnums = "TOTAL_CLICK_SIGNUP_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestSchemasValidEnumsTotalConversions                             SourcePinterestSchemasValidEnums = "TOTAL_CONVERSIONS"
	SourcePinterestSchemasValidEnumsTotalCustom                                  SourcePinterestSchemasValidEnums = "TOTAL_CUSTOM"
	SourcePinterestSchemasValidEnumsTotalEngagement                              SourcePinterestSchemasValidEnums = "TOTAL_ENGAGEMENT"
	SourcePinterestSchemasValidEnumsTotalEngagementCheckout                      SourcePinterestSchemasValidEnums = "TOTAL_ENGAGEMENT_CHECKOUT"
	SourcePinterestSchemasValidEnumsTotalEngagementCheckoutValueInMicroDollar    SourcePinterestSchemasValidEnums = "TOTAL_ENGAGEMENT_CHECKOUT_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestSchemasValidEnumsTotalEngagementLead                          SourcePinterestSchemasValidEnums = "TOTAL_ENGAGEMENT_LEAD"
	SourcePinterestSchemasValidEnumsTotalEngagementSignup                        SourcePinterestSchemasValidEnums = "TOTAL_ENGAGEMENT_SIGNUP"
	SourcePinterestSchemasValidEnumsTotalEngagementSignupValueInMicroDollar      SourcePinterestSchemasValidEnums = "TOTAL_ENGAGEMENT_SIGNUP_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestSchemasValidEnumsTotalIdeaPinProductTagVisit                  SourcePinterestSchemasValidEnums = "TOTAL_IDEA_PIN_PRODUCT_TAG_VISIT"
	SourcePinterestSchemasValidEnumsTotalImpressionFrequency                     SourcePinterestSchemasValidEnums = "TOTAL_IMPRESSION_FREQUENCY"
	SourcePinterestSchemasValidEnumsTotalImpressionUser                          SourcePinterestSchemasValidEnums = "TOTAL_IMPRESSION_USER"
	SourcePinterestSchemasValidEnumsTotalLead                                    SourcePinterestSchemasValidEnums = "TOTAL_LEAD"
	SourcePinterestSchemasValidEnumsTotalOfflineCheckout                         SourcePinterestSchemasValidEnums = "TOTAL_OFFLINE_CHECKOUT"
	SourcePinterestSchemasValidEnumsTotalPageVisit                               SourcePinterestSchemasValidEnums = "TOTAL_PAGE_VISIT"
	SourcePinterestSchemasValidEnumsTotalRepinRate                               SourcePinterestSchemasValidEnums = "TOTAL_REPIN_RATE"
	SourcePinterestSchemasValidEnumsTotalSignup                                  SourcePinterestSchemasValidEnums = "TOTAL_SIGNUP"
	SourcePinterestSchemasValidEnumsTotalSignupValueInMicroDollar                SourcePinterestSchemasValidEnums = "TOTAL_SIGNUP_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestSchemasValidEnumsTotalVideo3SecViews                          SourcePinterestSchemasValidEnums = "TOTAL_VIDEO_3SEC_VIEWS"
	SourcePinterestSchemasValidEnumsTotalVideoAvgWatchtimeInSecond               SourcePinterestSchemasValidEnums = "TOTAL_VIDEO_AVG_WATCHTIME_IN_SECOND"
	SourcePinterestSchemasValidEnumsTotalVideoMrcViews                           SourcePinterestSchemasValidEnums = "TOTAL_VIDEO_MRC_VIEWS"
	SourcePinterestSchemasValidEnumsTotalVideoP0Combined                         SourcePinterestSchemasValidEnums = "TOTAL_VIDEO_P0_COMBINED"
	SourcePinterestSchemasValidEnumsTotalVideoP100Complete                       SourcePinterestSchemasValidEnums = "TOTAL_VIDEO_P100_COMPLETE"
	SourcePinterestSchemasValidEnumsTotalVideoP25Combined                        SourcePinterestSchemasValidEnums = "TOTAL_VIDEO_P25_COMBINED"
	SourcePinterestSchemasValidEnumsTotalVideoP50Combined                        SourcePinterestSchemasValidEnums = "TOTAL_VIDEO_P50_COMBINED"
	SourcePinterestSchemasValidEnumsTotalVideoP75Combined                        SourcePinterestSchemasValidEnums = "TOTAL_VIDEO_P75_COMBINED"
	SourcePinterestSchemasValidEnumsTotalVideoP95Combined                        SourcePinterestSchemasValidEnums = "TOTAL_VIDEO_P95_COMBINED"
	SourcePinterestSchemasValidEnumsTotalViewAddToCart                           SourcePinterestSchemasValidEnums = "TOTAL_VIEW_ADD_TO_CART"
	SourcePinterestSchemasValidEnumsTotalViewCheckout                            SourcePinterestSchemasValidEnums = "TOTAL_VIEW_CHECKOUT"
	SourcePinterestSchemasValidEnumsTotalViewCheckoutValueInMicroDollar          SourcePinterestSchemasValidEnums = "TOTAL_VIEW_CHECKOUT_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestSchemasValidEnumsTotalViewLead                                SourcePinterestSchemasValidEnums = "TOTAL_VIEW_LEAD"
	SourcePinterestSchemasValidEnumsTotalViewSignup                              SourcePinterestSchemasValidEnums = "TOTAL_VIEW_SIGNUP"
	SourcePinterestSchemasValidEnumsTotalViewSignupValueInMicroDollar            SourcePinterestSchemasValidEnums = "TOTAL_VIEW_SIGNUP_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestSchemasValidEnumsTotalWebCheckout                             SourcePinterestSchemasValidEnums = "TOTAL_WEB_CHECKOUT"
	SourcePinterestSchemasValidEnumsTotalWebCheckoutValueInMicroDollar           SourcePinterestSchemasValidEnums = "TOTAL_WEB_CHECKOUT_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestSchemasValidEnumsTotalWebClickCheckout                        SourcePinterestSchemasValidEnums = "TOTAL_WEB_CLICK_CHECKOUT"
	SourcePinterestSchemasValidEnumsTotalWebClickCheckoutValueInMicroDollar      SourcePinterestSchemasValidEnums = "TOTAL_WEB_CLICK_CHECKOUT_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestSchemasValidEnumsTotalWebEngagementCheckout                   SourcePinterestSchemasValidEnums = "TOTAL_WEB_ENGAGEMENT_CHECKOUT"
	SourcePinterestSchemasValidEnumsTotalWebEngagementCheckoutValueInMicroDollar SourcePinterestSchemasValidEnums = "TOTAL_WEB_ENGAGEMENT_CHECKOUT_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestSchemasValidEnumsTotalWebSessions                             SourcePinterestSchemasValidEnums = "TOTAL_WEB_SESSIONS"
	SourcePinterestSchemasValidEnumsTotalWebViewCheckout                         SourcePinterestSchemasValidEnums = "TOTAL_WEB_VIEW_CHECKOUT"
	SourcePinterestSchemasValidEnumsTotalWebViewCheckoutValueInMicroDollar       SourcePinterestSchemasValidEnums = "TOTAL_WEB_VIEW_CHECKOUT_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestSchemasValidEnumsVideo3SecViews2                              SourcePinterestSchemasValidEnums = "VIDEO_3SEC_VIEWS_2"
	SourcePinterestSchemasValidEnumsVideoLength                                  SourcePinterestSchemasValidEnums = "VIDEO_LENGTH"
	SourcePinterestSchemasValidEnumsVideoMrcViews2                               SourcePinterestSchemasValidEnums = "VIDEO_MRC_VIEWS_2"
	SourcePinterestSchemasValidEnumsVideoP0Combined2                             SourcePinterestSchemasValidEnums = "VIDEO_P0_COMBINED_2"
	SourcePinterestSchemasValidEnumsVideoP100Complete2                           SourcePinterestSchemasValidEnums = "VIDEO_P100_COMPLETE_2"
	SourcePinterestSchemasValidEnumsVideoP25Combined2                            SourcePinterestSchemasValidEnums = "VIDEO_P25_COMBINED_2"
	SourcePinterestSchemasValidEnumsVideoP50Combined2                            SourcePinterestSchemasValidEnums = "VIDEO_P50_COMBINED_2"
	SourcePinterestSchemasValidEnumsVideoP75Combined2                            SourcePinterestSchemasValidEnums = "VIDEO_P75_COMBINED_2"
	SourcePinterestSchemasValidEnumsVideoP95Combined2                            SourcePinterestSchemasValidEnums = "VIDEO_P95_COMBINED_2"
	SourcePinterestSchemasValidEnumsWebCheckoutCostPerAction                     SourcePinterestSchemasValidEnums = "WEB_CHECKOUT_COST_PER_ACTION"
	SourcePinterestSchemasValidEnumsWebCheckoutRoas                              SourcePinterestSchemasValidEnums = "WEB_CHECKOUT_ROAS"
	SourcePinterestSchemasValidEnumsWebSessions1                                 SourcePinterestSchemasValidEnums = "WEB_SESSIONS_1"
	SourcePinterestSchemasValidEnumsWebSessions2                                 SourcePinterestSchemasValidEnums = "WEB_SESSIONS_2"
)

func (e SourcePinterestSchemasValidEnums) ToPointer() *SourcePinterestSchemasValidEnums {
	return &e
}

func (e *SourcePinterestSchemasValidEnums) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ADVERTISER_ID":
		fallthrough
	case "AD_ACCOUNT_ID":
		fallthrough
	case "AD_GROUP_ENTITY_STATUS":
		fallthrough
	case "AD_GROUP_ID":
		fallthrough
	case "AD_ID":
		fallthrough
	case "CAMPAIGN_DAILY_SPEND_CAP":
		fallthrough
	case "CAMPAIGN_ENTITY_STATUS":
		fallthrough
	case "CAMPAIGN_ID":
		fallthrough
	case "CAMPAIGN_LIFETIME_SPEND_CAP":
		fallthrough
	case "CAMPAIGN_NAME":
		fallthrough
	case "CHECKOUT_ROAS":
		fallthrough
	case "CLICKTHROUGH_1":
		fallthrough
	case "CLICKTHROUGH_1_GROSS":
		fallthrough
	case "CLICKTHROUGH_2":
		fallthrough
	case "CPC_IN_MICRO_DOLLAR":
		fallthrough
	case "CPM_IN_DOLLAR":
		fallthrough
	case "CPM_IN_MICRO_DOLLAR":
		fallthrough
	case "CTR":
		fallthrough
	case "CTR_2":
		fallthrough
	case "ECPCV_IN_DOLLAR":
		fallthrough
	case "ECPCV_P95_IN_DOLLAR":
		fallthrough
	case "ECPC_IN_DOLLAR":
		fallthrough
	case "ECPC_IN_MICRO_DOLLAR":
		fallthrough
	case "ECPE_IN_DOLLAR":
		fallthrough
	case "ECPM_IN_MICRO_DOLLAR":
		fallthrough
	case "ECPV_IN_DOLLAR":
		fallthrough
	case "ECTR":
		fallthrough
	case "EENGAGEMENT_RATE":
		fallthrough
	case "ENGAGEMENT_1":
		fallthrough
	case "ENGAGEMENT_2":
		fallthrough
	case "ENGAGEMENT_RATE":
		fallthrough
	case "IDEA_PIN_PRODUCT_TAG_VISIT_1":
		fallthrough
	case "IDEA_PIN_PRODUCT_TAG_VISIT_2":
		fallthrough
	case "IMPRESSION_1":
		fallthrough
	case "IMPRESSION_1_GROSS":
		fallthrough
	case "IMPRESSION_2":
		fallthrough
	case "INAPP_CHECKOUT_COST_PER_ACTION":
		fallthrough
	case "OUTBOUND_CLICK_1":
		fallthrough
	case "OUTBOUND_CLICK_2":
		fallthrough
	case "PAGE_VISIT_COST_PER_ACTION":
		fallthrough
	case "PAGE_VISIT_ROAS":
		fallthrough
	case "PAID_IMPRESSION":
		fallthrough
	case "PIN_ID":
		fallthrough
	case "PIN_PROMOTION_ID":
		fallthrough
	case "REPIN_1":
		fallthrough
	case "REPIN_2":
		fallthrough
	case "REPIN_RATE":
		fallthrough
	case "SPEND_IN_DOLLAR":
		fallthrough
	case "SPEND_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_CHECKOUT":
		fallthrough
	case "TOTAL_CHECKOUT_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_CLICKTHROUGH":
		fallthrough
	case "TOTAL_CLICK_ADD_TO_CART":
		fallthrough
	case "TOTAL_CLICK_CHECKOUT":
		fallthrough
	case "TOTAL_CLICK_CHECKOUT_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_CLICK_LEAD":
		fallthrough
	case "TOTAL_CLICK_SIGNUP":
		fallthrough
	case "TOTAL_CLICK_SIGNUP_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_CONVERSIONS":
		fallthrough
	case "TOTAL_CUSTOM":
		fallthrough
	case "TOTAL_ENGAGEMENT":
		fallthrough
	case "TOTAL_ENGAGEMENT_CHECKOUT":
		fallthrough
	case "TOTAL_ENGAGEMENT_CHECKOUT_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_ENGAGEMENT_LEAD":
		fallthrough
	case "TOTAL_ENGAGEMENT_SIGNUP":
		fallthrough
	case "TOTAL_ENGAGEMENT_SIGNUP_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_IDEA_PIN_PRODUCT_TAG_VISIT":
		fallthrough
	case "TOTAL_IMPRESSION_FREQUENCY":
		fallthrough
	case "TOTAL_IMPRESSION_USER":
		fallthrough
	case "TOTAL_LEAD":
		fallthrough
	case "TOTAL_OFFLINE_CHECKOUT":
		fallthrough
	case "TOTAL_PAGE_VISIT":
		fallthrough
	case "TOTAL_REPIN_RATE":
		fallthrough
	case "TOTAL_SIGNUP":
		fallthrough
	case "TOTAL_SIGNUP_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_VIDEO_3SEC_VIEWS":
		fallthrough
	case "TOTAL_VIDEO_AVG_WATCHTIME_IN_SECOND":
		fallthrough
	case "TOTAL_VIDEO_MRC_VIEWS":
		fallthrough
	case "TOTAL_VIDEO_P0_COMBINED":
		fallthrough
	case "TOTAL_VIDEO_P100_COMPLETE":
		fallthrough
	case "TOTAL_VIDEO_P25_COMBINED":
		fallthrough
	case "TOTAL_VIDEO_P50_COMBINED":
		fallthrough
	case "TOTAL_VIDEO_P75_COMBINED":
		fallthrough
	case "TOTAL_VIDEO_P95_COMBINED":
		fallthrough
	case "TOTAL_VIEW_ADD_TO_CART":
		fallthrough
	case "TOTAL_VIEW_CHECKOUT":
		fallthrough
	case "TOTAL_VIEW_CHECKOUT_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_VIEW_LEAD":
		fallthrough
	case "TOTAL_VIEW_SIGNUP":
		fallthrough
	case "TOTAL_VIEW_SIGNUP_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_WEB_CHECKOUT":
		fallthrough
	case "TOTAL_WEB_CHECKOUT_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_WEB_CLICK_CHECKOUT":
		fallthrough
	case "TOTAL_WEB_CLICK_CHECKOUT_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_WEB_ENGAGEMENT_CHECKOUT":
		fallthrough
	case "TOTAL_WEB_ENGAGEMENT_CHECKOUT_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_WEB_SESSIONS":
		fallthrough
	case "TOTAL_WEB_VIEW_CHECKOUT":
		fallthrough
	case "TOTAL_WEB_VIEW_CHECKOUT_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "VIDEO_3SEC_VIEWS_2":
		fallthrough
	case "VIDEO_LENGTH":
		fallthrough
	case "VIDEO_MRC_VIEWS_2":
		fallthrough
	case "VIDEO_P0_COMBINED_2":
		fallthrough
	case "VIDEO_P100_COMPLETE_2":
		fallthrough
	case "VIDEO_P25_COMBINED_2":
		fallthrough
	case "VIDEO_P50_COMBINED_2":
		fallthrough
	case "VIDEO_P75_COMBINED_2":
		fallthrough
	case "VIDEO_P95_COMBINED_2":
		fallthrough
	case "WEB_CHECKOUT_COST_PER_ACTION":
		fallthrough
	case "WEB_CHECKOUT_ROAS":
		fallthrough
	case "WEB_SESSIONS_1":
		fallthrough
	case "WEB_SESSIONS_2":
		*e = SourcePinterestSchemasValidEnums(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePinterestSchemasValidEnums: %v", v)
	}
}

// SourcePinterestConversionReportTime - The date by which the conversion metrics returned from this endpoint will be reported. There are two dates associated with a conversion event: the date that the user interacted with the ad, and the date that the user completed a conversion event..
type SourcePinterestConversionReportTime string

const (
	SourcePinterestConversionReportTimeTimeOfAdAction   SourcePinterestConversionReportTime = "TIME_OF_AD_ACTION"
	SourcePinterestConversionReportTimeTimeOfConversion SourcePinterestConversionReportTime = "TIME_OF_CONVERSION"
)

func (e SourcePinterestConversionReportTime) ToPointer() *SourcePinterestConversionReportTime {
	return &e
}

func (e *SourcePinterestConversionReportTime) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TIME_OF_AD_ACTION":
		fallthrough
	case "TIME_OF_CONVERSION":
		*e = SourcePinterestConversionReportTime(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePinterestConversionReportTime: %v", v)
	}
}

// SourcePinterestEngagementWindowDays - Number of days to use as the conversion attribution window for an engagement action.
type SourcePinterestEngagementWindowDays int64

const (
	SourcePinterestEngagementWindowDaysZero     SourcePinterestEngagementWindowDays = 0
	SourcePinterestEngagementWindowDaysOne      SourcePinterestEngagementWindowDays = 1
	SourcePinterestEngagementWindowDaysSeven    SourcePinterestEngagementWindowDays = 7
	SourcePinterestEngagementWindowDaysFourteen SourcePinterestEngagementWindowDays = 14
	SourcePinterestEngagementWindowDaysThirty   SourcePinterestEngagementWindowDays = 30
	SourcePinterestEngagementWindowDaysSixty    SourcePinterestEngagementWindowDays = 60
)

func (e SourcePinterestEngagementWindowDays) ToPointer() *SourcePinterestEngagementWindowDays {
	return &e
}

func (e *SourcePinterestEngagementWindowDays) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 7:
		fallthrough
	case 14:
		fallthrough
	case 30:
		fallthrough
	case 60:
		*e = SourcePinterestEngagementWindowDays(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePinterestEngagementWindowDays: %v", v)
	}
}

// SourcePinterestGranularity - Chosen granularity for API
type SourcePinterestGranularity string

const (
	SourcePinterestGranularityTotal SourcePinterestGranularity = "TOTAL"
	SourcePinterestGranularityDay   SourcePinterestGranularity = "DAY"
	SourcePinterestGranularityHour  SourcePinterestGranularity = "HOUR"
	SourcePinterestGranularityWeek  SourcePinterestGranularity = "WEEK"
	SourcePinterestGranularityMonth SourcePinterestGranularity = "MONTH"
)

func (e SourcePinterestGranularity) ToPointer() *SourcePinterestGranularity {
	return &e
}

func (e *SourcePinterestGranularity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TOTAL":
		fallthrough
	case "DAY":
		fallthrough
	case "HOUR":
		fallthrough
	case "WEEK":
		fallthrough
	case "MONTH":
		*e = SourcePinterestGranularity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePinterestGranularity: %v", v)
	}
}

// SourcePinterestLevel - Chosen level for API
type SourcePinterestLevel string

const (
	SourcePinterestLevelAdvertiser            SourcePinterestLevel = "ADVERTISER"
	SourcePinterestLevelAdvertiserTargeting   SourcePinterestLevel = "ADVERTISER_TARGETING"
	SourcePinterestLevelCampaign              SourcePinterestLevel = "CAMPAIGN"
	SourcePinterestLevelCampaignTargeting     SourcePinterestLevel = "CAMPAIGN_TARGETING"
	SourcePinterestLevelAdGroup               SourcePinterestLevel = "AD_GROUP"
	SourcePinterestLevelAdGroupTargeting      SourcePinterestLevel = "AD_GROUP_TARGETING"
	SourcePinterestLevelPinPromotion          SourcePinterestLevel = "PIN_PROMOTION"
	SourcePinterestLevelPinPromotionTargeting SourcePinterestLevel = "PIN_PROMOTION_TARGETING"
	SourcePinterestLevelKeyword               SourcePinterestLevel = "KEYWORD"
	SourcePinterestLevelProductGroup          SourcePinterestLevel = "PRODUCT_GROUP"
	SourcePinterestLevelProductGroupTargeting SourcePinterestLevel = "PRODUCT_GROUP_TARGETING"
	SourcePinterestLevelProductItem           SourcePinterestLevel = "PRODUCT_ITEM"
)

func (e SourcePinterestLevel) ToPointer() *SourcePinterestLevel {
	return &e
}

func (e *SourcePinterestLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ADVERTISER":
		fallthrough
	case "ADVERTISER_TARGETING":
		fallthrough
	case "CAMPAIGN":
		fallthrough
	case "CAMPAIGN_TARGETING":
		fallthrough
	case "AD_GROUP":
		fallthrough
	case "AD_GROUP_TARGETING":
		fallthrough
	case "PIN_PROMOTION":
		fallthrough
	case "PIN_PROMOTION_TARGETING":
		fallthrough
	case "KEYWORD":
		fallthrough
	case "PRODUCT_GROUP":
		fallthrough
	case "PRODUCT_GROUP_TARGETING":
		fallthrough
	case "PRODUCT_ITEM":
		*e = SourcePinterestLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePinterestLevel: %v", v)
	}
}

// SourcePinterestViewWindowDays - Number of days to use as the conversion attribution window for a view action.
type SourcePinterestViewWindowDays int64

const (
	SourcePinterestViewWindowDaysZero     SourcePinterestViewWindowDays = 0
	SourcePinterestViewWindowDaysOne      SourcePinterestViewWindowDays = 1
	SourcePinterestViewWindowDaysSeven    SourcePinterestViewWindowDays = 7
	SourcePinterestViewWindowDaysFourteen SourcePinterestViewWindowDays = 14
	SourcePinterestViewWindowDaysThirty   SourcePinterestViewWindowDays = 30
	SourcePinterestViewWindowDaysSixty    SourcePinterestViewWindowDays = 60
)

func (e SourcePinterestViewWindowDays) ToPointer() *SourcePinterestViewWindowDays {
	return &e
}

func (e *SourcePinterestViewWindowDays) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 7:
		fallthrough
	case 14:
		fallthrough
	case 30:
		fallthrough
	case 60:
		*e = SourcePinterestViewWindowDays(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePinterestViewWindowDays: %v", v)
	}
}

// SourcePinterestReportConfig - Config for custom report
type SourcePinterestReportConfig struct {
	// List of types of attribution for the conversion report
	AttributionTypes []SourcePinterestValidEnums `json:"attribution_types,omitempty"`
	// Number of days to use as the conversion attribution window for a pin click action.
	ClickWindowDays *SourcePinterestClickWindowDays `default:"30" json:"click_window_days"`
	// A list of chosen columns
	Columns []SourcePinterestSchemasValidEnums `json:"columns"`
	// The date by which the conversion metrics returned from this endpoint will be reported. There are two dates associated with a conversion event: the date that the user interacted with the ad, and the date that the user completed a conversion event..
	ConversionReportTime *SourcePinterestConversionReportTime `default:"TIME_OF_AD_ACTION" json:"conversion_report_time"`
	// Number of days to use as the conversion attribution window for an engagement action.
	EngagementWindowDays *SourcePinterestEngagementWindowDays `default:"30" json:"engagement_window_days"`
	// Chosen granularity for API
	Granularity *SourcePinterestGranularity `default:"TOTAL" json:"granularity"`
	// Chosen level for API
	Level *SourcePinterestLevel `default:"ADVERTISER" json:"level"`
	// The name value of report
	Name string `json:"name"`
	// A date in the format YYYY-MM-DD. If you have not set a date, it would be defaulted to latest allowed date by report api (913 days from today).
	StartDate *types.Date `json:"start_date,omitempty"`
	// Number of days to use as the conversion attribution window for a view action.
	ViewWindowDays *SourcePinterestViewWindowDays `default:"30" json:"view_window_days"`
}

func (s SourcePinterestReportConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourcePinterestReportConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourcePinterestReportConfig) GetAttributionTypes() []SourcePinterestValidEnums {
	if o == nil {
		return nil
	}
	return o.AttributionTypes
}

func (o *SourcePinterestReportConfig) GetClickWindowDays() *SourcePinterestClickWindowDays {
	if o == nil {
		return nil
	}
	return o.ClickWindowDays
}

func (o *SourcePinterestReportConfig) GetColumns() []SourcePinterestSchemasValidEnums {
	if o == nil {
		return []SourcePinterestSchemasValidEnums{}
	}
	return o.Columns
}

func (o *SourcePinterestReportConfig) GetConversionReportTime() *SourcePinterestConversionReportTime {
	if o == nil {
		return nil
	}
	return o.ConversionReportTime
}

func (o *SourcePinterestReportConfig) GetEngagementWindowDays() *SourcePinterestEngagementWindowDays {
	if o == nil {
		return nil
	}
	return o.EngagementWindowDays
}

func (o *SourcePinterestReportConfig) GetGranularity() *SourcePinterestGranularity {
	if o == nil {
		return nil
	}
	return o.Granularity
}

func (o *SourcePinterestReportConfig) GetLevel() *SourcePinterestLevel {
	if o == nil {
		return nil
	}
	return o.Level
}

func (o *SourcePinterestReportConfig) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourcePinterestReportConfig) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *SourcePinterestReportConfig) GetViewWindowDays() *SourcePinterestViewWindowDays {
	if o == nil {
		return nil
	}
	return o.ViewWindowDays
}

type Pinterest string

const (
	PinterestPinterest Pinterest = "pinterest"
)

func (e Pinterest) ToPointer() *Pinterest {
	return &e
}

func (e *Pinterest) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pinterest":
		*e = Pinterest(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Pinterest: %v", v)
	}
}

type SourcePinterestStatus string

const (
	SourcePinterestStatusActive   SourcePinterestStatus = "ACTIVE"
	SourcePinterestStatusPaused   SourcePinterestStatus = "PAUSED"
	SourcePinterestStatusArchived SourcePinterestStatus = "ARCHIVED"
)

func (e SourcePinterestStatus) ToPointer() *SourcePinterestStatus {
	return &e
}

func (e *SourcePinterestStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ACTIVE":
		fallthrough
	case "PAUSED":
		fallthrough
	case "ARCHIVED":
		*e = SourcePinterestStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePinterestStatus: %v", v)
	}
}

type SourcePinterest struct {
	Credentials *SourcePinterestOAuth20 `json:"credentials,omitempty"`
	// A list which contains ad statistics entries, each entry must have a name and can contains fields, breakdowns or action_breakdowns. Click on "add" to fill this field.
	CustomReports []SourcePinterestReportConfig `json:"custom_reports,omitempty"`
	sourceType    *Pinterest                    `const:"pinterest" json:"sourceType,omitempty"`
	// A date in the format YYYY-MM-DD. If you have not set a date, it would be defaulted to latest allowed date by api (89 days from today).
	StartDate *types.Date `json:"start_date,omitempty"`
	// For the ads, ad_groups, and campaigns streams, specifying a status will filter out records that do not match the specified ones. If a status is not specified, the source will default to records with a status of either ACTIVE or PAUSED.
	Status []SourcePinterestStatus `json:"status,omitempty"`
}

func (s SourcePinterest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourcePinterest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourcePinterest) GetCredentials() *SourcePinterestOAuth20 {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourcePinterest) GetCustomReports() []SourcePinterestReportConfig {
	if o == nil {
		return nil
	}
	return o.CustomReports
}

func (o *SourcePinterest) GetSourceType() *Pinterest {
	return PinterestPinterest.ToPointer()
}

func (o *SourcePinterest) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *SourcePinterest) GetStatus() []SourcePinterestStatus {
	if o == nil {
		return nil
	}
	return o.Status
}
