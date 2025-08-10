package middlewares

// * Access token

type userIDContext string
type tokenTypeContext string
type accessTokenClaimContext string

const UserIDKey userIDContext = "userID"
const TokenTypeIDKey tokenTypeContext = "token-type"
const AccessTokenClaimsIDKey accessTokenClaimContext = "access-token"

// * Refresh token

type refreshTokenClaimsContext string

const RefreshTokenClaimsKey refreshTokenClaimsContext = "refresh-token-claims"
