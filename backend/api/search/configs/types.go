package configs

import (
	"strings"

	"github.com/Alia5/steaminputdb.com/steamapi"
)

type ControllerType string

const (
	ControllerTypeXbox360                 ControllerType = "controller_xbox360"
	ControllerTypeXboxOne                 ControllerType = "controller_xboxone"
	ControllerTypeXboxElite               ControllerType = "controller_xboxelite"
	ControllerTypePS3                     ControllerType = "controller_ps3"
	ControllerTypePS4                     ControllerType = "controller_ps4"
	ControllerTypePS5                     ControllerType = "controller_ps5"
	ControllerTypePS5Edge                 ControllerType = "controller_ps5_edge"
	ControllerTypeSteamController2015     ControllerType = "controller_steamcontroller_gordon"
	ControllerTypeSteamController         ControllerType = "controller_triton"
	ControllerTypeSteamControllerHeadcrab ControllerType = "controller_steamcontroller_headcrab"
	ControllerTypeSwitchPro               ControllerType = "controller_switch_pro"
	ControllerTypeSteamDeck               ControllerType = "controller_neptune"
	ControllerType8BitDo                  ControllerType = "controller_8bitdo"
	ControllerTypeLegionGoS               ControllerType = "controller_legion_go_s"
	ControllerHoriSteamDeck               ControllerType = "controller_hori_steam"
	ControllerRogAlly                     ControllerType = "controller_rog_ally"
	//
	ControllerTypeGeneric ControllerType = "controller_generic"
	// ControllerTypeNative       ControllerType = "controller_native" --- IGNORE ---
	ControllerTypeMobileTouch ControllerType = "controller_mobile_touch"
	ControllerTypeAndroid     ControllerType = "controller_android"
)

func (c *ControllerType) NiceName() string {
	if name, ok := controllerNiceNames[*c]; ok {
		return name
	}
	return ""
}

var controllerNiceNames map[ControllerType]string = map[ControllerType]string{
	ControllerTypeXbox360:                 "Xbox 360",
	ControllerTypeXboxOne:                 "Xbox One",
	ControllerTypeXboxElite:               "Xbox Elite",
	ControllerTypePS3:                     "DualShock 3",
	ControllerTypePS4:                     "DualShock 4",
	ControllerTypePS5:                     "DualSense",
	ControllerTypePS5Edge:                 "DualSense Edge",
	ControllerTypeSteamController2015:     "Steam Controller (2015)",
	ControllerTypeSteamController:         "Steam Controller",
	ControllerTypeSteamControllerHeadcrab: "Steam Controller (Headcrab)",
	ControllerTypeSwitchPro:               "Nintendo Switch Pro",
	ControllerTypeSteamDeck:               "Steam Deck",
	ControllerType8BitDo:                  "8BitDo",
	ControllerTypeLegionGoS:               "Lenovo Legion Go S",
	ControllerHoriSteamDeck:               "Horipad Steam",
	ControllerRogAlly:                     "ASUS ROG Ally",
	//
	ControllerTypeGeneric: "Generic",
	// ControllerTypeNative:       "Native", --- IGNORE ---
	ControllerTypeMobileTouch: "Mobile Touch",
	ControllerTypeAndroid:     "Android",
}

type RankBy string

const (
	RankedByVote                     RankBy = "vote"
	RankedByPublication              RankBy = "publication"
	RankedByTrend                    RankBy = "trend"
	RankedByTotalUniqueSubscriptions RankBy = "subscriptions"
	RankedByTotalVotesAsc            RankBy = "votes_asc"
	RankedByVotesUp                  RankBy = "votes_up"
	RankedByTextSearch               RankBy = "text_search"
	RankedByPlaytimeTrend            RankBy = "playtime_trend"
	RankedByTotalPlaytime            RankBy = "total_playtime"
	RankedByAveragePlaytimeTrend     RankBy = "avg_playtime_trend"
	RankedByLifetimeAveragePlaytime  RankBy = "lifetime_avg_playtime"
	RankedByPlaytimeSessionsTrend    RankBy = "playtime_sessions_trend"
	RankedByLifetimePlaytimeSessions RankBy = "lifetime_playtime_sessions"
	RankedByLastUpdated              RankBy = "updated"
)

func (s *RankBy) PublishedFileQueryType() steamapi.EPublishedFileQueryType {
	switch strings.ToLower(string(*s)) {
	case "vote":
		return steamapi.EPublishedFileQueryType_k_PublishedFileQueryType_RankedByVote
	case "publication":
		return steamapi.EPublishedFileQueryType_k_PublishedFileQueryType_RankedByPublicationDate
	case "trend":
		return steamapi.EPublishedFileQueryType_k_PublishedFileQueryType_RankedByTrend
	case "subscriptions":
		return steamapi.EPublishedFileQueryType_k_PublishedFileQueryType_RankedByTotalUniqueSubscriptions
	case "votes_asc":
		return steamapi.EPublishedFileQueryType_k_PublishedFileQueryType_RankedByTotalVotesAsc
	case "votes_up":
		return steamapi.EPublishedFileQueryType_k_PublishedFileQueryType_RankedByVotesUp
	case "text_search":
		return steamapi.EPublishedFileQueryType_k_PublishedFileQueryType_RankedByTextSearch
	case "playtime_trend":
		return steamapi.EPublishedFileQueryType_k_PublishedFileQueryType_RankedByPlaytimeTrend
	case "total_playtime":
		return steamapi.EPublishedFileQueryType_k_PublishedFileQueryType_RankedByTotalPlaytime
	case "avg_playtime_trend":
		return steamapi.EPublishedFileQueryType_k_PublishedFileQueryType_RankedByAveragePlaytimeTrend
	case "lifetime_avg_playtime":
		return steamapi.EPublishedFileQueryType_k_PublishedFileQueryType_RankedByLifetimeAveragePlaytime
	case "playtime_sessions_trend":
		return steamapi.EPublishedFileQueryType_k_PublishedFileQueryType_RankedByPlaytimeSessionsTrend
	case "lifetime_playtime_sessions":
		return steamapi.EPublishedFileQueryType_k_PublishedFileQueryType_RankedByLifetimePlaytimeSessions
	case "updated":
		return steamapi.EPublishedFileQueryType_k_PublishedFileQueryType_RankedByLastUpdatedDate
	default:
		return steamapi.EPublishedFileQueryType_k_PublishedFileQueryType_RankedByVote
	}

}
