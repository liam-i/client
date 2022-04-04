// NOTE: This file is GENERATED from json files in actions/json. Run 'yarn build-actions' to regenerate

import type * as Types from '../constants/types/tracker2'

// Constants
export const resetStore = 'common:resetStore' // not a part of tracker2 but is handled by every reducer. NEVER dispatch this
export const typePrefix = 'tracker2:'
export const changeFollow = 'tracker2:changeFollow'
export const closeTracker = 'tracker2:closeTracker'
export const getProofSuggestions = 'tracker2:getProofSuggestions'
export const ignore = 'tracker2:ignore'
export const load = 'tracker2:load'
export const loadNonUserProfile = 'tracker2:loadNonUserProfile'
export const loadedNonUserProfile = 'tracker2:loadedNonUserProfile'
export const proofSuggestionsUpdated = 'tracker2:proofSuggestionsUpdated'
export const showUser = 'tracker2:showUser'
export const updateFollows = 'tracker2:updateFollows'
export const updateResult = 'tracker2:updateResult'
export const updateWotEntries = 'tracker2:updateWotEntries'
export const updatedDetails = 'tracker2:updatedDetails'

// Payload Types
type _ChangeFollowPayload = {readonly guiID: string; readonly follow: boolean}
type _CloseTrackerPayload = {readonly guiID: string}
type _GetProofSuggestionsPayload = undefined
type _IgnorePayload = {readonly guiID: string}
type _LoadNonUserProfilePayload = {readonly assertion: string}
type _LoadPayload = {
  readonly assertion: string
  readonly forceDisplay?: boolean
  readonly fromDaemon?: boolean
  readonly guiID: string
  readonly ignoreCache?: boolean
  readonly reason: string
  readonly inTracker: boolean
}
type _LoadedNonUserProfilePayload = {
  readonly assertion: string
  readonly assertionKey: string
  readonly assertionValue: string
  readonly formattedName?: string
  readonly bio?: string
  readonly description: string
  readonly fullName?: string
  readonly location?: string
  readonly pictureUrl?: string
  readonly siteIcon: Array<Types.SiteIcon>
  readonly siteIconDarkmode: Array<Types.SiteIcon>
  readonly siteIconFull: Array<Types.SiteIcon>
  readonly siteIconFullDarkmode: Array<Types.SiteIcon>
}
type _ProofSuggestionsUpdatedPayload = {readonly suggestions: ReadonlyArray<Types.Assertion>}
type _ShowUserPayload = {readonly asTracker: boolean; readonly username: string; readonly skipNav?: boolean}
type _UpdateFollowsPayload = {
  readonly username: string
  readonly following?: Array<{fullname: string; username: string}>
  readonly followers?: Array<{fullname: string; username: string}>
}
type _UpdateResultPayload = {
  readonly guiID: string
  readonly result: Types.DetailsState
  readonly reason?: string
}
type _UpdateWotEntriesPayload = {
  readonly voucheeUsername: string
  readonly entries: Array<Types.WebOfTrustEntry>
}
type _UpdatedDetailsPayload = {
  readonly guiID: string
  readonly bio: string
  readonly fullname: string
  readonly location: string
  readonly unverifiedFollowersCount: number
  readonly unverifiedFollowingCount: number
  readonly stellarHidden: boolean
  readonly username: string
  readonly teamShowcase: Array<Types.TeamShowcase>
  readonly blocked: boolean
  readonly hidFromFollowers: boolean
}

// Action Creators
/**
 * Short term until new nav, a convenience to show a card from inside the app
 */
export const createShowUser = (payload: _ShowUserPayload): ShowUserPayload => ({payload, type: showUser})
export const createChangeFollow = (payload: _ChangeFollowPayload): ChangeFollowPayload => ({
  payload,
  type: changeFollow,
})
export const createCloseTracker = (payload: _CloseTrackerPayload): CloseTrackerPayload => ({
  payload,
  type: closeTracker,
})
export const createGetProofSuggestions = (
  payload?: _GetProofSuggestionsPayload
): GetProofSuggestionsPayload => ({payload, type: getProofSuggestions})
export const createIgnore = (payload: _IgnorePayload): IgnorePayload => ({payload, type: ignore})
export const createLoad = (payload: _LoadPayload): LoadPayload => ({payload, type: load})
export const createLoadNonUserProfile = (payload: _LoadNonUserProfilePayload): LoadNonUserProfilePayload => ({
  payload,
  type: loadNonUserProfile,
})
export const createLoadedNonUserProfile = (
  payload: _LoadedNonUserProfilePayload
): LoadedNonUserProfilePayload => ({payload, type: loadedNonUserProfile})
export const createProofSuggestionsUpdated = (
  payload: _ProofSuggestionsUpdatedPayload
): ProofSuggestionsUpdatedPayload => ({payload, type: proofSuggestionsUpdated})
export const createUpdateFollows = (payload: _UpdateFollowsPayload): UpdateFollowsPayload => ({
  payload,
  type: updateFollows,
})
export const createUpdateResult = (payload: _UpdateResultPayload): UpdateResultPayload => ({
  payload,
  type: updateResult,
})
export const createUpdateWotEntries = (payload: _UpdateWotEntriesPayload): UpdateWotEntriesPayload => ({
  payload,
  type: updateWotEntries,
})
export const createUpdatedDetails = (payload: _UpdatedDetailsPayload): UpdatedDetailsPayload => ({
  payload,
  type: updatedDetails,
})

// Action Payloads
export type ChangeFollowPayload = {readonly payload: _ChangeFollowPayload; readonly type: typeof changeFollow}
export type CloseTrackerPayload = {readonly payload: _CloseTrackerPayload; readonly type: typeof closeTracker}
export type GetProofSuggestionsPayload = {
  readonly payload: _GetProofSuggestionsPayload
  readonly type: typeof getProofSuggestions
}
export type IgnorePayload = {readonly payload: _IgnorePayload; readonly type: typeof ignore}
export type LoadNonUserProfilePayload = {
  readonly payload: _LoadNonUserProfilePayload
  readonly type: typeof loadNonUserProfile
}
export type LoadPayload = {readonly payload: _LoadPayload; readonly type: typeof load}
export type LoadedNonUserProfilePayload = {
  readonly payload: _LoadedNonUserProfilePayload
  readonly type: typeof loadedNonUserProfile
}
export type ProofSuggestionsUpdatedPayload = {
  readonly payload: _ProofSuggestionsUpdatedPayload
  readonly type: typeof proofSuggestionsUpdated
}
export type ShowUserPayload = {readonly payload: _ShowUserPayload; readonly type: typeof showUser}
export type UpdateFollowsPayload = {
  readonly payload: _UpdateFollowsPayload
  readonly type: typeof updateFollows
}
export type UpdateResultPayload = {readonly payload: _UpdateResultPayload; readonly type: typeof updateResult}
export type UpdateWotEntriesPayload = {
  readonly payload: _UpdateWotEntriesPayload
  readonly type: typeof updateWotEntries
}
export type UpdatedDetailsPayload = {
  readonly payload: _UpdatedDetailsPayload
  readonly type: typeof updatedDetails
}

// All Actions
// prettier-ignore
export type Actions =
  | ChangeFollowPayload
  | CloseTrackerPayload
  | GetProofSuggestionsPayload
  | IgnorePayload
  | LoadNonUserProfilePayload
  | LoadPayload
  | LoadedNonUserProfilePayload
  | ProofSuggestionsUpdatedPayload
  | ShowUserPayload
  | UpdateFollowsPayload
  | UpdateResultPayload
  | UpdateWotEntriesPayload
  | UpdatedDetailsPayload
  | {type: 'common:resetStore', payload: {}}
