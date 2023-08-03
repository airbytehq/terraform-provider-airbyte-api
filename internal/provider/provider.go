// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"airbyte/internal/sdk/pkg/models/shared"
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &AirbyteProvider{}

type AirbyteProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// AirbyteProviderModel describes the provider data model.
type AirbyteProviderModel struct {
	ServerURL  types.String `tfsdk:"server_url"`
	Password   types.String `tfsdk:"password"`
	Username   types.String `tfsdk:"username"`
	BearerAuth types.String `tfsdk:"bearer_auth"`
}

func (p *AirbyteProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "airbyte"
	resp.Version = p.version
}

func (p *AirbyteProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: `airbyte-api: Programatically control Airbyte Cloud.`,
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Server URL (defaults to https://api.airbyte.com/v1)",
				Optional:            true,
				Required:            false,
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"username": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"bearer_auth": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *AirbyteProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data AirbyteProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ServerURL := data.ServerURL.ValueString()

	if ServerURL == "" {
		ServerURL = "https://api.airbyte.com/v1"
	}

	var basicAuth *shared.SchemeBasicAuth
	password := data.Password.ValueString()
	username := data.Username.ValueString()
	basicAuth = &shared.SchemeBasicAuth{
		Password: password,
		Username: username,
	}
	bearerAuth := new(string)
	if !data.BearerAuth.IsUnknown() && !data.BearerAuth.IsNull() {
		*bearerAuth = data.BearerAuth.ValueString()
	} else {
		bearerAuth = nil
	}
	security := shared.Security{
		BasicAuth:  basicAuth,
		BearerAuth: bearerAuth,
	}

	opts := []sdk.SDKOption{
		sdk.WithServerURL(ServerURL),
		sdk.WithSecurity(security),
	}
	client := sdk.New(opts...)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *AirbyteProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewConnectionResource,
		NewDestinationAmazonSqsResource,
		NewDestinationAwsDatalakeResource,
		NewDestinationAzureBlobStorageResource,
		NewDestinationBigqueryResource,
		NewDestinationBigqueryDenormalizedResource,
		NewDestinationCassandraResource,
		NewDestinationClickhouseResource,
		NewDestinationConvexResource,
		NewDestinationCumulioResource,
		NewDestinationDatabendResource,
		NewDestinationDatabricksResource,
		NewDestinationDevNullResource,
		NewDestinationDynamodbResource,
		NewDestinationElasticsearchResource,
		NewDestinationFireboltResource,
		NewDestinationFirestoreResource,
		NewDestinationGcsResource,
		NewDestinationGoogleSheetsResource,
		NewDestinationKeenResource,
		NewDestinationKinesisResource,
		NewDestinationMariadbColumnstoreResource,
		NewDestinationMeilisearchResource,
		NewDestinationMongodbResource,
		NewDestinationMssqlResource,
		NewDestinationMysqlResource,
		NewDestinationOracleResource,
		NewDestinationPostgresResource,
		NewDestinationPubsubResource,
		NewDestinationPulsarResource,
		NewDestinationRabbitmqResource,
		NewDestinationRedisResource,
		NewDestinationRedshiftResource,
		NewDestinationRocksetResource,
		NewDestinationS3Resource,
		NewDestinationS3GlueResource,
		NewDestinationScyllaResource,
		NewDestinationSftpJSONResource,
		NewDestinationSnowflakeResource,
		NewDestinationTypesenseResource,
		NewDestinationVerticaResource,
		NewSourceAircallResource,
		NewSourceAirtableResource,
		NewSourceAlloydbResource,
		NewSourceAmazonAdsResource,
		NewSourceAmazonSqsResource,
		NewSourceAmplitudeResource,
		NewSourceApifyDatasetResource,
		NewSourceAsanaResource,
		NewSourceAuth0Resource,
		NewSourceAwsCloudtrailResource,
		NewSourceAzureBlobStorageResource,
		NewSourceAzureTableResource,
		NewSourceBambooHrResource,
		NewSourceBigcommerceResource,
		NewSourceBigqueryResource,
		NewSourceBingAdsResource,
		NewSourceBraintreeResource,
		NewSourceBrazeResource,
		NewSourceChargebeeResource,
		NewSourceChartmogulResource,
		NewSourceClickhouseResource,
		NewSourceClickupAPIResource,
		NewSourceCloseComResource,
		NewSourceCodaResource,
		NewSourceCoinAPIResource,
		NewSourceCoinmarketcapResource,
		NewSourceConfigcatResource,
		NewSourceConfluenceResource,
		NewSourceDatascopeResource,
		NewSourceDelightedResource,
		NewSourceDixaResource,
		NewSourceDockerhubResource,
		NewSourceDremioResource,
		NewSourceDynamodbResource,
		NewSourceE2eTestCloudResource,
		NewSourceEmailoctopusResource,
		NewSourceExchangeRatesResource,
		NewSourceFacebookMarketingResource,
		NewSourceFacebookPagesResource,
		NewSourceFakerResource,
		NewSourceFaunaResource,
		NewSourceFileSecureResource,
		NewSourceFireboltResource,
		NewSourceFreshcallerResource,
		NewSourceFreshdeskResource,
		NewSourceFreshsalesResource,
		NewSourceFreshserviceResource,
		NewSourceGcsResource,
		NewSourceGetlagoResource,
		NewSourceGithubResource,
		NewSourceGitlabResource,
		NewSourceGlassfrogResource,
		NewSourceGnewsResource,
		NewSourceGoogleAdsResource,
		NewSourceGoogleAnalyticsDataAPIResource,
		NewSourceGoogleAnalyticsV4Resource,
		NewSourceGoogleDirectoryResource,
		NewSourceGooglePagespeedInsightsResource,
		NewSourceGoogleSearchConsoleResource,
		NewSourceGoogleSheetsResource,
		NewSourceGoogleWebfontsResource,
		NewSourceGoogleWorkspaceAdminReportsResource,
		NewSourceGreenhouseResource,
		NewSourceGridlyResource,
		NewSourceHarvestResource,
		NewSourceHubplannerResource,
		NewSourceHubspotResource,
		NewSourceInsightlyResource,
		NewSourceInstagramResource,
		NewSourceInstatusResource,
		NewSourceIntercomResource,
		NewSourceIp2whoisResource,
		NewSourceIterableResource,
		NewSourceJiraResource,
		NewSourceK6CloudResource,
		NewSourceKlarnaResource,
		NewSourceKlaviyoResource,
		NewSourceKustomerSingerResource,
		NewSourceLaunchdarklyResource,
		NewSourceLemlistResource,
		NewSourceLinkedinAdsResource,
		NewSourceLinkedinPagesResource,
		NewSourceLinnworksResource,
		NewSourceLokaliseResource,
		NewSourceMailchimpResource,
		NewSourceMailgunResource,
		NewSourceMailjetSmsResource,
		NewSourceMarketoResource,
		NewSourceMetabaseResource,
		NewSourceMicrosoftTeamsResource,
		NewSourceMixpanelResource,
		NewSourceMondayResource,
		NewSourceMongodbResource,
		NewSourceMssqlResource,
		NewSourceMyHoursResource,
		NewSourceMysqlResource,
		NewSourceNetsuiteResource,
		NewSourceNotionResource,
		NewSourceNytimesResource,
		NewSourceOktaResource,
		NewSourceOmnisendResource,
		NewSourceOnesignalResource,
		NewSourceOpenweatherResource,
		NewSourceOracleResource,
		NewSourceOrbResource,
		NewSourceOrbitResource,
		NewSourceOutreachResource,
		NewSourcePaypalTransactionResource,
		NewSourcePaystackResource,
		NewSourcePendoResource,
		NewSourcePersistiqResource,
		NewSourcePexelsAPIResource,
		NewSourcePinterestResource,
		NewSourcePipedriveResource,
		NewSourcePocketResource,
		NewSourcePokeapiResource,
		NewSourcePolygonStockAPIResource,
		NewSourcePostgresResource,
		NewSourcePosthogResource,
		NewSourcePostmarkappResource,
		NewSourcePrestashopResource,
		NewSourcePublicApisResource,
		NewSourcePunkAPIResource,
		NewSourcePypiResource,
		NewSourceQualarooResource,
		NewSourceQuickbooksResource,
		NewSourceRailzResource,
		NewSourceRechargeResource,
		NewSourceRecreationResource,
		NewSourceRecruiteeResource,
		NewSourceRecurlyResource,
		NewSourceRedshiftResource,
		NewSourceRetentlyResource,
		NewSourceRkiCovidResource,
		NewSourceRssResource,
		NewSourceS3Resource,
		NewSourceSalesforceResource,
		NewSourceSalesforceSingerResource,
		NewSourceSalesloftResource,
		NewSourceSapFieldglassResource,
		NewSourceSecodaResource,
		NewSourceSendgridResource,
		NewSourceSendinblueResource,
		NewSourceSenseforceResource,
		NewSourceSentryResource,
		NewSourceSftpResource,
		NewSourceSftpBulkResource,
		NewSourceShopifyResource,
		NewSourceShortioResource,
		NewSourceSlackResource,
		NewSourceSmailyResource,
		NewSourceSmartengageResource,
		NewSourceSmartsheetsResource,
		NewSourceSnapchatMarketingResource,
		NewSourceSnowflakeResource,
		NewSourceSonarCloudResource,
		NewSourceSpacexAPIResource,
		NewSourceSquareResource,
		NewSourceStravaResource,
		NewSourceStripeResource,
		NewSourceSurveymonkeyResource,
		NewSourceSurveySparrowResource,
		NewSourceTempoResource,
		NewSourceTheGuardianAPIResource,
		NewSourceTiktokMarketingResource,
		NewSourceTodoistResource,
		NewSourceTrelloResource,
		NewSourceTrustpilotResource,
		NewSourceTvmazeScheduleResource,
		NewSourceTwilioResource,
		NewSourceTwilioTaskrouterResource,
		NewSourceTwitterResource,
		NewSourceTypeformResource,
		NewSourceUsCensusResource,
		NewSourceVantageResource,
		NewSourceWebflowResource,
		NewSourceWhiskyHunterResource,
		NewSourceWikipediaPageviewsResource,
		NewSourceWoocommerceResource,
		NewSourceXeroResource,
		NewSourceXkcdResource,
		NewSourceYandexMetricaResource,
		NewSourceYotpoResource,
		NewSourceYouniumResource,
		NewSourceYoutubeAnalyticsResource,
		NewSourceZendeskChatResource,
		NewSourceZendeskSunshineResource,
		NewSourceZendeskSupportResource,
		NewSourceZendeskTalkResource,
		NewSourceZenloopResource,
		NewSourceZohoCrmResource,
		NewSourceZoomResource,
		NewSourceZuoraResource,
		NewWorkspaceResource,
	}
}

func (p *AirbyteProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &AirbyteProvider{
			version: version,
		}
	}
}
