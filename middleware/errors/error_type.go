package errors

type ErrorType int

const (
	Unauthorized ErrorType = iota
	DisallowReregistration
	BadParams
	InvalidRequest
	NotFound
	UniqueViolation
	DatabaseError
	InternalError
	AccessTokenExpired
	RefreshTokenExpired
	SignupRestricted

	InvalidPasscode
	InvalidPassword
	LockedLogin
	LockedSendSMS
	LockedVerifySMS
	LockedMigrateCard
	LockedPreChargeCard

	PhoneNumberExisted
	MemberIDExisted
	InvalidSMSCode
	SMSCodeExpired
	IncorrectPasscode
	IncorrectPassword
	SameNumbersProhibited
	ConsecutiveNumbersProhibited
	MemberCanceled
	MemberSuspended

	MemberNotFound
	MemberProfileNotFound
	ProfileExisted
	InvalidInstallationID

	CardNotFound
	InvalidCardInfo
	CardInactive
	InvalidCardStatus

	TerminalNotFound
	InvalidRefreshToken
	RefreshTokenRequired
	UnregisteredTerminal
	TerminalInactive

	CouponNotFound
	CouponExpired
	MemberCouponNotFound
	CouponUseHistoryExisted
)

var errorHTTPCode = map[ErrorType]int{
	Unauthorized:           401,
	DisallowReregistration: 403,
	BadParams:              400,
	InvalidRequest:         400,
	NotFound:               404,
	UniqueViolation:        409,
	DatabaseError:          500,
	InternalError:          500,
	AccessTokenExpired:     401,
	RefreshTokenExpired:    401,
	SignupRestricted:       400,

	InvalidPasscode:     400,
	InvalidPassword:     400,
	LockedLogin:         400,
	LockedSendSMS:       400,
	LockedVerifySMS:     400,
	LockedMigrateCard:   400,
	LockedPreChargeCard: 400,

	PhoneNumberExisted:           400,
	MemberIDExisted:              400,
	InvalidSMSCode:               400,
	SMSCodeExpired:               400,
	IncorrectPasscode:            400,
	IncorrectPassword:            400,
	SameNumbersProhibited:        400,
	ConsecutiveNumbersProhibited: 400,
	MemberCanceled:               400,
	MemberSuspended:              400,

	MemberNotFound:        404,
	MemberProfileNotFound: 404,
	ProfileExisted:        409,
	InvalidInstallationID: 400,

	CardNotFound:      404,
	InvalidCardInfo:   400,
	CardInactive:      400,
	InvalidCardStatus: 400,

	TerminalNotFound:     404,
	InvalidRefreshToken:  400,
	RefreshTokenRequired: 400,
	UnregisteredTerminal: 400,
	TerminalInactive:     400,

	CouponNotFound:          404,
	CouponExpired:           400,
	MemberCouponNotFound:    404,
	CouponUseHistoryExisted: 400,
}

var errorCode = map[ErrorType]string{
	Unauthorized:           "UNAUTHORIZED",
	DisallowReregistration: "DISALLOW_REREGISTRATION",
	BadParams:              "BAD_PARAMS",
	InvalidRequest:         "INVALID_REQUEST",
	NotFound:               "NOT_FOUND",
	UniqueViolation:        "UNIQUE_VIOLATION",
	DatabaseError:          "DATABASE_ERROR",
	InternalError:          "INTERNAL_SERVER_ERROR",
	AccessTokenExpired:     "ACCESS_TOKEN_EXPIRED",
	RefreshTokenExpired:    "REFRESH_TOKEN_EXPIRED",
	SignupRestricted:       "SIGNUP_RESTRICTED",

	InvalidPasscode:     "INVALID_PASSCODE",
	InvalidPassword:     "INVALID_PASSWORD",
	LockedLogin:         "LOGIN_RESTRICTED",
	LockedSendSMS:       "SMS_SENDING_RESTRICTED",
	LockedVerifySMS:     "SMS_VERIFY_LOCKED",
	LockedMigrateCard:   "LOCKED_MIGRATE_CARD",
	LockedPreChargeCard: "LOCKED_PRECHARGE_CARD",

	PhoneNumberExisted:           "PHONE_NUMBER_EXISTED",
	MemberIDExisted:              "MEMBER_ID_EXISTED",
	InvalidSMSCode:               "INVALID_SMS",
	SMSCodeExpired:               "SMS_EXPIRED",
	IncorrectPasscode:            "INCORRECT_PASSCODE",
	IncorrectPassword:            "INCORRECT_PASSWORD",
	SameNumbersProhibited:        "SAME_NUMBERS_PROHIBITED",
	ConsecutiveNumbersProhibited: "CONSECUTIVE_NUMBERS_PROHIBITED",
	MemberCanceled:               "MEMBER_CANCELED",
	MemberSuspended:              "MEMBER_SUSPENDED",

	MemberNotFound:        "MEMBER_NOT_FOUND",
	MemberProfileNotFound: "MEMBER_PROFILE_NOT_FOUND",
	ProfileExisted:        "EXISTED_PROFILE",
	InvalidInstallationID: "INVALID_INSTALLATION_ID",

	CardNotFound:      "CARD_NOT_FOUND",
	InvalidCardInfo:   "INVALID_CARD_INFO",
	CardInactive:      "CARD_INACTIVE",
	InvalidCardStatus: "INVALID_CARD_STATUS",

	TerminalNotFound:     "TERMINAL_NOT_FOUND",
	InvalidRefreshToken:  "INVALID_REFRESH_TOKEN",
	RefreshTokenRequired: "REFRESH_TOKEN_REQUIRED",
	UnregisteredTerminal: "UNREGISTERED_TERMINAL",
	TerminalInactive:     "TERMINAL_INACTIVE",

	CouponNotFound:          "COUPON_NOT_FOUND",
	CouponExpired:           "COUPON_EXPIRED",
	MemberCouponNotFound:    "MEMBER_COUPON_NOT_FOUND",
	CouponUseHistoryExisted: "COUPON_USE_HISTORY_EXISTED",
}

func (t ErrorType) HTTPCode() int {
	return errorHTTPCode[t]
}

func (t ErrorType) Code() string {
	return errorCode[t]
}
