// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourcePexelsAPIUpdate struct {
	// API key is required to access pexels api, For getting your's goto https://www.pexels.com/api/documentation and create account for free.
	APIKey string `json:"api_key"`
	// Optional, Desired photo color. Supported colors red, orange, yellow, green, turquoise, blue, violet, pink, brown, black, gray, white or any hexidecimal color code.
	Color *string `json:"color,omitempty"`
	// Optional, The locale of the search you are performing. The current supported locales are 'en-US' 'pt-BR' 'es-ES' 'ca-ES' 'de-DE' 'it-IT' 'fr-FR' 'sv-SE' 'id-ID' 'pl-PL' 'ja-JP' 'zh-TW' 'zh-CN' 'ko-KR' 'th-TH' 'nl-NL' 'hu-HU' 'vi-VN' 'cs-CZ' 'da-DK' 'fi-FI' 'uk-UA' 'el-GR' 'ro-RO' 'nb-NO' 'sk-SK' 'tr-TR' 'ru-RU'.
	Locale *string `json:"locale,omitempty"`
	// Optional, Desired photo orientation. The current supported orientations are landscape, portrait or square
	Orientation *string `json:"orientation,omitempty"`
	// Optional, the search query, Example Ocean, Tigers, Pears, etc.
	Query string `json:"query"`
	// Optional, Minimum photo size. The current supported sizes are large(24MP), medium(12MP) or small(4MP).
	Size *string `json:"size,omitempty"`
}

func (o *SourcePexelsAPIUpdate) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourcePexelsAPIUpdate) GetColor() *string {
	if o == nil {
		return nil
	}
	return o.Color
}

func (o *SourcePexelsAPIUpdate) GetLocale() *string {
	if o == nil {
		return nil
	}
	return o.Locale
}

func (o *SourcePexelsAPIUpdate) GetOrientation() *string {
	if o == nil {
		return nil
	}
	return o.Orientation
}

func (o *SourcePexelsAPIUpdate) GetQuery() string {
	if o == nil {
		return ""
	}
	return o.Query
}

func (o *SourcePexelsAPIUpdate) GetSize() *string {
	if o == nil {
		return nil
	}
	return o.Size
}
